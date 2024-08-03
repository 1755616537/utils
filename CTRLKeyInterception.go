package utils

import (
	"os"
	"os/signal"
)

// CTRL+C拦截
func CTRLKeyInterception() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	_ = <-c
}
