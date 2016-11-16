package matchers_test

import (
	"testing"

	"github.com/apoydence/onpar/matchers"
)

func TestAbove(t *testing.T) {
	t.Parallel()

	m := matchers.Above(101)

	_, err := m.Match(99.0)
	if err == nil {
		t.Error("expected err to not be nil")
	}

	_, err = m.Match(103.0)
	if err != nil {
		t.Error("expected err to be nil")
	}

	_, err = m.Match("invalid")
	if err == nil {
		t.Error("expected err to not be nil")
	}
}
