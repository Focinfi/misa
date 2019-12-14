package handlerbuilders

import (
	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/confparam"
	"github.com/Focinfi/misa/handlerbuilders/generators"
	"github.com/Focinfi/misa/handlerbuilders/gui"
	"github.com/Focinfi/misa/handlerbuilders/interpreters"
	"github.com/Focinfi/misa/handlerbuilders/io"
	"github.com/Focinfi/misa/handlerbuilders/iterators"
	"github.com/Focinfi/misa/handlerbuilders/net"
	"github.com/Focinfi/misa/handlerbuilders/parsers"
	"github.com/Focinfi/misa/handlerbuilders/strings"
)

type Builder interface {
	Build() (pipeline.Handler, error)
	ConfParams() map[string]confparam.ConfParam
	InitByConf(conf map[string]interface{}) error
}

type TypedBuilderMap map[string]Builder

func (m TypedBuilderMap) GetHandlerBuilderOK(id string) (pipeline.HandlerBuilder, bool) {
	tb, ok := m[id]
	if !ok {
		return nil, false
	}
	return pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		if err := tb.InitByConf(conf); err != nil {
			return nil, err
		}
		return tb.Build()
	}), true
}

func (m TypedBuilderMap) GetConfParamsOK(id string) (*confparam.ConfParam, bool) {
	return nil, false
}

type SingletonBuilder struct {
	Handler pipeline.Handler
}

func (b SingletonBuilder) Build() (pipeline.Handler, error)           { return b.Handler, nil }
func (SingletonBuilder) ConfParams() map[string]confparam.ConfParam   { return nil }
func (SingletonBuilder) InitByConf(conf map[string]interface{}) error { return nil }

var Builders = TypedBuilderMap{
	// interpreters
	"interpreter": &interpreters.Conf{},
	// generators
	"generator-int-range":  SingletonBuilder{Handler: generators.DefaultIntRange},
	"generator-time-range": SingletonBuilder{Handler: generators.DefaultTimeRange},
	// io
	"io-reader-clipboard": SingletonBuilder{Handler: io.DefaultReaderClipboard},
	"io-writer-clipboard": SingletonBuilder{Handler: io.DefaultWriterClipboard},
	"io-writer-stdout":    SingletonBuilder{Handler: io.DefaultWriterStdOut},
	"io-writer-file":      &io.WriterFile{},
	// parser
	"parser-unix": SingletonBuilder{Handler: parsers.DefaultUnixParser},
	"parser-json": SingletonBuilder{Handler: parsers.DefaultParserJSON},

	// iterators
	"iterator":        &iterators.Conf{},
	"iterator-map":    &iterators.Conf{Type: "map"},
	"iterator-reduce": &iterators.Conf{Type: "reduce"},
	"iterator-select": &iterators.Conf{Type: "select"},

	// string
	"string-split":     &strings.SeparatorConf{Type: "split"},
	"string-join":      &strings.SeparatorConf{Type: "join"},
	"finder-json-attr": &strings.FinderJSONAttr{},

	// net
	"net-http": SingletonBuilder{Handler: net.DefaultHTTP},

	// gui
	"notify-desktop": &gui.DesktopNotificator{},
}
