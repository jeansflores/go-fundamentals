package sampledata

import (
	"time"

	"example.com/platform/data"
)

func GetData() []data.Printable {
	var courses []data.Printable

	jean := data.Instructor{Name: "Jean Flores"}

	goCourse := data.Course{Name: "Go Fundamentals", Duration: 4}
	goCourse.Instructor = jean

	kotlinWorkshop := data.Workshop{Date: time.Now().AddDate(0, 0, 15)}
	kotlinWorkshop.Name = "Kotlin for Android"
	kotlinWorkshop.Instructor = jean

	swiftWorkshop, _ := data.NewWorkshop("Swift", 3, jean, time.Now())

	return append(courses, goCourse, kotlinWorkshop, swiftWorkshop)
}
