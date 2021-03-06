// This file was generated by counterfeiter
package fake_namespacer

import (
	"sync"

	"github.com/concourse/baggageclaim/uidjunk"
)

type FakeNamespacer struct {
	CacheKeyStub        func() string
	cacheKeyMutex       sync.RWMutex
	cacheKeyArgsForCall []struct{}
	cacheKeyReturns     struct {
		result1 string
	}
	NamespaceStub        func(rootfsPath string) error
	namespaceMutex       sync.RWMutex
	namespaceArgsForCall []struct {
		rootfsPath string
	}
	namespaceReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNamespacer) CacheKey() string {
	fake.cacheKeyMutex.Lock()
	fake.cacheKeyArgsForCall = append(fake.cacheKeyArgsForCall, struct{}{})
	fake.recordInvocation("CacheKey", []interface{}{})
	fake.cacheKeyMutex.Unlock()
	if fake.CacheKeyStub != nil {
		return fake.CacheKeyStub()
	} else {
		return fake.cacheKeyReturns.result1
	}
}

func (fake *FakeNamespacer) CacheKeyCallCount() int {
	fake.cacheKeyMutex.RLock()
	defer fake.cacheKeyMutex.RUnlock()
	return len(fake.cacheKeyArgsForCall)
}

func (fake *FakeNamespacer) CacheKeyReturns(result1 string) {
	fake.CacheKeyStub = nil
	fake.cacheKeyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeNamespacer) Namespace(rootfsPath string) error {
	fake.namespaceMutex.Lock()
	fake.namespaceArgsForCall = append(fake.namespaceArgsForCall, struct {
		rootfsPath string
	}{rootfsPath})
	fake.recordInvocation("Namespace", []interface{}{rootfsPath})
	fake.namespaceMutex.Unlock()
	if fake.NamespaceStub != nil {
		return fake.NamespaceStub(rootfsPath)
	} else {
		return fake.namespaceReturns.result1
	}
}

func (fake *FakeNamespacer) NamespaceCallCount() int {
	fake.namespaceMutex.RLock()
	defer fake.namespaceMutex.RUnlock()
	return len(fake.namespaceArgsForCall)
}

func (fake *FakeNamespacer) NamespaceArgsForCall(i int) string {
	fake.namespaceMutex.RLock()
	defer fake.namespaceMutex.RUnlock()
	return fake.namespaceArgsForCall[i].rootfsPath
}

func (fake *FakeNamespacer) NamespaceReturns(result1 error) {
	fake.NamespaceStub = nil
	fake.namespaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNamespacer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cacheKeyMutex.RLock()
	defer fake.cacheKeyMutex.RUnlock()
	fake.namespaceMutex.RLock()
	defer fake.namespaceMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeNamespacer) recordInvocation(key string, args []interface{}) {
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

var _ uidjunk.Namespacer = new(FakeNamespacer)
