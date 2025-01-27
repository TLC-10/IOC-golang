/*
 * Copyright (c) 2022, Alibaba Group;
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package trace

import (
	"bytes"
	"sort"

	"github.com/fatih/color"
	"github.com/jaegertracing/jaeger/model"

	"github.com/alibaba/ioc-golang/aop/common"
	tracePB "github.com/alibaba/ioc-golang/extension/aop/trace/api/ioc_golang/aop/trace"
)

type traceServiceImpl struct {
	tracePB.UnimplementedTraceServiceServer
	traceInterceptor *methodTraceInterceptor
}

func (d *traceServiceImpl) Trace(req *tracePB.TraceRequest, traceServer tracePB.TraceService_TraceServer) error {
	color.Red("[Debug Server] Receive trace request %+v\n", req.String())
	defer color.Red("[Debug Server] Trace request %s finished \n", req.String())
	sdid := req.GetSdid()
	method := req.GetMethod()
	var fieldMatcher *common.FieldMatcher
	for _, matcher := range req.GetMatchers() {
		// todo multi match support
		fieldMatcher = &common.FieldMatcher{
			FieldIndex: int(matcher.Index),
			MatchRule:  matcher.GetMatchPath() + "=" + matcher.GetMatchValue(),
		}
	}

	traceCtx := newTraceByMethodContext(sdid, method, fieldMatcher)
	d.traceInterceptor.StartTraceByMethod(traceCtx)

	done := traceServer.Context().Done()
	if err := traceServer.Send(&tracePB.TraceResponse{
		CollectorAddress: getCollectorAddress(),
	}); err != nil {
		return err
	}

	if req.GetPushToCollectorAddress() != "" {
		// start subscribing batch buffer
		outBatchBuffer := make(chan *bytes.Buffer, 100)
		getGlobalTracer().subscribeBatchBuffer(outBatchBuffer)
		go func() {
			for {
				select {
				case <-done:
					getGlobalTracer().removeSubscribingBatchBuffer()
					return
				case batchBuffer := <-outBatchBuffer:
					_ = traceServer.Send(&tracePB.TraceResponse{
						ThriftSerializedSpans: batchBuffer.Bytes(),
					})
				}
			}
		}()
	}

	outTraceCh := make(chan []*model.Trace, 100)
	// start subscribing traces info
	getGlobalTracer().subscribeTrace(outTraceCh)
	go func() {
		for {
			select {
			case <-done:
				getGlobalTracer().removeSubscribingTrace()
				return
			case traces := <-outTraceCh:
				sortableTraces := traceSorter(traces)
				sort.Sort(sortableTraces)
				_ = traceServer.Send(&tracePB.TraceResponse{
					Traces: sortableTraces,
				})
			}
		}
	}()
	<-done
	d.traceInterceptor.StopTraceByMethod()
	return nil
}

func newTraceGRPCService() *traceServiceImpl {
	return &traceServiceImpl{
		traceInterceptor: getTraceInterceptorSingleton(),
	}
}
