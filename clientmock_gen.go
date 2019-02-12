/* This file is auto-generated. Do not edit it! */

package kivikmock

import "github.com/go-kivik/kivik"

var _ = kivik.EndKeySuffix // To ensure a reference to kivik package

// ExpectAllDBs queues an expectation that AllDBs will be called.
func (c *MockClient) ExpectAllDBs() *ExpectedAllDBs {
	e := &ExpectedAllDBs{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectClose queues an expectation that Close will be called.
func (c *MockClient) ExpectClose() *ExpectedClose {
	e := &ExpectedClose{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectClusterSetup queues an expectation that ClusterSetup will be called.
func (c *MockClient) ExpectClusterSetup() *ExpectedClusterSetup {
	e := &ExpectedClusterSetup{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectClusterStatus queues an expectation that ClusterStatus will be called.
func (c *MockClient) ExpectClusterStatus() *ExpectedClusterStatus {
	e := &ExpectedClusterStatus{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectDBExists queues an expectation that DBExists will be called.
func (c *MockClient) ExpectDBExists() *ExpectedDBExists {
	e := &ExpectedDBExists{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectDestroyDB queues an expectation that DestroyDB will be called.
func (c *MockClient) ExpectDestroyDB() *ExpectedDestroyDB {
	e := &ExpectedDestroyDB{}
	c.expected = append(c.expected, e)
	return e
}

// ExpectPing queues an expectation that Ping will be called.
func (c *MockClient) ExpectPing() *ExpectedPing {
	e := &ExpectedPing{}
	c.expected = append(c.expected, e)
	return e
}
