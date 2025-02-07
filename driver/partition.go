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

import (
	"context"

	jsoniter "github.com/json-iterator/go"
)

// PartitionedDB is an optional interface that may be satisfied by a DB to
// support querying partitoin-specific information.
type PartitionedDB interface {
	// PartitionStats returns information about the named partition.
	PartitionStats(ctx context.Context, name string) (*PartitionStats, error)
}

// PartitionStats contains partition statistics.
type PartitionStats struct {
	DBName          string
	DocCount        int64
	DeletedDocCount int64
	Partition       string
	ActiveSize      int64
	ExternalSize    int64
	RawResponse     jsoniter.RawMessage
}
