/*
 * JumpCloud Project API
 *
 * This is a simple API for JumpCloud Project
 *
 * API version: 1.0.0
 * Contact: birk.pete@gmail.com
 */
 package main

import (
	"strings"
	"path"
)

/**
   This is just to show a placeholder for some
   validation rules.  Ideally, we would have a helper
   class that we always use for these kinds of things.
**/
func ValidateString (input string) (bool) {
	if (strings.Contains(input, "<script>")) {
		return false;
	}
	// won't find a string this large, etc.
	if (len(input) > 1000) {
		return false;
	}

	return true;
}

func ValidatePositiveInteger (input int) (bool) {
	if (input < 0) {
		return false;
	}
	// some large # too big for use case
	if (input > 10000000) {
		return false;
	}

	return true;
}

func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	var i int;
	if (strings.Contains(p[1:], "api/v1")) {
		i = strings.Index(p[1:], "api/v1") + 7
	 } else {
		i = strings.Index(p[1:], "/") + 1
	 }

	if i <= 0 {
        return p[1:], "/"
    }
    return p[1:i], p[i:]
}
