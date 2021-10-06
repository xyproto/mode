package mode

import (
	"testing"
)

func TestDetect(t *testing.T) {
	if Detect("test.sh") != Shell {
		t.Fail()
	}
	if Detect(".zshrc") != Shell {
		t.Fail()
	}
	if Detect("main.go") != Go {
		t.Fail()
	}
	if Detect("main.s") != Assembly {
		t.Fail()
	}
	if Detect("main.asm") != Assembly {
		t.Fail()
	}
}