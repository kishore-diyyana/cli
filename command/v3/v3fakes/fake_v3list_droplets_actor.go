// Code generated by counterfeiter. DO NOT EDIT.
package v3fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v3action"
	"code.cloudfoundry.org/cli/command/v3"
)

type FakeV3ListDropletsActor struct {
	GetApplicationDropletsStub        func(appName string, spaceGUID string) ([]v3action.Droplet, v3action.Warnings, error)
	getApplicationDropletsMutex       sync.RWMutex
	getApplicationDropletsArgsForCall []struct {
		appName   string
		spaceGUID string
	}
	getApplicationDropletsReturns struct {
		result1 []v3action.Droplet
		result2 v3action.Warnings
		result3 error
	}
	getApplicationDropletsReturnsOnCall map[int]struct {
		result1 []v3action.Droplet
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3ListDropletsActor) GetApplicationDroplets(appName string, spaceGUID string) ([]v3action.Droplet, v3action.Warnings, error) {
	fake.getApplicationDropletsMutex.Lock()
	ret, specificReturn := fake.getApplicationDropletsReturnsOnCall[len(fake.getApplicationDropletsArgsForCall)]
	fake.getApplicationDropletsArgsForCall = append(fake.getApplicationDropletsArgsForCall, struct {
		appName   string
		spaceGUID string
	}{appName, spaceGUID})
	fake.recordInvocation("GetApplicationDroplets", []interface{}{appName, spaceGUID})
	fake.getApplicationDropletsMutex.Unlock()
	if fake.GetApplicationDropletsStub != nil {
		return fake.GetApplicationDropletsStub(appName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationDropletsReturns.result1, fake.getApplicationDropletsReturns.result2, fake.getApplicationDropletsReturns.result3
}

func (fake *FakeV3ListDropletsActor) GetApplicationDropletsCallCount() int {
	fake.getApplicationDropletsMutex.RLock()
	defer fake.getApplicationDropletsMutex.RUnlock()
	return len(fake.getApplicationDropletsArgsForCall)
}

func (fake *FakeV3ListDropletsActor) GetApplicationDropletsArgsForCall(i int) (string, string) {
	fake.getApplicationDropletsMutex.RLock()
	defer fake.getApplicationDropletsMutex.RUnlock()
	return fake.getApplicationDropletsArgsForCall[i].appName, fake.getApplicationDropletsArgsForCall[i].spaceGUID
}

func (fake *FakeV3ListDropletsActor) GetApplicationDropletsReturns(result1 []v3action.Droplet, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationDropletsStub = nil
	fake.getApplicationDropletsReturns = struct {
		result1 []v3action.Droplet
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3ListDropletsActor) GetApplicationDropletsReturnsOnCall(i int, result1 []v3action.Droplet, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationDropletsStub = nil
	if fake.getApplicationDropletsReturnsOnCall == nil {
		fake.getApplicationDropletsReturnsOnCall = make(map[int]struct {
			result1 []v3action.Droplet
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationDropletsReturnsOnCall[i] = struct {
		result1 []v3action.Droplet
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3ListDropletsActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationDropletsMutex.RLock()
	defer fake.getApplicationDropletsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3ListDropletsActor) recordInvocation(key string, args []interface{}) {
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

var _ v3.V3ListDropletsActor = new(FakeV3ListDropletsActor)