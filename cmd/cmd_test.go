package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidDate(t *testing.T) {
	validDateFormat("12.12.1221")
	assert.Equal(t, validDateFormat("20.20.2020"), true)
	assert.Equal(t, validDateFormat("2.10.1991"), false)
	assert.Equal(t, validDateFormat("12.1.1991"), false)
	assert.Equal(t, validDateFormat("12.10.19911"), false)
	assert.Equal(t, validDateFormat("2.10.1991.12"), false)
	assert.Equal(t, validDateFormat("1991.10.02"), false)
	assert.Equal(t, validDateFormat("02.1991.10"), false)
	assert.Equal(t, validDateFormat("aa.aaaa.%%"), false)
	assert.Equal(t, validDateFormat("23.A2.2013"), false)

}
