package types

import (
	"encoding/json"
	"math/rand"
	"time"
)

//Skills contains a string array of all possible skills
type Skills struct {
	Names []string `json:"skills"`
}

//GetSkills returns a random selection of three skills in the form of a CharacterSkills struct
func (s *Skills) GetSkills() CharacterSkills {
	var skills CharacterSkills
	rand.Seed(time.Now().Unix())
	key := rand.Intn(len(s.Names))
	skills.First = s.Names[key]
	key2 := rand.Intn(len(s.Names))
	skills.Second = s.Names[key2]
	key3 := rand.Intn(len(s.Names))
	skills.Third = s.Names[key3]
	return skills
}

//LoadSkills returns embedded json of skills into a struct
func LoadSkills(b []byte) Skills {
	s := new(Skills)
	_ = json.Unmarshal(b, &s)
	return *s
}
