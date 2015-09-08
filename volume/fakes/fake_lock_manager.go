// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/baggageclaim/volume"
)

type FakeLockManager struct {
	LockStub        func(key string)
	lockMutex       sync.RWMutex
	lockArgsForCall []struct {
		key string
	}
	UnlockStub        func(key string)
	unlockMutex       sync.RWMutex
	unlockArgsForCall []struct {
		key string
	}
}

func (fake *FakeLockManager) Lock(key string) {
	fake.lockMutex.Lock()
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct {
		key string
	}{key})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		fake.LockStub(key)
	}
}

func (fake *FakeLockManager) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *FakeLockManager) LockArgsForCall(i int) string {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return fake.lockArgsForCall[i].key
}

func (fake *FakeLockManager) Unlock(key string) {
	fake.unlockMutex.Lock()
	fake.unlockArgsForCall = append(fake.unlockArgsForCall, struct {
		key string
	}{key})
	fake.unlockMutex.Unlock()
	if fake.UnlockStub != nil {
		fake.UnlockStub(key)
	}
}

func (fake *FakeLockManager) UnlockCallCount() int {
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	return len(fake.unlockArgsForCall)
}

func (fake *FakeLockManager) UnlockArgsForCall(i int) string {
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	return fake.unlockArgsForCall[i].key
}

var _ volume.LockManager = new(FakeLockManager)
