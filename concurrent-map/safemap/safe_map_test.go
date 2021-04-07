package safemap

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

// copy from https://github.com/golang/go/blob/b345a306a0258085b65081cf2dadc238dc7e26ee/src/sync/map_reference_test.go#L15
type mapInterface interface {
	Load(interface{}) (interface{}, bool)
	Store(key, value interface{})
	LoadOrStore(key, value interface{}) (actual interface{}, loaded bool)
	LoadAndDelete(key interface{}) (value interface{}, loaded bool)
	Delete(interface{})
	Range(func(key, value interface{}) (shouldContinue bool))
}

type bench struct {
	setup func(*testing.B, mapInterface)
	perG  func(b *testing.B, pb *testing.PB, i int, m mapInterface)
}

func benchMap(b *testing.B, bench bench) {
	for _, m := range [...]mapInterface{NewWrapConcurrentMap(), NewSafeMap(), &sync.Map{}} {
		b.Run(fmt.Sprintf("%T", m), func(b *testing.B) {
			// m = reflect.New(reflect.TypeOf(m).Elem()).Interface().(mapInterface)
			if bench.setup != nil {
				bench.setup(b, m)
			}

			b.ResetTimer()

			var i int64
			b.RunParallel(func(pb *testing.PB) {
				id := int(atomic.AddInt64(&i, 1) - 1)
				bench.perG(b, pb, id*b.N, m)
			})
		})
	}
}

func BenchmarkLoadMostlyHits(b *testing.B) {
	const hits, misses = 1023, 1

	benchMap(b, bench{
		setup: func(_ *testing.B, m mapInterface) {
			for i := 0; i < hits; i++ {
				m.Store(strconv.Itoa(i), i)
			}
			// Prime the map to get it into a steady state.
			for i := 0; i < hits*2; i++ {
				m.Load(strconv.Itoa(i % hits))
			}
		},

		perG: func(b *testing.B, pb *testing.PB, i int, m mapInterface) {
			for ; pb.Next(); i++ {
				m.Load(strconv.Itoa(i % (hits + misses)))
			}
		},
	})
}
func BenchmarkLoadMostlyMisses(b *testing.B) {
	const hits, misses = 1, 1023

	benchMap(b, bench{
		setup: func(_ *testing.B, m mapInterface) {
			for i := 0; i < hits; i++ {
				m.Store(strconv.Itoa(i), i)
			}
			// Prime the map to get it into a steady state.
			for i := 0; i < hits*2; i++ {
				m.Load(strconv.Itoa(i % hits))
			}
		},

		perG: func(b *testing.B, pb *testing.PB, i int, m mapInterface) {
			for ; pb.Next(); i++ {
				m.Load(strconv.Itoa(i % (hits + misses)))
			}
		},
	})
}

func BenchmarkLoadOrStoreBalanced(b *testing.B) {
	const hits, misses = 128, 128

	benchMap(b, bench{
		setup: func(b *testing.B, m mapInterface) {
			for i := 0; i < hits; i++ {
				m.Store(strconv.Itoa(i), i)
			}
			// Prime the map to get it into a steady state.
			for i := 0; i < hits*2; i++ {
				m.Load(strconv.Itoa(i % hits))
			}
		},

		perG: func(b *testing.B, pb *testing.PB, i int, m mapInterface) {
			for ; pb.Next(); i++ {
				j := i % (hits + misses)
				if j < hits {
					m.Load(strconv.Itoa((j)))
				} else {
					m.Store(strconv.Itoa(j), i)
				}
			}
		},
	})
}
