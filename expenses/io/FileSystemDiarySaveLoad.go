package io

import (
	"bufio"
	"container/list"
	"expenses/expenses"
	"fmt"
	"os"
	"strconv"
	"time"
)

const DefaultDateFormat = time.RFC3339Nano

type FileSystemDiarySaveLoad struct {
	Path       string
	DateFormat string
}

func (f FileSystemDiarySaveLoad) Save(d *expenses.Diary) {
	file, err := os.Create(f.Path)
	if err != nil {
		panic(err)
	}

	dateFormat := DefaultDateFormat
	if len(f.DateFormat) > 0 {
		dateFormat = f.DateFormat
	}

	for e := d.Entries.Front(); e != nil; e = e.Next() {
		buf := fmt.Sprintln(e.Value.(expenses.Expense).Date.Format(dateFormat))
		buf += fmt.Sprintln(e.Value.(expenses.Expense).Sum)
		buf += fmt.Sprintln(e.Value.(expenses.Expense).Comment)
		if e.Next() != nil {
			buf += "\n"
		}

		_, err := file.WriteString(buf)
		if err != nil {
			panic(err)
		}
	}
	err = file.Close()
}

func (f FileSystemDiarySaveLoad) Load() *expenses.Diary {
	file, err := os.Open(f.Path)
	if err != nil {
		panic(err)
	}

	dateFormat := DefaultDateFormat
	if len(f.DateFormat) > 0 {
		dateFormat = f.DateFormat
	}

	scanner := bufio.NewScanner(file)
	entries := new(list.List)
	for scanner.Scan() {
		var date time.Time
		var sum float32
		var comment string

		date, err = time.Parse(dateFormat, scanner.Text())
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		buf, err2 := strconv.ParseFloat(scanner.Text(), 32)
		if err2 != nil {
			panic(err2)
		}
		sum = float32(buf)
		scanner.Scan()
		comment = scanner.Text()
		entry := expenses.Create(date, sum, comment)
		entries.PushBack(entry)
		scanner.Scan() // empty line
	}

	d := new(expenses.Diary)
	d.Entries = entries

	return d
}
