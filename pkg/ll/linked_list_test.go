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

package ll_test

import (
	"testing"

	"github.com/overflowingd/emlb/pkg/ll"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	items := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := ll.New(items)

	require.NotNil(t, list.Head)
	require.NotNil(t, list.Tail)
	require.Equal(t, list.Head.Prev.I, list.Tail.I)
	require.Equal(t, list.Head.Prev.Value, list.Tail.Value)
	require.Equal(t, list.Head.I, list.Tail.Next.I)
	require.Equal(t, list.Head.Value, list.Tail.Next.Value)

	current := list.Head
	for i := 0; i < len(items); i++ {
		require.Equal(t, current.Value, items[i])
		require.Equal(t, current.I, uint64(i))
		current = current.Next
	}
}

func TestNewEmpty(t *testing.T) {
	items := []uint64{}
	list := ll.New(items)

	require.Nil(t, list.Head)
	require.Nil(t, list.Tail)
}

func TestNewLen1(t *testing.T) {
	items := []uint64{1}
	list := ll.New(items)

	require.NotNil(t, list.Head)
	require.NotNil(t, list.Tail)

	require.Equal(t, list.Head.I, uint64(0))
	require.Equal(t, list.Head.Value, uint64(1))
	require.Equal(t, *list.Head, *list.Tail)
}

const MakeCapacity = 10

func TestMake(t *testing.T) {
	list := ll.Make[uint64](MakeCapacity)

	require.NotNil(t, list.Head)
	require.NotNil(t, list.Tail)
	require.Equal(t, list.Head.Prev.I, list.Tail.I)
	require.Equal(t, list.Head.Prev.Value, list.Tail.Value)
	require.Equal(t, list.Head.I, list.Tail.Next.I)
	require.Equal(t, list.Head.Value, list.Tail.Next.Value)

	current := list.Head
	for i := 0; i < MakeCapacity; i++ {
		require.Equal(t, current.Value, uint64(0))
		require.Equal(t, current.I, uint64(i))
		current = current.Next
	}
}

func TestMakeEmpty(t *testing.T) {
	list := ll.Make[uint64](0)

	require.Nil(t, list.Head)
	require.Nil(t, list.Tail)
}

func TestMakeCap1(t *testing.T) {
	list := ll.Make[uint64](1)

	require.NotNil(t, list.Head)
	require.NotNil(t, list.Tail)

	require.Equal(t, list.Head.I, uint64(0))
	require.Equal(t, list.Head.Value, uint64(0))
	require.Equal(t, *list.Head, *list.Tail)
}
