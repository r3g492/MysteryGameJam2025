package game

import "time"

var (
	startTime        time.Time
	CountDownBegin   bool = false
	countDownSeconds int
)

func StartCountdown(
	countDownSecondsInput int,
) {
	if CountDownBegin {
		return
	}
	startTime = time.Now()
	CountDownBegin = true
	countDownSeconds = countDownSecondsInput
}

func GetCountdown() int {
	if CountDownBegin {
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
	if CountDownBegin {
		elapsed := time.Since(startTime)
		remaining := countDownSeconds - int(elapsed.Seconds())
		if remaining <= 0 {
			return true
		}
	}
	return false
}
