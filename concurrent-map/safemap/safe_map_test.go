package safemap

import (
	"sync"
	"testing"
)

const hits = 16

func BenchmarkSafeMapMostlyHits(b *testing.B) {
	safeMap := NewSafeMap()
	for i := 0; i < hits; i++ {
		safeMap.Put(i, i)
	}
	for i := 0; i < hits*100; i++ {
		safeMap.Get(i % hits)
	}
}

func BenchmarkSyncMapMostlyHits(b *testing.B) {
	syncMap := sync.Map{}
	for i := 0; i < hits; i++ {
		syncMap.Store(i, i)
	}
	for i := 0; i < hits*100; i++ {
		syncMap.Load(i % hits)
	}
}
