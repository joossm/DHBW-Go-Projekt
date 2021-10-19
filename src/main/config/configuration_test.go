package config

import (
	"testing"
)

func TestInit(t *testing.T){
	if XmlFilePath ==nil {
		t.Error("FilePath is nil")
	}
	if *XmlFilePath =="" {
		t.Error("FilePath is null")
	}
	if port1 !=8444 {
		t.Error("port 1 is ", port1, "but it should be 8444")
	}
	if port2 !=8443{
		t.Error("port 2 is " , port2, "but it should be 8443")
	}
	if tokenDuration !=60 {
		t.Error("token duration is ", tokenDuration, "but it should be 60")
	}

}

