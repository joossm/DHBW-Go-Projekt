package main

import (
	"testing"
)

func TestRead(t *testing.T) {
	a := read()
	if a.getLength() <= 0 {
		t.Error("The locations List is empty")
	} else {
		for i := 0; i < len(a.Locations); i++ {
			if a.Locations[i].Name == "" {
				t.Error("Location name may not be empty")
			}
		}
	}

}
