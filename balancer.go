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

import "fmt"

type Balancer[I any] interface {
	// Len returns the number of elements in balancer set
	Len() uint64
	// Next returns next variant from a balancing set using an underlying algorithm
	Next() (I, uint64, error)
	// Retain excludes item from balancing algorithm until recover call.
	// Returns true if item retentained ok
	Retain(uint64) (bool, error)
	// Recover item after retention.
	// Returns true if item recovered ok.
	Recover(uint64) bool
}

type balancer[I any] struct {
	algorithm Algorithm
	items     []I
	len       uint64
}

func New[I any](
	algorithm Algorithm,
	items []I,
) (Balancer[I], error) {
	return &balancer[I]{
		algorithm: algorithm,
		items:     items,
		len:       uint64(len(items)),
	}, nil
}

func (b *balancer[I]) Len() uint64 {
	return b.len
}

func (b *balancer[I]) Next() (I, uint64, error) {
	next, err := b.algorithm.Next()
	if err != nil {
		return *new(I), 0, fmt.Errorf("algorithm: next: %w", err)
	}

	return b.items[next], next, err
}

func (b *balancer[I]) Retain(i uint64) (bool, error) {
	ok, err := b.algorithm.Retain(i)
	if err != nil {
		return false, fmt.Errorf("algorithm: retain: %w", err)
	}

	return ok, nil
}

func (b *balancer[I]) Recover(i uint64) bool {
	return b.algorithm.Recover(i)
}
