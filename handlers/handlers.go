package handlers

import (
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders"
)

var Handlers = pipeline.MapHandlerGetter{}

func InitHandlers(confPath string) error {
	confs, err := parseConfJSON(confPath)
	if err != nil {
		return err
	}

	handlers := pipeline.MapHandlerGetter{}
	for _, conf := range confs {
		if _, ok := handlers[conf.ID]; ok {
			return fmt.Errorf("handler with id(%s) already exists", conf.ID)
		}
		line, err := pipeline.NewLineByJSON(string(conf.Conf), handlerbuilders.DefaultBuilders, handlers)
		if err != nil {
			return fmt.Errorf("new line with conf id(%s) by json failed, err: %v", conf.ID, err)
		}
		handlers[conf.ID] = line
	}
	Handlers = handlers
	return nil
}

func AddHandler(confJSON string) error {
	// copy old handlers
	return nil
}
