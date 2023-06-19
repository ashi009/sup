package main

import (
	"context"
	"os"
	"os/signal"

	dnssdlog "github.com/brutella/dnssd/log"
	"github.com/brutella/hap"
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/log"
	"golang.org/x/sys/unix"
)

func main() {
	log.Debug.Enable()
	dnssdlog.Debug.Enable()

	ab := accessory.NewBridge(accessory.Info{
		Name: "HAP Bridge",
	})

	// Store the data in the "./db" directory.
	fs := hap.NewFsStore("./db")

	// Create the hap server.
	server, err := hap.NewServer(fs, ab.A, ConnectAirConditioner().A, ConnectTV().A)
	if err != nil {
		// stop if an error happens
		log.Info.Panic(err)
	}

	// Setup a listener for interrupts and SIGTERM signals
	// to stop the server.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, unix.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-c
		// Stop delivering signals.
		signal.Stop(c)
		// Cancel the context to stop the server.
		cancel()
	}()

	// Run the server.
	server.ListenAndServe(ctx)
}
