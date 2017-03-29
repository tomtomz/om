// This file was generated by counterfeiter
package fakes

import "sync"

type LiveWriter struct {
	WriteStub        func(p []byte) (n int, err error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		p []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
	writeReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	StartStub        func()
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	StopStub         func()
	stopMutex        sync.RWMutex
	stopArgsForCall  []struct{}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LiveWriter) Write(p []byte) (n int, err error) {
	var pCopy []byte
	if p != nil {
		pCopy = make([]byte, len(p))
		copy(pCopy, p)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		p []byte
	}{pCopy})
	fake.recordInvocation("Write", []interface{}{pCopy})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(p)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.writeReturns.result1, fake.writeReturns.result2
}

func (fake *LiveWriter) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *LiveWriter) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].p
}

func (fake *LiveWriter) WriteReturns(result1 int, result2 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *LiveWriter) WriteReturnsOnCall(i int, result1 int, result2 error) {
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *LiveWriter) Start() {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		fake.StartStub()
	}
}

func (fake *LiveWriter) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *LiveWriter) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		fake.StopStub()
	}
}

func (fake *LiveWriter) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *LiveWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.invocations
}

func (fake *LiveWriter) recordInvocation(key string, args []interface{}) {
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
