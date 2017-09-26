package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ardanlabs/kit/cfg"
	"github.com/ardanlabs/kit/tcp"
)

//configuration setting
const (
	configKey = "CHAT"
)

var evtTypes = []string{
	"unknown",
	"Accept",
	"Join",
	"Read",
	"Remove",
	"Drop",
}

//set of event sub types
var typTypes = []string{
	"unknown",
	"Error",
	"Info",
	"Trigger",
}

//Event writes tcp events
func Event(evt, typ int, ipAddress string, format string, a ...interface{}) {
	log.Printf("*****> Event : IP[ %s ] : EVT [ %s ] : TYP [ %s ]  %s", ipAddress, evtTypes[evt], typTypes[typ], fmt.Sprintf(format, a...))
}
func init() {
	os.Setenv("CHAT_HOST", ":6000")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds)
}

func main() {

	// init the configuration system
	if err := cfg.Init(cfg.EnvProvider{Namespace: configKey}); err != nil {
		log.Println("Error initializing configuration system", err)
		os.Exit(1)
	}

	log.Println("Configuration\nl", cfg.Log())

	// get configuration
	host := cfg.MustString("HOST")

	cfg := tcp.Config{
		NetType: "tcp4",
		Addr:    host,

		ConnHandler: connHandler{},
		ReqHandler:  reqHandler{},
		RespHandler: respHandler{},

		OptEvnet: tcp.OptEvnet{
			Event: Event,
		},
	}

	//Create a new tcp value
	t, err := tcp.New("Sample", cfg)
	if err != nil {
		log.Printf("main : %s", err)
		return
	}

	//start accepting client data
	if err := t.Start(); err != nil {
		log.Printf("main : %s", err)
		return

	}

	defer t.Stop()

	log.Printf("main : Waiting for data on : %s", t.Addr())

	// Listen for interrupt from the OS

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan

}
