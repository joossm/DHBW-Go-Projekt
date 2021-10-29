package log

import (
	"GoProjekt/src/model/config"
	"flag"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestGetTimeInFormat(t *testing.T) {
	assert.Equal(t, time.Now().Format(time.RFC3339)[0:10], getDateInFormat())
}
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
	assert.Equal(t, []byte("LOGIN, Oct 28 18:00:28, test;\n")[0:4], check[0:4])
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

func TestWriteToFile(t *testing.T) {
	config.InitByMatthias()
	var before, _ = ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + getDateInFormat() + ".txt")
	beforeLine := strings.Split(string(before), "\n")
	WriteToFile(true, "Test")
	var check, _ = ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + getDateInFormat() + ".txt")
	assert.Equal(t, []byte("LOGIN, Oct 28 18:00:28, test;\n")[0:4], check[0:4])
	assert.FileExists(t, "../../src/log/files/"+getDateInFormat()+".txt")
	lines := strings.Split(string(check), "\n")
	var data []byte

	for i := 0; i < len(beforeLine)-2; i++ {
		data = append(data, []byte(lines[i]+"\n")...)
	}

	err := ioutil.WriteFile(flag.Lookup("logfilePath").Value.String()+getDateInFormat()+".txt", []byte(data), 0644)
	panicHandling(err)
	_ = os.Remove(flag.Lookup("logfilePath").Value.String() + getDateInFormat() + ".txt")
	assert.NoFileExists(t, flag.Lookup("logfilePath").Value.String()+getDateInFormat()+".txt")

	WriteToFile(true, "Test")
	var checkN, _ = ioutil.ReadFile(flag.Lookup("logfilePath").Value.String() + getDateInFormat() + ".txt")
	assert.Equal(t, []byte("LOGIN, Oct 28 18:00:28, test;\n")[0:4], checkN[0:4])

	_ = os.Remove(flag.Lookup("logfilePath").Value.String() + getDateInFormat() + ".txt")
	assert.NoFileExists(t, flag.Lookup("logfilePath").Value.String()+getDateInFormat()+".txt")

	err = ioutil.WriteFile(flag.Lookup("logfilePath").Value.String()+getDateInFormat()+".txt", []byte(data), 0644)
	assert.FileExists(t, flag.Lookup("logfilePath").Value.String()+getDateInFormat()+".txt")
}
