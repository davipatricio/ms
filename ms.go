package ms

import (
	"regexp"
	"strconv"
)

const secondsMs int = 1000
const minutesMs int = secondsMs * 60
const hoursMs int = minutesMs * 60
const daysMs int = hoursMs * 24
const weeksMs int = daysMs * 7
const yearsMs float64 = float64(daysMs) * 365.25

var regex = regexp.MustCompile(`^(-?(?:\d+)?\.?\d+) *(milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|weeks?|w|years?|yrs?|y)?$`)

// Converts a string to milisseconds
// If a invalid string is passed, 0 is returned
// 	ms.ConvertInt("1m") // Returns: 60000
// 	ms.ConvertInt("1 m") // Returns: 60000
// 	ms.ConvertInt("1s") // Returns: 1000
// 	ms.ConvertInt("-5h") // Returns: -18000000
// 	ms.ConvertInt("1.5s") // Returns: 0 (use ConvertFloat instead)
// 	ms.ConvertInt("2w")
// 	ms.ConvertInt("5y")
func ConvertInt(val string) int {
	matches := regex.FindAllStringSubmatch(val, -1)

	if len(matches) < 1 {
		return 0
	}

	number, numberErr := strconv.Atoi(matches[0][1])

	if numberErr != nil {
		return 0
	}

	return parseInt(number, matches)
}

// Converts a both float and int strings to milisseconds
// If a invalid string is passed, 0 is returned
// 	ms.ConvertInt("2.5h") // Returns: 9000000
// 	ms.ConvertInt("1s") // Returns: 1000
// 	ms.ConvertInt("-60s") // Returns: -60000
func ConvertFloat(val string) float64 {
	matches := regex.FindAllStringSubmatch(val, -1)

	if len(matches) < 1 {
		return 0
	}

	float, floatErr := strconv.ParseFloat(matches[0][1], 64)

	if floatErr != nil {
		return 0
	}

	return parseFloat(float, matches)
}

func parseFloat(num float64, matches [][]string) float64 {
	// Get the unit
	unit := matches[0][2]

	switch unit {
	case "ms", "msec", "msecs", "millisecond", "milliseconds":
		return num
	case "s", "sec", "secs", "second", "seconds":
		return num * float64(secondsMs)
	case "m", "min", "mins", "minute", "minutes":
		return num * float64(minutesMs)
	case "h", "hour", "hours":
		return num * float64(hoursMs)
	case "d", "day", "days":
		return num * float64(daysMs)
	case "w", "week", "weeks":
		return num * float64(weeksMs)
	case "y", "yr", "yrs", "year", "years":
		return num * yearsMs
	default:
		return 0
	}
}

func parseInt(n int, matches [][]string) int {
	// Get the unit
	unit := matches[0][2]

	switch unit {
	case "ms", "msec", "msecs", "millisecond", "milliseconds":
		return n
	case "s", "sec", "secs", "second", "seconds":
		return n * secondsMs
	case "m", "min", "mins", "minute", "minutes":
		return n * minutesMs
	case "h", "hour", "hours":
		return n * hoursMs
	case "d", "day", "days":
		return n * daysMs
	case "w", "week", "weeks":
		return n * weeksMs
	case "y", "yr", "yrs", "year", "years":
		return n * int(yearsMs)
	default:
		return 0
	}
}
