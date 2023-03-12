package roaringcompany

import (
	"strconv"
	"time"
)

const DatabaseTimeout = time.Second * 5

func GetTimeStamp() (mmdd uint16, year uint8, timeofDayInSeconds int) {
	nowis := time.Now()
	return uint16(MMDD_From_Month_Date(int(nowis.Month()), int(nowis.Day()))), uint8(YY_From_YYYY(nowis.Year())), TimeInSeconds(nowis)
}

func CurrentTimeStampV2() (day uint8, month uint8, year uint8, timeofDayInSeconds int) {
	nowis := time.Now()
	return uint8(nowis.Day()), uint8(int(nowis.Month())), uint8(YY_From_YYYY(nowis.Year())), TimeInSeconds(nowis)
}

func CurrentTimeStamp() (day int, month int, year int, timeofDayInSeconds int) {
	nowis := time.Now()
	return nowis.Day(), int(nowis.Month()), YY_From_YYYY(nowis.Year()), TimeInSeconds(nowis)
}

func TimeInSeconds(t time.Time) int {
	return 60*60*t.Hour() + 60*t.Minute() + t.Second()
}

func YY_From_YYYY(passedYYYY int) int {
	stringTargetYear := strconv.Itoa(passedYYYY)
	yy := stringTargetYear[2:]
	returnYY, _ := strconv.Atoi(yy)

	return returnYY
}

func MMDD_From_Month_Date(month, date int) int {

	stringTargetMonth := Digit_NINE
	monthStr := strconv.Itoa(month)
	if month < TEN {
		stringTargetMonth += monthStr
	} else {
		stringTargetMonth = monthStr
	}

	stringTargetDate := Digit_ZERO
	dateStr := strconv.Itoa(date)
	if date < TEN {
		stringTargetDate += dateStr
	} else {
		stringTargetDate = dateStr
	}

	returnYY, _ := strconv.Atoi(stringTargetMonth + stringTargetDate)

	return returnYY
}
