package iterators

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Focinfi/misa/handlers/utils"

	"github.com/Focinfi/misa/handlers/combination"

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
		return nil, errors.New("unsupported iterator type")
	}
}

func BuildIterator(conf map[string]interface{}) pipeline.Handler {
	i, err := NewIterator(conf["type"].(string), conf["interpreter_name"].(string), conf["script"].(string))
	if err != nil {
		panic(err)
	}
	return i
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

func BuildIterators(conf map[string]interface{}) pipeline.Handler {
	iteratorConfs := make([]Conf, 0)
	confStr, err := utils.AnyTypeToString(conf["iterators"])
	if err != nil {
		panic(err.Error())
	}
	if err := json.Unmarshal([]byte(confStr), &iteratorConfs); err != nil {
		panic(err.Error())
	}
	h, err := NewIterators(iteratorConfs)
	if err != nil {
		panic(err.Error())
	}
	return h
}

func (it Iterators) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	return it.handler.Handle(ctx, reqRes)
}
