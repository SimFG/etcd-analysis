package core

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var (
	byteUnits = []string{"b", "kb", "mb", "gb", "tb", "pb", "eb"}
)

func Exit(err error) {
	fmt.Fprintln(os.Stderr, "Error:", err)
	os.Exit(-1)
}

func EmptyChar() string {
	return string([]byte{0})
}

func Min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func Max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func ReadableSize(s int) string {
	sf := float64(s)
	i := 0
	for sf > 1024 {
		i++
		sf /= 1024
	}
	return fmt.Sprintf("%.1f%s", sf, byteUnits[i])
}

func Interrupt(f func()) {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c)
		s := <-c
		if s == syscall.SIGINT {
			fmt.Println("interrupt")
			f()
			Exit(errors.New("signal: interrupt"))
		}
	}()
}
