package expenses

import (
	"time"
)

// Diary entry itself
type Expense struct {
	Date    time.Time
	Sum     float32
	Comment string
}

func Create(date time.Time, sum float32, comment string) Expense {
	return Expense{Date: date.Truncate(time.Second), Sum: sum, Comment: comment}
}
