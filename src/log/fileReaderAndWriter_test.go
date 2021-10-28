package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDateInFormat(t *testing.T) {
	assert.Equal(t, "2021-10-28", getDateInFormat())
}
func TestWriteToFile(t *testing.T) {

}

/*func TestGetTimeStamp(t *testing.T) {
	assert.Equal(t, "2021-10-27", getDateInFormat())
}*/
func TestCreateText(t *testing.T) {
	assert.Equal(t, "LOGIN, "+getTimeStamp()+", "+"name, address, location"+";\n", createText(true, "name, address, location"))
	assert.Equal(t, "LOGOUT, "+getTimeStamp()+", "+"name, address, location"+";\n", createText(false, "name, address, location"))
}
func TestPanicHandling(t *testing.T) {

}
func TestDoesFileExists(t *testing.T) {
	assert.Equal(t, true, doesFileExists("../../server.go"))
	assert.Equal(t, false, doesFileExists("../../party.go"))
}
func TestReadAndWriteAndSave(t *testing.T) {

}
func TestCreateFile(t *testing.T) {

}
