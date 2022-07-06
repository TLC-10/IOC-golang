package trace

import (
	"sync"

	"github.com/petermattis/goid"

	"github.com/alibaba/ioc-golang/aop/interceptor"
	"github.com/alibaba/ioc-golang/aop/interceptor/common"
)

type goRoutineTraceInterceptor struct {
	tracingGrIDMap sync.Map // tracingGrIDMap stores goroutine-id -> goRoutineTracingContext
}

func (g *goRoutineTraceInterceptor) BeforeInvoke(_ *interceptor.InvocationContext) {
	// 1. if current goroutine is watched?
	grID := goid.Get()
	if val, ok := g.tracingGrIDMap.Load(grID); ok {
		// this goRoutine is watched, add new child node
		val.(*goRoutineTracingContext).getTrace().addChildSpan(common.CurrentCallingMethodName())
		return
	}
}

func (g *goRoutineTraceInterceptor) AfterInvoke(_ *interceptor.InvocationContext) {
	// if current goroutine is watched?
	grID := goid.Get()
	if val, ok := g.tracingGrIDMap.Load(grID); ok {
		// this goRoutine is watched, return span
		traceCtx := val.(*goRoutineTracingContext)
		traceCtx.getTrace().returnSpan()

		// calculate level
		if common.TraceLevel(traceCtx.getTrace().entranceMethod) == 0 {
			g.tracingGrIDMap.Delete(grID)
		}
	}
}

func (g *goRoutineTraceInterceptor) AddCurrentGRTracingContext(ctx *goRoutineTracingContext) {
	grID := goid.Get()
	g.tracingGrIDMap.Store(grID, ctx)
}

func (g *goRoutineTraceInterceptor) DeleteCurrentGRTracingContext() {
	grID := goid.Get()
	g.tracingGrIDMap.Delete(grID)
}

func (g *goRoutineTraceInterceptor) IsCurrentGRTracing() bool {
	grID := goid.Get()
	_, ok := g.tracingGrIDMap.Load(grID)
	return ok
}

func (g *goRoutineTraceInterceptor) GetCurrentGRTracingContext() *goRoutineTracingContext {
	grID := goid.Get()
	val, ok := g.tracingGrIDMap.Load(grID)
	if ok {
		return val.(*goRoutineTracingContext)
	}
	return nil
}

var goRoutineTraceInterceptorSingleton *goRoutineTraceInterceptor

func GetGoRoutineTraceInterceptor() *goRoutineTraceInterceptor {
	if goRoutineTraceInterceptorSingleton == nil {
		goRoutineTraceInterceptorSingleton = &goRoutineTraceInterceptor{}
	}
	return goRoutineTraceInterceptorSingleton
}