package types

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var skills = Skills{
	Names: []string{
		"Biting",
		"Bow Hunting",
		"Cowboy Whip",
		"Crafting",
		"Dance",
		"Disguise",
		"Distracting People",
		"Driving",
		"Electrical",
		"Escape Artist",
		"Finding Lost Items",
		"Guns",
		"Gymnastics",
		"Hiding",
		"Intimidation",
		"Law",
		"Lying",
		"Magic Tricks",
		"Martial Arts",
		"Organizing",
		"Parkour",
		"Persuade",
		"Plumbing",
		"Running",
		"Seduction",
		"Theft",
		"Vomit On Command",
	},
}

func TestGetSkills(t *testing.T) {
	s := skills.GetSkills(0)
	fmt.Println(s)
	assert.Equal(t, len(s), 2)
	s2 := skills.GetSkills(3)
	fmt.Println(s2)
	assert.Equal(t, len(s2), 3)
}
