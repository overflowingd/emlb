// Copyright 2024 Vadim Kharitonovich
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package emlb

import "sync/atomic"

type roundRobin struct {
	i   uint64
	cap uint64
}

func NewRoundRobin(cap uint64) (Algorithm, error) {
	if cap < 1 {
		return nil, ErrNoVariant
	}

	return &roundRobin{
		i:   0,
		cap: cap,
	}, nil
}

// Next makes a round across items returning every item sequentially if they were not omitted
func (rr *roundRobin) Next() (uint64, error) {
	t := rr.i
	atomic.AddUint64(&rr.i, 1)
	return t % rr.cap, nil
}
