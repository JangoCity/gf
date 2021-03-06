// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// go test *.go -bench=".*" -benchmem

package glist

import (
    "testing"
)

var l = New()

func Benchmark_PushBack(b *testing.B) {
    for i := 0; i < b.N; i++ {
        l.PushBack(i)
    }
}

func Benchmark_PopFront(b *testing.B) {
    for i := 0; i < b.N; i++ {
        l.PopFront()
    }
}

func Benchmark_PushFront(b *testing.B) {
    for i := 0; i < b.N; i++ {
        l.PushFront(i)
    }
}

func Benchmark_PopBack(b *testing.B) {
    for i := 0; i < b.N; i++ {
        l.PopBack()
    }
}

func Benchmark_Len(b *testing.B) {
    for i := 0; i < b.N; i++ {
        l.Len()
    }
}

