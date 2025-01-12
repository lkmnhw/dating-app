package main

import (
	"os"
	"os/signal"
	"syscall"

	"dating-app/internal/servers"
)

func main() {
	s := servers.Init()               // initialize server
	exitCh := make(chan os.Signal, 1) // buffered channel to avoid missing signals
	signal.Notify(exitCh,
		syscall.SIGTERM, // terminate: stopped by `kill -TERM PID`
		syscall.SIGINT,  // interrupt: stopped by Ctrl + C
	)

	go func() {
		defer func() {
			exitCh <- syscall.SIGTERM // send terminate signal when
			// application stop naturally
		}()
		s.Run() // run server / start the application
	}()

	<-exitCh  // blocking until receive exit signal
	s.Close() // close server
}
