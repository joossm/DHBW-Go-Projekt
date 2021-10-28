package model



import (
	"GoProjekt/src/model/config"
	"github.com/stretchr/testify/assert"
	"testing"
)
var loc1 = Location{Name: "Warburg"}
var loc2 = Location{Name: "Berlin"}
var list = createTestLocations()

func createTestLocations() LocationsList{
	list := LocationsList{}
	list.Locations = append(list.Locations, loc1)
	list.Locations = append(list.Locations, loc2)
	return list
}

func TestRegisterLocations(t *testing.T) {
	config.Init()
	loc := RegisterLocations("../../assets/locations.xml")
	assert.Equal(t, loc.getLength()>0, true)
		for i := 0; i < len(loc.Locations); i++ {
			assert.Equal(t, loc.Locations[i].Name != "", true)
		}
	}
func TestLocationsList_GetLength(t *testing.T) {
	assert.Equal(t,list.getLength()==2 ,true)
}
func TestLocationsList_ToStrings(t *testing.T) {
	assert.Equal(t, list.ToStrings()[0]==list.Locations[0].Name,true)
	assert.Equal(t, list.ToStrings()[1]==list.Locations[1].Name,true)
}


func TestListContains(t *testing.T){

}



