package utils

import (
	"fmt"
	"time"
)

func Spinner(msg string, done chan bool) {
	chars := `-\|/`
	i := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("\r%s %c", msg, chars[i%len(chars)])
			time.Sleep(100 * time.Millisecond)
			i++
		}
	}
}
