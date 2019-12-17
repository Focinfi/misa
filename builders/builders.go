package builders

import (
	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/builders/generators"
	"github.com/Focinfi/misa/builders/gui"
	"github.com/Focinfi/misa/builders/interpreters"
	"github.com/Focinfi/misa/builders/io"
	"github.com/Focinfi/misa/builders/iterators"
	"github.com/Focinfi/misa/builders/net"
	"github.com/Focinfi/misa/builders/nosql"
	"github.com/Focinfi/misa/builders/parsers"
	"github.com/Focinfi/misa/builders/sql"
	"github.com/Focinfi/misa/builders/strings"
	"github.com/Focinfi/misa/builders/utils"
)

type Builder interface {
	Build() (pipeline.Handler, error)
}

type BuilderMap map[string]Builder

func (m BuilderMap) GetHandlerBuilderOK(id string) (pipeline.HandlerBuilder, bool) {
	tb, ok := m[id]
	if !ok {
		return nil, false
	}
	return pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		if len(conf) > 0 {
			if err := utils.JSONUnmarshalWithMap(conf, tb); err != nil {
				return nil, err
			}
		}
		return tb.Build()
	}), true
}

type SingletonBuilder struct {
	Handler pipeline.Handler
}

func (b SingletonBuilder) Build() (pipeline.Handler, error) { return b.Handler, nil }

var Builders = BuilderMap{
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

	// sql
	"mysql": &sql.MySQL{},

	// nosql
	"redis": &nosql.Redis{},
}
