package kivikmock

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-kivik/kivik"
	"github.com/go-kivik/kivik/driver"
)

var pool *mockDriver

func init() {
	pool = &mockDriver{
		clients: make(map[string]*kivikmock),
	}
	kivik.Register("kivikmock", pool)
}

type mockDriver struct {
	sync.Mutex
	counter int
	clients map[string]*kivikmock
}

var _ driver.Driver = &mockDriver{}

func (d *mockDriver) NewClient(dsn string) (driver.Client, error) {
	d.Lock()
	defer d.Unlock()

	c, ok := d.clients[dsn]
	if !ok {
		return nil, errors.New("kivikmock: no available connection found")
	}
	c.opened++
	return c, nil
}

// New creates a kivik client connection and a mock to manage expectations.
func New() (*kivik.Client, Mock, error) {
	pool.Lock()
	dsn := fmt.Sprintf("kivikmock_%d", pool.counter)
	pool.counter++

	kmock := &kivikmock{dsn: dsn, drv: pool, ordered: true}
	pool.clients[dsn] = kmock
	pool.Unlock()

	return kmock.open()
}
