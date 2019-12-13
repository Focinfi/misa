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
	"generator-int-range": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return generators.DefaultIntRange, nil
	}),
	"generator-time-range": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return generators.DefaultTimeRange, nil
	}),

	// interpreters
	"interpreter-tengo": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return &interpreters.Tengo{
			Meta: interpreters.Meta{
				Script:     conf["script"].(string),
				InitVarMap: conf["init_var_map"].(map[string]interface{}),
				RtVarName:  conf["rt_var_name"].(string),
			},
		}, nil
	}),

	// io
	"io-reader-clipboard": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return io.DefaultReaderClipboard, nil
	}),
	"io-writer-clipboard": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return io.DefaultWriterClipboard, nil
	}),
	"io-writer-stdout": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return io.DefaultWriterStdOut, nil
	}),
	"io-writer-file": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return io.BuildFile(conf)
	}),

	// parser
	"parser-unix": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return parsers.DefaultUnixParser, nil
	}),
	"parser-json": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return parsers.DefaultParserJSON, nil
	}),

	// iterators
	"iterator": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return iterators.BuildIterator(conf)
	}),
	"iterator-map": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return iterators.BuildIteratorByType("map", conf)
	}),
	"iterator-reduce": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return iterators.BuildIteratorByType("reduce", conf)
	}),
	"iterator-select": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return iterators.BuildIteratorByType("select", conf)
	}),
	"iterators": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return iterators.BuildIterators(conf)
	}),

	// string
	"string-split": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return strings.BuildSplitter(conf)
	}),
	"string-join": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return strings.BuildJoin(conf)
	}),
	"finder-json-attr": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return strings.BuildFinderJSONAttr(conf)
	}),

	// net
	"net-http": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return net.DefaultHTTP, nil
	}),

	// gui
	"notify-desktop": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return gui.BuildDesktopNotificator(conf)
	}),
}
