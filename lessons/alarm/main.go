package main

import (
	"fmt"
	"time"
)

type data struct {
	time time.Time
}

func main() {
	d := data{}
	fmt.Print(d)

}
