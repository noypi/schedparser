package schedparser

import (
	"github.com/noypi/schedparser/gen"
)

var _emptystruct struct{} = struct{}{}

//{"every 21 hrs from 06:30 to 08:30", 0, &SchedParseRes{C_RecurByTime, []RecurByTime{"every", 21, "hrs", []RecurByTime2{"from", "06:30", "to", "08:30"}}}},
func ToByTime(res *gen.SchedParseRes) (byTime *SchedRecurByTime) {
	if (gen.C_RecurByTime != res.Type) || (nil == res.Item) {
		return
	}

	var rbt gen.RecurByTime
	var ok bool
	rbt, ok = res.Item.(gen.RecurByTime)
	if !ok {
		return
	}

	e := new(SchedRecurByTime)
	e.N = rbt.N
	e.ClockLbl = rbt.ClockLbl

	byTime = e

	// optional
	if 0 < len(rbt.ClockRange.From) && 0 < len(rbt.ClockRange.To) {
		e.FromClock = rbt.ClockRange.From
		e.ToClock = rbt.ClockRange.To
	}
	return
}

//{"312th sun of feb,jan,mar 17:23", 0, &SchedParseRes{C_RecurByDay, []RecurByDay{312, []Days{"sunday", nildays}, []RecurByDay2{"of", []Months{"february", []Months1{"january", "march"}}}, "17:23"}}},
func ToByDay(res *gen.SchedParseRes) (byDay *SchedRecurByDay) {

	if (gen.C_RecurByDay != res.Type) || (nil == res.Item) {
		return
	}

	var rbd gen.RecurByDay
	var ok bool
	rbd, ok = res.Item.(gen.RecurByDay)
	if !ok {
		return
	}

	e := new(SchedRecurByDay)
	e.Days = map[string]struct{}{}
	e.Months = map[string]struct{}{}

	e.Ord = rbd.Ord
	e.Clock = rbd.Clock

	for _, v := range rbd.Days {
		e.Days[v] = _emptystruct
	}

	byDay = e

	// optional
	for _, v := range rbd.Months {
		e.Months[v] = _emptystruct
	}

	return
}
