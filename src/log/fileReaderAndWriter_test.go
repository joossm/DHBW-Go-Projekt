package log

import (
	"GoProjekt/src/model/config"
	"flag"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateText(t *testing.T) {
	assert.Equal(t, "LOGIN, "+getTimeStamp()+", "+"name, address, location"+";\n", createText(true, "name, address, location"))
	assert.Equal(t, "LOGOUT, "+getTimeStamp()+", "+"name, address, location"+";\n", createText(false, "name, address, location"))
}

func TestDoesFileExists(t *testing.T) {
	assert.Equal(t, true, doesFileExists("../../server.go"))
	assert.Equal(t, false, doesFileExists("../../party.go"))
}
func TestReadAndWriteAndSave(t *testing.T) {
	config.InitByMatthias()
	createFile("fileTest")
	readAndWriteAndSave(true, "test", "fileTest")
	var check, _ = ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + "fileTest" + ".txt")
	assert.Equal(t, []byte("LOGIN, Oct 28 18:00:28, test;\n")[0:16], check[0:16])
	_ = os.Remove("../../src/log/files/fileTest.txt")
	assert.NoFileExists(t, "../../src/log/files/fileTest.txt")
}
func TestCreateFile(t *testing.T) {
	config.InitByMatthias()
	createFile("fileTest")
	assert.FileExists(t, "../../src/log/files/fileTest.txt")
	_ = os.Remove("../../src/log/files/fileTest.txt")
	assert.NoFileExists(t, "../../src/log/files/fileTest.txt")
}
