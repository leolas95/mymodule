package greetings

import "fmt"

func Greet(name string) string {
    return fmt.Sprintf("hello %s", name)
}

func Bye() string {
    return "bye"
}

func Ask() string {
    return "???"
}

