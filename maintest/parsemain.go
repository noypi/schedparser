package main

import (
	"bufio"
	"fmt"
	. "github.com/noypi/schedparser"
	"os"
	"strings"
)

func readline(fi *bufio.Reader) (string, bool) {
	s, err := fi.ReadString('\n')
	if err != nil {
		return "", false
	}
	return s, true
}

func main() {

	fi := bufio.NewReader(os.Stdin)

	var token string
	var ok bool
	for {
		fmt.Printf(": ")
		if token, ok = readline(fi); ok {
			ret := ParseSched(strings.TrimSpace(token), func(bytime *SchedRecurByTime, byday *SchedRecurByDay) {
				fmt.Println("bytime=", bytime)
				fmt.Println("byday=", byday)
			})
			fmt.Println("SchedParse:", ret)
		} else {
			break
		}
	}
}
