package model

import (
	"testing"
)

func TestRead(t *testing.T) {
	main()
	loc := Read()
	if loc.getLength() <= 0 {
		t.Error("The locations List is empty")
	} else {
		for i := 0; i < len(loc.Locations); i++ {
			if loc.Locations[i].Name == "" {
				t.Error("Location name may not be empty")
			}
		}
	}

}
