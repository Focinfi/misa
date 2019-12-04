package handlers

import (
	"context"
	"testing"
	"time"

	"github.com/Focinfi/misa/handlers/iterators"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlers/combination"
	"github.com/Focinfi/misa/handlers/generators"
)

func Test(t *testing.T) {
	mapperDate, _ := iterators.NewMapper("tengo", `import("times").time_format(v, "2006-01-02")`)
	mapperTemplate, _ := iterators.NewMapper("tengo", `"hello at " + v + " morning!"`)
	h := combination.HandlerList{
		Handlers: []pipeline.Handler{
			generators.TimeRange{},
			mapperDate,
			mapperTemplate,
		},
	}
	resp, err := h.Handle(context.Background(), &pipeline.HandleRes{
		Data: map[string]interface{}{
			"start": time.Date(2019, 12, 1, 0, 0, 0, 0, time.Local),
			"end":   time.Date(2019, 12, 3, 0, 0, 0, 0, time.Local),
			"step":  time.Hour * 24,
		},
	})
	t.Log(resp, err)
}
