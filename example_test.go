package schedparser_test

import (
	"fmt"
)

func ExampleParseSched() {

	// parses a token
	token := "every 3 secs from 06:30 to 08:30"
	ret := ParseSched(v.token, func(bytime *SchedRecurByTime, byday *SchedRecurByDay) {
		// do something
	})
	
	fmt.Println(token, ret)

	// RecurByTime = every decimal_digit (hrs|mins|secs)  ["from" time "to" time] .
	// RecurByDay = (every|ordinal) Days ["of" (Months)] (time) .

	// VALID input examples
	// "every 21 hrs"
	// "every 3 secs from 06:30 to 08:30"
	// "every 05 mins from 06:30 to 09:30"
	// "every monday 05:30"
	// "every mon 05:30"
	// "every wed 15:40"
	// "every sat 05:10"
	// "every tue of january 17:23"
	// "1st thu of sept 17:23"
	// "321st sat of dec 17:23"
	// "343rd sun of oct 17:23"
	// "313th sun of feb 17:23"
	// "312th sun of feb 17:23"
	// "232nd sun 17:23"
	// "232nd mon,sun 17:23"
	// "312th sun of feb,jan,mar 17:23"
	//
	// INVALID input examples:
	// "every tuesday 25:30"
	// "every friday 24:70"
	// "0th thu of sept 17:23"
	// "11st thu of january 17:23"
	// "211st thu of january 17:23"
	// "321rd sat of dec 17:23"
	// "313rd sun of oct 17:23"

}
