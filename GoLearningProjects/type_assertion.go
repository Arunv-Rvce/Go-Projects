package main

import (
	"fmt"
	"time"
)

func timeMap(y interface{}) {
	fmt.Printf("%T", &y)
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_at"] = time.Now()
	}
}

type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

func (f *fakeString) String() string {
	return f.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case Stringer:
		fmt.Println(str.String())
	case string:
		fmt.Println(str)
	}
}
