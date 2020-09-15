package dateutils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNowString : get utc current time in string
func GetNowString() string {
	now := time.Now().UTC()
	return now.Format(apiDateLayout)
}
