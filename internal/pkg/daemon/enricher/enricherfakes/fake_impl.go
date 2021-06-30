/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package enricherfakes

import (
	"os"
	"sync"

	"github.com/hpcloud/tail"
	"google.golang.org/grpc"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	api "sigs.k8s.io/security-profiles-operator/api/server"
)

type FakeImpl struct {
	CloseStub        func(*grpc.ClientConn) error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
		arg1 *grpc.ClientConn
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	DialStub        func() (*grpc.ClientConn, error)
	dialMutex       sync.RWMutex
	dialArgsForCall []struct {
	}
	dialReturns struct {
		result1 *grpc.ClientConn
		result2 error
	}
	dialReturnsOnCall map[int]struct {
		result1 *grpc.ClientConn
		result2 error
	}
	GetenvStub        func(string) string
	getenvMutex       sync.RWMutex
	getenvArgsForCall []struct {
		arg1 string
	}
	getenvReturns struct {
		result1 string
	}
	getenvReturnsOnCall map[int]struct {
		result1 string
	}
	InClusterConfigStub        func() (*rest.Config, error)
	inClusterConfigMutex       sync.RWMutex
	inClusterConfigArgsForCall []struct {
	}
	inClusterConfigReturns struct {
		result1 *rest.Config
		result2 error
	}
	inClusterConfigReturnsOnCall map[int]struct {
		result1 *rest.Config
		result2 error
	}
	LinesStub        func(*tail.Tail) chan *tail.Line
	linesMutex       sync.RWMutex
	linesArgsForCall []struct {
		arg1 *tail.Tail
	}
	linesReturns struct {
		result1 chan *tail.Line
	}
	linesReturnsOnCall map[int]struct {
		result1 chan *tail.Line
	}
	ListPodsStub        func(*kubernetes.Clientset, string) (*v1.PodList, error)
	listPodsMutex       sync.RWMutex
	listPodsArgsForCall []struct {
		arg1 *kubernetes.Clientset
		arg2 string
	}
	listPodsReturns struct {
		result1 *v1.PodList
		result2 error
	}
	listPodsReturnsOnCall map[int]struct {
		result1 *v1.PodList
		result2 error
	}
	MetricsAuditIncStub        func(api.SecurityProfilesOperatorClient, *api.MetricsAuditRequest) (*api.EmptyResponse, error)
	metricsAuditIncMutex       sync.RWMutex
	metricsAuditIncArgsForCall []struct {
		arg1 api.SecurityProfilesOperatorClient
		arg2 *api.MetricsAuditRequest
	}
	metricsAuditIncReturns struct {
		result1 *api.EmptyResponse
		result2 error
	}
	metricsAuditIncReturnsOnCall map[int]struct {
		result1 *api.EmptyResponse
		result2 error
	}
	NewForConfigStub        func(*rest.Config) (*kubernetes.Clientset, error)
	newForConfigMutex       sync.RWMutex
	newForConfigArgsForCall []struct {
		arg1 *rest.Config
	}
	newForConfigReturns struct {
		result1 *kubernetes.Clientset
		result2 error
	}
	newForConfigReturnsOnCall map[int]struct {
		result1 *kubernetes.Clientset
		result2 error
	}
	OpenStub        func(string) (*os.File, error)
	openMutex       sync.RWMutex
	openArgsForCall []struct {
		arg1 string
	}
	openReturns struct {
		result1 *os.File
		result2 error
	}
	openReturnsOnCall map[int]struct {
		result1 *os.File
		result2 error
	}
	RecordSyscallStub        func(api.SecurityProfilesOperatorClient, *api.RecordSyscallRequest) (*api.EmptyResponse, error)
	recordSyscallMutex       sync.RWMutex
	recordSyscallArgsForCall []struct {
		arg1 api.SecurityProfilesOperatorClient
		arg2 *api.RecordSyscallRequest
	}
	recordSyscallReturns struct {
		result1 *api.EmptyResponse
		result2 error
	}
	recordSyscallReturnsOnCall map[int]struct {
		result1 *api.EmptyResponse
		result2 error
	}
	TailFileStub        func(string, tail.Config) (*tail.Tail, error)
	tailFileMutex       sync.RWMutex
	tailFileArgsForCall []struct {
		arg1 string
		arg2 tail.Config
	}
	tailFileReturns struct {
		result1 *tail.Tail
		result2 error
	}
	tailFileReturnsOnCall map[int]struct {
		result1 *tail.Tail
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImpl) Close(arg1 *grpc.ClientConn) error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
		arg1 *grpc.ClientConn
	}{arg1})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{arg1})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeImpl) CloseCalls(stub func(*grpc.ClientConn) error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeImpl) CloseArgsForCall(i int) *grpc.ClientConn {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	argsForCall := fake.closeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) Dial() (*grpc.ClientConn, error) {
	fake.dialMutex.Lock()
	ret, specificReturn := fake.dialReturnsOnCall[len(fake.dialArgsForCall)]
	fake.dialArgsForCall = append(fake.dialArgsForCall, struct {
	}{})
	stub := fake.DialStub
	fakeReturns := fake.dialReturns
	fake.recordInvocation("Dial", []interface{}{})
	fake.dialMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) DialCallCount() int {
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	return len(fake.dialArgsForCall)
}

func (fake *FakeImpl) DialCalls(stub func() (*grpc.ClientConn, error)) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = stub
}

func (fake *FakeImpl) DialReturns(result1 *grpc.ClientConn, result2 error) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = nil
	fake.dialReturns = struct {
		result1 *grpc.ClientConn
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) DialReturnsOnCall(i int, result1 *grpc.ClientConn, result2 error) {
	fake.dialMutex.Lock()
	defer fake.dialMutex.Unlock()
	fake.DialStub = nil
	if fake.dialReturnsOnCall == nil {
		fake.dialReturnsOnCall = make(map[int]struct {
			result1 *grpc.ClientConn
			result2 error
		})
	}
	fake.dialReturnsOnCall[i] = struct {
		result1 *grpc.ClientConn
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Getenv(arg1 string) string {
	fake.getenvMutex.Lock()
	ret, specificReturn := fake.getenvReturnsOnCall[len(fake.getenvArgsForCall)]
	fake.getenvArgsForCall = append(fake.getenvArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetenvStub
	fakeReturns := fake.getenvReturns
	fake.recordInvocation("Getenv", []interface{}{arg1})
	fake.getenvMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) GetenvCallCount() int {
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	return len(fake.getenvArgsForCall)
}

func (fake *FakeImpl) GetenvCalls(stub func(string) string) {
	fake.getenvMutex.Lock()
	defer fake.getenvMutex.Unlock()
	fake.GetenvStub = stub
}

func (fake *FakeImpl) GetenvArgsForCall(i int) string {
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	argsForCall := fake.getenvArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) GetenvReturns(result1 string) {
	fake.getenvMutex.Lock()
	defer fake.getenvMutex.Unlock()
	fake.GetenvStub = nil
	fake.getenvReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeImpl) GetenvReturnsOnCall(i int, result1 string) {
	fake.getenvMutex.Lock()
	defer fake.getenvMutex.Unlock()
	fake.GetenvStub = nil
	if fake.getenvReturnsOnCall == nil {
		fake.getenvReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getenvReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeImpl) InClusterConfig() (*rest.Config, error) {
	fake.inClusterConfigMutex.Lock()
	ret, specificReturn := fake.inClusterConfigReturnsOnCall[len(fake.inClusterConfigArgsForCall)]
	fake.inClusterConfigArgsForCall = append(fake.inClusterConfigArgsForCall, struct {
	}{})
	stub := fake.InClusterConfigStub
	fakeReturns := fake.inClusterConfigReturns
	fake.recordInvocation("InClusterConfig", []interface{}{})
	fake.inClusterConfigMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) InClusterConfigCallCount() int {
	fake.inClusterConfigMutex.RLock()
	defer fake.inClusterConfigMutex.RUnlock()
	return len(fake.inClusterConfigArgsForCall)
}

func (fake *FakeImpl) InClusterConfigCalls(stub func() (*rest.Config, error)) {
	fake.inClusterConfigMutex.Lock()
	defer fake.inClusterConfigMutex.Unlock()
	fake.InClusterConfigStub = stub
}

func (fake *FakeImpl) InClusterConfigReturns(result1 *rest.Config, result2 error) {
	fake.inClusterConfigMutex.Lock()
	defer fake.inClusterConfigMutex.Unlock()
	fake.InClusterConfigStub = nil
	fake.inClusterConfigReturns = struct {
		result1 *rest.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) InClusterConfigReturnsOnCall(i int, result1 *rest.Config, result2 error) {
	fake.inClusterConfigMutex.Lock()
	defer fake.inClusterConfigMutex.Unlock()
	fake.InClusterConfigStub = nil
	if fake.inClusterConfigReturnsOnCall == nil {
		fake.inClusterConfigReturnsOnCall = make(map[int]struct {
			result1 *rest.Config
			result2 error
		})
	}
	fake.inClusterConfigReturnsOnCall[i] = struct {
		result1 *rest.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Lines(arg1 *tail.Tail) chan *tail.Line {
	fake.linesMutex.Lock()
	ret, specificReturn := fake.linesReturnsOnCall[len(fake.linesArgsForCall)]
	fake.linesArgsForCall = append(fake.linesArgsForCall, struct {
		arg1 *tail.Tail
	}{arg1})
	stub := fake.LinesStub
	fakeReturns := fake.linesReturns
	fake.recordInvocation("Lines", []interface{}{arg1})
	fake.linesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) LinesCallCount() int {
	fake.linesMutex.RLock()
	defer fake.linesMutex.RUnlock()
	return len(fake.linesArgsForCall)
}

func (fake *FakeImpl) LinesCalls(stub func(*tail.Tail) chan *tail.Line) {
	fake.linesMutex.Lock()
	defer fake.linesMutex.Unlock()
	fake.LinesStub = stub
}

func (fake *FakeImpl) LinesArgsForCall(i int) *tail.Tail {
	fake.linesMutex.RLock()
	defer fake.linesMutex.RUnlock()
	argsForCall := fake.linesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) LinesReturns(result1 chan *tail.Line) {
	fake.linesMutex.Lock()
	defer fake.linesMutex.Unlock()
	fake.LinesStub = nil
	fake.linesReturns = struct {
		result1 chan *tail.Line
	}{result1}
}

func (fake *FakeImpl) LinesReturnsOnCall(i int, result1 chan *tail.Line) {
	fake.linesMutex.Lock()
	defer fake.linesMutex.Unlock()
	fake.LinesStub = nil
	if fake.linesReturnsOnCall == nil {
		fake.linesReturnsOnCall = make(map[int]struct {
			result1 chan *tail.Line
		})
	}
	fake.linesReturnsOnCall[i] = struct {
		result1 chan *tail.Line
	}{result1}
}

func (fake *FakeImpl) ListPods(arg1 *kubernetes.Clientset, arg2 string) (*v1.PodList, error) {
	fake.listPodsMutex.Lock()
	ret, specificReturn := fake.listPodsReturnsOnCall[len(fake.listPodsArgsForCall)]
	fake.listPodsArgsForCall = append(fake.listPodsArgsForCall, struct {
		arg1 *kubernetes.Clientset
		arg2 string
	}{arg1, arg2})
	stub := fake.ListPodsStub
	fakeReturns := fake.listPodsReturns
	fake.recordInvocation("ListPods", []interface{}{arg1, arg2})
	fake.listPodsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) ListPodsCallCount() int {
	fake.listPodsMutex.RLock()
	defer fake.listPodsMutex.RUnlock()
	return len(fake.listPodsArgsForCall)
}

func (fake *FakeImpl) ListPodsCalls(stub func(*kubernetes.Clientset, string) (*v1.PodList, error)) {
	fake.listPodsMutex.Lock()
	defer fake.listPodsMutex.Unlock()
	fake.ListPodsStub = stub
}

func (fake *FakeImpl) ListPodsArgsForCall(i int) (*kubernetes.Clientset, string) {
	fake.listPodsMutex.RLock()
	defer fake.listPodsMutex.RUnlock()
	argsForCall := fake.listPodsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) ListPodsReturns(result1 *v1.PodList, result2 error) {
	fake.listPodsMutex.Lock()
	defer fake.listPodsMutex.Unlock()
	fake.ListPodsStub = nil
	fake.listPodsReturns = struct {
		result1 *v1.PodList
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) ListPodsReturnsOnCall(i int, result1 *v1.PodList, result2 error) {
	fake.listPodsMutex.Lock()
	defer fake.listPodsMutex.Unlock()
	fake.ListPodsStub = nil
	if fake.listPodsReturnsOnCall == nil {
		fake.listPodsReturnsOnCall = make(map[int]struct {
			result1 *v1.PodList
			result2 error
		})
	}
	fake.listPodsReturnsOnCall[i] = struct {
		result1 *v1.PodList
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) MetricsAuditInc(arg1 api.SecurityProfilesOperatorClient, arg2 *api.MetricsAuditRequest) (*api.EmptyResponse, error) {
	fake.metricsAuditIncMutex.Lock()
	ret, specificReturn := fake.metricsAuditIncReturnsOnCall[len(fake.metricsAuditIncArgsForCall)]
	fake.metricsAuditIncArgsForCall = append(fake.metricsAuditIncArgsForCall, struct {
		arg1 api.SecurityProfilesOperatorClient
		arg2 *api.MetricsAuditRequest
	}{arg1, arg2})
	stub := fake.MetricsAuditIncStub
	fakeReturns := fake.metricsAuditIncReturns
	fake.recordInvocation("MetricsAuditInc", []interface{}{arg1, arg2})
	fake.metricsAuditIncMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) MetricsAuditIncCallCount() int {
	fake.metricsAuditIncMutex.RLock()
	defer fake.metricsAuditIncMutex.RUnlock()
	return len(fake.metricsAuditIncArgsForCall)
}

func (fake *FakeImpl) MetricsAuditIncCalls(stub func(api.SecurityProfilesOperatorClient, *api.MetricsAuditRequest) (*api.EmptyResponse, error)) {
	fake.metricsAuditIncMutex.Lock()
	defer fake.metricsAuditIncMutex.Unlock()
	fake.MetricsAuditIncStub = stub
}

func (fake *FakeImpl) MetricsAuditIncArgsForCall(i int) (api.SecurityProfilesOperatorClient, *api.MetricsAuditRequest) {
	fake.metricsAuditIncMutex.RLock()
	defer fake.metricsAuditIncMutex.RUnlock()
	argsForCall := fake.metricsAuditIncArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) MetricsAuditIncReturns(result1 *api.EmptyResponse, result2 error) {
	fake.metricsAuditIncMutex.Lock()
	defer fake.metricsAuditIncMutex.Unlock()
	fake.MetricsAuditIncStub = nil
	fake.metricsAuditIncReturns = struct {
		result1 *api.EmptyResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) MetricsAuditIncReturnsOnCall(i int, result1 *api.EmptyResponse, result2 error) {
	fake.metricsAuditIncMutex.Lock()
	defer fake.metricsAuditIncMutex.Unlock()
	fake.MetricsAuditIncStub = nil
	if fake.metricsAuditIncReturnsOnCall == nil {
		fake.metricsAuditIncReturnsOnCall = make(map[int]struct {
			result1 *api.EmptyResponse
			result2 error
		})
	}
	fake.metricsAuditIncReturnsOnCall[i] = struct {
		result1 *api.EmptyResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) NewForConfig(arg1 *rest.Config) (*kubernetes.Clientset, error) {
	fake.newForConfigMutex.Lock()
	ret, specificReturn := fake.newForConfigReturnsOnCall[len(fake.newForConfigArgsForCall)]
	fake.newForConfigArgsForCall = append(fake.newForConfigArgsForCall, struct {
		arg1 *rest.Config
	}{arg1})
	stub := fake.NewForConfigStub
	fakeReturns := fake.newForConfigReturns
	fake.recordInvocation("NewForConfig", []interface{}{arg1})
	fake.newForConfigMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) NewForConfigCallCount() int {
	fake.newForConfigMutex.RLock()
	defer fake.newForConfigMutex.RUnlock()
	return len(fake.newForConfigArgsForCall)
}

func (fake *FakeImpl) NewForConfigCalls(stub func(*rest.Config) (*kubernetes.Clientset, error)) {
	fake.newForConfigMutex.Lock()
	defer fake.newForConfigMutex.Unlock()
	fake.NewForConfigStub = stub
}

func (fake *FakeImpl) NewForConfigArgsForCall(i int) *rest.Config {
	fake.newForConfigMutex.RLock()
	defer fake.newForConfigMutex.RUnlock()
	argsForCall := fake.newForConfigArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) NewForConfigReturns(result1 *kubernetes.Clientset, result2 error) {
	fake.newForConfigMutex.Lock()
	defer fake.newForConfigMutex.Unlock()
	fake.NewForConfigStub = nil
	fake.newForConfigReturns = struct {
		result1 *kubernetes.Clientset
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) NewForConfigReturnsOnCall(i int, result1 *kubernetes.Clientset, result2 error) {
	fake.newForConfigMutex.Lock()
	defer fake.newForConfigMutex.Unlock()
	fake.NewForConfigStub = nil
	if fake.newForConfigReturnsOnCall == nil {
		fake.newForConfigReturnsOnCall = make(map[int]struct {
			result1 *kubernetes.Clientset
			result2 error
		})
	}
	fake.newForConfigReturnsOnCall[i] = struct {
		result1 *kubernetes.Clientset
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Open(arg1 string) (*os.File, error) {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.OpenStub
	fakeReturns := fake.openReturns
	fake.recordInvocation("Open", []interface{}{arg1})
	fake.openMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *FakeImpl) OpenCalls(stub func(string) (*os.File, error)) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = stub
}

func (fake *FakeImpl) OpenArgsForCall(i int) string {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	argsForCall := fake.openArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) OpenReturns(result1 *os.File, result2 error) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) OpenReturnsOnCall(i int, result1 *os.File, result2 error) {
	fake.openMutex.Lock()
	defer fake.openMutex.Unlock()
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 *os.File
			result2 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RecordSyscall(arg1 api.SecurityProfilesOperatorClient, arg2 *api.RecordSyscallRequest) (*api.EmptyResponse, error) {
	fake.recordSyscallMutex.Lock()
	ret, specificReturn := fake.recordSyscallReturnsOnCall[len(fake.recordSyscallArgsForCall)]
	fake.recordSyscallArgsForCall = append(fake.recordSyscallArgsForCall, struct {
		arg1 api.SecurityProfilesOperatorClient
		arg2 *api.RecordSyscallRequest
	}{arg1, arg2})
	stub := fake.RecordSyscallStub
	fakeReturns := fake.recordSyscallReturns
	fake.recordInvocation("RecordSyscall", []interface{}{arg1, arg2})
	fake.recordSyscallMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) RecordSyscallCallCount() int {
	fake.recordSyscallMutex.RLock()
	defer fake.recordSyscallMutex.RUnlock()
	return len(fake.recordSyscallArgsForCall)
}

func (fake *FakeImpl) RecordSyscallCalls(stub func(api.SecurityProfilesOperatorClient, *api.RecordSyscallRequest) (*api.EmptyResponse, error)) {
	fake.recordSyscallMutex.Lock()
	defer fake.recordSyscallMutex.Unlock()
	fake.RecordSyscallStub = stub
}

func (fake *FakeImpl) RecordSyscallArgsForCall(i int) (api.SecurityProfilesOperatorClient, *api.RecordSyscallRequest) {
	fake.recordSyscallMutex.RLock()
	defer fake.recordSyscallMutex.RUnlock()
	argsForCall := fake.recordSyscallArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) RecordSyscallReturns(result1 *api.EmptyResponse, result2 error) {
	fake.recordSyscallMutex.Lock()
	defer fake.recordSyscallMutex.Unlock()
	fake.RecordSyscallStub = nil
	fake.recordSyscallReturns = struct {
		result1 *api.EmptyResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) RecordSyscallReturnsOnCall(i int, result1 *api.EmptyResponse, result2 error) {
	fake.recordSyscallMutex.Lock()
	defer fake.recordSyscallMutex.Unlock()
	fake.RecordSyscallStub = nil
	if fake.recordSyscallReturnsOnCall == nil {
		fake.recordSyscallReturnsOnCall = make(map[int]struct {
			result1 *api.EmptyResponse
			result2 error
		})
	}
	fake.recordSyscallReturnsOnCall[i] = struct {
		result1 *api.EmptyResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) TailFile(arg1 string, arg2 tail.Config) (*tail.Tail, error) {
	fake.tailFileMutex.Lock()
	ret, specificReturn := fake.tailFileReturnsOnCall[len(fake.tailFileArgsForCall)]
	fake.tailFileArgsForCall = append(fake.tailFileArgsForCall, struct {
		arg1 string
		arg2 tail.Config
	}{arg1, arg2})
	stub := fake.TailFileStub
	fakeReturns := fake.tailFileReturns
	fake.recordInvocation("TailFile", []interface{}{arg1, arg2})
	fake.tailFileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) TailFileCallCount() int {
	fake.tailFileMutex.RLock()
	defer fake.tailFileMutex.RUnlock()
	return len(fake.tailFileArgsForCall)
}

func (fake *FakeImpl) TailFileCalls(stub func(string, tail.Config) (*tail.Tail, error)) {
	fake.tailFileMutex.Lock()
	defer fake.tailFileMutex.Unlock()
	fake.TailFileStub = stub
}

func (fake *FakeImpl) TailFileArgsForCall(i int) (string, tail.Config) {
	fake.tailFileMutex.RLock()
	defer fake.tailFileMutex.RUnlock()
	argsForCall := fake.tailFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) TailFileReturns(result1 *tail.Tail, result2 error) {
	fake.tailFileMutex.Lock()
	defer fake.tailFileMutex.Unlock()
	fake.TailFileStub = nil
	fake.tailFileReturns = struct {
		result1 *tail.Tail
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) TailFileReturnsOnCall(i int, result1 *tail.Tail, result2 error) {
	fake.tailFileMutex.Lock()
	defer fake.tailFileMutex.Unlock()
	fake.TailFileStub = nil
	if fake.tailFileReturnsOnCall == nil {
		fake.tailFileReturnsOnCall = make(map[int]struct {
			result1 *tail.Tail
			result2 error
		})
	}
	fake.tailFileReturnsOnCall[i] = struct {
		result1 *tail.Tail
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.dialMutex.RLock()
	defer fake.dialMutex.RUnlock()
	fake.getenvMutex.RLock()
	defer fake.getenvMutex.RUnlock()
	fake.inClusterConfigMutex.RLock()
	defer fake.inClusterConfigMutex.RUnlock()
	fake.linesMutex.RLock()
	defer fake.linesMutex.RUnlock()
	fake.listPodsMutex.RLock()
	defer fake.listPodsMutex.RUnlock()
	fake.metricsAuditIncMutex.RLock()
	defer fake.metricsAuditIncMutex.RUnlock()
	fake.newForConfigMutex.RLock()
	defer fake.newForConfigMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.recordSyscallMutex.RLock()
	defer fake.recordSyscallMutex.RUnlock()
	fake.tailFileMutex.RLock()
	defer fake.tailFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImpl) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
