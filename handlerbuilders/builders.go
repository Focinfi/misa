package handlerbuilders

import (
	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/generators"
	"github.com/Focinfi/misa/handlerbuilders/gui"
	"github.com/Focinfi/misa/handlerbuilders/interpreters"
	"github.com/Focinfi/misa/handlerbuilders/io"
	"github.com/Focinfi/misa/handlerbuilders/iterators"
	"github.com/Focinfi/misa/handlerbuilders/net"
	"github.com/Focinfi/misa/handlerbuilders/parsers"
	"github.com/Focinfi/misa/handlerbuilders/strings"
)

var DefaultBuilders = pipeline.MapHandlerBuilderGetter{
	// generators
	"generator-int-range": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return generators.DefaultIntRange
	}),
	"generator-time-range": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return generators.DefaultTimeRange
	}),

	// interpreters
	"interpreter-tengo": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return &interpreters.Tengo{
			Meta: interpreters.Meta{
				Script:     conf["script"].(string),
				InitVarMap: conf["init_var_map"].(map[string]interface{}),
				RtVarName:  conf["rt_var_name"].(string),
			},
		}
	}),

	// io
	"io-reader-clipboard": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return io.DefaultReaderClipboard
	}),
	"io-writer-clipboard": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return io.DefaultWriterClipboard
	}),
	"io-writer-stdout": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return io.DefaultWriterStdOut
	}),
	"io-writer-file": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return io.BuildFile(conf)
	}),

	// parser
	"parser-unix": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return parsers.DefaultUnixParser
	}),
	"parser-json": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return parsers.DefaultParserJSON
	}),

	// iterators
	"iterator": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return iterators.BuildIterator(conf)
	}),
	"iterator-map": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return iterators.BuildIteratorByType("map", conf)
	}),
	"iterator-reduce": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return iterators.BuildIteratorByType("reduce", conf)
	}),
	"iterator-select": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return iterators.BuildIteratorByType("select", conf)
	}),
	"iterators": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return iterators.BuildIterators(conf)
	}),

	// string
	"string-split": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return strings.BuildSplitter(conf)
	}),
	"string-join": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return strings.BuildJoin(conf)
	}),
	"finder-json-attr": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return strings.BuildFinderJSONAttr(conf)
	}),

	// net
	"net-http": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return net.DefaultHTTP
	}),

	// gui
	"notify-desktop": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return gui.BuildDesktopNotificator(conf)
	}),
}
