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
	Next() (I, uint64, error)
}

type balancer[I any] struct {
	algorithm Algorithm
	items     []I
}

func New[I any](
	algorithm Algorithm,
	items []I,
) (Balancer[I], error) {
	return &balancer[I]{
		algorithm: algorithm,
		items:     items,
	}, nil
}

// Next returns a next item according to a load balancing algorithm
func (b *balancer[I]) Next() (I, uint64, error) {
	next, err := b.algorithm.Next()
	if err != nil {
		return *new(I), 0, fmt.Errorf("algorithm: next: %w", err)
	}

	return b.items[next], next, err
}