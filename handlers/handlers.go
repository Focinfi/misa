package handlers

import (
	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlers/generators"
	"github.com/Focinfi/misa/handlers/interpreters"
	"github.com/Focinfi/misa/handlers/io"
	"github.com/Focinfi/misa/handlers/iterators"
	"github.com/Focinfi/misa/handlers/strings"
)

type builders map[string]pipeline.HandlerBuilder

func (b builders) GetOK(id string) (pipeline.HandlerBuilder, bool) {
	builder, ok := DefaultBuilders[id]
	return builder, ok
}

var DefaultBuilders = builders{
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
	"io-clipboard-reader": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return io.DefaultClipboardReader
	}),
	"io-clipboard-writer": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return io.DefaultClipboardWriter
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
