package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)
var loc1 = Location{Name: "Warburg"}
var loc2 = Location{Name: "Berlin"}
var list = createTestLocations()

func createTestLocations() LocationsList{
	list := LocationsList{}
	list.Locations = append(list.Locations, &loc1)
	list.Locations = append(list.Locations, &loc2)
	return list
}

func TestLocationsList_equals(t *testing.T) {
	var loc3 = Location{Name: "Warburg"}
	var loc4 = Location{Name: "Berlin"}
	list2 := LocationsList{}
	list2.Locations = append(list2.Locations, &loc3)
	list2.Locations = append(list2.Locations, &loc4)
	assert.True(t, list.equals(list2))
	list2.Locations = append(list2.Locations, &Location{Name:"Darmstadt"})
	assert.False(t, list.equals(list2))
	list3 := LocationsList{}
	list3.Locations = append(list3.Locations, &Location{Name:"Volkmarsen"})
	list3.Locations = append(list3.Locations,&Location{Name : "München"})
	assert.False(t, list.equals(list3))
}



func TestLocationsList_GetLength(t *testing.T) {
	assert.Equal(t,list.getLength()==2 ,true)
}

func TestLocationsList_ToStrings(t *testing.T) {
	assert.Equal(t, list.ToStrings()[0]==list.Locations[0].Name,true)
	assert.Equal(t, list.ToStrings()[1]==list.Locations[1].Name,true)
}

func TestShowAllLocations(t *testing.T) {
	assert.Equal(t, list.ShowAllLoc()[0].Name == loc1.Name,true)
	assert.Equal(t, list.ShowAllLoc()[1].Name == loc2.Name,true)
}

func TestLocationsList_ShowAllLoc(t *testing.T) {
	assert.True(t, list.Locations[0]==list.ShowAllLoc()[0])
	assert.True(t, list.Locations[1]==list.ShowAllLoc()[1])
	assert.False(t, list.Locations[1]==list.ShowAllLoc()[0])
}

func TestRead(t *testing.T) {
	loc := ReadXmlFile("../../assets/locations.xml")
	assert.True(t, loc.getLength()>0)
	assert.True(t, loc.Locations[3].Name == "Warburg")
	for i := 0; i < len(loc.Locations); i++ {
		assert.True(t, loc.Locations[i].Name != "")
	}
}
func TestGetList(t *testing.T) {
	ReadXmlFile("../../assets/locations.xml")
	assert.Equal(t, GetList().Locations[0].Name == "Mosbach",true)
	assert.Equal(t, GetList().Locations[1].Name == "Dresden",true)
}



