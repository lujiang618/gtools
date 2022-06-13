package utils

import (
	"context"
	"sync"
	"testing"
	"time"
)

// go test -timeout 30s -run ^TestNewDistributedLock$ gitlab.dxbim.com/products/dxyp-common/utils -v
func TestNewDistributedLock(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(t *testing.T, i int) {
			defer wg.Done()
			var endPoints = []string{"127.0.0.1:2379"}
			lock, err := NewDistributedLock(endPoints, "test")
			if err != nil {
				t.Log(err)
			}

			defer lock.Close()
			if err := lock.TryLock(context.TODO()); err != nil {
				t.Logf("try lock failed,err:%v", err)
				return
			}

			t.Logf("%d get Lock..............................", i)
			time.Sleep(1 * time.Second)

			if err := lock.Unlock(context.TODO()); err != nil {
				t.Log(err)
			}

			t.Logf("%d execute finish....", i)
		}(t, i)
	}

	wg.Wait()
}

//  go test -timeout 30s -run ^TestDistributedLock_Close$ gitlab.dxbim.com/products/dxyp-common/utils -v
func TestDistributedLock_Close(t *testing.T) {
	var endPoints = []string{"127.0.0.1:2379"}
	lock, err := NewDistributedLock(endPoints, "test")
	if err != nil {
		t.Fatal(err)
	}

	defer lock.Close()
}
