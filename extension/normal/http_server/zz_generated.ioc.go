//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package http_server

import (
	contextx "context"
	autowire "github.com/alibaba/ioc-golang/autowire"
	normal "github.com/alibaba/ioc-golang/autowire/normal"
	util "github.com/alibaba/ioc-golang/autowire/util"
	"github.com/alibaba/ioc-golang/extension/normal/http_server/ghttp"
	"github.com/urfave/negroni"
	"net/http"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &impl_{}
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &Impl{}
		},
		ParamFactory: func() interface{} {
			var _ hTTPServerConfigInterface = &HTTPServerConfig{}
			return &HTTPServerConfig{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(hTTPServerConfigInterface)
			impl := i.(*Impl)
			return param.Create(impl)
		},
	})
}

type hTTPServerConfigInterface interface {
	Create(impl *Impl) (*Impl, error)
}
type impl_ struct {
	UseMW_                            func(filters ...negroni.Handler)
	UseIOCGolangMW_                   func(filters ...ghttp.Filter)
	Run_                              func(ctx contextx.Context)
	RegisterRouterWithRawHttpHandler_ func(path string, handler func(w http.ResponseWriter, r *http.Request), method string)
	RegisterWSRouter_                 func(path string, handler func(*ghttp.GRegisterWSController))
}

func (i *impl_) UseMW(filters ...negroni.Handler) {
	i.UseMW_(filters...)
}
func (i *impl_) UseIOCGolangMW(filters ...ghttp.Filter) {
	i.UseIOCGolangMW_(filters...)
}
func (i *impl_) Run(ctx contextx.Context) {
	i.Run_(ctx)
}
func (i *impl_) RegisterRouterWithRawHttpHandler(path string, handler func(w http.ResponseWriter, r *http.Request), method string) {
	i.RegisterRouterWithRawHttpHandler_(path, handler, method)
}
func (i *impl_) RegisterWSRouter(path string, handler func(*ghttp.GRegisterWSController)) {
	i.RegisterWSRouter_(path, handler)
}
func GetImpl(p *HTTPServerConfig) (*Impl, error) {
	i, err := normal.GetImpl(util.GetSDIDByStructPtr(new(Impl)), p)
	if err != nil {
		return nil, err
	}
	impl := i.(*Impl)
	return impl, nil
}