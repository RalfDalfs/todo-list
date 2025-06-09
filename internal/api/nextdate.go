package api

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type Weekday int

const TimeFormat = "20060102"

// NextDate вычисляет следующую дату задачи
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("repeat is required")
	}

	rep := strings.Split(repeat, " ")

	date, err := time.Parse(TimeFormat, dstart)
	if err != nil {
		return "", err
	}

	switch rep[0] {
	case "d":
		return handleDaily(now, date, rep)
	case "y":
		return handleYearly(now, date)
	case "w":
		return handleWeekly(now, date, rep)
	case "m":
		return handleMonthly(now, date, rep)
	default:
		return "", errors.New("invalid tags")
	}
}

// checkCorrectDay проверяет на корректность день
func checkCorrectDay(d []int) bool {
	for _, i := range d {
		if i > 31 || i == 0 || i < -2 {
			return false
		}
	}
	return true
}

// convertStringToInt преобразует строку в слайс чисел
func convertStringToInt(s string) ([]int, error) {
	slStr := strings.Split(s, ",")
	var result []int
	for _, d := range slStr {
		i, err := strconv.Atoi(d)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil

}

// contains проверяет есть ли число в слайсе чисел
func contains(weekdays []int, day int) bool {
	for _, weekday := range weekdays {
		if weekday == day {
			return true
		}
	}
	return false
}

// afterNow смотрит не является ли текущая дата больше чем передаваемая
func afterNow(date, now time.Time) bool {
	if date.After(now) {
		return true
	} else {
		return false
	}
}

// handleDaily повторяет задачу каждые n дней
func handleDaily(now, date time.Time, rep []string) (string, error) {
	if len(rep) < 2 {
		return "", errors.New("missing days value")
	}
	days, err := strconv.Atoi(rep[1])
	if err != nil {
		return "", err
	}
	if days > 400 {
		return "", errors.New("days must be less than 400 days")
	}
	if days < 1 {
		return "", errors.New("days must be greater than 0")
	}
	for {
		date = date.AddDate(0, 0, days)
		if afterNow(date, now) {
			break
		}
	}
	return date.Format(TimeFormat), nil
}

// handleYearly повторяет задачу каждый год
func handleYearly(now, date time.Time) (string, error) {
	for {
		date = date.AddDate(1, 0, 0)
		if afterNow(date, now) {
			break
		}
	}
	return date.Format(TimeFormat), nil
}

// handleWeekly повторяет задачу в указанные дни недели
func handleWeekly(now, date time.Time, rep []string) (string, error) {
	if len(rep) < 2 {
		return "", errors.New("missing weekday value")
	}
	weekdays, err := convertStringToInt(rep[1])
	if err != nil {
		return "", err
	}

	for _, wd := range weekdays {
		if wd < 1 || wd > 7 {
			return "", errors.New("invalid weekday")
		}
	}

	adjustedWeekdays := make([]int, 0, len(weekdays))
	for _, wd := range weekdays {
		if wd == 7 {
			adjustedWeekdays = append(adjustedWeekdays, 0) // Воскресенье
		} else {
			adjustedWeekdays = append(adjustedWeekdays, wd)
		}
	}

	for {
		if afterNow(date, now) && contains(adjustedWeekdays, int(date.Weekday())) {
			break
		}
		date = date.AddDate(0, 0, 1)
	}
	return date.Format(TimeFormat), nil
}

// handleMonthly повторяет задачу в указанные дни месяца
func handleMonthly(now, date time.Time, rep []string) (string, error) {
	if len(rep) < 2 {
		return "", errors.New("missing days value")
	}

	days, err := convertStringToInt(rep[1])
	if err != nil {
		return "", err
	}

	if !checkCorrectDay(days) {
		return "", errors.New("invalid day")
	}

	var months []int
	if len(rep) >= 3 {
		months, err = convertStringToInt(rep[2])
		if err != nil {
			return "", err
		}
	}

	current := date

	for {
		currentMonth := int(current.Month())
		currentDay := current.Day()

		monthValid := len(months) == 0 || contains(months, currentMonth)

		dayValid := false
		for _, d := range days {
			if d > 0 {
				if d == currentDay && afterNow(current, now) {
					dayValid = true
					break
				}
			} else {
				lastDay := time.Date(
					current.Year(),
					current.Month()+1,
					0,
					0, 0, 0, 0, time.UTC,
				).Day()

				actualDay := lastDay + d + 1
				if actualDay == currentDay && afterNow(current, now) {
					dayValid = true
					break
				}
			}
		}

		if monthValid && dayValid {
			return current.Format(TimeFormat), nil
		}

		current = current.AddDate(0, 0, 1)
	}
}
