package schedparser

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

type MySuite struct{}

var _ = Suite(&MySuite{})

func Test(t *testing.T) { TestingT(t) }

func (suite *MySuite) TestToByDay(c *C) {

	type _test struct {
		token        string
		expected_ret int
		expected_res interface{}
	}

	var nilday *SchedRecurByDay
	var emptymap map[string]struct{} = map[string]struct{}{}

	tests := []_test{
		{"every monday 05:30", 0, &SchedRecurByDay{Ord: -1, Days: map[string]struct{}{"monday": _emptystruct}, Months: emptymap, Clock: "05:30"}},
		{"every tuesday 25:30", 1, nilday},
		{"every friday 24:70", 1, nilday},
		{"every mon 05:30", 0, &SchedRecurByDay{Ord: -1, Days: map[string]struct{}{"monday": _emptystruct}, Months: emptymap, Clock: "05:30"}},
		{"every wed 15:40", 0, &SchedRecurByDay{Ord: -1, Days: map[string]struct{}{"wednesday": _emptystruct}, Months: emptymap, Clock: "15:40"}},
		{"every sat 05:10", 0, &SchedRecurByDay{Ord: -1, Days: map[string]struct{}{"saturday": _emptystruct}, Months: emptymap, Clock: "05:10"}},
		{"every tue of january 17:23", 0, &SchedRecurByDay{Ord: -1, Days: map[string]struct{}{"tuesday": _emptystruct}, Months: map[string]struct{}{"january": _emptystruct}, Clock: "17:23"}},
		{"1st thu of sept 17:23", 0, &SchedRecurByDay{Ord: 1, Days: map[string]struct{}{"thursday": _emptystruct}, Months: map[string]struct{}{"september": _emptystruct}, Clock: "17:23"}},
		{"0th thu of sept 17:23", 1, nilday},
		{"11st thu of january 17:23", 1, nilday},
		{"211st thu of january 17:23", 1, nilday},
		{"321st sat of dec 17:23", 0, &SchedRecurByDay{Ord: 321, Days: map[string]struct{}{"saturday": _emptystruct}, Months: map[string]struct{}{"december": _emptystruct}, Clock: "17:23"}},
		{"321rd sat of dec 17:23", 1, nilday},
		{"343rd sun of oct 17:23", 0, &SchedRecurByDay{Ord: 343, Days: map[string]struct{}{"sunday": _emptystruct}, Months: map[string]struct{}{"october": _emptystruct}, Clock: "17:23"}},
		{"313rd sun of oct 17:23", 1, nilday},
		{"313th sun of feb 17:23", 0, &SchedRecurByDay{Ord: 313, Days: map[string]struct{}{"sunday": _emptystruct}, Months: map[string]struct{}{"february": _emptystruct}, Clock: "17:23"}},
		{"312th sun of feb 17:23", 0, &SchedRecurByDay{Ord: 312, Days: map[string]struct{}{"sunday": _emptystruct}, Months: map[string]struct{}{"february": _emptystruct}, Clock: "17:23"}},
		{"232nd sun 17:23", 0, &SchedRecurByDay{Ord: 232, Days: map[string]struct{}{"sunday": _emptystruct}, Months: emptymap, Clock: "17:23"}},
		{"232nd mon,sun 17:23", 0, &SchedRecurByDay{Ord: 232, Days: map[string]struct{}{"monday": _emptystruct, "sunday": _emptystruct}, Months: emptymap, Clock: "17:23"}},
		{"312th sun of feb,jan,mar 17:23", 0, &SchedRecurByDay{Ord: 312, Days: map[string]struct{}{"sunday": _emptystruct}, Months: map[string]struct{}{"february": _emptystruct, "january": _emptystruct, "march": _emptystruct}, Clock: "17:23"}},
	}

	for _, v := range tests {
		fmt.Println("testing token=", v.token)
		var byday *SchedRecurByDay
		var bytime *SchedRecurByTime
		ret := ParseSched(v.token, func(_time *SchedRecurByTime, _day *SchedRecurByDay) {
			bytime, byday = _time, _day
		})

		fmt.Println("bytime=", bytime, ";byday=", byday)

		c.Assert(bytime, IsNil)
		c.Assert(ret, Equals, v.expected_ret)
		c.Assert(byday, DeepEquals, v.expected_res)
	}

}

func (suite *MySuite) TestToByTime(c *C) {

	type _test struct {
		token        string
		expected_ret int
		expected_res interface{}
	}

	var niltime *SchedRecurByTime

	tests := []_test{
		{"every 21 hrs", 0, &SchedRecurByTime{N: 21, ClockLbl: "hours", FromClock: "", ToClock: ""}},
		{"every 3 secs from 06:30 to 08:30", 0, &SchedRecurByTime{N: 3, ClockLbl: "seconds", FromClock: "06:30", ToClock: "08:30"}},
		{"every 05 mins from 06:30 to 09:30", 0, &SchedRecurByTime{N: 5, ClockLbl: "minutes", FromClock: "06:30", ToClock: "09:30"}},
		{"every 05 mins from 06:30 to ", 1, niltime},
		{"every 05 mins from 06:30 to 25:30", 1, niltime},
	}

	for _, v := range tests {
		fmt.Println("testing token=", v.token)
		var byday *SchedRecurByDay
		var bytime *SchedRecurByTime
		ret := ParseSched(v.token, func(_time *SchedRecurByTime, _day *SchedRecurByDay) {
			bytime, byday = _time, _day
		})
		c.Assert(byday, IsNil)
		c.Assert(ret, Equals, v.expected_ret)
		c.Assert(bytime, DeepEquals, v.expected_res)
	}

}

func (suite *MySuite) TestMultiline(c *C) {

	token := `
every 21 hrs
321st sat of dec 17:23
every 05 mins from 06:30 to 09:30`

	var count int
	byday := []SchedRecurByDay{}
	bytime := []SchedRecurByTime{}
	ParseSched(token, func(_time *SchedRecurByTime, _day *SchedRecurByDay) {
		count++
		if nil != _time {
			bytime = append(bytime, *_time)
		}
		if nil != _day {
			byday = append(byday, *_day)
		}
	})
	c.Assert(count, Equals, 3)
	c.Assert(len(bytime), Equals, 2)
	c.Assert(len(byday), Equals, 1)

}
