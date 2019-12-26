package os

import (
	"context"
	"os/exec"

	"github.com/Focinfi/go-pipeline"
)

type CmdExec struct {
	Name string    `json:"name" validate:"required"`
	Args []string  `json:"args" validate:"-"`
	cmd  *exec.Cmd `json:"-"`
}

func (e CmdExec) Build() (pipeline.Handler, error) {
	return &CmdExec{
		Name: e.Name,
		Args: e.Args,
		cmd:  exec.Command(e.Name, e.Args...),
	}, nil
}

func (e *CmdExec) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	if reqRes != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}
	}
	respRes = &pipeline.HandleRes{}
	if err := e.cmd.Run(); err != nil {
		return nil, err
	}
	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}
