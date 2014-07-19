package schedparser

import (
	"github.com/noypi/schedparser/gen"
)

type SchedRecurByDay struct {
	Ord int

	// value is just an empty struct = struct{}{}
	Days map[string]struct{}

	// value is just an empty struct = struct{}{}
	Months map[string]struct{}

	Clock string
}

type SchedRecurByTime struct {
	N         int
	ClockLbl  string
	FromClock string
	ToClock   string
}

type OnSchedParseCb func(*SchedRecurByTime, *SchedRecurByDay)

func ParseSched(token string, cb OnSchedParseCb) int {
	return gen.SchedParse(&gen.SchedLex{S: token, OnSchedParse: func(res *gen.SchedParseRes) {
		bytime, byday := ToEasy(res)
		cb(bytime, byday)
	}})
}

func ToEasy(res *gen.SchedParseRes) (byTime *SchedRecurByTime, byDay *SchedRecurByDay) {
	byTime = ToByTime(res)
	byDay = ToByDay(res)
	return
}

func (this SchedRecurByDay) HasDay(s string) bool {
	_, b := this.Days[s]
	return b
}

func (this SchedRecurByDay) HasMonth(s string) bool {
	_, b := this.Months[s]
	return b
}
