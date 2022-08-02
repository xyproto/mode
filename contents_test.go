package mode

import (
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
}

func TestSimpleDetect(t *testing.T) {
	for s, targetMode := range examples {
		if m := SimpleDetect(s); m != targetMode {
			t.Fatalf("Expected %s got %s", targetMode.String(), m.String())
		}
	}
}

func TestSimpleDetectFromBytes(t *testing.T) {
	for s, targetMode := range examples {
		if m := SimpleDetectBytes([]byte(s)); m != targetMode {
			t.Fatalf("Expected %s got %s", targetMode.String(), m.String())
		}
	}
}
