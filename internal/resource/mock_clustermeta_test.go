// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/derailed/k9s/internal/resource (interfaces: ClusterMeta)

package resource_test

import (
	k8s "github.com/derailed/k9s/internal/k8s"
	pegomock "github.com/petergtz/pegomock"
	dynamic "k8s.io/client-go/dynamic"
	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	versioned "k8s.io/metrics/pkg/client/clientset/versioned"
	"reflect"
	"time"
)

type MockClusterMeta struct {
	fail func(message string, callerSkip ...int)
}

func NewMockClusterMeta() *MockClusterMeta {
	return &MockClusterMeta{fail: pegomock.GlobalFailHandler}
}

func (mock *MockClusterMeta) ClusterName() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ClusterName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) Config() *k8s.Config {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Config", params, []reflect.Type{reflect.TypeOf((**k8s.Config)(nil)).Elem()})
	var ret0 *k8s.Config
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*k8s.Config)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) ContextName() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ContextName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) DialOrDie() kubernetes.Interface {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DialOrDie", params, []reflect.Type{reflect.TypeOf((*kubernetes.Interface)(nil)).Elem()})
	var ret0 kubernetes.Interface
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(kubernetes.Interface)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) DynDialOrDie() dynamic.Interface {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DynDialOrDie", params, []reflect.Type{reflect.TypeOf((*dynamic.Interface)(nil)).Elem()})
	var ret0 dynamic.Interface
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(dynamic.Interface)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) HasMetrics() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("HasMetrics", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) MXDial() (*versioned.Clientset, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("MXDial", params, []reflect.Type{reflect.TypeOf((**versioned.Clientset)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *versioned.Clientset
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*versioned.Clientset)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterMeta) NSDialOrDie() dynamic.NamespaceableResourceInterface {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("NSDialOrDie", params, []reflect.Type{reflect.TypeOf((*dynamic.NamespaceableResourceInterface)(nil)).Elem()})
	var ret0 dynamic.NamespaceableResourceInterface
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(dynamic.NamespaceableResourceInterface)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) RestConfigOrDie() *rest.Config {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("RestConfigOrDie", params, []reflect.Type{reflect.TypeOf((**rest.Config)(nil)).Elem()})
	var ret0 *rest.Config
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*rest.Config)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) SwitchContextOrDie(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SwitchContextOrDie", params, []reflect.Type{})
}

func (mock *MockClusterMeta) UserName() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UserName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockClusterMeta) Version() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockClusterMeta().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Version", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockClusterMeta) VerifyWasCalledOnce() *VerifierClusterMeta {
	return &VerifierClusterMeta{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockClusterMeta) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierClusterMeta {
	return &VerifierClusterMeta{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockClusterMeta) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierClusterMeta {
	return &VerifierClusterMeta{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockClusterMeta) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierClusterMeta {
	return &VerifierClusterMeta{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierClusterMeta struct {
	mock                   *MockClusterMeta
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierClusterMeta) ClusterName() *ClusterMeta_ClusterName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ClusterName", params, verifier.timeout)
	return &ClusterMeta_ClusterName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_ClusterName_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_ClusterName_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterMeta_ClusterName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterMeta) Config() *ClusterMeta_Config_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Config", params, verifier.timeout)
	return &ClusterMeta_Config_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_Config_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_Config_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterMeta_Config_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterMeta) ContextName() *ClusterMeta_ContextName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ContextName", params, verifier.timeout)
	return &ClusterMeta_ContextName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_ContextName_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_ContextName_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterMeta_ContextName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterMeta) DialOrDie() *ClusterMeta_DialOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DialOrDie", params, verifier.timeout)
	return &ClusterMeta_DialOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_DialOrDie_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_DialOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterMeta_DialOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterMeta) DynDialOrDie() *ClusterMeta_DynDialOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DynDialOrDie", params, verifier.timeout)
	return &ClusterMeta_DynDialOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_DynDialOrDie_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_DynDialOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterMeta_DynDialOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterMeta) HasMetrics() *ClusterMeta_HasMetrics_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "HasMetrics", params, verifier.timeout)
	return &ClusterMeta_HasMetrics_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_HasMetrics_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_HasMetrics_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterMeta_HasMetrics_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterMeta) MXDial() *ClusterMeta_MXDial_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "MXDial", params, verifier.timeout)
	return &ClusterMeta_MXDial_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_MXDial_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_MXDial_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterMeta_MXDial_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterMeta) NSDialOrDie() *ClusterMeta_NSDialOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NSDialOrDie", params, verifier.timeout)
	return &ClusterMeta_NSDialOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_NSDialOrDie_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_NSDialOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterMeta_NSDialOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterMeta) RestConfigOrDie() *ClusterMeta_RestConfigOrDie_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RestConfigOrDie", params, verifier.timeout)
	return &ClusterMeta_RestConfigOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_RestConfigOrDie_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_RestConfigOrDie_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterMeta_RestConfigOrDie_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterMeta) SwitchContextOrDie(_param0 string) *ClusterMeta_SwitchContextOrDie_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SwitchContextOrDie", params, verifier.timeout)
	return &ClusterMeta_SwitchContextOrDie_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_SwitchContextOrDie_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_SwitchContextOrDie_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *ClusterMeta_SwitchContextOrDie_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierClusterMeta) UserName() *ClusterMeta_UserName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UserName", params, verifier.timeout)
	return &ClusterMeta_UserName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_UserName_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_UserName_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterMeta_UserName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierClusterMeta) Version() *ClusterMeta_Version_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Version", params, verifier.timeout)
	return &ClusterMeta_Version_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type ClusterMeta_Version_OngoingVerification struct {
	mock              *MockClusterMeta
	methodInvocations []pegomock.MethodInvocation
}

func (c *ClusterMeta_Version_OngoingVerification) GetCapturedArguments() {
}

func (c *ClusterMeta_Version_OngoingVerification) GetAllCapturedArguments() {
}
