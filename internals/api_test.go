package internals

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := Hello_world()
	if result["hello"] != "world" {
		t.Errorf("Result was incorrect, expected: %s, got: %s.", "world", result["hello"])
	}
}
