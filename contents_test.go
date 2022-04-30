package mode

import (
	"testing"
)

func TestSimpleDetect(t *testing.T) {
	examples := map[string]Mode{
		"#!/bin/bash\necho Hello\n":              Shell,
		"#!/usr/bin/env python\nprint(\"hi\")\n": Python,
		"":                                       Blank,
		"# Hello":                                Config,
		"#include <stdio.h>":                     Blank, // hard to tell if this is C or C++, should be detected by filename, not by contents
		"<?xml version=\"1.0\" encoding=":        XML,
	}
	for s, targetMode := range examples {
		if m := SimpleDetect(s); m != targetMode {
			t.Fatalf("Expected %s got %s", targetMode.String(), m.String())
		}
	}
}
