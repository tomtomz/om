// This file was generated by counterfeiter
package fakes

import "sync"

type AvailableProductChecker struct {
	CheckProductAvailabilityStub        func(productName string, productVersion string) (bool, error)
	checkProductAvailabilityMutex       sync.RWMutex
	checkProductAvailabilityArgsForCall []struct {
		productName    string
		productVersion string
	}
	checkProductAvailabilityReturns struct {
		result1 bool
		result2 error
	}
	checkProductAvailabilityReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *AvailableProductChecker) CheckProductAvailability(productName string, productVersion string) (bool, error) {
	fake.checkProductAvailabilityMutex.Lock()
	ret, specificReturn := fake.checkProductAvailabilityReturnsOnCall[len(fake.checkProductAvailabilityArgsForCall)]
	fake.checkProductAvailabilityArgsForCall = append(fake.checkProductAvailabilityArgsForCall, struct {
		productName    string
		productVersion string
	}{productName, productVersion})
	fake.recordInvocation("CheckProductAvailability", []interface{}{productName, productVersion})
	fake.checkProductAvailabilityMutex.Unlock()
	if fake.CheckProductAvailabilityStub != nil {
		return fake.CheckProductAvailabilityStub(productName, productVersion)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.checkProductAvailabilityReturns.result1, fake.checkProductAvailabilityReturns.result2
}

func (fake *AvailableProductChecker) CheckProductAvailabilityCallCount() int {
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	return len(fake.checkProductAvailabilityArgsForCall)
}

func (fake *AvailableProductChecker) CheckProductAvailabilityArgsForCall(i int) (string, string) {
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	return fake.checkProductAvailabilityArgsForCall[i].productName, fake.checkProductAvailabilityArgsForCall[i].productVersion
}

func (fake *AvailableProductChecker) CheckProductAvailabilityReturns(result1 bool, result2 error) {
	fake.CheckProductAvailabilityStub = nil
	fake.checkProductAvailabilityReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *AvailableProductChecker) CheckProductAvailabilityReturnsOnCall(i int, result1 bool, result2 error) {
	fake.CheckProductAvailabilityStub = nil
	if fake.checkProductAvailabilityReturnsOnCall == nil {
		fake.checkProductAvailabilityReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.checkProductAvailabilityReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *AvailableProductChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkProductAvailabilityMutex.RLock()
	defer fake.checkProductAvailabilityMutex.RUnlock()
	return fake.invocations
}

func (fake *AvailableProductChecker) recordInvocation(key string, args []interface{}) {
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
