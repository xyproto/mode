package mode

import (
	"strings"
	"testing"
)

func TestDetectFromContents(t *testing.T) {
	examples := map[string]Mode{
		"#!/bin/bash\necho Hello\n":              Shell,
		"#!/usr/bin/env python\nprint(\"hi\")\n": Python,
		"":                                       Blank,
		"# Hello":                                Config,
		"#include <stdio.h>":                     Blank, // hard to tell if this is C or C++, should be detected by filename, not by contents
	}
	for s, targetMode := range examples {
		if m, _ := DetectFromContents(strings.Split(s, "\n")[0], func() string { return s }); m != targetMode {
			t.Fatalf("Expected %s got %s", targetMode.String(), m.String())
		}
	}
}
