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
	if guests == nil {
		return []Load{}
	}
	var tmpSlice []Load
	var tmpStruct Load
	var tmpDateSlice []int
	var lastValue int = -1
	common_cal := make(map[int]int, len(guests))

	for _, v := range guests {
		days := v.CheckOutDate - v.CheckInDate
		for y := range days + 1 {
			someDate := v.CheckInDate + y
			_, ok := common_cal[someDate]
			if !ok {
				common_cal[someDate] = 1
				tmpDateSlice = append(tmpDateSlice, someDate)
			} else {
				common_cal[someDate] += 1
			}
		}
	}

	for _, v := range guests {
		_, ok := common_cal[v.CheckOutDate]
		if ok {
			common_cal[v.CheckOutDate]--
		}
	}
	sort.Ints(tmpDateSlice)
	for _, v := range tmpDateSlice {
		if common_cal[v] != lastValue {

			tmpStruct.StartDate = v
			tmpStruct.GuestCount = common_cal[v]
			tmpSlice = append(tmpSlice, tmpStruct)
		}
		lastValue = common_cal[v]
	}


	return tmpSlice
}
