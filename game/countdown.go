package game

import "time"

var (
	startTime        time.Time
	countDownBegin   bool = false
	countDownSeconds int
)

func StartCountdown(
	countDownSecondsInput int,
) {
	if countDownBegin {
		return
	}
	startTime = time.Now()
	countDownBegin = true
	countDownSeconds = countDownSecondsInput
}

func GetCountdown() int {
	if countDownBegin {
		elapsed := time.Since(startTime)
		remaining := countDownSeconds - int(elapsed.Seconds())
		if remaining <= 0 {
			return 0
		}
		return remaining
	}
	return 999
}

func IsCountdownFinished() bool {
	if countDownBegin {
		elapsed := time.Since(startTime)
		remaining := countDownSeconds - int(elapsed.Seconds())
		if remaining <= 0 {
			return true
		}
	}
	return false
}
