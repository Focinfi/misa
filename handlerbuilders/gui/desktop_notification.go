package gui

import (
	"context"
	"fmt"

	"github.com/Focinfi/misa/handlerbuilders/utils"

	"github.com/Focinfi/misa/handlerbuilders/confparam"

	"github.com/0xAX/notificator"
	"github.com/Focinfi/go-pipeline"
)

var desktopNotificatorParams = make(map[string]confparam.ConfParam)

func init() {
	params, err := confparam.GetConfParams(DesktopNotificator{})
	if err != nil {
		panic(err)
	}
	desktopNotificatorParams = params
}

type DesktopNotificator struct {
	AppName         string `json:"app_name" validate:"required"`
	DefaultIconPath string `json:"default_icon_path" validate:"-"`
	notificator     *notificator.Notificator
}

func (n DesktopNotificator) Build() (pipeline.Handler, error) {
	return NewDesktopNotificator(n.AppName, n.DefaultIconPath), nil
}

func (n DesktopNotificator) ConfParams() map[string]confparam.ConfParam {
	return desktopNotificatorParams
}

func (n *DesktopNotificator) InitByConf(conf map[string]interface{}) error {
	return utils.JSONUnmarshalWithMap(conf, n)
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

func BuildDesktopNotificator(conf map[string]interface{}) (pipeline.Handler, error) {
	appName := fmt.Sprint(conf["app_name"])
	defaultIconPath := fmt.Sprint(conf["default_icon_path"])
	return NewDesktopNotificator(appName, defaultIconPath), nil
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
