package datetime

import (
	"time"
)

func StrParseToTime(dateValue, dateLayout string) time.Time {
	dateResult, _ := time.Parse(dateLayout, dateValue)

	return dateResult
}
