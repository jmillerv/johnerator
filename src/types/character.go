package types

const (
	willpower = 10
)

//Character is the struct that we populate for to pass into the character template
type Character struct {
	Name        string              `json:"name"`
	Skills      CharacterSkills     `json:"skills"`
	Obsessions  CharacterObsessions `json:"obsessions"`
	Willpower   int                 `json:"willpower"`
	ThreeSkills bool                `json:"three_skills"`
}

//CharacterObsessions hold three strings for a voice indicating whether they are easy/medium/or hard
type CharacterObsessions struct {
	Easy   string
	Medium string
	Hard   string
}

//CharacterSkills holds a given character's skills. First and Second are required Third is optional
type CharacterSkills struct {
	First  string
	Second string
	Third  string
}

//CreateNewCharacter generates a new voice and returns it in the form of a Character
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
