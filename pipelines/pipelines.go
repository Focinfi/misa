package pipelines

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/builders"
)

var mux sync.Mutex

type Line struct {
	Index   int              `json:"index"`
	ID      string           `json:"id"`
	Handler pipeline.Handler `json:"handler"`
}

type LineMap map[string]Line

func (m LineMap) GetHandlerOK(name string) (pipeline.Handler, bool) {
	handler, ok := m[name]
	return handler.Handler, ok
}

type lines struct {
	LineMap LineMap
	DepMap  map[string]*Dep
}

func (l lines) GetHandlerOK(id string) (pipeline.Handler, bool) {
	handler, ok := l.LineMap[id]
	return handler.Handler, ok
}

func (l lines) GetDep(id string) (*Dep, bool) {
	dep, ok := l.DepMap[id]
	return dep, ok
}

func InitLines(confPath string) (*lines, error) {
	confs, err := parseConfJSON(confPath)
	if err != nil {
		return nil, err
	}

	m := make(LineMap)
	for i, conf := range confs {
		if _, ok := m[conf.ID]; ok {
			return nil, fmt.Errorf("pipeline with id(%s) already exists", conf.ID)
		}
		line, err := pipeline.NewLineByJSON(string(conf.Conf), builders.Builders, m)
		if err != nil {
			return nil, fmt.Errorf("new pipeline with conf id(%s) by json err: %v", conf.ID, err)
		}
		m[conf.ID] = Line{
			Index:   i,
			ID:      conf.ID,
			Handler: line,
		}
	}
	depMap, err := BuildDepMap(m)
	if err != nil {
		return nil, err
	}

	return &lines{
		LineMap: m,
		DepMap:  depMap,
	}, nil
}

func (l *lines) AddByConfJSON(id, confJSON string) error {
	line, err := pipeline.NewLineByJSON(confJSON, builders.Builders, l)
	if err != nil {
		return fmt.Errorf("add pipeline with conf id by json err: %v", err)
	}

	newLineMap := make(map[string]Line, len(l.LineMap)+1)
	for id, handler := range l.LineMap {
		newLineMap[id] = handler
	}
	newLineMap[id] = Line{
		Index:   len(l.LineMap),
		ID:      id,
		Handler: line,
	}

	depMap, err := BuildDepMap(newLineMap)
	if err != nil {
		return err
	}
	mux.Lock()
	l.LineMap = newLineMap
	l.DepMap = depMap
	mux.Unlock()
	return nil
}

func (l *lines) UpdateByConfJSON(id, confJSON string) error {
	line, ok := l.LineMap[id]
	if !ok {
		return fmt.Errorf("pipeline with id(%#v) not found", id)
	}
	newLineMap := make(LineMap, len(l.LineMap)+1)
	for id, handler := range l.LineMap {
		newLineMap[id] = handler
	}
	delete(newLineMap, id)

	updated, err := pipeline.NewLineByJSON(confJSON, builders.Builders, newLineMap)
	if err != nil {
		return fmt.Errorf("add pipeline with conf id by json err: %v", err)
	}

	deps, err := FindDeps(updated, DepMap)
	if err != nil {
		return err
	}

	for _, dep := range deps {
		if ids, ok := dep.HasDepID(id); ok {
			return fmt.Errorf("circular dependency, %v", strings.Join(append([]string{id}, ids...), "->"))
		}
	}

	newLineMap[id] = Line{
		ID:      id,
		Index:   line.Index,
		Handler: updated,
	}
	depMap, err := BuildDepMap(newLineMap)
	if err != nil {
		return err
	}

	mux.Lock()
	l.LineMap = newLineMap
	l.DepMap = depMap
	mux.Unlock()
	return nil
}

func (l *lines) Delete(id string) error {
	_, ok := l.LineMap[id]
	if !ok {
		return fmt.Errorf("pipeline with id(%#v) not found", id)
	}

	dep, ok := l.GetDep(id)
	if !ok {
		return fmt.Errorf("pipeline dep with id(%#v) not found", id)
	}
	if len(dep.Deped) != 0 {
		ids := strings.Join(Deps(dep.Deped).IDs(), ",")
		return fmt.Errorf("pipeline with id(%#v) is depended on pipelines(%v)", id, ids)
	}
	newLineMap := make(LineMap, len(l.LineMap))
	for id, handler := range l.LineMap {
		newLineMap[id] = handler
	}
	delete(newLineMap, id)
	depMap, err := BuildDepMap(newLineMap)
	if err != nil {
		return err
	}

	mux.Lock()
	l.LineMap = newLineMap
	l.DepMap = depMap
	mux.Unlock()
	return nil
}
