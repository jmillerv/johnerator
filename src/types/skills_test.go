package types

import (
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
	s := skills.GetSkills()
	assert.NotEmpty(t, s.First)
	assert.NotEmpty(t, s.Second)
	assert.NotEmpty(t, s.Third)
}
