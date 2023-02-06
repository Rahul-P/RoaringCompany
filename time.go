package roaringcompany

import (
	"strconv"
	"time"
)

const dbTimeout = time.Second * 5

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
