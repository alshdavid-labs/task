# Task

Inspired by JavaScript's `Promise.all`


```Go
package main

import (
	"github.com/qkgo/task"
	"fmt"
	"time"
)

type item struct {
	key        string
	value      string
	otherValue string
}

func todo(input string) *item {
	time.Sleep(1 * time.Second)
	return &item{
		key: input,
	}
}

func main() {
	var results *item

	err := task.Async(
		task.New(todo).Params("one"),
		task.New(todo).Params("two"),
		task.New(todo).Params("three").Results(&results),
	)
	if err != nil {
		return
	}

	fmt.Println(results)
}

```
