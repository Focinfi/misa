package pipelines

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders"
)

var mux sync.Mutex

type Pipeline struct {
	Index            int    `json:"index"`
	ID               string `json:"id"`
	pipeline.Handler `json:"handler"`
}

type pipelineMap map[string]Pipeline

func (m pipelineMap) GetHandlerOK(name string) (pipeline.Handler, bool) {
	handler, ok := m[name]
	return handler, ok
}

var PipelineMap = &pipelineMap{}

func InitHandlers(confPath string) error {
	confs, err := parseConfJSON(confPath)
	if err != nil {
		return err
	}

	pipelines := make(pipelineMap)
	for i, conf := range confs {
		if _, ok := pipelines[conf.ID]; ok {
			return fmt.Errorf("handler with id(%s) already exists", conf.ID)
		}
		line, err := pipeline.NewLineByJSON(string(conf.Conf), handlerbuilders.Builders, pipelines)
		if err != nil {
			return fmt.Errorf("new pipeline with conf id(%s) by json err: %v", conf.ID, err)
		}
		pipelines[conf.ID] = Pipeline{
			Index:   i,
			ID:      conf.ID,
			Handler: line,
		}
	}
	DepMap, err = BuildDepMap(pipelines)
	if err != nil {
		return err
	}
	PipelineMap = &pipelines
	return nil
}

func (m *pipelineMap) AddPipeline(id, confJSON string) error {
	line, err := pipeline.NewLineByJSON(confJSON, handlerbuilders.Builders, m)
	if err != nil {
		return fmt.Errorf("add pipeline with conf id by json err: %v", err)
	}

	newPipelineMap := make(pipelineMap, len(*m)+1)
	for id, handler := range *m {
		newPipelineMap[id] = handler
	}
	newPipelineMap[id] = Pipeline{
		Index:   len(*m),
		ID:      id,
		Handler: line,
	}

	DepMap, err = BuildDepMap(newPipelineMap)
	if err != nil {
		return err
	}
	mux.Lock()
	*m = newPipelineMap
	mux.Unlock()
	return nil
}

func (m *pipelineMap) UpdatePipeline(id, confJSON string) error {
	line, ok := (*m)[id]
	if !ok {
		return fmt.Errorf("pipeline with id(%#v) not found", id)
	}

	updated, err := pipeline.NewLineByJSON(confJSON, handlerbuilders.Builders, m)
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

	mux.Lock()
	(*m)[id] = Pipeline{
		ID:      id,
		Index:   line.Index,
		Handler: updated,
	}
	mux.Unlock()
	return nil
}
