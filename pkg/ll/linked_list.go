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

package ll

type List[V any] struct {
	Head *Node[V]
	Tail *Node[V]
}

func Make[V any](capacity uint64) *List[V] {
	var (
		list = List[V]{}
	)

	if capacity < 1 {
		return &list
	}

	list.Head = &Node[V]{}
	if capacity < 2 {
		list.Tail = list.Head
		list.Head.Next = list.Head
		list.Head.Prev = list.Head
		return &list
	}

	capacity--

	current := list.Head
	for i := uint64(0); i < capacity; i++ {
		current.Next = &Node[V]{
			I:    i + 1,
			Prev: current,
		}

		current = current.Next
	}

	list.Tail = current
	list.Tail.Next = list.Head
	list.Head.Prev = list.Tail
	return &list
}

func New[V any](items []V) *List[V] {
	var (
		size = uint64(len(items))
		list = List[V]{}
	)

	if size < 1 {
		return &list
	}

	list.Head = &Node[V]{
		Value: items[0],
	}

	if size < 2 {
		list.Tail = list.Head
		list.Head.Next = list.Head
		list.Head.Prev = list.Head
		return &list
	}

	items = items[1:]
	size--

	current := list.Head
	for i := uint64(0); i < size; i++ {
		current.Next = &Node[V]{
			I:     i + 1,
			Value: items[i],
			Prev:  current,
		}

		current = current.Next
	}

	list.Tail = current
	list.Tail.Next = list.Head
	list.Head.Prev = list.Tail
	return &list
}

type Node[V any] struct {
	Value V
	I     uint64
	Next  *Node[V]
	Prev  *Node[V]
}
