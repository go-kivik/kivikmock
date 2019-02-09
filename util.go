package kivikmock

import (
	"fmt"

	"github.com/go-kivik/kivik"
	"github.com/go-kivik/kivik/driver"
)

func optionsString(opt map[string]interface{}) string {
	if opt == nil {
		return "\n\t- has any options"
	}
	return fmt.Sprintf("\n\t- has options: %v", opt)
}

func errorString(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("\n\t- should return error: %s", err)
}

func kivikStats2driverStats(i *kivik.DBStats) *driver.DBStats {
	var cluster *driver.ClusterStats
	if i.Cluster != nil {
		c := driver.ClusterStats(*i.Cluster)
		cluster = &c
	}
	return &driver.DBStats{
		Name:           i.Name,
		CompactRunning: i.CompactRunning,
		DocCount:       i.DocCount,
		DeletedCount:   i.DeletedCount,
		UpdateSeq:      i.UpdateSeq,
		DiskSize:       i.DiskSize,
		ActiveSize:     i.ActiveSize,
		ExternalSize:   i.ExternalSize,
		Cluster:        cluster,
		RawResponse:    i.RawResponse,
	}
}
