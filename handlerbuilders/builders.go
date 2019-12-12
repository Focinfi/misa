package handlerbuilders

import (
	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/generators"
	"github.com/Focinfi/misa/handlerbuilders/interpreters"
	"github.com/Focinfi/misa/handlerbuilders/io"
	"github.com/Focinfi/misa/handlerbuilders/iterators"
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
		return io.DefaultClipboardReader
	}),
	"io-writer-clipboard": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return io.DefaultClipboardWriter
	}),
	"io-writer-stdout": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return io.DefaultStdOutWriter
	}),
	"io-writer-file": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return io.BuildFile(conf)
	}),

	"parser-unix": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return parsers.DefaultUnixParser
	}),

	// iterators
	"iterator-map": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return iterators.BuildIterator(conf)
	}),
	"iterator-reducer": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return iterators.BuildIterator(conf)
	}),
	"iterator-selector": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return iterators.BuildIterator(conf)
	}),
	"iterators": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return iterators.BuildIterators(conf)
	}),

	// string
	"string-splitter": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return strings.BuildSplitter(conf)
	}),
	"string-join": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return strings.BuildJoin(conf)
	}),
}
