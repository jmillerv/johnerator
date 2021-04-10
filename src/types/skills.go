package types

import (
	_ "embed"
	"encoding/json"
	"math/rand"
	"time"
)


type Skills struct {
	Names []string `json:"skills"`
}



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

func LoadSkills(b []byte) Skills {
	s := new(Skills)
	_ = json.Unmarshal(b, &s)
	return *s
}
