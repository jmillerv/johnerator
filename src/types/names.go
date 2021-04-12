package types

import (
	"encoding/json"
	"math/rand"
	"time"
)
//Names contains a string array of all possible names
type Names struct {
	List []string `json:"names"`
}

//GetName returns a random name from the Names struct
func (n *Names) GetName() string {
	rand.Seed(time.Now().Unix())
	key := rand.Intn(len(n.List))
	return n.List[key]
}

//LoadNames returns embedded json of names into a struct
func LoadNames(b []byte) Names {
	n := new(Names)
	_ = json.Unmarshal(b, &n)
	return *n
}