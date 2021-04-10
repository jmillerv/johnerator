package types

import (
	_ "embed"
	"encoding/json"
	"math/rand"
	"time"
)

type Obsessions struct {
	List struct {
		Easy   []string `json:"easy"`
		Medium []string `json:"medium"`
		Hard   []string `json:"hard"`
	} `json:"obsessions"`
}

func (o *Obsessions) GetObsessions() CharacterObsessions {
	rand.Seed(time.Now().Unix())
	var obs CharacterObsessions
	key := rand.Intn(len(o.List.Easy))
	obs.Easy = o.List.Easy[key]
	key2 := rand.Intn(len(o.List.Medium))
	obs.Medium = o.List.Medium[key2]
	key3 := rand.Intn(len(o.List.Hard))
	obs.Hard = o.List.Hard[key3]
	return obs
}

func LoadObsessions(b []byte) Obsessions {
	o := new(Obsessions)
	_ = json.Unmarshal(b, &o)
	return *o
}