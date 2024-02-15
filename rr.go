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

import (
	"sync"

	"github.com/overflowingd/emlb/pkg/ll"
)

type roundRobinNode[V any] struct {
	*ll.Node[V]
	retained bool
}

type roundRobin struct {
	sync.Mutex

	current *ll.Node[struct{}]
	nodes   []*roundRobinNode[struct{}]
}

func NewRoundRobin(cap uint64) (Algorithm, error) {
	if cap < 1 {
		return nil, ErrNoVariant
	}

	var (
		list    = ll.Make[struct{}](cap)
		current = list.Head
		nodes   = make([]*roundRobinNode[struct{}], cap)
	)

	for i := uint64(0); i < cap; i++ {
		nodes[i] = &roundRobinNode[struct{}]{
			Node: current,
		}

		current = current.Next
	}

	return &roundRobin{
		current: list.Head,
		nodes:   nodes,
	}, nil
}

// Next makes a round across items returning every item sequentially if they were not omitted
func (r *roundRobin) Next() (uint64, error) {
	r.Lock()
	defer r.Unlock()

	if r.current.Next == nil {
		return 0, ErrNoVariant
	}

	current := r.current
	r.current = r.current.Next
	return current.I, nil
}

func (r *roundRobin) Retain(i uint64) (bool, error) {
	r.Lock()
	defer r.Unlock()

	current := r.nodes[i]
	if current.retained {
		return false, nil
	}

	current.retained = true

	prev, next := current.Prev, current.Next
	if prev.I == current.I && next.I == current.I {
		prev.Next = nil
		next.Prev = nil
		return true, nil
	}

	prev.Next = next
	next.Prev = prev

	if r.current.I == current.I {
		r.current = next
	}

	return true, nil
}
