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
	result := &item{
		key: input,
	}
	return result
}

func main() {
	var results *item

	task.Go(
		task.New(todo).Params("one"),
		task.New(todo).Params("two"),
		task.New(todo).Params("three").Results(&results),
	)

	fmt.Println(results)
}

```
