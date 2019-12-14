package gui

import (
	"context"
	"fmt"

	"github.com/0xAX/notificator"
	"github.com/Focinfi/go-pipeline"
)

type DesktopNotificator struct {
	AppName         string `json:"app_name" validate:"required"`
	DefaultIconPath string `json:"default_icon_path" validate:"-"`
	notificator     *notificator.Notificator
}

func (n DesktopNotificator) Build() (pipeline.Handler, error) {
	return NewDesktopNotificator(n.AppName, n.DefaultIconPath), nil
}

func NewDesktopNotificator(appName string, defaultIconPath string) *DesktopNotificator {
	return &DesktopNotificator{
		AppName:         appName,
		DefaultIconPath: defaultIconPath,
		notificator: notificator.New(notificator.Options{
			AppName:     appName,
			DefaultIcon: defaultIconPath,
		}),
	}
}

func (n DesktopNotificator) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	respRes = &pipeline.HandleRes{}
	if reqRes != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		if reqRes.Data != nil {
			data := reqRes.Data.(map[string]interface{})

			title := fmt.Sprint(data["title"])
			text := fmt.Sprint(data["text"])
			iconPath := fmt.Sprint(data["icon_path"])
			urgency := fmt.Sprint(data["urgency"])
			if err := n.notificator.Push(title, text, iconPath, urgency); err != nil {
				return nil, err
			}
		}
	}
	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}
