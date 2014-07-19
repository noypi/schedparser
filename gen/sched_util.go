package gen

import (
	"fmt"
	"github.com/noypi/logger"
	"strconv"
	"strings"
)

const (
	C_RecurByTime = iota
	C_RecurByDay
)

var (
	_, _, _, _, _, _, TRACE = logger.StartLogger("sp")
)

var _emptystruct struct{} = struct{}{}

type SchedParseRes struct {
	Type int
	Item interface{}
}

func (this SchedLex) OnSched(item interface{}) {

	res := &SchedParseRes{Item: item}

	switch t := item.(type) {
	default:
		fmt.Printf("OnSched() err, invalid type, %T=", t)
		return
	case RecurByTime:
		res.Type = C_RecurByTime
	case RecurByDay:
		res.Type = C_RecurByDay
	}

	if nil != this.OnSchedParse {
		this.OnSchedParse(res)
	}

	if nil != this.Itemchan {
		this.Itemchan <- res
	}
}

func (this *SchedLex) peek() (bret byte) {
	bret = this.next()
	this.pos -= 1
	return
}

func (this *SchedLex) back() {
	if 0 < this.pos {
		this.pos -= 1
	}
	return
}

func (this *SchedLex) next() (bret byte) {

	if this.pos < len(this.S) {
		bret = byte(this.S[this.pos])
		this.buf = append(this.buf, bret)
	} else {
		bret = 0
	}
	this.pos += 1
	return
}

func (this SchedLex) data() (bb []byte) {
	if this.pos < len(this.S) {
		bb = this.buf[:len(this.buf)-1]
	} else {
		bb = this.buf
	}
	return
}

func (this SchedLex) Error(s string) {
	fmt.Println("error: " + s)
}

func (this SchedLex) getInt() (n int, err error) {
	s := string(this.data())
	if n, err = strconv.Atoi(strings.TrimSpace(s)); nil != err {
		this.Error(err.Error() + ";s=" + s)
	}
	return
}

func (this SchedLex) getClockLbl() (s string) {
	lbl := string(this.data())
	switch {
	case "hrs" == lbl || "hours" == lbl:
		s = "hours"
	case "mins" == lbl || "minutes" == lbl:
		s = "minutes"
	case "secs" == lbl || "seconds" == lbl:
		s = "seconds"
	}
	return
}

func (this SchedLex) getDay() (s string) {
	day := string(this.data())
	switch {
	case "mon" == day || "monday" == day:
		s = "monday"
	case "tue" == day || "tuesday" == day:
		s = "tuesday"
	case "wed" == day || "wednesday" == day:
		s = "wednesday"
	case "thu" == day || "thursday" == day:
		s = "thursday"
	case "fri" == day || "friday" == day:
		s = "friday"
	case "sat" == day || "saturday" == day:
		s = "saturday"
	case "sun" == day || "sunday" == day:
		s = "sunday"
	}
	return
}

func (this SchedLex) getMonth() (s string) {
	mon := string(this.data())
	switch {
	case "jan" == mon || "january" == mon:
		s = "january"
	case "feb" == mon || "february" == mon:
		s = "february"
	case "mar" == mon || "march" == mon:
		s = "march"
	case "apr" == mon || "april" == mon:
		s = "april"
	case "may" == mon:
		s = "may"
	case "jun" == mon || "june" == mon:
		s = "june"
	case "jul" == mon || "july" == mon:
		s = "july"
	case "aug" == mon || "august" == mon:
		s = "august"
	case "sep" == mon || "sept" == mon || "september" == mon:
		s = "september"
	case "oct" == mon || "october" == mon:
		s = "october"
	case "nov" == mon || "november" == mon:
		s = "november"
	case "dec" == mon || "december" == mon:
		s = "december"
	}
	return
}

func (this SchedLex) getStr() (s string) {
	s = strings.TrimSpace(string(this.data()))
	return
}

func (this SchedLex) getOrdinal() (n int, err error) {
	s := strings.TrimRight(string(this.data()), "st|nd|rd|th")
	if n, err = strconv.Atoi(strings.TrimSpace(s)); nil != err {
		this.Error(err.Error() + ";s=" + s)
	}
	return
}
