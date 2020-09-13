package io

import (
	"container/list"
	"expenses/expenses"
	"math/rand"
	"testing"
	"time"
)

func TestConsistentSaveLoad(t *testing.T) {
	d := getSampleDiary()
	saver := new(FileSystemDiarySaveLoad)
	saver.Path = "./test.diary"
	saver.Save(d)

	loader := new(FileSystemDiarySaveLoad)
	loader.Path = "./test.diary"
	d2 := loader.Load()

	var e, e2 *list.Element
	var i int

	for e, e2, i = d.Entries.Front(), d2.Entries.Front(), 0; e != nil && e2 != nil; e, e2, i = e.Next(), e2.Next(), i+1 {
		_e := e.Value.(expenses.Expense)
		_e2 := e2.Value.(expenses.Expense)

		if _e.Date.Truncate(time.Second) != _e2.Date.Truncate(time.Second) {
			t.Errorf("Data mismatch for entry %d for the 'Date' field: expected %s, got %s", i, _e.Date.String(), _e2.Date.String())
		}
		if _e.Sum != _e2.Sum {
			t.Errorf("Data mismatch for entry %d for the 'Sum' field: expected %f, got %f", i, _e.Sum, _e2.Sum)
		}
		if _e.Comment != _e2.Comment {
			t.Errorf("Data mismatch for entry %d for the 'Comment' field: expected '%s', got '%s'", i, _e.Comment, _e2.Comment)
		}
	}

	if e == nil && e2 != nil {
		t.Error("Loaded diary is longer than initial")
	} else if e != nil && e2 == nil {
		t.Error("Loaded diary is shorter than initial")
	}
}

func getSampleDiary() *expenses.Diary {
	testList := new(list.List)

	var expense expenses.Expense

	expense = expenses.Expense{
		Date:    time.Now(),
		Sum:     rand.Float32() * 100,
		Comment: "First expense",
	}
	testList.PushBack(expense)

	expense = expenses.Expense{
		Date:    time.Now(),
		Sum:     rand.Float32() * 50,
		Comment: "Second expense",
	}
	testList.PushBack(expense)

	expense = expenses.Expense{
		Date:    time.Now(),
		Sum:     rand.Float32() * 300,
		Comment: "Third expense",
	}
	testList.PushBack(expense)

	d := new(expenses.Diary)
	d.Entries = testList

	return d
}
