package main

import (
	"fmt"
	"time"
	_ "time/tzdata"
)

func timeAndDate1() {
	fmt.Println("timeAndDate1()")

	current := time.Now()
	specific := time.Date(1999, 12, 14, 13, 45, 23, 00, time.Local) // Winter time
	unixTime := time.Unix(1433228090, 0).Local()                    // Summer time

	Printfln("Current time: %v", current)
	Printfln("Specific time: %v", specific)
	Printfln("Unix time: %v", unixTime)
}

func timeAndDate2() {
	fmt.Println("timeAndDate2()")

	layout := "02.01.2006 03:04:05" // Values should be used as in documentation ("01/02 03:04:05PM '06 -0700"), but format can be any

	current := time.Now()
	specific := time.Date(1999, 12, 14, 13, 45, 23, 00, time.Local) // Winter time
	unixTime := time.Unix(1433228090, 0).Local()                    // Summer time

	Printfln("Current time: %v", current.Format(layout))
	Printfln("Current time: %v", current.Format(time.RFC3339))
	Printfln("Specific time: %v", specific.Format(layout))
	Printfln("Unix time: %v", unixTime.Format(layout))
}

func timeAndDate3() {
	fmt.Println("timeAndDate3()")

	layout := "02|01|2006"

	dates := []string{
		"25|03|2023",
		"32|03|2023",
	}

	for _, dateAsString := range dates {
		Printfln("Date as string: %v", dateAsString)

		if date, err := time.Parse(layout, dateAsString); err == nil {
			Printfln("Parsed date: %v", date)
		} else {
			Printfln("Parsing error: %v", err)
		}

		fmt.Println()
	}
}

func timeAndDate4() {
	fmt.Println("timeAndDate4()")

	dateAsString := "2020-01-01T12:12:12Z"
	date, _ := time.Parse(time.RFC3339, dateAsString)

	Printfln("Date as string: %v", dateAsString)
	Printfln("Parsed date: %v", date)
}

func timeAndDate5() {
	fmt.Println("timeAndDate5()")

	layout := "02|01|2006"

	dates := []string{
		"25|03|2023",
		"32|03|2023",
	}

	for _, dateAsString := range dates {
		Printfln("Date as string: %v", dateAsString)

		loc, _ := time.LoadLocation("America/Ensenada")

		if date, err := time.ParseInLocation(layout, dateAsString, loc); err == nil {
			Printfln("Parsed date: %v", date)
		} else {
			Printfln("Parsing error: %v", err)
		}

		fmt.Println()
	}
}

func timeAndDate6() {
	fmt.Println("timeAndDate6()")

	layout := "02|01|2006"

	dates := []string{
		"25|03|2023",
		"32|03|2023",
	}

	for _, dateAsString := range dates {
		Printfln("Date as string: %v", dateAsString)

		loc := time.FixedZone("PST", -8*60*60)

		if date, err := time.ParseInLocation(layout, dateAsString, loc); err == nil {
			Printfln("Parsed date: %v", date)
		} else {
			Printfln("Parsing error: %v", err)
		}

		fmt.Println()
	}
}

func timeAndDate7() {
	fmt.Println("timeAndDate7()")

	currentTime := time.Now()

	timeAdd := currentTime.Add(time.Minute * 10)
	timeSub := timeAdd.Sub(currentTime)
	timeAddDate := currentTime.AddDate(1, 2, 3)

	Printfln("Current time: %v", currentTime)
	Printfln("Current time + 10 min: %v", timeAdd)
	Printfln("Duration between timeAdd and current: %v", timeSub)
	Printfln("Current time + 1 year, 2 month and 3 days: %v", timeAddDate)

	Printfln("Is timeAdd after current time: %v", timeAdd.After(currentTime))
	Printfln("Is timeAdd before current time: %v", timeAdd.Before(currentTime))
	Printfln("Is timeAdd equal current time: %v", timeAdd.Equal(currentTime))
	Printfln("Is timeAdd zero: %v", timeAdd.IsZero())
	Printfln("Location for timeAdd: %v", timeAdd.Location())
	Printfln("TimeAdd in PST time zone: %v", timeAdd.In(time.FixedZone("PST", -8*60*60)))

	Printfln("TimeAdd rounded to hour: %v", timeAdd.Round(time.Hour))
	Printfln("TimeAdd truncated to hour: %v", timeAdd.Truncate(time.Hour))
}

func timeAndDate8() {
	fmt.Println("timeAndDate8()")

	date1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
	date2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")

	Printfln("date1 equals date2: %v", date1.Equal(date2))
	Printfln("date1 == date2: %v", date1 == date2)
}

func timeAndDate9() {
	fmt.Println("timeAndDate9()")

	d := time.Hour*6 + time.Minute*23 + time.Second*45

	Printfln("Original: %v", d)
	Printfln("Total hours: %v", d.Hours())
	Printfln("Total minutes: %v", d.Minutes())
	Printfln("Total seconds: %v", d.Seconds())
	Printfln("Total milliseconds: %v", d.Milliseconds())
	Printfln("Total nanoseconds: %v", d.Nanoseconds())
	Printfln("Round: %v", d.Round(time.Hour))
	Printfln("Truncate: %v", d.Truncate(time.Hour))

}

func timeAndDate10_1(d time.Duration) string {
	sign := 1

	if d < 0 {
		sign = -1
		d = -1 * d
	}

	years := int(d.Hours() / (24 * 365))
	extraDays := years / 4 // Amount of extra days due to leap years (366 days) (1 time per 4 years)
	d = d - (time.Hour * time.Duration(24*365*years+24*extraDays))

	days := int(d.Hours() / 24)
	d = d - (time.Hour * 24 * time.Duration(days))

	return fmt.Sprintf("%v year(s), %v day(s), %v hour(s)", sign*years, days, d.Truncate(time.Hour).Hours())
}

func timeAndDate10() {
	fmt.Println("timeAndDate10()")

	date1 := time.Date(1984, 1, 1, 0, 0, 0, 0, time.Local)
	date2 := time.Date(2030, 1, 1, 0, 0, 0, 0, time.Local)

	Printfln("Since 1984: %v", timeAndDate10_1(time.Since(date1)))
	Printfln("Until 2030: %v", timeAndDate10_1(time.Until(date2)))
	Printfln("Until 1984: %v", timeAndDate10_1(time.Until(date1)))
}

func timeAndDate11() {
	fmt.Println("timeAndDate11()")

	durationAsString := "15h33m44s"
	d, _ := time.ParseDuration(durationAsString)

	Printfln("Parsed duration: %v", d)
}

func timeAndDate12_1(ch chan<- string) {
	names := []string{"Alice", "Bob", "Charlie"}

	for _, name := range names {
		ch <- name
		time.Sleep(time.Second * 1)
	}
	close(ch)
}

func timeAndDate12() {
	fmt.Println("timeAndDate12()")

	ch := make(chan string)

	go timeAndDate12_1(ch)

	for name := range ch {
		Printfln("Received name from channel: %v", name)
	}
}

func timeAndDate13_1(ch chan<- string) {
	names := []string{"Alice", "Bob", "Charlie"}

	for _, name := range names {
		ch <- name
	}
	close(ch)
}

func timeAndDate13() {
	fmt.Println("timeAndDate13()")

	ch := make(chan string)

	time.AfterFunc(time.Second*2, func() {
		timeAndDate13_1(ch)
	})

	for name := range ch {
		Printfln("Received name from channel: %v", name)
	}
}

func timeAndDate14_1(ch chan<- string) {
	Printfln("Waiting for initial duration")
	delayCh := time.After(time.Second)
	<-delayCh
	Printfln("Initial duration elapsed")

	names := []string{"Alice", "Bob", "Charlie"}

	for _, name := range names {
		ch <- name
		time.Sleep(time.Second * 1)
	}
	close(ch)
}

func timeAndDate14() {
	fmt.Println("timeAndDate14()")

	ch := make(chan string)

	go timeAndDate14_1(ch)

	for name := range ch {
		Printfln("Received name from channel: %v", name)
	}
}

func timeAndDate15_1(ch chan<- string) {
	Printfln("Waiting for initial duration")
	//_ = <-time.After(time.Second * 2)
	Printfln("Initial duration elapsed")

	names := []string{"Alice", "Bob", "Charlie"}

	for _, name := range names {
		ch <- name
		time.Sleep(time.Second * 3)
	}
	close(ch)
}

func timeAndDate15() {
	fmt.Println("timeAndDate15()")

	ch := make(chan string)

	go timeAndDate15_1(ch)

	channelOpen := true

	for channelOpen {
		select {
		case name, ok := <-ch:
			if !ok {
				channelOpen = false
				break
			} else {
				Printfln("Received name from channel: %v", name)
			}
		case <-time.After(time.Second * 3):
			fmt.Println("Timeout")
		}
	}
}

func timeAndDate16() {
	fmt.Println("timeAndDate16()")

	select {
	case <-time.After(time.Second * 2):
		Printfln("Timeout, current time: %v", time.Now())
	}
}

func timeAndDate17_1(ch chan<- string) {
	t := time.NewTimer(time.Minute * 10)

	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Resetting timer")
		t.Reset(time.Second)
	}()

	Printfln("Waiting for initial duration")
	_ = <-t.C
	Printfln("Initial duration elapsed")

	names := []string{"Alice", "Bob", "Charlie"}

	for _, name := range names {
		ch <- name
	}
	close(ch)
}

func timeAndDate17() {
	fmt.Println("timeAndDate17()")

	ch := make(chan string)

	go timeAndDate17_1(ch)

	for name := range ch {
		Printfln("Received name from channel: %v", name)
	}
}

func timeAndDate18_1(ch chan<- string) {
	names := []string{"Alice", "Bob", "Charlie"}

	tickCh := time.Tick(time.Second)
	index := 0

	for range tickCh {
		ch <- names[index]
		index++
		if index == len(names) {
			index = 0
		}
	}

	close(ch)
}

func timeAndDate18() {
	fmt.Println("timeAndDate18()")

	ch := make(chan string)

	go timeAndDate18_1(ch)

	for name := range ch {
		Printfln("Received name from channel: %v", name)
	}

}

func timeAndDate19_1(ch chan<- string) {
	names := []string{"Alice", "Bob", "Charlie"}

	ticker := time.NewTicker(time.Second)

	index := 0

	for range ticker.C {
		ch <- names[index]
		index++
		if index == len(names) {
			index = 0
			ticker.Stop()
			break
		}
	}

	close(ch)
}

func timeAndDate19() {
	fmt.Println("timeAndDate19()")

	ch := make(chan string)

	go timeAndDate19_1(ch)

	for name := range ch {
		Printfln("Received name from channel: %v", name)
	}

}
