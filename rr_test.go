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
	"testing"

	"github.com/stretchr/testify/require"
)

const RRCap = uint64(10)

func TestNew(t *testing.T) {
	rr, err := NewRoundRobin(RRCap)
	require.Nil(t, err)
	require.NotNil(t, rr)
}

func TestNewCap0(t *testing.T) {
	rr, err := NewRoundRobin(0)
	require.ErrorIs(t, err, ErrNoVariant)
	require.Nil(t, rr)
}

func TestNext(t *testing.T) {
	rr, _ := NewRoundRobin(RRCap)

	i, err := rr.Next()
	require.Nil(t, err)
	require.Equal(t, i, uint64(0))
}

func TestRetainTail(t *testing.T) {
	rr, _ := NewRoundRobin(RRCap)

	ok, err := rr.Retain(RRCap - 1)
	require.Nil(t, err)
	require.True(t, ok)

	ok, err = rr.Retain(RRCap - 1)
	require.Nil(t, err)
	require.False(t, ok)
}

func TestRetainCurrent(t *testing.T) {
	rr, _ := NewRoundRobin(RRCap)

	i, _ := rr.Next()
	ok, err := rr.Retain(i + 1)
	require.Nil(t, err)
	require.True(t, ok)

	j, _ := rr.Next()
	require.Nil(t, err)
	require.NotEqual(t, j, i+1)
	require.Equal(t, j, uint64(i+2))
}

func TestRetainAllLeft(t *testing.T) {
	rr, _ := NewRoundRobin(RRCap)
	for i := RRCap - 1; i < RRCap; i-- {
		ok, err := rr.Retain(i)
		require.Nil(t, err)
		require.True(t, ok)
	}

	i, err := rr.Next()
	require.Equal(t, i, uint64(0))
	require.ErrorIs(t, err, ErrNoVariant)
}

func TestRetainAllRight(t *testing.T) {
	rr, _ := NewRoundRobin(RRCap)
	for i := uint64(0); i < RRCap; i++ {
		ok, err := rr.Retain(i)
		require.Nil(t, err)
		require.True(t, ok)
	}

	i, err := rr.Next()
	require.Equal(t, i, uint64(0))
	require.ErrorIs(t, err, ErrNoVariant)
}

func TestRecoverMid(t *testing.T) {
	rr, _ := NewRoundRobin(RRCap)

	rr.Retain(RRCap / 2)

	ok := rr.Recover(RRCap / 2)
	require.True(t, ok)
}
