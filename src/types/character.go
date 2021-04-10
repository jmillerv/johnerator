package types

import (
	_ "embed"
)



type Character struct {
	Name       string              `json:"name"`
	Skills     CharacterSkills           `json:"skills"`
	Obsessions CharacterObsessions `json:"obsessions"`
	ThreeSkills bool
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
	if skillCount != 0 {
		c.ThreeSkills = true
	}
	return c
}

