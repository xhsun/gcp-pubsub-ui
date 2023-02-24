package main

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/config"
	"github.com/xhsun/gcp-pubsub-ui/pubsub-ui-server/internal/registry"
)

func main() {
	var config config.Config
	configPath, exist := os.LookupEnv("PUBSUB_SERVER_CONFIG_PATH")
	if !exist {
		configPath = "config/config.json"
	}
	cleanenv.ReadConfig(configPath, &config)
	cleanenv.ReadEnv(&config)

	if config.VerboseLog {
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
		log.SetFormatter(&log.TextFormatter{
			FullTimestamp: true,
		})
	} else {
		log.SetLevel(log.WarnLevel)
		log.SetFormatter(&log.JSONFormatter{})
	}

	log.WithField("config", config).Debug("Attempt to start PubSub UI server")

	// Intialize services
	appServers, err := registry.InitializeServer(&config)
	if err != nil {
		log.WithError(err).Error("Failed to initialize PubSub UI server")
		os.Exit(2)
	}

	// Start services
	err = appServers.Start()
	if err != nil {
		log.WithError(err).Error("Failed to start PubSub UI server")
		os.Exit(2)
	}
}
