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

package driver

import jsoniter "github.com/json-iterator/go"

// Changes is an iterator of the database changes feed.
type Changes interface {
	// Next is called to populate *Change with the next value in the changes
	// feed.
	//
	// Next should return io.EOF when the changes feed is closed by request.
	Next(*Change) error
	// Close closes the changes feed iterator.
	Close() error
	// LastSeq returns the last change update sequence.
	LastSeq() string
	// Pending returns the count of remaining items in the feed
	Pending() int64
	// ETag returns the unquoted ETag header, if present.
	ETag() string
}

// Change represents the changes to a single document.
type Change struct {
	// ID is the document ID to which the change relates.
	ID string `json:"id"`
	// Seq is the update sequence for the changes feed.
	Seq string `json:"seq"`
	// Deleted is set to true for the changes feed, if the document has been
	// deleted.
	Deleted bool `json:"deleted"`
	// Changes represents a list of document leaf revisions for the /_changes
	// endpoint.
	Changes ChangedRevs `json:"changes"`
	// Doc is the raw, un-decoded JSON document. This is only populated when
	// include_docs=true is set.
	Doc jsoniter.RawMessage `json:"doc"`
}

// ChangedRevs represents a "changes" field of a result in the /_changes stream.
type ChangedRevs []string

// UnmarshalJSON satisfies the json.Unmarshaler interface
func (c *ChangedRevs) UnmarshalJSON(data []byte) error {
	var changes []struct {
		Rev string `json:"rev"`
	}
	if err := json.Unmarshal(data, &changes); err != nil {
		return err
	}
	revs := ChangedRevs(make([]string, len(changes)))
	for i, change := range changes {
		revs[i] = change.Rev
	}
	*c = revs
	return nil
}
