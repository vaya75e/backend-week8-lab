package counter

type Counter struct {
    n int
}

func (c *Counter) Add()       { c.n = c.n + 1 }
func (c *Counter) Value() int { return c.n }
counter/counter_test.go
package counter

import (
    "sync"
    "testing"
)

func TestCounter100(t *testing.T) {
    var c Counter
    var wg sync.WaitGroup
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            c.Add()
        }()
    }
    wg.Wait()
    if c.Value() != 100 {
        t.Errorf("อยากได้ 100 แต่ได้ %d", c.Value())
    }
}