package kivikmock

import (
	"encoding/json"
	"fmt"

	"github.com/go-kivik/kivik/v4/driver"
)

// ExpectedDBClose is used to manage *kivik.Client.Close expectation returned
// by Mock.ExpectClose.
type ExpectedDBClose struct {
	commonExpectation
	callback func() error
}

func (e *ExpectedDBClose) method(v bool) string {
	if v {
		return "DB.Close(ctx)"
	}
	return "DB.Close()"
}

func (e *ExpectedDBClose) met(ex expectation) bool {
	return true
}

// WillReturnError allows setting an error for *kivik.Client.Close action.
func (e *ExpectedDBClose) WillReturnError(err error) *ExpectedDBClose {
	e.err = err
	return e
}

// WillExecute sets a callback function to be called with any inputs to the
// original function. Any values returned by the callback will be returned as
// if generated by the driver.
func (e *ExpectedDBClose) WillExecute(cb func() error) *ExpectedDBClose {
	e.callback = cb
	return e
}

func (e *ExpectedDBClose) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).Close()", e.dbo().name, e.dbo().id)
	extra := delayString(e.delay)
	extra += errorString(e.err)
	if extra != "" {
		msg += " which:" + extra
	}
	return msg
}

func (e *ExpectedAllDocs) String() string {
	var rets []string
	if e.ret0 != nil {
		rets = []string{fmt.Sprintf("should return: %d results", e.ret0.count())}
	}
	return dbStringer("AllDocs", &e.commonExpectation, withOptions, nil, rets)
}

func jsonDoc(i interface{}) string {
	jsonText, err := json.Marshal(i)
	if err != nil {
		return fmt.Sprintf("<invalid json:%s>", err)
	}
	return string(jsonText)
}

func (e *ExpectedBulkGet) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).BulkGet() which:", e.dbo().name, e.dbo().id)
	if e.arg0 == nil {
		msg += "\n\t- has any doc references"
	} else {
		msg += fmt.Sprintf("\n\t- has doc references: %v", jsonDoc(e.arg0))
	}
	msg += optionsString(e.options)
	if e.ret0 != nil {
		msg += fmt.Sprintf("\n\t- should return: %d results", e.ret0.count())
	}
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

func (e *ExpectedFind) String() string {
	var opts, rets []string
	if e.arg0 == nil {
		opts = append(opts, "has any query")
	} else {
		opts = append(opts, fmt.Sprintf("has query: %v", e.arg0))
	}
	if e.ret0 != nil {
		rets = []string{fmt.Sprintf("should return: %d results", e.ret0.count())}
	}
	return dbStringer("Find", &e.commonExpectation, withOptions, opts, rets)
}

// WithQuery sets the expected query for the Find() call.
func (e *ExpectedFind) WithQuery(query interface{}) *ExpectedFind {
	e.arg0 = query
	return e
}

func (e *ExpectedCreateIndex) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).CreateIndex() which:", e.dbo().name, e.dbo().id)
	if e.arg0 == "" {
		msg += "\n\t- has any ddoc"
	} else {
		msg += "\n\t- has ddoc: " + e.arg0
	}
	if e.arg1 == "" {
		msg += "\n\t- has any name"
	} else {
		msg += "\n\t- has name: " + e.arg1
	}
	if e.arg2 == nil {
		msg += "\n\t- has any index"
	} else {
		msg += fmt.Sprintf("\n\t- has index: %v", e.arg2)
	}
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

// WithDDocID sets the expected ddocID value for the DB.CreateIndex() call.
func (e *ExpectedCreateIndex) WithDDocID(ddocID string) *ExpectedCreateIndex {
	e.arg0 = ddocID
	return e
}

// WithName sets the expected name value for the DB.CreateIndex() call.
func (e *ExpectedCreateIndex) WithName(name string) *ExpectedCreateIndex {
	e.arg1 = name
	return e
}

// WithIndex sets the expected index value for the DB.CreateIndex() call.
func (e *ExpectedCreateIndex) WithIndex(index interface{}) *ExpectedCreateIndex {
	e.arg2 = index
	return e
}

func (e *ExpectedGetIndexes) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).GetIndexes()", e.dbo().name, e.dbo().id)
	var extra string
	if e.ret0 != nil {
		extra += fmt.Sprintf("\n\t- should return indexes: %v", e.ret0)
	}
	extra += delayString(e.delay)
	extra += errorString(e.err)
	if extra != "" {
		msg += " which:" + extra
	}
	return msg
}

func (e *ExpectedDeleteIndex) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).DeleteIndex() which:", e.dbo().name, e.dbo().id)
	msg += fieldString("ddoc", e.arg0)
	msg += fieldString("name", e.arg1)
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

// WithDDoc sets the expected ddoc to be passed to the DB.DeleteIndex() call.
func (e *ExpectedDeleteIndex) WithDDoc(ddoc string) *ExpectedDeleteIndex {
	e.arg0 = ddoc
	return e
}

// WithName sets the expected name to be passed to the DB.DeleteIndex() call.
func (e *ExpectedDeleteIndex) WithName(name string) *ExpectedDeleteIndex {
	e.arg1 = name
	return e
}

func (e *ExpectedExplain) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).Explain() which:", e.dbo().name, e.dbo().id)
	if e.arg0 == nil {
		msg += "\n\t- has any query"
	} else {
		msg += fmt.Sprintf("\n\t- has query: %v", e.arg0)
	}
	if e.ret0 != nil {
		msg += fmt.Sprintf("\n\t- should return query plan: %v", e.ret0)
	}
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

// WithQuery sets the expected query for the Explain() call.
func (e *ExpectedExplain) WithQuery(query interface{}) *ExpectedExplain {
	e.arg0 = query
	return e
}

func (e *ExpectedCreateDoc) String() string {
	msg := fmt.Sprintf("call to DB(%s#%d).CreateDoc() which:", e.dbo().name, e.dbo().id)
	if e.arg0 == nil {
		msg += "\n\t- has any doc"
	} else {
		msg += fmt.Sprintf("\n\t- has doc: %s", jsonDoc(e.arg0))
	}
	msg += optionsString(e.options)
	if e.ret0 != "" {
		msg += "\n\t- should return docID: " + e.ret0
	}
	if e.ret1 != "" {
		msg += "\n\t- should return rev: " + e.ret1
	}
	msg += delayString(e.delay)
	msg += errorString(e.err)
	return msg
}

// WithDoc sets the expected doc for the call to CreateDoc().
func (e *ExpectedCreateDoc) WithDoc(doc interface{}) *ExpectedCreateDoc {
	e.arg0 = doc
	return e
}

const (
	withOptions = 1 << iota
)

func dbStringer(methodName string, e *commonExpectation, flags int, opts []string, rets []string) string {
	msg := fmt.Sprintf("call to DB(%s#%d).%s()", e.db.name, e.db.id, methodName)
	var extra string
	for _, c := range opts {
		extra += "\n\t- " + c
	}
	if flags&withOptions > 0 {
		extra += optionsString(e.options)
	}
	for _, c := range rets {
		extra += "\n\t- " + c
	}
	extra += delayString(e.delay)
	extra += errorString(e.err)
	if extra != "" {
		msg += " which:" + extra
	}
	return msg
}

func (e *ExpectedCompact) String() string {
	return dbStringer("Compact", &e.commonExpectation, 0, nil, nil)
}

func (e *ExpectedViewCleanup) String() string {
	return dbStringer("ViewCleanup", &e.commonExpectation, 0, nil, nil)
}

func (e *ExpectedPut) String() string {
	custom := []string{}
	if e.arg0 == "" {
		custom = append(custom, "has any docID")
	} else {
		custom = append(custom, fmt.Sprintf("has docID: %s", e.arg0))
	}
	if e.arg1 == nil {
		custom = append(custom, "has any doc")
	} else {
		custom = append(custom, fmt.Sprintf("has doc: %s", jsonDoc(e.arg1)))
	}
	return dbStringer("Put", &e.commonExpectation, withOptions, custom, nil)
}

// WithDocID sets the expectation for the docID passed to the DB.Put() call.
func (e *ExpectedPut) WithDocID(docID string) *ExpectedPut {
	e.arg0 = docID
	return e
}

// WithDoc sets the expectation for the doc passed to the DB.Put() call.
func (e *ExpectedPut) WithDoc(doc interface{}) *ExpectedPut {
	e.arg1 = doc
	return e
}

func (e *ExpectedGetRev) String() string {
	var opts []string
	if e.arg0 == "" {
		opts = append(opts, "has any docID")
	} else {
		opts = append(opts, fmt.Sprintf("has docID: %s", e.arg0))
	}
	var rets []string
	if e.ret0 != "" {
		rets = append(rets, "should return rev: "+e.ret0)
	}
	return dbStringer("GetRev", &e.commonExpectation, withOptions, opts, rets)
}

// WithDocID sets the expectation for the docID passed to the DB.GetRev call.
func (e *ExpectedGetRev) WithDocID(docID string) *ExpectedGetRev {
	e.arg0 = docID
	return e
}

func (e *ExpectedFlush) String() string {
	return dbStringer("Flush", &e.commonExpectation, 0, nil, nil)
}

func (e *ExpectedDeleteAttachment) String() string {
	var opts, rets []string
	if e.arg0 == "" {
		opts = append(opts, "has any docID")
	} else {
		opts = append(opts, "has docID: "+e.arg0)
	}
	if e.arg1 == "" {
		opts = append(opts, "has any filename")
	} else {
		opts = append(opts, "has filename: "+e.arg1)
	}
	if e.ret0 != "" {
		rets = append(rets, "should return rev: "+e.ret0)
	}
	return dbStringer("DeleteAttachment", &e.commonExpectation, withOptions, opts, rets)
}

// WithDocID sets the expectation for the docID passed to the DB.DeleteAttachment() call.
func (e *ExpectedDeleteAttachment) WithDocID(docID string) *ExpectedDeleteAttachment {
	e.arg0 = docID
	return e
}

// WithFilename sets the expectation for the filename passed to the DB.DeleteAttachment() call.
func (e *ExpectedDeleteAttachment) WithFilename(filename string) *ExpectedDeleteAttachment {
	e.arg1 = filename
	return e
}

func (e *ExpectedDelete) String() string {
	var opts, rets []string
	if e.arg0 == "" {
		opts = append(opts, "has any docID")
	} else {
		opts = append(opts, "has docID: "+e.arg0)
	}
	if e.ret0 != "" {
		rets = append(rets, "should return rev: "+e.ret0)
	}
	return dbStringer("Delete", &e.commonExpectation, withOptions, opts, rets)
}

// WithDocID sets the expectation for the docID passed to the DB.Delete() call.
func (e *ExpectedDelete) WithDocID(docID string) *ExpectedDelete {
	e.arg0 = docID
	return e
}

func (e *ExpectedCopy) String() string {
	var opts, rets []string
	if e.arg0 == "" {
		opts = append(opts, "has any targetID")
	} else {
		opts = append(opts, "has targetID: "+e.arg0)
	}
	if e.arg1 == "" {
		opts = append(opts, "has any sourceID")
	} else {
		opts = append(opts, "has sourceID: "+e.arg1)
	}
	if e.ret0 != "" {
		rets = append(rets, "should return rev: "+e.ret0)
	}
	return dbStringer("Copy", &e.commonExpectation, withOptions, opts, rets)
}

// WithTargetID sets the expectation for the docID passed to the DB.Copy() call.
func (e *ExpectedCopy) WithTargetID(docID string) *ExpectedCopy {
	e.arg0 = docID
	return e
}

// WithSourceID sets the expectation for the docID passed to the DB.Copy() call.
func (e *ExpectedCopy) WithSourceID(docID string) *ExpectedCopy {
	e.arg1 = docID
	return e
}

func (e *ExpectedCompactView) String() string {
	var opts []string
	if e.arg0 == "" {
		opts = []string{"has any ddocID"}
	} else {
		opts = []string{"has ddocID: " + e.arg0}
	}
	return dbStringer("CompactView", &e.commonExpectation, 0, opts, nil)
}

// WithDDoc sets the expected design doc name for the call to DB.CompactView().
func (e *ExpectedCompactView) WithDDoc(ddocID string) *ExpectedCompactView {
	e.arg0 = ddocID
	return e
}

func (e *ExpectedGet) String() string {
	var opts, rets []string
	if e.arg0 == "" {
		opts = []string{"has any docID"}
	} else {
		opts = []string{"has docID: " + e.arg0}
	}
	if e.ret0 != nil {
		rets = []string{fmt.Sprintf("should return document with rev: %s", e.ret0.Rev)}
	}
	return dbStringer("Get", &e.commonExpectation, withOptions, opts, rets)
}

// WithDocID sets the expected docID for the DB.Get() call.
func (e *ExpectedGet) WithDocID(docID string) *ExpectedGet {
	e.arg0 = docID
	return e
}

func (e *ExpectedGetAttachmentMeta) String() string {
	var opts, rets []string
	if e.arg0 == "" {
		opts = []string{"has any docID"}
	} else {
		opts = []string{"has docID: " + e.arg0}
	}
	if e.arg1 == "" {
		opts = append(opts, "has any filename")
	} else {
		opts = append(opts, "has filename: "+e.arg1)
	}
	if e.ret0 != nil {
		rets = []string{fmt.Sprintf("should return attachment: %s", e.ret0.Filename)}
	}
	return dbStringer("GetAttachmentMeta", &e.commonExpectation, withOptions, opts, rets)
}

// WithDocID sets the expectation for the docID passed to the DB.GetAttachmentMeta() call.
func (e *ExpectedGetAttachmentMeta) WithDocID(docID string) *ExpectedGetAttachmentMeta {
	e.arg0 = docID
	return e
}

// WithFilename sets the expectation for the doc passed to the DB.GetAttachmentMeta() call.
func (e *ExpectedGetAttachmentMeta) WithFilename(filename string) *ExpectedGetAttachmentMeta {
	e.arg1 = filename
	return e
}

func (e *ExpectedLocalDocs) String() string {
	var rets []string
	if e.ret0 != nil {
		rets = []string{fmt.Sprintf("should return: %d results", e.ret0.count())}
	}
	return dbStringer("LocalDocs", &e.commonExpectation, withOptions, nil, rets)
}

func (e *ExpectedPurge) String() string {
	var opts, rets []string
	if e.arg0 == nil {
		opts = []string{"has any docRevMap"}
	} else {
		opts = []string{fmt.Sprintf("has docRevMap: %v", e.arg0)}
	}
	if e.ret0 != nil {
		rets = []string{fmt.Sprintf("should return result: %v", e.ret0)}
	}
	return dbStringer("Purge", &e.commonExpectation, 0, opts, rets)
}

// WithDocRevMap sets the expected docRevMap for the call to DB.Purge().
func (e *ExpectedPurge) WithDocRevMap(docRevMap map[string][]string) *ExpectedPurge {
	e.arg0 = docRevMap
	return e
}

func (e *ExpectedPutAttachment) String() string {
	var opts, rets []string
	if e.arg0 == "" {
		opts = append(opts, "has any docID")
	} else {
		opts = append(opts, fmt.Sprintf("has docID: %s", e.arg0))
	}
	if e.arg1 == nil {
		opts = append(opts, "has any attachment")
	} else {
		opts = append(opts, fmt.Sprintf("has attachment: %s", e.arg1.Filename))
	}
	if e.ret0 != "" {
		rets = append(rets, "should return rev: "+e.ret0)
	}
	return dbStringer("PutAttachment", &e.commonExpectation, withOptions, opts, rets)
}

// WithDocID sets the expectation for the docID passed to the DB.PutAttachment() call.
func (e *ExpectedPutAttachment) WithDocID(docID string) *ExpectedPutAttachment {
	e.arg0 = docID
	return e
}

// WithAttachment sets the expectation for the rev passed to the DB.PutAttachment() call.
func (e *ExpectedPutAttachment) WithAttachment(att *driver.Attachment) *ExpectedPutAttachment {
	e.arg1 = att
	return e
}

func (e *ExpectedQuery) String() string {
	var opts, rets []string
	if e.arg0 == "" {
		opts = append(opts, "has any ddocID")
	} else {
		opts = append(opts, "has ddocID: "+e.arg0)
	}
	if e.arg1 == "" {
		opts = append(opts, "has any view")
	} else {
		opts = append(opts, "has view: "+e.arg1)
	}
	if e.ret0 != nil {
		rets = []string{fmt.Sprintf("should return: %d results", e.ret0.count())}
	}
	return dbStringer("Query", &e.commonExpectation, withOptions, opts, rets)
}

// WithDDocID sets the expected ddocID value for the DB.Query() call.
func (e *ExpectedQuery) WithDDocID(ddocID string) *ExpectedQuery {
	e.arg0 = ddocID
	return e
}

// WithView sets the expected view value for the DB.Query() call.
func (e *ExpectedQuery) WithView(view string) *ExpectedQuery {
	e.arg1 = view
	return e
}

func (e *ExpectedSecurity) String() string {
	var rets []string
	if e.ret0 != nil {
		rets = append(rets, fmt.Sprintf("should return: %s", jsonDoc(e.ret0)))
	}
	return dbStringer("Security", &e.commonExpectation, 0, nil, rets)
}

func (e *ExpectedSetSecurity) String() string {
	var opts []string
	if e.arg0 == nil {
		opts = append(opts, "has any security object")
	} else {
		opts = append(opts, fmt.Sprintf("has security object: %v", e.arg0))
	}
	return dbStringer("SetSecurity", &e.commonExpectation, 0, opts, nil)
}

// WithSecurity sets the expected security object for the DB.SetSecurity() call.
func (e *ExpectedSetSecurity) WithSecurity(sec *driver.Security) *ExpectedSetSecurity {
	e.arg0 = sec
	return e
}

func (e *ExpectedStats) String() string {
	var rets []string
	if e.ret0 != nil {
		rets = append(rets, fmt.Sprintf("should return stats: %v", e.ret0))
	}
	return dbStringer("Stats", &e.commonExpectation, 0, nil, rets)
}

func (e *ExpectedBulkDocs) String() string {
	var opts, rets []string
	if e.arg0 == nil {
		opts = append(opts, "has any docs")
	} else {
		opts = append(opts, fmt.Sprintf("has: %d docs", len(e.arg0)))
	}
	if e.ret0 != nil {
		rets = append(rets, fmt.Sprintf("should return: %d results", len(e.ret0)))
	}
	return dbStringer("BulkDocs", &e.commonExpectation, withOptions, opts, rets)
}

func (e *ExpectedGetAttachment) String() string {
	var opts, rets []string
	if e.arg0 == "" {
		opts = append(opts, "has any docID")
	} else {
		opts = append(opts, "has docID: "+e.arg0)
	}
	if e.arg1 == "" {
		opts = append(opts, "has any filename")
	} else {
		opts = append(opts, "has filename: "+e.arg1)
	}
	if e.ret0 != nil {
		rets = append(rets, "should return attachment: "+e.ret0.Filename)
	}
	return dbStringer("GetAttachment", &e.commonExpectation, withOptions, opts, rets)
}

// WithDocID sets the expectation for the docID passed to the DB.GetAttachment() call.
func (e *ExpectedGetAttachment) WithDocID(docID string) *ExpectedGetAttachment {
	e.arg0 = docID
	return e
}

// WithFilename sets the expectation for the filename passed to the DB.GetAttachment() call.
func (e *ExpectedGetAttachment) WithFilename(docID string) *ExpectedGetAttachment {
	e.arg1 = docID
	return e
}

func (e *ExpectedDesignDocs) String() string {
	var rets []string
	if e.ret0 != nil {
		rets = []string{fmt.Sprintf("should return: %d results", e.ret0.count())}
	}
	return dbStringer("DesignDocs", &e.commonExpectation, withOptions, nil, rets)
}

func (e *ExpectedChanges) String() string {
	var rets []string
	if e.ret0 != nil {
		rets = []string{fmt.Sprintf("should return: %d results", e.ret0.count())}
	}
	return dbStringer("Changes", &e.commonExpectation, withOptions, nil, rets)
}

func (e *ExpectedRevsDiff) String() string {
	var rets, opts []string
	if e.ret0 != nil {
		rets = []string{fmt.Sprintf("should return: %d results", e.ret0.count())}
	}
	if e.arg0 != nil {
		opts = []string{fmt.Sprintf("with revMap: %v", e.arg0)}
	} else {
		opts = []string{"has any revMap"}
	}
	return dbStringer("RevsDiff", &e.commonExpectation, 0, opts, rets)
}

// WithRevLookup sets the expectation for the rev lookup passed to the
// DB.RevsDiff() call.
func (e *ExpectedRevsDiff) WithRevLookup(revLookup interface{}) *ExpectedRevsDiff {
	e.arg0 = revLookup
	return e
}

func (e *ExpectedPartitionStats) String() string {
	var rets, opts []string
	if e.ret0 != nil {
		rets = []string{fmt.Sprintf("should return: %s", jsonDoc(e.ret0))}
	}
	if e.arg0 != "" {
		opts = []string{fmt.Sprintf("with name: %v", e.arg0)}
	} else {
		opts = []string{"has any name"}
	}
	return dbStringer("PartitionStats", &e.commonExpectation, 0, opts, rets)
}

// WithName sets the expectation for the partition name passed to the
// DB.PartitionStats() call.
func (e *ExpectedPartitionStats) WithName(name string) *ExpectedPartitionStats {
	e.arg0 = name
	return e
}
