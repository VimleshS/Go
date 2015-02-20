package main

import (
	"fmt"
	"os"
	"strings"
)

type CustomerType int

const (
	Regular CustomerType = iota
	Rewarded
)

type RateDays int

const (
	WeekDay RateDays = iota
	WeekEnd
)

type Hotel struct {
	rating       int
	name         string
	weekDayRates []float64
	weekEndRates []float64
}

func (h *Hotel) PushRates(dayType RateDays, rates []float64) {
	switch dayType {
	case WeekDay:
		h.weekDayRates = rates
	case WeekEnd:
		h.weekEndRates = rates
	}
}

func (h *Hotel) AddInWeekDayRate(rt1 float64) {
	h.weekDayRates[Regular] = rt1
}

func (h Hotel) GetRate(dayType RateDays, custType CustomerType) float64 {
	switch dayType {
	case WeekDay:
		return h.weekDayRates[custType]
	case WeekEnd:
		return h.weekEndRates[custType]
	default:
		return 0.0
	}
}

func (h *Hotel) String() string {
	return fmt.Sprintf("%s having rating %d", h.name, h.rating)
}

//http://stackoverflow.com/questions/4498998/how-to-initialize-members-in-go-struct
func NewHotel(_name string, _rating int) Hotel {
	sm := Hotel{
		name:         _name,
		rating:       _rating,
		weekDayRates: make([]float64, 2),
		weekEndRates: make([]float64, 2),
	}
	return sm
}

func GetDaysType(daystr string) RateDays {
	daystr = strings.ToUpper(daystr)
	if daystr == "MON" || daystr == "TUES" || daystr == "WED" || daystr == "THUR" || daystr == "FRI" {
		return WeekDay
	} else {
		return WeekEnd
	}
}

func main() {
	Lakewood := NewHotel("Lakewood", 3)
	Lakewood.PushRates(WeekDay, []float64{110, 80})
	Lakewood.PushRates(WeekEnd, []float64{90, 80})

	Bridgewood := NewHotel("Bridgewood", 4)
	Bridgewood.PushRates(WeekDay, []float64{160, 110})
	Bridgewood.PushRates(WeekEnd, []float64{60, 50})

	Ridgewood := NewHotel("Ridgewood", 5)
	Ridgewood.PushRates(WeekDay, []float64{220, 100})
	Ridgewood.PushRates(WeekEnd, []float64{150, 40})

	//input := "Regular: 16Mar2009(mon), 17Mar2009(tues), 18Mar2009(wed)"
	//input := "Rewards: 26Mar2009(thur), 27Mar2009(fri), 28Mar2009(sat)"
	//input := "Regular: 20Mar2009(fri), 21Mar2009(sat), 22Mar2009(sun)"

	fmt.Println("----------------------------------------")
	arg := os.Args[1:]
	//fmt.Println(arg)
	//input := arg
	input := strings.Join(arg, " ")
	fmt.Println(input)
	fmt.Println("----------------------------------------")

	minRateHotelMap := make(map[*Hotel]float64)
	minRateHotelMap[&Lakewood] = 0.0
	minRateHotelMap[&Bridgewood] = 0.0
	minRateHotelMap[&Ridgewood] = 0.0

	usersInput := strings.Fields(input)

	var custType CustomerType
	_days := make([]RateDays, 0)
	for i, r := range usersInput {
		if i == 0 {
			if r == "Regular:" {
				custType = Regular
			} else {
				custType = Rewarded
			}
		}

		st := strings.Index(r, "(")
		ed := strings.Index(r, ")")
		if st > 0 {
			_dstr := r[st+1 : ed]
			_days = append(_days, GetDaysType(_dstr))
		}
	}

	//Find the min Rate
	for hotel := range minRateHotelMap {
		mRate := 0.0
		for _, rateday := range _days {
			t := hotel.GetRate(rateday, custType)
			//fmt.Println(custType)
			mRate = mRate + t
		}
		minRateHotelMap[hotel] = mRate
	}
	//fmt.Println(minRateHotelMap)
	//finding minimum
	var _minHotel *Hotel
	_minAmt := 0.0
	for hotel, currentAmt := range minRateHotelMap {
		if _minHotel == nil {
			_minHotel = hotel
			_minAmt = currentAmt
			continue
		}

		//if currentAmt < _minAmt {
		//	_minHotel = hotel
		//	_minAmt = currentAmt
		//}
		//if (currentAmt == _minAmt) && (_minHotel.rating < hotel.rating) {
		//	_minHotel = hotel
		//	_minAmt = currentAmt
		//}

		if (currentAmt < _minAmt) || ((currentAmt == _minAmt) && (_minHotel.rating < hotel.rating)) {
			_minHotel = hotel
			_minAmt = currentAmt
		}
	}

	fmt.Println(_minHotel)
}
