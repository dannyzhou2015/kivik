// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package mock

import (
	"context"

	"github.com/dannyzhou2015/kivik/v4/driver"
)

// DB mocks a driver.DB
type DB struct {
	// ID is a unique identifier for the DB instance.
	ID                   string
	AllDocsFunc          func(ctx context.Context, options map[string]interface{}) (driver.Rows, error)
	GetFunc              func(ctx context.Context, docID string, options map[string]interface{}) (*driver.Document, error)
	CreateDocFunc        func(ctx context.Context, doc interface{}, options map[string]interface{}) (docID, rev string, err error)
	PutFunc              func(ctx context.Context, docID string, doc interface{}, options map[string]interface{}) (rev string, err error)
	DeleteFunc           func(ctx context.Context, docID, rev string, options map[string]interface{}) (newRev string, err error)
	StatsFunc            func(ctx context.Context) (*driver.DBStats, error)
	CompactFunc          func(ctx context.Context) error
	CompactViewFunc      func(ctx context.Context, ddocID string) error
	ViewCleanupFunc      func(ctx context.Context) error
	SecurityFunc         func(ctx context.Context) (*driver.Security, error)
	SetSecurityFunc      func(ctx context.Context, security *driver.Security) error
	ChangesFunc          func(ctx context.Context, options map[string]interface{}) (driver.Changes, error)
	PutAttachmentFunc    func(ctx context.Context, docID, rev string, att *driver.Attachment, options map[string]interface{}) (newRev string, err error)
	GetAttachmentFunc    func(ctx context.Context, docID, filename string, options map[string]interface{}) (*driver.Attachment, error)
	DeleteAttachmentFunc func(ctx context.Context, docID, rev, filename string, options map[string]interface{}) (newRev string, err error)
	QueryFunc            func(context.Context, string, string, map[string]interface{}) (driver.Rows, error)
}

var _ driver.DB = &DB{}

// AllDocs calls db.AllDocsFunc
func (db *DB) AllDocs(ctx context.Context, options map[string]interface{}) (driver.Rows, error) {
	return db.AllDocsFunc(ctx, options)
}

// Get calls db.GetFunc
func (db *DB) Get(ctx context.Context, docID string, opts map[string]interface{}) (*driver.Document, error) {
	return db.GetFunc(ctx, docID, opts)
}

// CreateDoc calls db.CreateDocFunc
func (db *DB) CreateDoc(ctx context.Context, doc interface{}, opts map[string]interface{}) (string, string, error) {
	return db.CreateDocFunc(ctx, doc, opts)
}

// Put calls db.PutFunc
func (db *DB) Put(ctx context.Context, docID string, doc interface{}, opts map[string]interface{}) (string, error) {
	return db.PutFunc(ctx, docID, doc, opts)
}

// Delete calls db.DeleteFunc
func (db *DB) Delete(ctx context.Context, docID, rev string, opts map[string]interface{}) (string, error) {
	return db.DeleteFunc(ctx, docID, rev, opts)
}

// Stats calls db.StatsFunc
func (db *DB) Stats(ctx context.Context) (*driver.DBStats, error) {
	return db.StatsFunc(ctx)
}

// Compact calls db.CompactFunc
func (db *DB) Compact(ctx context.Context) error {
	return db.CompactFunc(ctx)
}

// CompactView calls db.CompactViewFunc
func (db *DB) CompactView(ctx context.Context, docID string) error {
	return db.CompactViewFunc(ctx, docID)
}

// ViewCleanup calls db.ViewCleanupFunc
func (db *DB) ViewCleanup(ctx context.Context) error {
	return db.ViewCleanupFunc(ctx)
}

// Security calls db.SecurityFunc
func (db *DB) Security(ctx context.Context) (*driver.Security, error) {
	return db.SecurityFunc(ctx)
}

// SetSecurity calls db.SetSecurityFunc
func (db *DB) SetSecurity(ctx context.Context, security *driver.Security) error {
	return db.SetSecurityFunc(ctx, security)
}

// Changes calls db.ChangesFunc
func (db *DB) Changes(ctx context.Context, opts map[string]interface{}) (driver.Changes, error) {
	return db.ChangesFunc(ctx, opts)
}

// PutAttachment calls db.PutAttachmentFunc
func (db *DB) PutAttachment(ctx context.Context, docID, rev string, att *driver.Attachment, opts map[string]interface{}) (string, error) {
	return db.PutAttachmentFunc(ctx, docID, rev, att, opts)
}

// GetAttachment calls db.GetAttachmentFunc
func (db *DB) GetAttachment(ctx context.Context, docID, filename string, opts map[string]interface{}) (*driver.Attachment, error) {
	return db.GetAttachmentFunc(ctx, docID, filename, opts)
}

// DeleteAttachment calls db.DeleteAttachmentFunc
func (db *DB) DeleteAttachment(ctx context.Context, docID, rev, filename string, opts map[string]interface{}) (string, error) {
	return db.DeleteAttachmentFunc(ctx, docID, rev, filename, opts)
}

// Query calls db.QueryFunc
func (db *DB) Query(ctx context.Context, ddoc, view string, opts map[string]interface{}) (driver.Rows, error) {
	return db.QueryFunc(ctx, ddoc, view, opts)
}

// Finder mocks a driver.DB and driver.Finder
type Finder struct {
	*DB
	CreateIndexFunc func(context.Context, string, string, interface{}) error
	DeleteIndexFunc func(context.Context, string, string) error
	FindFunc        func(context.Context, interface{}) (driver.Rows, error)
	GetIndexesFunc  func(context.Context) ([]driver.Index, error)
	ExplainFunc     func(context.Context, interface{}) (*driver.QueryPlan, error)
}

// nolint:staticcheck
var _ driver.Finder = &Finder{}

// CreateIndex calls db.CreateIndexFunc
func (db *Finder) CreateIndex(ctx context.Context, ddoc, name string, index interface{}) error {
	return db.CreateIndexFunc(ctx, ddoc, name, index)
}

// DeleteIndex calls db.DeleteIndexFunc
func (db *Finder) DeleteIndex(ctx context.Context, ddoc, name string) error {
	return db.DeleteIndexFunc(ctx, ddoc, name)
}

// Find calls db.FindFunc
func (db *Finder) Find(ctx context.Context, query interface{}) (driver.Rows, error) {
	return db.FindFunc(ctx, query)
}

// GetIndexes calls db.GetIndexesFunc
func (db *Finder) GetIndexes(ctx context.Context) ([]driver.Index, error) {
	return db.GetIndexesFunc(ctx)
}

// Explain calls db.ExplainFunc
func (db *Finder) Explain(ctx context.Context, query interface{}) (*driver.QueryPlan, error) {
	return db.ExplainFunc(ctx, query)
}

// OptsFinder mocks a driver.DB and driver.Finder
type OptsFinder struct {
	*DB
	CreateIndexFunc func(context.Context, string, string, interface{}, map[string]interface{}) error
	DeleteIndexFunc func(context.Context, string, string, map[string]interface{}) error
	FindFunc        func(context.Context, interface{}, map[string]interface{}) (driver.Rows, error)
	GetIndexesFunc  func(context.Context, map[string]interface{}) ([]driver.Index, error)
	ExplainFunc     func(context.Context, interface{}, map[string]interface{}) (*driver.QueryPlan, error)
}

var _ driver.OptsFinder = &OptsFinder{}

// CreateIndex calls db.CreateIndexFunc
func (db *OptsFinder) CreateIndex(ctx context.Context, ddoc, name string, index interface{}, opts map[string]interface{}) error {
	return db.CreateIndexFunc(ctx, ddoc, name, index, opts)
}

// DeleteIndex calls db.DeleteIndexFunc
func (db *OptsFinder) DeleteIndex(ctx context.Context, ddoc, name string, opts map[string]interface{}) error {
	return db.DeleteIndexFunc(ctx, ddoc, name, opts)
}

// Find calls db.FindFunc
func (db *OptsFinder) Find(ctx context.Context, query interface{}, opts map[string]interface{}) (driver.Rows, error) {
	return db.FindFunc(ctx, query, opts)
}

// GetIndexes calls db.GetIndexesFunc
func (db *OptsFinder) GetIndexes(ctx context.Context, opts map[string]interface{}) ([]driver.Index, error) {
	return db.GetIndexesFunc(ctx, opts)
}

// Explain calls db.ExplainFunc
func (db *OptsFinder) Explain(ctx context.Context, query interface{}, opts map[string]interface{}) (*driver.QueryPlan, error) {
	return db.ExplainFunc(ctx, query, opts)
}

// Flusher mocks a driver.DB and driver.Flusher
type Flusher struct {
	*DB
	FlushFunc func(context.Context) error
}

var _ driver.Flusher = &Flusher{}

// Flush calls db.FlushFunc
func (db *Flusher) Flush(ctx context.Context) error {
	return db.FlushFunc(ctx)
}

// MetaGetter mocks a driver.DB and driver.MetaGetter
type MetaGetter struct {
	*DB
	GetMetaFunc func(context.Context, string, map[string]interface{}) (int64, string, error)
}

var _ driver.MetaGetter = &MetaGetter{}

// GetMeta calls db.GetMetaFunc
func (db *MetaGetter) GetMeta(ctx context.Context, docID string, opts map[string]interface{}) (int64, string, error) {
	return db.GetMetaFunc(ctx, docID, opts)
}

// Copier mocks a driver.DB and driver.Copier.
type Copier struct {
	*DB
	CopyFunc func(context.Context, string, string, map[string]interface{}) (string, error)
}

var _ driver.Copier = &Copier{}

// Copy calls db.CopyFunc
func (db *Copier) Copy(ctx context.Context, target, source string, options map[string]interface{}) (string, error) {
	return db.CopyFunc(ctx, target, source, options)
}

// AttachmentMetaGetter mocks a driver.DB and driver.AttachmentMetaGetter
type AttachmentMetaGetter struct {
	*DB
	GetAttachmentMetaFunc func(ctx context.Context, docID, filename string, options map[string]interface{}) (*driver.Attachment, error)
}

var _ driver.AttachmentMetaGetter = &AttachmentMetaGetter{}

// GetAttachmentMeta calls db.GetAttachmentMetaFunc
func (db *AttachmentMetaGetter) GetAttachmentMeta(ctx context.Context, docID, filename string, options map[string]interface{}) (*driver.Attachment, error) {
	return db.GetAttachmentMetaFunc(ctx, docID, filename, options)
}

// DesignDocer mocks a driver.DB and driver.DesignDocer
type DesignDocer struct {
	*DB
	DesignDocsFunc func(context.Context, map[string]interface{}) (driver.Rows, error)
}

var _ driver.DesignDocer = &DesignDocer{}

// DesignDocs calls db.DesignDocsFunc
func (db *DesignDocer) DesignDocs(ctx context.Context, options map[string]interface{}) (driver.Rows, error) {
	return db.DesignDocsFunc(ctx, options)
}

// LocalDocer mocks a driver.DB and driver.DesignDocer
type LocalDocer struct {
	*DB
	LocalDocsFunc func(context.Context, map[string]interface{}) (driver.Rows, error)
}

var _ driver.LocalDocer = &LocalDocer{}

// LocalDocs calls db.LocalDocsFunc
func (db *LocalDocer) LocalDocs(ctx context.Context, options map[string]interface{}) (driver.Rows, error) {
	return db.LocalDocsFunc(ctx, options)
}

// Purger mocks a driver.DB and driver.Purger
type Purger struct {
	*DB
	PurgeFunc func(context.Context, map[string][]string) (*driver.PurgeResult, error)
}

var _ driver.Purger = &Purger{}

// Purge calls db.PurgeFunc
func (db *Purger) Purge(ctx context.Context, docMap map[string][]string) (*driver.PurgeResult, error) {
	return db.PurgeFunc(ctx, docMap)
}

// BulkGetter mocks a driver.DB and driver.BulkGetter
type BulkGetter struct {
	*DB
	BulkGetFunc func(context.Context, []driver.BulkGetReference, map[string]interface{}) (driver.Rows, error)
}

var _ driver.BulkGetter = &BulkGetter{}

// BulkGet calls db.BulkGetFunc
func (db *BulkGetter) BulkGet(ctx context.Context, docs []driver.BulkGetReference, opts map[string]interface{}) (driver.Rows, error) {
	return db.BulkGetFunc(ctx, docs, opts)
}

// DBCloser mocks driver.DB and driver.DBCloser
type DBCloser struct {
	*DB
	CloseFunc func(context.Context) error
}

var _ driver.DBCloser = &DBCloser{}

// Close calls db.CloseFunc
func (db *DBCloser) Close(ctx context.Context) error {
	return db.CloseFunc(ctx)
}

// RevsDiffer mocks a driver.DB and driver.RevsDiffer.
type RevsDiffer struct {
	*BulkDocer
	RevsDiffFunc func(context.Context, interface{}) (driver.Rows, error)
}

var _ driver.RevsDiffer = &RevsDiffer{}

// RevsDiff calls db.RevsDiffFunc
func (db *RevsDiffer) RevsDiff(ctx context.Context, revMap interface{}) (driver.Rows, error) {
	return db.RevsDiffFunc(ctx, revMap)
}

// PartitionedDB mocks a driver.DB and a driver.PartitionedDB.
type PartitionedDB struct {
	*DB
	PartitionStatsFunc func(context.Context, string) (*driver.PartitionStats, error)
}

// PartitionStats calls db.PartitionStatsFunc.
func (db *PartitionedDB) PartitionStats(ctx context.Context, name string) (*driver.PartitionStats, error) {
	return db.PartitionStatsFunc(ctx, name)
}
