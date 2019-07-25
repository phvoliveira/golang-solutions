package main

import (
	"fmt"
	"strconv"
	"strings"
)

type work struct {
	Count                   int
	StartHour               int
	StartMinute             int
	EndHour                 int
	EndMinute               int
	MorningInterval         int
	HighestMidRangeInterval int
	NightInterval           int
}

func (w *work) setStartTime(hour, minute int) {
	if w.StartHour >= hour {
		w.StartHour = hour

		if w.StartMinute > minute {
			w.StartMinute = minute
		}
	}
}

func (w *work) setEndTime(hour, minute int) {
	if w.EndHour <= hour {
		w.EndHour = hour

		if w.EndMinute < minute {
			w.StartMinute = minute
		}
	}
}

func (w *work) checkMidRangeInterval(startHour, startMinute, endHour, endMinute int) {
	if w.Count > 0 {
		midRangeInterval := 0
		if w.EndHour < startHour || (w.EndHour == startHour && w.EndMinute < startMinute) {
			mMin := (startMinute - w.EndMinute)
			if mMin < 0 {
				mMin *= -1
			}
			midRangeInterval = ((startHour - w.EndHour) * 60) + mMin
		} else if w.StartHour > endHour || (w.StartHour == endHour && w.StartMinute > endMinute) {
			mMin := (w.StartMinute - endMinute)
			if mMin < 0 {
				mMin *= -1
			}
			midRangeInterval = ((w.StartHour - endHour) * 60) + mMin
		}

		if w.HighestMidRangeInterval == 0 || w.HighestMidRangeInterval < midRangeInterval {
			w.HighestMidRangeInterval = midRangeInterval
		}
	}
}

func (w *work) checkDayInterval() {
	w.MorningInterval = (w.StartHour * 60) + w.StartMinute
	w.NightInterval = ((24 - w.EndHour) * 60)
	if w.EndHour < 24 {
		w.NightInterval += w.EndMinute
	}
}

func solution(S string) int {
	dayInterval := 0
	highestSleepTime := 0
	dayMap := map[string]*work{}
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	r := strings.Split(S, "\n")
	for _, v := range r {
		rs := strings.Split(v, " ")
		w := dayMap[rs[0]]
		if w == nil {
			w = &work{StartHour: 24, StartMinute: 60}
		}
		if len(rs) == 2 {
			ts := strings.Split(rs[1], "-")
			h1, m1, err := retrieveHourAndMinute(ts[0])
			if err != nil {
				fmt.Println(err)
			}
			h2, m2, err := retrieveHourAndMinute(ts[1])
			if err != nil {
				fmt.Println(err)
			}
			w.checkMidRangeInterval(h1, m1, h2, m2)
			w.setStartTime(h1, m1)
			w.setEndTime(h2, m2)
			w.checkDayInterval()
		}
		w.Count++
		dayMap[rs[0]] = w
	}

	for i := 0; i < len(days); i++ {
		if i == 0 {
			highestSleepTime = dayMap[days[i]].MorningInterval
		}

		if highestSleepTime < dayMap[days[i]].HighestMidRangeInterval {
			highestSleepTime = dayMap[days[i]].HighestMidRangeInterval
		}

		dayInterval = dayMap[days[i]].NightInterval

		if i < len(days)-1 {
			dayInterval += dayMap[days[i+1]].MorningInterval
		}

		if highestSleepTime < dayInterval {
			highestSleepTime = dayInterval
		}
	}

	return highestSleepTime
}

func retrieveHourAndMinute(t string) (int, int, error) {
	ts := strings.Split(t, ":")
	h, err := strconv.Atoi(ts[0])
	if err != nil {
		return 0, 0, err
	}
	m, err := strconv.Atoi(ts[1])
	if err != nil {
		return 0, 0, err
	}

	return h, m, nil
}

func main() {
	fmt.Printf("Primeiro Resultado: %v\n", solution("Sun 10:00-20:00\nFri 05:00-10:00\nFri 16:30-23:50\nSat 10:00-24:00\nSun 01:00-04:00\nSat 02:00-06:00\nTue 03:30-18:15\nTue 19:00-20:00\nWed 04:25-15:14\nWed 15:14-22:40\nThu 00:00-23:59\nMon 05:00-13:00\nMon 15:00-21:00"))
	fmt.Printf("Segundo Resultado: %v\n", solution("Mon 01:00-23:00\nTue 01:00-23:00\nWed 01:00-23:00\nThu 01:00-23:00\nFri 01:00-23:00\nSat 01:00-23:00\nSun 01:00-21:00"))
}
