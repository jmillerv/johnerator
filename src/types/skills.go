package types

import (
	"encoding/json"
	"math/rand"
	"time"
)

//Skills contains a string array of all possible skills
type Skills struct {
	CallOfCthulhu      []string `json:"call_of_cthulhu"`
	DeltaGreen         []string `json:"delta_green"`
	DungeonsAndDragons []string `json:"dungeons_and_dragons"`
	Fantasy            []string `json:"fantasy"`
	Future             []string `json:"future"`
	Heroics            []string `json:"heroics"`
	Horror             []string `json:"horror"`
	Magic              []string `json:"magic"`
	Pathfinder         []string `json:"pathfinder"`
	Standard           []string `json:"standard"`
	Superpowers        []string `json:"superpowers"`
}

//GetSkills returns a random selection of three skills in the form of a CharacterSkills struct
func (s *Skills) GetSkills() CharacterSkills {
	var skills CharacterSkills
	rand.Seed(time.Now().Unix())
	key := rand.Intn(len(s.Standard))
	skills.First = s.Standard[key]
	key2 := rand.Intn(len(s.Standard))
	skills.Second = s.Standard[key2]
	key3 := rand.Intn(len(s.Standard))
	skills.Third = s.Standard[key3]
	return skills
}

//LoadSkills returns embedded json of skills into a struct
func LoadSkills(b []byte) Skills {
	s := struct {
		Skills `json:"skills"`
	}{}
	_ = json.Unmarshal(b, &s)
	skills := &Skills{
		Fantasy:     s.Fantasy,
		Future:      s.Future,
		Heroics:     s.Heroics,
		Horror:      s.Horror,
		Magic:       s.Magic,
		Standard:    s.Standard,
		Superpowers: s.Superpowers,
	}
	return *skills
}
