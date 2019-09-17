package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	component "github.com/andreylm/nats-component"
	"github.com/andreylm/yii2-nats-test/pkg/server"
	nats "github.com/nats-io/go-nats"
)

var (
	showHelp    bool
	showVersion bool

	natsServers string
	systemTopic string
	natsUser    string
	natsSecret  string
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: api-server [options...]\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
	}
	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.StringVar(&natsServers, "nats", nats.DefaultURL, "Network host:port to listen to")
	flag.StringVar(&systemTopic, "nats-system-topic", "_NATS_SYSTEM_TOPIC", "Main NATS topic for discover and status usage")
	flag.StringVar(&natsUser, "nats-user", "", "NATS user")
	flag.StringVar(&natsSecret, "nats-secret", "", "NATS secret")

	flag.Parse()

	switch {
	case showHelp:
		flag.Usage()
		os.Exit(0)
	case showVersion:
		fmt.Fprintf(os.Stderr, "NATS Test Service %s", server.Version)
		os.Exit(0)
	}
}

func main() {
	log.Printf("Starting NATS Service version %s", server.Version)
	component := component.NewComponent("test-service")
	component.SetupConnectionToNATS(natsServers, systemTopic)

	server := apiserver.Server{
		Component: component,
	}

	if err := server.ListenAndServe(serverListen); err != nil {
		log.Fatal(err)
	}
}
