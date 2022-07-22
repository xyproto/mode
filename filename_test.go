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
	if Detect("90-libvirt-mydevice") != Config {
		t.Fail()
	}
	if Detect("a90-libvirt-mydevice") == Config {
		t.Fail()
	}
	if Detect("/tmp/man.XXXXtweZrK") != ManPage {
		t.Fail()
	}
}
