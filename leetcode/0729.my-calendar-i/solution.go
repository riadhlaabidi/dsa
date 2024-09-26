package main

import (
	"fmt"
)

type Event struct {
	start int
	end   int
}

type Calendar struct {
	calendar []Event
}

func NewCalendar() Calendar {
	return Calendar{}
}

func (c *Calendar) Book(start, end int) bool {
	for _, event := range c.calendar {
		if start < event.end && end > event.start {
			return false
		}
	}

	c.calendar = append(c.calendar, Event{start: start, end: end})
	return true
}

func main() {
	cal := NewCalendar()
	events := [3]Event{{10, 20}, {15, 25}, {20, 30}}

	expected := [3]bool{true, false, true}
	var actual [3]bool

	for i := range actual {
		actual[i] = cal.Book(events[i].start, events[i].end)
	}

	if actual != expected {
		fmt.Printf("Expected %v, but got %v", expected, actual)
		return
	}

	fmt.Printf("Correct %v", actual)
}
