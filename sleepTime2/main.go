package main

import (
	"fmt"
	"strings"
)

type work struct {
	earlyTime          string
	laterTime          string
	biggestDayInterval int
}

func solution(S string) int {
	highestSleepTime := 0
	dayMap := map[string]*work{}

	r := strings.Split(S, "\n")
	for _, v := range r {
		rs := strings.Split(v, " ")
		w := dayMap[rs[0]]
		if w == nil {
			w = &work{
				earlyTime: "24:00",
				laterTime: "00:00",
			}
		}
		if len(rs) == 2 {
			ts := strings.Split(rs[1], "-")

			if w.earlyTime > ts[0] {
				w.earlyTime = ts[0]
			}

			if w.laterTime < ts[1] {
				w.laterTime = ts[1]
			}
		}

		dayMap[rs[0]] = w
	}

	return highestSleepTime
}

func main() {
	fmt.Printf("Primeiro Resultado: %v\n", solution("Sun 10:00-20:00\nFri 05:00-10:00\nFri 16:30-23:50\nSat 10:00-24:00\nSun 01:00-04:00\nSat 02:00-06:00\nTue 03:30-18:15\nTue 19:00-20:00\nWed 04:25-15:14\nWed 15:14-22:40\nThu 00:00-23:59\nMon 05:00-13:00\nMon 15:00-21:00"))
	fmt.Printf("Segundo Resultado: %v\n", solution("Mon 01:00-23:00\nTue 01:00-23:00\nWed 01:00-23:00\nThu 01:00-23:00\nFri 01:00-23:00\nSat 01:00-23:00\nSun 01:00-21:00"))
}
