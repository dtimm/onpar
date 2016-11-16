package matchers_test

import (
	"testing"

	"github.com/apoydence/onpar/matchers"
)

func TestLen(t *testing.T) {
	m := matchers.Len(5)

	_, err := m.Match([]int{1, 2, 3})
	if err == nil {
		t.Error("expected err to not be nil")
	}

	_, err = m.Match([]int{1, 2, 3, 4, 5})
	if err != nil {
		t.Error("expected err to be nil")
	}

	_, err = m.Match(make(chan int, 3))
	if err == nil {
		t.Error("expected err to not be nil")
	}

	c := make(chan int, 10)
	for i := 0; i < 5; i++ {
		c <- i
	}
	_, err = m.Match(c)
	if err != nil {
		t.Error("expected err to be nil")
	}

	_, err = m.Match(map[int]bool{1: true, 2: true, 3: true})
	if err == nil {
		t.Error("expected err to not be nil")
	}

	_, err = m.Match(map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true})
	if err != nil {
		t.Error("expected err to be nil")
	}

	_, err = m.Match("123")
	if err == nil {
		t.Error("expected err to not be nil")
	}

	_, err = m.Match("12345")
	if err != nil {
		t.Error("expected err to be nil")
	}

	_, err = m.Match(123)
	if err == nil {
		t.Error("expected err to not be nil")
	}
}
