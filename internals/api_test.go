package internals

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	_, err := Hello_world()

	if err != nil {
		t.Fail()
	}
}
