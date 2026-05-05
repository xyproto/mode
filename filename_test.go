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
	if Detect("justfile") != Just {
		t.Fail()
	}
	if Detect("local.ini") != Ini {
		t.Fail()
	}
	if Detect("test.abw") != Abiword {
		t.Fail()
	}
	if Detect("test.docx") != DOCX {
		t.Fail()
	}
	if Detect("test.odt") != LibreOffice {
		t.Fail()
	}
	if Detect("test.rtf") != RTF {
		t.Fail()
	}
	if Detect("test.wg") != WordGrinder {
		t.Fail()
	}
	if Detect("config.yml") != YAML {
		t.Fail()
	}
	if Detect("config.yaml") != YAML {
		t.Fail()
	}
	if Detect("Cargo.toml") != TOML {
		t.Fail()
	}
	if Detect("main.tf") != HCL {
		t.Fail()
	}
	if Detect("vars.tfvars") != HCL {
		t.Fail()
	}
	if Detect("service.proto") != Protobuf {
		t.Fail()
	}
	if Detect("data.csv") != CSV {
		t.Fail()
	}
	if Detect("data.tsv") != CSV {
		t.Fail()
	}
	if Detect("config.dhall") != Dhall {
		t.Fail()
	}
	if Detect("script.janet") != Janet {
		t.Fail()
	}
	if Detect("script.nu") != Nushell {
		t.Fail()
	}
	if Detect("config.pkl") != Pkl {
		t.Fail()
	}
	if Detect("shader.wgsl") != WGSL {
		t.Fail()
	}
	if Detect("BUILD") != Bazel {
		t.Fail()
	}
	if Detect("WORKSPACE") != Bazel {
		t.Fail()
	}
	if Detect("BUILD.bazel") != Bazel {
		t.Fail()
	}
	if Detect("shader.hlsl") != Shader {
		t.Fail()
	}
}
