package model

import (
	"GoProjekt/src/model/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterLocations(t *testing.T) {
	//Test nicht möglich da Path falsch
	config.Init()
	loc := RegisterLocations()
	assert.Equal(t, loc.getLength()!=0, true)
		for i := 0; i < len(loc.Locations); i++ {
			assert.Equal(t, loc.Locations[i].Name != "", true)
		}
	}

