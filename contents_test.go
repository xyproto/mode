package mode

import (
	"os"
	"testing"
)

var examples = map[string]Mode{
	"#!/bin/bash\necho Hello\n":              Shell,
	"#!/usr/bin/env python\nprint(\"hi\")\n": Python,
	"":                                       Blank,
	"# Hello":                                Config,
	"#include <stdio.h>":                     Blank, // hard to tell if this is C or C++, should be detected by filename, not by contents
	"<?xml version=\"1.0\" encoding=":        XML,
	"::\n[source,C]":                         ReStructured,
	"\" This file is blabla\nand":            Vim,
	"x = 42\ny = 32\nalso (\n  z = 5\n)\n":   Config,
	"#!/usr/bin/env python3\n\n":             Python,
	"\n\n<asdf\n\n    >   \n\n":              XML, // if code starts with "<" and ends with ">"
	"{\\rtf1\\ansi\\deff0":                   RTF,
	"WordGrinder: uncompressed document":     WordGrinder,
	"<abiword version=\"2.1\">":              Abiword,
	"PK\x03\x04\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00mimetypeapplication/vnd.oasis.opendocument.text": LibreOffice,
	"%YAML 1.2\n---\nfoo: bar\n":                        YAML,
	"---\nfoo: bar\n":                                   YAML,
	"syntax = \"proto3\";\npackage foo;\n":              Protobuf,
	"terraform {\n  required_version = \">= 1.0\"\n}\n": HCL,
	"resource \"aws_instance\" \"web\" {\n}\n":          HCL,
	"#!/usr/bin/env nu\necho hello\n":                   Nushell,
	"amends \"package://example.com/config.pkl\"\n":     Pkl,
	"@vertex\nfn vs_main() -> vec4<f32> {\n}\n":         WGSL,
	"@fragment\nfn fs_main() -> vec4<f32> {\n}\n":       WGSL,
}

var exampleFiles = map[string]Mode{
	"testfiles/META": Config,
}

func TestGoAssembly(t *testing.T) {
	exampleString := "// func getisar0() uint64\nTEXT ·getisar0(SB),NOSPLIT,$0\n  // get Instruction Set Attributes 0 into R0\n  MRS	ID_AA64ISAR0_EL1, R0\n  MOVD	R0, ret+0(FP)\n  RET"
	m := SimpleDetectBytes([]byte(exampleString))
	if m != GoAssembly {
		t.Fatalf("Expected %s got %s for example Go/Plan9 style Assembly", Mode(GoAssembly).String(), m.String())
	}
}

func TestSimpleDetect(t *testing.T) {
	for s, targetMode := range examples {
		if m := SimpleDetect(s); m != targetMode {
			t.Fatalf("Expected %s got %s", targetMode.String(), m.String())
		}
	}
	for filename, targetMode := range exampleFiles {
		data, err := os.ReadFile(filename)
		if err != nil {
			t.Fatalf("Could not read %s: %v\n", filename, err)
		}
		if m := SimpleDetect(string(data)); m != targetMode {
			t.Fatalf("Expected %s got %s for %s", targetMode.String(), m.String(), filename)
		}
	}
}

func TestSimpleDetectFromBytes(t *testing.T) {
	for s, targetMode := range examples {
		if m := SimpleDetectBytes([]byte(s)); m != targetMode {
			t.Fatalf("Expected %s got %s", targetMode.String(), m.String())
		}
	}
	for filename, targetMode := range exampleFiles {
		data, err := os.ReadFile(filename)
		if err != nil {
			t.Fatalf("Could not read %s: %v\n", filename, err)
		}
		if m := SimpleDetectBytes(data); m != targetMode {
			t.Fatalf("Expected %s got %s for %s", targetMode.String(), m.String(), filename)
		}
	}
}
