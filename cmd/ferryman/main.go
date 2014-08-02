package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	. "github.com/marconi/ferryman"
	"github.com/marconi/rivers"
)

func main() {
	// parse args
	urgentName := flag.String("urgent", "styx", "a string")
	delayedName := flag.String("delayed", "acheron", "a string")
	flag.Parse()

	// create the queues
	log.Println("creating queues...")
	urgent := rivers.NewQueue(*urgentName, "urgent")
	delayed := rivers.NewQueue(*delayedName, "delayed")

	// start rowing!
	log.Println("running Ferryman...")
	quit := Run(urgent, delayed)

	// listen for kill and ^C
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// stop Run goroutine
	close(quit)

	// destroy queues
	urgent.Destroy()
	delayed.Destroy()

	// wait for the quit log
	time.Sleep(1 * time.Second)
}
