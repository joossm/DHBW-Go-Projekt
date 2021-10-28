package config

import (
	"flag"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Init()
	if port1 != "8443" {
		t.Error("port 1 is ", port1, "but it should be :8443")
	}
	if port2 != "8444" {
		t.Error("port 2 is ", port2, "but it should be :8444")
	}
	if tokenDuration != time.Minute*5 {
		t.Error("token duration is ", tokenDuration, "but it should be 60")
	}
	if xmlPath == "" {
		t.Error("path is empty")
	}
	assert.Equal(t, flag.Parsed(), true)

	assert.False(t, flag.Lookup("port1") == nil)
	assert.False(t, flag.Lookup("port2") == nil)
	assert.False(t, flag.Lookup("tokenDuration") == nil)
	assert.False(t, flag.Lookup("xmlPath") == nil)

	assert.Equal(t, port1, flag.Lookup("port1").Value.String(), true)
	assert.Equal(t, port2, flag.Lookup("port2").Value.String(), true)
	assert.Equal(t, tokenDuration, flag.Lookup("tokenDuration").Value.String(), true)
	assert.Equal(t, xmlPath, flag.Lookup("xmlPath").Value.String(), true)

}
