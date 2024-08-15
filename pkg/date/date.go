package date

import (
	"time"
)

func TimeToString(t string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, t)
	if err != nil {
		return time.Now(), err
	}
	return parsedTime, nil
}
