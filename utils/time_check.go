package utils

import (
	"errors"
	"log"
	"strings"
	"time"
)

func TimeCheck(timeRangs []string) bool {
	for _, timeRange := range timeRangs {
		res, err := timeCheck(timeRange)
		if err != nil {
			log.Println(err)
			return false
		}
		if res {
			return true
		}
	}
	return false
}

func timeCheck(timeRange string) (bool, error) {
	tr := strings.Split(timeRange, "-")
	if len(tr) != 2 {
		err := errors.New("Time Range Format Wrong")
		return false, err
	}
	timeStart, err := time.Parse("15:04", tr[0])
	timeEnd, err := time.Parse("15:04", tr[1])
	if err != nil {
		return false, err
	}
	timeNow := time.Now()
	hour, min, sec := timeNow.Clock()
	timeNowHour := time.Date(0, 1, 1, hour, min, sec, 0, time.UTC)
	if timeNowHour.Before(timeEnd) && timeNowHour.After(timeStart) {
		return true, nil
	}
	return false, nil

}
