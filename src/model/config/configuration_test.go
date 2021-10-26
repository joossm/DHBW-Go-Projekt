package config

import (
	"testing"
)

func TestInit(t *testing.T) {
	Init()
	if port1 != 8444 {
		t.Error("port 1 is ", port1, "but it should be :8443")
	}
	if port2 != 8443 {
		t.Error("port 2 is ", port2, "but it should be :8444")
	}
	if tokenDuration != 60 {
		t.Error("token duration is ", tokenDuration, "but it should be 60")
	}
	if xmlPath == "" {
		t.Error("path is empty")
	}

}
