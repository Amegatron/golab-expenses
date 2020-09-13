package io

import "expenses/expenses"

type DiarySaveLoad interface {
	Save(diary *expenses.Diary)
	Load() *expenses.Diary
}
