package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDateInFormat(t *testing.T) {
	assert.Equal(t, "2021-10-27", getDateInFormat())
}
func TestWriteToFile(t *testing.T) {

}
func TestGetTimeStamp(t *testing.T) {
	assert.Equal(t, "2021-10-27", getDateInFormat())
}
func TestCreateText(t *testing.T) {
	assert.Equal(t, "Oct 27 21:28:10", getTimeStamp())
}
func TestPanicHandling(t *testing.T) {

}
func TestDoesFileExists(t *testing.T) {

}
func TestReadAndWriteAndSave(t *testing.T) {

}
func TestCreateFile(t *testing.T) {

}
