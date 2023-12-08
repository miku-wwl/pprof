package data_test

import (
	"pprof/data/block"
	"pprof/data/cpu"
	"pprof/data/goroutine"
	"pprof/data/mem"
	"pprof/data/mutex"
	"testing"
)

func BenchmarkData(b *testing.B) {
	b.Run("block", func(b *testing.B) {
		o := block.Block{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})

	b.Run("cpu", func(b *testing.B) {
		o := cpu.Cpu{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})

	b.Run("mem", func(b *testing.B) {
		o := mem.Mem{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})

	b.Run("goroutine", func(b *testing.B) {
		o := goroutine.Goroutine{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})
	b.Run("mutex", func(b *testing.B) {
		o := mutex.Mutex{}
		for i := 0; i < b.N; i++ {
			o.Run()
		}
	})
}
