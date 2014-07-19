package gen

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func Test(t *testing.T) { TestingT(t) }

func (suite *MySuite) TestParser(c *C) {

	type _test struct {
		token        string
		expected_ret int
		expected_res *SchedParseRes
	}

	var emptystrs []string = []string{}
	var emptyclockrange RecurByTime2 = RecurByTime2{}

	tests := []_test{
		{"every 21 hrs", 0, &SchedParseRes{C_RecurByTime, RecurByTime{21, "hours", emptyclockrange}}},
		{"every 21 hrs from 06:30 to 08:30", 0, &SchedParseRes{C_RecurByTime, RecurByTime{21, "hours", RecurByTime2{From: "06:30", To: "08:30"}}}},
		{"every monday 05:30", 0, &SchedParseRes{C_RecurByDay, RecurByDay{-1, []string{"monday"}, emptystrs, "05:30"}}},
		{"every tuesday 25:30", 1, nil},
		{"every friday 24:70", 1, nil},
		{"every mon 05:30", 0, &SchedParseRes{C_RecurByDay, RecurByDay{-1, []string{"monday"}, emptystrs, "05:30"}}},
		{"every wed 15:40", 0, &SchedParseRes{C_RecurByDay, RecurByDay{-1, []string{"wednesday"}, emptystrs, "15:40"}}},
		{"every sat 05:10", 0, &SchedParseRes{C_RecurByDay, RecurByDay{-1, []string{"saturday"}, emptystrs, "05:10"}}},
		{"every tue of january 17:23", 0, &SchedParseRes{C_RecurByDay, RecurByDay{-1, []string{"tuesday"}, []string{"january"}, "17:23"}}},
		{"1st thu of sept 17:23", 0, &SchedParseRes{C_RecurByDay, RecurByDay{1, []string{"thursday"}, []string{"september"}, "17:23"}}},
		{"0th thu of sept 17:23", 1, nil},
		{"11st thu of january 17:23", 1, nil},
		{"211st thu of january 17:23", 1, nil},
		{"321st sat of dec 17:23", 0, &SchedParseRes{C_RecurByDay, RecurByDay{321, []string{"saturday"}, []string{"december"}, "17:23"}}},
		{"321rd sat of dec 17:23", 1, nil},
		{"343rd sun of oct 17:23", 0, &SchedParseRes{C_RecurByDay, RecurByDay{343, []string{"sunday"}, []string{"october"}, "17:23"}}},
		{"313rd sun of oct 17:23", 1, nil},
		{"313th sun of feb 17:23", 0, &SchedParseRes{C_RecurByDay, RecurByDay{313, []string{"sunday"}, []string{"february"}, "17:23"}}},
		{"312th sun of feb 17:23", 0, &SchedParseRes{C_RecurByDay, RecurByDay{312, []string{"sunday"}, []string{"february"}, "17:23"}}},
		{"232nd sun 17:23", 0, &SchedParseRes{C_RecurByDay, RecurByDay{232, []string{"sunday"}, emptystrs, "17:23"}}},
		{"232nd mon,sun 17:23", 0, &SchedParseRes{C_RecurByDay, RecurByDay{232, []string{"monday", "sunday"}, emptystrs, "17:23"}}},
		{"312th sun of feb,jan,mar 17:23", 0, &SchedParseRes{C_RecurByDay, RecurByDay{312, []string{"sunday"}, []string{"february", "january", "march"}, "17:23"}}},
	}

	for _, v := range tests {
		fmt.Println("testing token=", v.token)
		var obtained_res *SchedParseRes = nil

		ret := SchedParse(&SchedLex{S: v.token, OnSchedParse: func(res *SchedParseRes) {
			obtained_res = res
		}})
		c.Assert(ret, Equals, v.expected_ret)
		c.Assert(obtained_res, DeepEquals, v.expected_res)
	}

}
