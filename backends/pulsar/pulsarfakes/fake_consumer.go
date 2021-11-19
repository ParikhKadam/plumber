// Code generated by counterfeiter. DO NOT EDIT.
package pulsarfakes

import (
	"context"
	"sync"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

type FakeConsumer struct {
	AckStub        func(pulsar.Message)
	ackMutex       sync.RWMutex
	ackArgsForCall []struct {
		arg1 pulsar.Message
	}
	AckIDStub        func(pulsar.MessageID)
	ackIDMutex       sync.RWMutex
	ackIDArgsForCall []struct {
		arg1 pulsar.MessageID
	}
	ChanStub        func() <-chan pulsar.ConsumerMessage
	chanMutex       sync.RWMutex
	chanArgsForCall []struct {
	}
	chanReturns struct {
		result1 <-chan pulsar.ConsumerMessage
	}
	chanReturnsOnCall map[int]struct {
		result1 <-chan pulsar.ConsumerMessage
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	NackStub        func(pulsar.Message)
	nackMutex       sync.RWMutex
	nackArgsForCall []struct {
		arg1 pulsar.Message
	}
	NackIDStub        func(pulsar.MessageID)
	nackIDMutex       sync.RWMutex
	nackIDArgsForCall []struct {
		arg1 pulsar.MessageID
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	ReceiveStub        func(context.Context) (pulsar.Message, error)
	receiveMutex       sync.RWMutex
	receiveArgsForCall []struct {
		arg1 context.Context
	}
	receiveReturns struct {
		result1 pulsar.Message
		result2 error
	}
	receiveReturnsOnCall map[int]struct {
		result1 pulsar.Message
		result2 error
	}
	ReconsumeLaterStub        func(pulsar.Message, time.Duration)
	reconsumeLaterMutex       sync.RWMutex
	reconsumeLaterArgsForCall []struct {
		arg1 pulsar.Message
		arg2 time.Duration
	}
	SeekStub        func(pulsar.MessageID) error
	seekMutex       sync.RWMutex
	seekArgsForCall []struct {
		arg1 pulsar.MessageID
	}
	seekReturns struct {
		result1 error
	}
	seekReturnsOnCall map[int]struct {
		result1 error
	}
	SeekByTimeStub        func(time.Time) error
	seekByTimeMutex       sync.RWMutex
	seekByTimeArgsForCall []struct {
		arg1 time.Time
	}
	seekByTimeReturns struct {
		result1 error
	}
	seekByTimeReturnsOnCall map[int]struct {
		result1 error
	}
	SubscriptionStub        func() string
	subscriptionMutex       sync.RWMutex
	subscriptionArgsForCall []struct {
	}
	subscriptionReturns struct {
		result1 string
	}
	subscriptionReturnsOnCall map[int]struct {
		result1 string
	}
	UnsubscribeStub        func() error
	unsubscribeMutex       sync.RWMutex
	unsubscribeArgsForCall []struct {
	}
	unsubscribeReturns struct {
		result1 error
	}
	unsubscribeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConsumer) Ack(arg1 pulsar.Message) {
	fake.ackMutex.Lock()
	fake.ackArgsForCall = append(fake.ackArgsForCall, struct {
		arg1 pulsar.Message
	}{arg1})
	stub := fake.AckStub
	fake.recordInvocation("Ack", []interface{}{arg1})
	fake.ackMutex.Unlock()
	if stub != nil {
		fake.AckStub(arg1)
	}
}

func (fake *FakeConsumer) AckCallCount() int {
	fake.ackMutex.RLock()
	defer fake.ackMutex.RUnlock()
	return len(fake.ackArgsForCall)
}

func (fake *FakeConsumer) AckCalls(stub func(pulsar.Message)) {
	fake.ackMutex.Lock()
	defer fake.ackMutex.Unlock()
	fake.AckStub = stub
}

func (fake *FakeConsumer) AckArgsForCall(i int) pulsar.Message {
	fake.ackMutex.RLock()
	defer fake.ackMutex.RUnlock()
	argsForCall := fake.ackArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConsumer) AckID(arg1 pulsar.MessageID) {
	fake.ackIDMutex.Lock()
	fake.ackIDArgsForCall = append(fake.ackIDArgsForCall, struct {
		arg1 pulsar.MessageID
	}{arg1})
	stub := fake.AckIDStub
	fake.recordInvocation("AckID", []interface{}{arg1})
	fake.ackIDMutex.Unlock()
	if stub != nil {
		fake.AckIDStub(arg1)
	}
}

func (fake *FakeConsumer) AckIDCallCount() int {
	fake.ackIDMutex.RLock()
	defer fake.ackIDMutex.RUnlock()
	return len(fake.ackIDArgsForCall)
}

func (fake *FakeConsumer) AckIDCalls(stub func(pulsar.MessageID)) {
	fake.ackIDMutex.Lock()
	defer fake.ackIDMutex.Unlock()
	fake.AckIDStub = stub
}

func (fake *FakeConsumer) AckIDArgsForCall(i int) pulsar.MessageID {
	fake.ackIDMutex.RLock()
	defer fake.ackIDMutex.RUnlock()
	argsForCall := fake.ackIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConsumer) Chan() <-chan pulsar.ConsumerMessage {
	fake.chanMutex.Lock()
	ret, specificReturn := fake.chanReturnsOnCall[len(fake.chanArgsForCall)]
	fake.chanArgsForCall = append(fake.chanArgsForCall, struct {
	}{})
	stub := fake.ChanStub
	fakeReturns := fake.chanReturns
	fake.recordInvocation("Chan", []interface{}{})
	fake.chanMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConsumer) ChanCallCount() int {
	fake.chanMutex.RLock()
	defer fake.chanMutex.RUnlock()
	return len(fake.chanArgsForCall)
}

func (fake *FakeConsumer) ChanCalls(stub func() <-chan pulsar.ConsumerMessage) {
	fake.chanMutex.Lock()
	defer fake.chanMutex.Unlock()
	fake.ChanStub = stub
}

func (fake *FakeConsumer) ChanReturns(result1 <-chan pulsar.ConsumerMessage) {
	fake.chanMutex.Lock()
	defer fake.chanMutex.Unlock()
	fake.ChanStub = nil
	fake.chanReturns = struct {
		result1 <-chan pulsar.ConsumerMessage
	}{result1}
}

func (fake *FakeConsumer) ChanReturnsOnCall(i int, result1 <-chan pulsar.ConsumerMessage) {
	fake.chanMutex.Lock()
	defer fake.chanMutex.Unlock()
	fake.ChanStub = nil
	if fake.chanReturnsOnCall == nil {
		fake.chanReturnsOnCall = make(map[int]struct {
			result1 <-chan pulsar.ConsumerMessage
		})
	}
	fake.chanReturnsOnCall[i] = struct {
		result1 <-chan pulsar.ConsumerMessage
	}{result1}
}

func (fake *FakeConsumer) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeConsumer) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeConsumer) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeConsumer) Nack(arg1 pulsar.Message) {
	fake.nackMutex.Lock()
	fake.nackArgsForCall = append(fake.nackArgsForCall, struct {
		arg1 pulsar.Message
	}{arg1})
	stub := fake.NackStub
	fake.recordInvocation("Nack", []interface{}{arg1})
	fake.nackMutex.Unlock()
	if stub != nil {
		fake.NackStub(arg1)
	}
}

func (fake *FakeConsumer) NackCallCount() int {
	fake.nackMutex.RLock()
	defer fake.nackMutex.RUnlock()
	return len(fake.nackArgsForCall)
}

func (fake *FakeConsumer) NackCalls(stub func(pulsar.Message)) {
	fake.nackMutex.Lock()
	defer fake.nackMutex.Unlock()
	fake.NackStub = stub
}

func (fake *FakeConsumer) NackArgsForCall(i int) pulsar.Message {
	fake.nackMutex.RLock()
	defer fake.nackMutex.RUnlock()
	argsForCall := fake.nackArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConsumer) NackID(arg1 pulsar.MessageID) {
	fake.nackIDMutex.Lock()
	fake.nackIDArgsForCall = append(fake.nackIDArgsForCall, struct {
		arg1 pulsar.MessageID
	}{arg1})
	stub := fake.NackIDStub
	fake.recordInvocation("NackID", []interface{}{arg1})
	fake.nackIDMutex.Unlock()
	if stub != nil {
		fake.NackIDStub(arg1)
	}
}

func (fake *FakeConsumer) NackIDCallCount() int {
	fake.nackIDMutex.RLock()
	defer fake.nackIDMutex.RUnlock()
	return len(fake.nackIDArgsForCall)
}

func (fake *FakeConsumer) NackIDCalls(stub func(pulsar.MessageID)) {
	fake.nackIDMutex.Lock()
	defer fake.nackIDMutex.Unlock()
	fake.NackIDStub = stub
}

func (fake *FakeConsumer) NackIDArgsForCall(i int) pulsar.MessageID {
	fake.nackIDMutex.RLock()
	defer fake.nackIDMutex.RUnlock()
	argsForCall := fake.nackIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConsumer) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConsumer) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeConsumer) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeConsumer) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConsumer) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConsumer) Receive(arg1 context.Context) (pulsar.Message, error) {
	fake.receiveMutex.Lock()
	ret, specificReturn := fake.receiveReturnsOnCall[len(fake.receiveArgsForCall)]
	fake.receiveArgsForCall = append(fake.receiveArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ReceiveStub
	fakeReturns := fake.receiveReturns
	fake.recordInvocation("Receive", []interface{}{arg1})
	fake.receiveMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeConsumer) ReceiveCallCount() int {
	fake.receiveMutex.RLock()
	defer fake.receiveMutex.RUnlock()
	return len(fake.receiveArgsForCall)
}

func (fake *FakeConsumer) ReceiveCalls(stub func(context.Context) (pulsar.Message, error)) {
	fake.receiveMutex.Lock()
	defer fake.receiveMutex.Unlock()
	fake.ReceiveStub = stub
}

func (fake *FakeConsumer) ReceiveArgsForCall(i int) context.Context {
	fake.receiveMutex.RLock()
	defer fake.receiveMutex.RUnlock()
	argsForCall := fake.receiveArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConsumer) ReceiveReturns(result1 pulsar.Message, result2 error) {
	fake.receiveMutex.Lock()
	defer fake.receiveMutex.Unlock()
	fake.ReceiveStub = nil
	fake.receiveReturns = struct {
		result1 pulsar.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeConsumer) ReceiveReturnsOnCall(i int, result1 pulsar.Message, result2 error) {
	fake.receiveMutex.Lock()
	defer fake.receiveMutex.Unlock()
	fake.ReceiveStub = nil
	if fake.receiveReturnsOnCall == nil {
		fake.receiveReturnsOnCall = make(map[int]struct {
			result1 pulsar.Message
			result2 error
		})
	}
	fake.receiveReturnsOnCall[i] = struct {
		result1 pulsar.Message
		result2 error
	}{result1, result2}
}

func (fake *FakeConsumer) ReconsumeLater(arg1 pulsar.Message, arg2 time.Duration) {
	fake.reconsumeLaterMutex.Lock()
	fake.reconsumeLaterArgsForCall = append(fake.reconsumeLaterArgsForCall, struct {
		arg1 pulsar.Message
		arg2 time.Duration
	}{arg1, arg2})
	stub := fake.ReconsumeLaterStub
	fake.recordInvocation("ReconsumeLater", []interface{}{arg1, arg2})
	fake.reconsumeLaterMutex.Unlock()
	if stub != nil {
		fake.ReconsumeLaterStub(arg1, arg2)
	}
}

func (fake *FakeConsumer) ReconsumeLaterCallCount() int {
	fake.reconsumeLaterMutex.RLock()
	defer fake.reconsumeLaterMutex.RUnlock()
	return len(fake.reconsumeLaterArgsForCall)
}

func (fake *FakeConsumer) ReconsumeLaterCalls(stub func(pulsar.Message, time.Duration)) {
	fake.reconsumeLaterMutex.Lock()
	defer fake.reconsumeLaterMutex.Unlock()
	fake.ReconsumeLaterStub = stub
}

func (fake *FakeConsumer) ReconsumeLaterArgsForCall(i int) (pulsar.Message, time.Duration) {
	fake.reconsumeLaterMutex.RLock()
	defer fake.reconsumeLaterMutex.RUnlock()
	argsForCall := fake.reconsumeLaterArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeConsumer) Seek(arg1 pulsar.MessageID) error {
	fake.seekMutex.Lock()
	ret, specificReturn := fake.seekReturnsOnCall[len(fake.seekArgsForCall)]
	fake.seekArgsForCall = append(fake.seekArgsForCall, struct {
		arg1 pulsar.MessageID
	}{arg1})
	stub := fake.SeekStub
	fakeReturns := fake.seekReturns
	fake.recordInvocation("Seek", []interface{}{arg1})
	fake.seekMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConsumer) SeekCallCount() int {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	return len(fake.seekArgsForCall)
}

func (fake *FakeConsumer) SeekCalls(stub func(pulsar.MessageID) error) {
	fake.seekMutex.Lock()
	defer fake.seekMutex.Unlock()
	fake.SeekStub = stub
}

func (fake *FakeConsumer) SeekArgsForCall(i int) pulsar.MessageID {
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	argsForCall := fake.seekArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConsumer) SeekReturns(result1 error) {
	fake.seekMutex.Lock()
	defer fake.seekMutex.Unlock()
	fake.SeekStub = nil
	fake.seekReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumer) SeekReturnsOnCall(i int, result1 error) {
	fake.seekMutex.Lock()
	defer fake.seekMutex.Unlock()
	fake.SeekStub = nil
	if fake.seekReturnsOnCall == nil {
		fake.seekReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.seekReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumer) SeekByTime(arg1 time.Time) error {
	fake.seekByTimeMutex.Lock()
	ret, specificReturn := fake.seekByTimeReturnsOnCall[len(fake.seekByTimeArgsForCall)]
	fake.seekByTimeArgsForCall = append(fake.seekByTimeArgsForCall, struct {
		arg1 time.Time
	}{arg1})
	stub := fake.SeekByTimeStub
	fakeReturns := fake.seekByTimeReturns
	fake.recordInvocation("SeekByTime", []interface{}{arg1})
	fake.seekByTimeMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConsumer) SeekByTimeCallCount() int {
	fake.seekByTimeMutex.RLock()
	defer fake.seekByTimeMutex.RUnlock()
	return len(fake.seekByTimeArgsForCall)
}

func (fake *FakeConsumer) SeekByTimeCalls(stub func(time.Time) error) {
	fake.seekByTimeMutex.Lock()
	defer fake.seekByTimeMutex.Unlock()
	fake.SeekByTimeStub = stub
}

func (fake *FakeConsumer) SeekByTimeArgsForCall(i int) time.Time {
	fake.seekByTimeMutex.RLock()
	defer fake.seekByTimeMutex.RUnlock()
	argsForCall := fake.seekByTimeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeConsumer) SeekByTimeReturns(result1 error) {
	fake.seekByTimeMutex.Lock()
	defer fake.seekByTimeMutex.Unlock()
	fake.SeekByTimeStub = nil
	fake.seekByTimeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumer) SeekByTimeReturnsOnCall(i int, result1 error) {
	fake.seekByTimeMutex.Lock()
	defer fake.seekByTimeMutex.Unlock()
	fake.SeekByTimeStub = nil
	if fake.seekByTimeReturnsOnCall == nil {
		fake.seekByTimeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.seekByTimeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumer) Subscription() string {
	fake.subscriptionMutex.Lock()
	ret, specificReturn := fake.subscriptionReturnsOnCall[len(fake.subscriptionArgsForCall)]
	fake.subscriptionArgsForCall = append(fake.subscriptionArgsForCall, struct {
	}{})
	stub := fake.SubscriptionStub
	fakeReturns := fake.subscriptionReturns
	fake.recordInvocation("Subscription", []interface{}{})
	fake.subscriptionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConsumer) SubscriptionCallCount() int {
	fake.subscriptionMutex.RLock()
	defer fake.subscriptionMutex.RUnlock()
	return len(fake.subscriptionArgsForCall)
}

func (fake *FakeConsumer) SubscriptionCalls(stub func() string) {
	fake.subscriptionMutex.Lock()
	defer fake.subscriptionMutex.Unlock()
	fake.SubscriptionStub = stub
}

func (fake *FakeConsumer) SubscriptionReturns(result1 string) {
	fake.subscriptionMutex.Lock()
	defer fake.subscriptionMutex.Unlock()
	fake.SubscriptionStub = nil
	fake.subscriptionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConsumer) SubscriptionReturnsOnCall(i int, result1 string) {
	fake.subscriptionMutex.Lock()
	defer fake.subscriptionMutex.Unlock()
	fake.SubscriptionStub = nil
	if fake.subscriptionReturnsOnCall == nil {
		fake.subscriptionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.subscriptionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConsumer) Unsubscribe() error {
	fake.unsubscribeMutex.Lock()
	ret, specificReturn := fake.unsubscribeReturnsOnCall[len(fake.unsubscribeArgsForCall)]
	fake.unsubscribeArgsForCall = append(fake.unsubscribeArgsForCall, struct {
	}{})
	stub := fake.UnsubscribeStub
	fakeReturns := fake.unsubscribeReturns
	fake.recordInvocation("Unsubscribe", []interface{}{})
	fake.unsubscribeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeConsumer) UnsubscribeCallCount() int {
	fake.unsubscribeMutex.RLock()
	defer fake.unsubscribeMutex.RUnlock()
	return len(fake.unsubscribeArgsForCall)
}

func (fake *FakeConsumer) UnsubscribeCalls(stub func() error) {
	fake.unsubscribeMutex.Lock()
	defer fake.unsubscribeMutex.Unlock()
	fake.UnsubscribeStub = stub
}

func (fake *FakeConsumer) UnsubscribeReturns(result1 error) {
	fake.unsubscribeMutex.Lock()
	defer fake.unsubscribeMutex.Unlock()
	fake.UnsubscribeStub = nil
	fake.unsubscribeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumer) UnsubscribeReturnsOnCall(i int, result1 error) {
	fake.unsubscribeMutex.Lock()
	defer fake.unsubscribeMutex.Unlock()
	fake.UnsubscribeStub = nil
	if fake.unsubscribeReturnsOnCall == nil {
		fake.unsubscribeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unsubscribeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeConsumer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.ackMutex.RLock()
	defer fake.ackMutex.RUnlock()
	fake.ackIDMutex.RLock()
	defer fake.ackIDMutex.RUnlock()
	fake.chanMutex.RLock()
	defer fake.chanMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.nackMutex.RLock()
	defer fake.nackMutex.RUnlock()
	fake.nackIDMutex.RLock()
	defer fake.nackIDMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.receiveMutex.RLock()
	defer fake.receiveMutex.RUnlock()
	fake.reconsumeLaterMutex.RLock()
	defer fake.reconsumeLaterMutex.RUnlock()
	fake.seekMutex.RLock()
	defer fake.seekMutex.RUnlock()
	fake.seekByTimeMutex.RLock()
	defer fake.seekByTimeMutex.RUnlock()
	fake.subscriptionMutex.RLock()
	defer fake.subscriptionMutex.RUnlock()
	fake.unsubscribeMutex.RLock()
	defer fake.unsubscribeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConsumer) recordInvocation(key string, args []interface{}) {
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

var _ pulsar.Consumer = new(FakeConsumer)
