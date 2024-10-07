package main

import (
	"fmt"
)

// Good morning! : When the time is between 6:00:01 am to 11:00:00 am
// Good afternoon! : When the time is between 11:00:01 am to 4:00:00 pm
// Good evening! : When the time is between 4:00:01 pm to 9:00:00 pm
// Good night! : When the time is between 9:00:01 pm to 6:00:00 am

func main() {
	const goodMorning string = "Good morning!"
	const goodAfternoon string = "Good afternoon!"
	const goodEvening string = "Good evening!"
	const goodNight string = "Good night!"
	fmt.Print("Enter the Time Of the Day (HH:mm:ss am/pm):")
	var meridian string
	var hour, min, sec int
	fmt.Scanf("%d:%d:%d %s", &hour, &min, &sec, &meridian)
	fmt.Scanln(&meridian)

	if meridian == "pm" && hour != 12 {
		hour += 12
	}
	if meridian == "am" && hour == 12 {
		hour = 0
	}
	timeInSeconds := hour*3600 + min*60 + sec
	if timeInSeconds <= 6*3600 {
		fmt.Println(goodNight)
	} else if timeInSeconds <= 11*3600 {
		fmt.Println(goodMorning)
	} else if timeInSeconds <= 16*3600 {
		fmt.Println(goodAfternoon)
	} else if timeInSeconds <= 21*3600 {
		fmt.Println(goodEvening)
	} else {
		fmt.Println(goodNight)
	}

}
