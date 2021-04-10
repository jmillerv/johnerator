package types

import (
	_ "embed"
	"encoding/json"
	"math/rand"
	"time"
)




type Names struct {
	List []string `json:"names"`
}

func (n *Names) GetName() string {
	rand.Seed(time.Now().Unix())
	key := rand.Intn(len(n.List))
	return n.List[key]
}


func LoadNames(b []byte) Names {
	n := new(Names)
	_ = json.Unmarshal(b, &n)
	return *n
}