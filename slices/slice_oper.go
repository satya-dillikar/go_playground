package main

import (
	"fmt"
)

type event struct {
	ID    int
	Title string
}

type allEvents []event

var events = allEvents{
	{
		ID:    0,
		Title: "Satya",
	},
	{
		ID:    1,
		Title: "Mike",
	},
	{
		ID:    2,
		Title: "Paul",
	},
	{
		ID:    3,
		Title: "Frank",
	},
}

func createEvent(newEvent event) {
	events = append(events, newEvent)
}

func getOneEvent(eventID int) {

	for _, singleEvent := range events {
		if singleEvent.ID == eventID {
			fmt.Println("getOneEvent", singleEvent)
			break
		}
	}
}

func updateEvent(eventID int, newTitle string) {

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			//events = append(events[:i], events[i:]...)
			events[i].Title = newTitle
			fmt.Println("updateEvent", singleEvent)
			break
		}
	}
}

func deleteEvent(eventID int) {

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Println("deleteEvent", singleEvent)
			break
		}
	}
}

func getAllEvents() {
	fmt.Println("Events:", events)
}

func main() {
	getAllEvents()
	newEvent := event{
		ID:    5,
		Title: "Foo",
	}

	createEvent(newEvent)
	getAllEvents()
	updateEvent(3, "Savir")
	getAllEvents()
	deleteEvent(3)
	getAllEvents()
	getOneEvent(2)
	getAllEvents()
}
