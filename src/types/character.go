package types

import (
	_ "embed"
)

const (
	willpower = 10
)

type Character struct {
	Name       string              `json:"name"`
	Skills     CharacterSkills           `json:"skills"`
	Obsessions CharacterObsessions `json:"obsessions"`
	Willpower int `json:"willpower"`
	ThreeSkills bool `json:"three_skills"`
}

type CharacterObsessions struct {
	Easy   string
	Medium string
	Hard   string
}

type CharacterSkills struct {
	First string
	Second string
	Third string
}


func CreateNewCharacter(n Names, s Skills, o Obsessions, skillCount int) Character {
	var c Character
	c.Name = n.GetName()
	c.Skills = s.GetSkills()
	c.Obsessions = o.GetObsessions()
	c.Willpower = willpower
	if skillCount != 0 {
		c.ThreeSkills = true
		c.Willpower = c.Willpower - 3
	}
	return c
}

