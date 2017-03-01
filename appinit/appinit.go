package appinit

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"

	"go.uber.org/fx/dig"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/transport/http"
	"go.uber.org/yarpc/transport/tchannel"
	"go.uber.org/yarpc/x/config"
	"go.uber.org/zap"
)

// New creates an app framework service
func New(config string) *Service {
	return &Service{config: config, container: dig.New()}
}

// Service is the service being bootstrapped
type Service struct {
	config    string
	container dig.Graph
}

// Procedures is a wrapper for []transport.Procedure
// since the container cant resolve lists
type Procedures struct {
	Register []transport.Procedure
}

// Provide adds a userland type to the container
func (s *Service) Provide(t interface{}) {
	s.container.Register(t)
}

// Start and starts the messaging framework
func (s *Service) Start() {
	confData := parseConfData(s.config)

	// delegate config keys to all participating components
	//for module, moduleConfig := range {
	// get configurator for $module
	// }

	dispatcher := newDispatcher(confData["yarpc"])
	s.container.Register(dispatcher)

	// register framework types
	logger := newLogger()
	s.container.Register(logger)

	// resolve and register procs
	// note we have to use an internal type here,
	// which we wouldnt have to if there was named deps support
	var procedures *Procedures
	s.container.ResolveAll(&procedures)
	if procedures != nil {
		logger.Info("Registering procedures.", zap.Any("procedures", procedures))
		dispatcher.Register(procedures.Register)
	} else {
		logger.Fatal("found no procs, exiting.")
	}

	if err := dispatcher.Start(); err != nil {
		log.Fatal(err)
	}
}

// Stop stops the service
func (s *Service) Stop() {
	var d *yarpc.Dispatcher
	s.container.ResolveAll(&d)
	d.Stop()
}

func parseConfData(confPath string) map[string]interface{} {
	confFile, err := os.Open(confPath)
	if err != nil {
		log.Fatal(err)
	}
	defer confFile.Close()

	confData, err := ioutil.ReadAll(confFile)
	if err != nil {
		log.Fatal(err)
	}

	var data map[string]interface{}
	if err := yaml.Unmarshal(confData, &data); err != nil {
		log.Fatal(err)
	}

	return data
}

func newLogger() *zap.Logger {
	logger, _ := zap.NewProduction()
	return logger
}

func newDispatcher(data interface{}) *yarpc.Dispatcher {
	cfg := config.New()
	if err := http.RegisterTransport(cfg); err != nil {
		log.Fatal(err)
	}
	if err := tchannel.RegisterTransport(cfg); err != nil {
		log.Fatal(err)
	}

	// yarpc:

	builder, err := cfg.Load(data)
	if err != nil {
		log.Fatal(err)
	}

	dispatcher, err := builder.BuildDispatcher()
	if err != nil {
		log.Fatal(err)
	}
	return dispatcher
}
