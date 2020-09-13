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
