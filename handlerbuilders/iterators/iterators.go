package iterators

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Focinfi/misa/handlerbuilders/utils"

	"github.com/Focinfi/misa/handlerbuilders/combination"

	"github.com/Focinfi/go-pipeline"
)

func NewIterator(t, interpreterName, script string) (pipeline.Handler, error) {
	switch t {
	case "select":
		return NewSelector(interpreterName, script)
	case "map":
		return NewMapper(interpreterName, script)
	case "reduce":
		return NewReducer(interpreterName, script)
	default:
		return nil, fmt.Errorf("unsupported iterator type: %v", t)
	}
}

func BuildIterator(conf map[string]interface{}) (pipeline.Handler, error) {
	t := fmt.Sprint(conf["type"])
	name := fmt.Sprint(conf["interpreter_name"])
	script := fmt.Sprint(conf["script"])
	return NewIterator(t, name, script)
}

func BuildIteratorByType(t string, conf map[string]interface{}) (pipeline.Handler, error) {
	if conf == nil {
		conf = make(map[string]interface{})
	}
	conf["type"] = t
	return BuildIterator(conf)
}

type Conf struct {
	Type            string `json:"type"`
	InterpreterName string `json:"interpreter_name"`
	Script          string `json:"script"`
}

type Iterators struct {
	Confs   []Conf                  `json:"confs"`
	handler combination.HandlerList `json:"-"`
}

func NewIterators(iteratorConfs []Conf) (*Iterators, error) {
	handlers := make([]pipeline.Handler, 0, len(iteratorConfs))
	for _, conf := range iteratorConfs {
		iterator, err := NewIterator(conf.Type, conf.InterpreterName, conf.Script)
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, iterator)
	}
	return &Iterators{
		Confs:   iteratorConfs,
		handler: combination.HandlerList{Handlers: handlers},
	}, nil
}

func BuildIterators(conf map[string]interface{}) (pipeline.Handler, error) {
	iteratorConfs := make([]Conf, 0)
	confStr, err := utils.AnyTypeToString(conf["iterators"])
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(confStr), &iteratorConfs); err != nil {
		return nil, err
	}
	return NewIterators(iteratorConfs)
}

func (it Iterators) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	return it.handler.Handle(ctx, reqRes)
}
