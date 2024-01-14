package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"kurt.blackwell.id.au/xyzmas/animation"
	"kurt.blackwell.id.au/xyzmas/configuration"
)

func main() {
	config, err := configuration.Load()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading config: ", err)
		os.Exit(1)
		return
	}

	if config.Verbose {
		fmt.Printf("%+v\n", config)
	}

	stop := make(chan bool, 1)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		signal := <-interrupt
		fmt.Printf("Stopping (%s)\n", signal.String())
		stop <- true

		// Second time, force exit
		signal = <-interrupt
		fmt.Printf("Force stop (%s)\n", signal.String())
		os.Exit(1)
	}()

	err = animation.Start(config, stop)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
