package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidDate(t *testing.T) {

	assert.Equal(t, validDateFormat("2020-20-20"), true)
	assert.Equal(t, validDateFormat("1991-10-1"), false)
	assert.Equal(t, validDateFormat("1991-1-19"), false)
	assert.Equal(t, validDateFormat("12-10-19911"), false)
	assert.Equal(t, validDateFormat("02-10-1991-12"), false)
	assert.Equal(t, validDateFormat("02-10-1991"), false)
	assert.Equal(t, validDateFormat("02-1991-10"), false)
	assert.Equal(t, validDateFormat("aaaa-aa-%%"), false)
	assert.Equal(t, validDateFormat("2023-A2-20"), false)

}
func TestLauft(t *testing.T) {
	Lauft()
}
