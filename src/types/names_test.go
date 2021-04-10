package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var names = Names{
	List: []string{
		"Indiana Johns",
		"Janeane",
		"Jean",
		"Joan",
		"Joanne",
		"John",
		"John Candy",
		"Johnny",
		"Jon",
		"Jonathan",
		"Juan",
		"June",
	},
}

func TestNames_GetName(t *testing.T) {
	name := names.GetName()
	fmt.Println(name)
	assert.NotEmpty(t, name)
}
