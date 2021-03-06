// Code generated by counterfeiter. DO NOT EDIT.
package cfnetworkingactionfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/cfnetworkingaction"
	"code.cloudfoundry.org/cli/actor/v3action"
)

type FakeV3Actor struct {
	GetApplicationByNameAndSpaceStub        func(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		appName   string
		spaceGUID string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	GetApplicationsBySpaceStub        func(spaceGUID string) ([]v3action.Application, v3action.Warnings, error)
	getApplicationsBySpaceMutex       sync.RWMutex
	getApplicationsBySpaceArgsForCall []struct {
		spaceGUID string
	}
	getApplicationsBySpaceReturns struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	getApplicationsBySpaceReturnsOnCall map[int]struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpace(appName string, spaceGUID string) (v3action.Application, v3action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		appName   string
		spaceGUID string
	}{appName, spaceGUID})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{appName, spaceGUID})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(appName, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationByNameAndSpaceReturns.result1, fake.getApplicationByNameAndSpaceReturns.result2, fake.getApplicationByNameAndSpaceReturns.result3
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return fake.getApplicationByNameAndSpaceArgsForCall[i].appName, fake.getApplicationByNameAndSpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceReturns(result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationsBySpace(spaceGUID string) ([]v3action.Application, v3action.Warnings, error) {
	fake.getApplicationsBySpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationsBySpaceReturnsOnCall[len(fake.getApplicationsBySpaceArgsForCall)]
	fake.getApplicationsBySpaceArgsForCall = append(fake.getApplicationsBySpaceArgsForCall, struct {
		spaceGUID string
	}{spaceGUID})
	fake.recordInvocation("GetApplicationsBySpace", []interface{}{spaceGUID})
	fake.getApplicationsBySpaceMutex.Unlock()
	if fake.GetApplicationsBySpaceStub != nil {
		return fake.GetApplicationsBySpaceStub(spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.getApplicationsBySpaceReturns.result1, fake.getApplicationsBySpaceReturns.result2, fake.getApplicationsBySpaceReturns.result3
}

func (fake *FakeV3Actor) GetApplicationsBySpaceCallCount() int {
	fake.getApplicationsBySpaceMutex.RLock()
	defer fake.getApplicationsBySpaceMutex.RUnlock()
	return len(fake.getApplicationsBySpaceArgsForCall)
}

func (fake *FakeV3Actor) GetApplicationsBySpaceArgsForCall(i int) string {
	fake.getApplicationsBySpaceMutex.RLock()
	defer fake.getApplicationsBySpaceMutex.RUnlock()
	return fake.getApplicationsBySpaceArgsForCall[i].spaceGUID
}

func (fake *FakeV3Actor) GetApplicationsBySpaceReturns(result1 []v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationsBySpaceStub = nil
	fake.getApplicationsBySpaceReturns = struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) GetApplicationsBySpaceReturnsOnCall(i int, result1 []v3action.Application, result2 v3action.Warnings, result3 error) {
	fake.GetApplicationsBySpaceStub = nil
	if fake.getApplicationsBySpaceReturnsOnCall == nil {
		fake.getApplicationsBySpaceReturnsOnCall = make(map[int]struct {
			result1 []v3action.Application
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getApplicationsBySpaceReturnsOnCall[i] = struct {
		result1 []v3action.Application
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3Actor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.getApplicationsBySpaceMutex.RLock()
	defer fake.getApplicationsBySpaceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3Actor) recordInvocation(key string, args []interface{}) {
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

var _ cfnetworkingaction.V3Actor = new(FakeV3Actor)
