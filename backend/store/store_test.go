package store

import (
    "sync"
    "testing"
)

func TestAddConcurrent(t *testing.T) {
    s := NewMemoryStore()
    var wg sync.WaitGroup
    for i := 0; i < 500; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            s.Add(Menu{Name: "จานใหม่", Price: 50, Type: "test"})
        }()
    }
    wg.Wait()

    got := len(s.List(""))
    want := 3 + 500          // 3 จานตั้งต้น บวกจานใหม่ 500 จาน
    if got != want {
        t.Errorf("อยากได้ %d จาน แต่ได้ %d จาน", want, got)
    }
}