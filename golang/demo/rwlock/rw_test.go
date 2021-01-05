package rwlock

import (
	"sync"
	"testing"
)

func TestRWLock(t *testing.T) {
	x := sync.RWMutex{}
	x.RLocker()
}
