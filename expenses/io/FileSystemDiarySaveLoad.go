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

type FileSystemDiarySaveLoad struct {
	Path string
}

func (f FileSystemDiarySaveLoad) Save(d *expenses.Diary) {
	file, err := os.Create(f.Path)
	if err != nil {
		panic(err)
	}

	for e := d.Entries.Front(); e != nil; e = e.Next() {
		buf := fmt.Sprintln(e.Value.(expenses.Expense).Date.Format(time.RFC1123))
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

	scanner := bufio.NewScanner(file)
	entries := new(list.List)
	var entry *expenses.Expense
	for scanner.Scan() {
		entry = new(expenses.Expense)
		entry.Date, err = time.Parse(time.RFC1123, scanner.Text())
		if err != nil {
			panic(err)
		}
		scanner.Scan()
		buf, err2 := strconv.ParseFloat(scanner.Text(), 32)
		if err2 != nil {
			panic(err2)
		}
		entry.Sum = float32(buf)
		scanner.Scan()
		entry.Comment = scanner.Text()
		entries.PushBack(*entry)
		entry = nil
		scanner.Scan() // empty line
	}

	d := new(expenses.Diary)
	d.Entries = entries

	return d
}
