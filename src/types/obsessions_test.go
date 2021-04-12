package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var obsessions = Obsessions{
	List: struct {
		Easy   []string `json:"easy"`
		Medium []string `json:"medium"`
		Hard   []string `json:"hard"`
	}(struct {
		Easy   []string
		Medium []string
		Hard   []string
	}{
		Easy: []string{
			"Buying groceries",
			"Changing clothes",
			"Collecting autographs",
			"Doing drugs",
			"Eating candy",
			"Exiting rooms unconventionally",
			"Giving out presents",
			"Defacing government owned buildings",
			"Overturning shelves",
			"Petting dogs",
			"Stealing hats off of people's heads",
			"Stealing ripe fruits from groceries and markets",
			"Yelling at children",
		},
		Medium: []string{
			"Land a job",
			"Throwing bags of candy on heads",
			"Whipping objects out of people's hands",
			"participating in flash mobs",
			"Siphoning gas out of cars",
			"Getting on live television",
			"Defeating people in trading card games",
			"Dueling other people named John",
			"Tackling large men",
			"Slow stripping on public transit",
			"Acquiring used bathrobes",
			"Teaching classes",
			"Winning dance offs",
			"Starting food fights",
			"Organizing protests",
		},
		Hard: []string{
			"Converting people to my religion",
			"Convincing people my surname is Candy",
			"Getting a doctorate",
			"Convincing someone to buy something fresh from the trash",
			"Having people solve my riddles",
			"Being referred to as 'me boy'",
			"Joining cults",
			"Liberating animals from cages",
			"Stealing artifacts",
			"Taking other people's pets",
			"Winning awards",
		},
	}),
}

func TestGetObsessions(t *testing.T) {
	obs := obsessions.GetObsessions()
	assert.NotEmpty(t, obs.Easy)
	assert.NotEmpty(t, obs.Medium)
	assert.NotEmpty(t, obs.Hard)
}
