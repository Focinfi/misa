package pipelines

import (
	"fmt"
	"sort"

	"github.com/Focinfi/go-pipeline"
)

var DepMap = make(map[string]*Dep)

type Dep struct {
	ID    string
	Deps  []*Dep
	Deped []*Dep
}

type Deps []*Dep

func (deps Deps) IDs() []string {
	ids := make([]string, 0, len(deps))
	for _, dep := range deps {
		if dep == nil {
			continue
		}
		ids = append(ids, dep.ID)
	}
	return ids
}

func (dep Dep) HasDepID(id string) ([]string, bool) {
	for _, d := range dep.Deps {
		if d != nil {
			if ids, ok := d.hasDepID([]string{dep.ID}, id); ok {
				return ids, ok
			}
		}
	}
	return nil, false
}

func (dep Dep) hasDepID(parents []string, id string) ([]string, bool) {
	if dep.ID == id {
		return append(parents, dep.ID), true
	}
	for _, d := range dep.Deps {
		if d == nil {
			continue
		}
		if ids, ok := d.hasDepID(append(parents, dep.ID), id); ok {
			return ids, ok
		}
	}
	return nil, false
}

func FindDeps(handler pipeline.Handler, depMap map[string]*Dep) ([]*Dep, error) {
	line := handler.(*pipeline.Line)
	deps := make([]*Dep, 0)
	for _, pipe := range line.Pipes {
		if pipe.Type == pipeline.PipeTypeSingle {
			if pipe.Conf.RefHandlerID == "" {
				continue
			}
			refHandlerID := pipe.Conf.RefHandlerID
			refHandler, ok := depMap[refHandlerID]
			if !ok {
				return nil, fmt.Errorf("ref handler with id(%#v) not found", refHandlerID)
			}
			deps = append(deps, refHandler)
			continue
		}

		for _, pipe := range pipe.Handler.(*pipeline.Parallel).Pipes {
			if pipe.Conf.RefHandlerID == "" {
				continue
			}
			refHandlerID := pipe.Conf.RefHandlerID
			refHandler := depMap[refHandlerID]
			deps = append(deps, refHandler)
		}
	}
	return deps, nil
}

func BuildDepMap(m map[string]Line) (map[string]*Dep, error) {
	depMap := make(map[string]*Dep, len(m))
	pipelines := make([]Line, 0, len(m))
	for _, line := range m {
		pipelines = append(pipelines, line)
	}
	sort.Slice(pipelines, func(i, j int) bool {
		return pipelines[i].Index < pipelines[j].Index
	})

	for _, l := range pipelines {
		line := l.Handler.(*pipeline.Line)
		dep := &Dep{
			ID: l.ID,
		}
		for _, pipe := range line.Pipes {
			if pipe.Conf.RefHandlerID == "" {
				continue
			}

			if pipe.Type == pipeline.PipeTypeSingle {
				refHandlerID := pipe.Conf.RefHandlerID
				refHandler, ok := depMap[refHandlerID]
				if !ok {
					return nil, fmt.Errorf("ref handler with id(%#v) of handler(%#v) not found", refHandlerID, l.ID)
				}
				dep.Deps = append(dep.Deps, refHandler)
				refHandler.Deped = append(refHandler.Deped, dep)
				continue
			}

			for _, pipe := range pipe.Handler.(*pipeline.Parallel).Pipes {
				if pipe.Conf.RefHandlerID == "" {
					continue
				}
				refHandlerID := pipe.Conf.RefHandlerID
				refHandler := depMap[refHandlerID]
				dep.Deps = append(dep.Deps, refHandler)
				refHandler.Deped = append(refHandler.Deped, dep)
			}
		}
		depMap[l.ID] = dep
	}
	return depMap, nil
}
