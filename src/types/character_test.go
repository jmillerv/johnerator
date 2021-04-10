package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateNewCharacter(t *testing.T) {
	c := CreateNewCharacter(names, skills, obsessions, 2)
	fmt.Println(c)
	assert.NotEmpty(t, c.Name)
	assert.NotEmpty(t, c.Skills.First)
	assert.NotEmpty(t, c.Skills.Second)
	assert.NotEmpty(t, c.Obsessions.Easy)
	assert.NotEmpty(t, c.Obsessions.Medium)
	assert.NotEmpty(t, c.Obsessions.Hard)
}
