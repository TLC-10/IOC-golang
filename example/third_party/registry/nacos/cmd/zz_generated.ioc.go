//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli, run 'iocli gen' to re-generate

package main

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	singleton "github.com/alibaba/ioc-golang/autowire/singleton"
	util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &app_{}
		},
	})
	appStructDescriptor := &autowire.StructDescriptor{
		Alias: "AppAlias",
		Factory: func() interface{} {
			return &App{}
		},
		ParamFactory: func() interface{} {
			var _ paramInterface = &Param{}
			return &Param{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(paramInterface)
			impl := i.(*App)
			return param.Init(impl)
		},
	}
	singleton.RegisterStructDescriptor(appStructDescriptor)
}

type paramInterface interface {
	Init(impl *App) (*App, error)
}
type app_ struct {
	Run_ func()
}

func (a *app_) Run() {
	a.Run_()
}

type AppIOCInterface interface {
	Run()
}

var _appSDID string

func GetAppSingleton(p *Param) (*App, error) {
	if _appSDID == "" {
		_appSDID = util.GetSDIDByStructPtr(new(App))
	}
	i, err := singleton.GetImpl(_appSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(*App)
	return impl, nil
}

func GetAppIOCInterfaceSingleton(p *Param) (AppIOCInterface, error) {
	if _appSDID == "" {
		_appSDID = util.GetSDIDByStructPtr(new(App))
	}
	i, err := singleton.GetImplWithProxy(_appSDID, p)
	if err != nil {
		return nil, err
	}
	impl := i.(AppIOCInterface)
	return impl, nil
}
