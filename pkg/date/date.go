package date

import (
	"os"
	"time"
)

func TimeToString(t string) (time.Time, error) {
	if os.Getenv("ENV") == "test" {
		layout := "2006-01-02T15:04:05Z"
		parsedTime, err := time.Parse(layout, t)
		if err != nil {
			return time.Now(), err
		}
		return parsedTime, nil
	}

	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, t)
	if err != nil {
		return time.Now(), err
	}
	return parsedTime, nil
}
