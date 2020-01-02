package moments

import (
	"testing"
)

type M [][]int

func TestFindCircleNum(t *testing.T) {
	testCases := []struct {
		name     string
		M        M
		wantNums int
	}{{
		"general case 1",
		M{
			{1, 1, 0},
			{1, 1, 0},
			{0, 0, 1},
		},
		2,
	}, {
		"general case 2",
		M{
			{1, 1, 0},
			{1, 1, 1},
			{0, 1, 1},
		},
		1,
	},
	}

	for _, tc := range testCases {
		num := FindCircleNum(tc.M)
		if num != tc.wantNums {
			t.Logf("start case %s",tc.name)
			t.Errorf("Got %d num, expected %d", num, tc.wantNums)
		}
	}
}
