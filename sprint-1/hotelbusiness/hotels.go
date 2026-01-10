//go:build !solution

package hotelbusiness

import (
	"sort"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	calendar := make(map[int]int, len(guests))
	dates := make([]int, 0)
	result := make([]Load, 0, len(guests))
	lastDayGuest := int(0)
	for _, guest := range guests {
		for i := guest.CheckInDate; i < guest.CheckOutDate; i++ {
			calendar[i]++
		}
		_, ok := calendar[guest.CheckOutDate]
		if !ok {
			calendar[guest.CheckOutDate] = 0
		}
	}
	for day, _ := range calendar {
		dates = append(dates, day)
	}
	sort.Ints(dates)

	for _, value := range dates {
		if calendar[value] != lastDayGuest {
			result = append(result, Load{StartDate: value, GuestCount: calendar[value]})
		}
		lastDayGuest = calendar[value]
	}
	return result
}
