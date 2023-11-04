package mock

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

import (
	"context"
	"route256/loms/internal/models"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// StatusNotifierMock implements order.StatusNotifier
type StatusNotifierMock struct {
	t minimock.Tester

	funcNotifyOrderStatus          func(ctx context.Context, m models.OrderID, status models.Status) (err error)
	inspectFuncNotifyOrderStatus   func(ctx context.Context, m models.OrderID, status models.Status)
	afterNotifyOrderStatusCounter  uint64
	beforeNotifyOrderStatusCounter uint64
	NotifyOrderStatusMock          mStatusNotifierMockNotifyOrderStatus
}

// NewStatusNotifierMock returns a mock for order.StatusNotifier
func NewStatusNotifierMock(t minimock.Tester) *StatusNotifierMock {
	m := &StatusNotifierMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.NotifyOrderStatusMock = mStatusNotifierMockNotifyOrderStatus{mock: m}
	m.NotifyOrderStatusMock.callArgs = []*StatusNotifierMockNotifyOrderStatusParams{}

	return m
}

type mStatusNotifierMockNotifyOrderStatus struct {
	mock               *StatusNotifierMock
	defaultExpectation *StatusNotifierMockNotifyOrderStatusExpectation
	expectations       []*StatusNotifierMockNotifyOrderStatusExpectation

	callArgs []*StatusNotifierMockNotifyOrderStatusParams
	mutex    sync.RWMutex
}

// StatusNotifierMockNotifyOrderStatusExpectation specifies expectation struct of the StatusNotifier.NotifyOrderStatus
type StatusNotifierMockNotifyOrderStatusExpectation struct {
	mock    *StatusNotifierMock
	params  *StatusNotifierMockNotifyOrderStatusParams
	results *StatusNotifierMockNotifyOrderStatusResults
	Counter uint64
}

// StatusNotifierMockNotifyOrderStatusParams contains parameters of the StatusNotifier.NotifyOrderStatus
type StatusNotifierMockNotifyOrderStatusParams struct {
	ctx    context.Context
	m      models.OrderID
	status models.Status
}

// StatusNotifierMockNotifyOrderStatusResults contains results of the StatusNotifier.NotifyOrderStatus
type StatusNotifierMockNotifyOrderStatusResults struct {
	err error
}

// Expect sets up expected params for StatusNotifier.NotifyOrderStatus
func (mmNotifyOrderStatus *mStatusNotifierMockNotifyOrderStatus) Expect(ctx context.Context, m models.OrderID, status models.Status) *mStatusNotifierMockNotifyOrderStatus {
	if mmNotifyOrderStatus.mock.funcNotifyOrderStatus != nil {
		mmNotifyOrderStatus.mock.t.Fatalf("StatusNotifierMock.NotifyOrderStatus mock is already set by Set")
	}

	if mmNotifyOrderStatus.defaultExpectation == nil {
		mmNotifyOrderStatus.defaultExpectation = &StatusNotifierMockNotifyOrderStatusExpectation{}
	}

	mmNotifyOrderStatus.defaultExpectation.params = &StatusNotifierMockNotifyOrderStatusParams{ctx, m, status}
	for _, e := range mmNotifyOrderStatus.expectations {
		if minimock.Equal(e.params, mmNotifyOrderStatus.defaultExpectation.params) {
			mmNotifyOrderStatus.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmNotifyOrderStatus.defaultExpectation.params)
		}
	}

	return mmNotifyOrderStatus
}

// Inspect accepts an inspector function that has same arguments as the StatusNotifier.NotifyOrderStatus
func (mmNotifyOrderStatus *mStatusNotifierMockNotifyOrderStatus) Inspect(f func(ctx context.Context, m models.OrderID, status models.Status)) *mStatusNotifierMockNotifyOrderStatus {
	if mmNotifyOrderStatus.mock.inspectFuncNotifyOrderStatus != nil {
		mmNotifyOrderStatus.mock.t.Fatalf("Inspect function is already set for StatusNotifierMock.NotifyOrderStatus")
	}

	mmNotifyOrderStatus.mock.inspectFuncNotifyOrderStatus = f

	return mmNotifyOrderStatus
}

// Return sets up results that will be returned by StatusNotifier.NotifyOrderStatus
func (mmNotifyOrderStatus *mStatusNotifierMockNotifyOrderStatus) Return(err error) *StatusNotifierMock {
	if mmNotifyOrderStatus.mock.funcNotifyOrderStatus != nil {
		mmNotifyOrderStatus.mock.t.Fatalf("StatusNotifierMock.NotifyOrderStatus mock is already set by Set")
	}

	if mmNotifyOrderStatus.defaultExpectation == nil {
		mmNotifyOrderStatus.defaultExpectation = &StatusNotifierMockNotifyOrderStatusExpectation{mock: mmNotifyOrderStatus.mock}
	}
	mmNotifyOrderStatus.defaultExpectation.results = &StatusNotifierMockNotifyOrderStatusResults{err}
	return mmNotifyOrderStatus.mock
}

// Set uses given function f to mock the StatusNotifier.NotifyOrderStatus method
func (mmNotifyOrderStatus *mStatusNotifierMockNotifyOrderStatus) Set(f func(ctx context.Context, m models.OrderID, status models.Status) (err error)) *StatusNotifierMock {
	if mmNotifyOrderStatus.defaultExpectation != nil {
		mmNotifyOrderStatus.mock.t.Fatalf("Default expectation is already set for the StatusNotifier.NotifyOrderStatus method")
	}

	if len(mmNotifyOrderStatus.expectations) > 0 {
		mmNotifyOrderStatus.mock.t.Fatalf("Some expectations are already set for the StatusNotifier.NotifyOrderStatus method")
	}

	mmNotifyOrderStatus.mock.funcNotifyOrderStatus = f
	return mmNotifyOrderStatus.mock
}

// When sets expectation for the StatusNotifier.NotifyOrderStatus which will trigger the result defined by the following
// Then helper
func (mmNotifyOrderStatus *mStatusNotifierMockNotifyOrderStatus) When(ctx context.Context, m models.OrderID, status models.Status) *StatusNotifierMockNotifyOrderStatusExpectation {
	if mmNotifyOrderStatus.mock.funcNotifyOrderStatus != nil {
		mmNotifyOrderStatus.mock.t.Fatalf("StatusNotifierMock.NotifyOrderStatus mock is already set by Set")
	}

	expectation := &StatusNotifierMockNotifyOrderStatusExpectation{
		mock:   mmNotifyOrderStatus.mock,
		params: &StatusNotifierMockNotifyOrderStatusParams{ctx, m, status},
	}
	mmNotifyOrderStatus.expectations = append(mmNotifyOrderStatus.expectations, expectation)
	return expectation
}

// Then sets up StatusNotifier.NotifyOrderStatus return parameters for the expectation previously defined by the When method
func (e *StatusNotifierMockNotifyOrderStatusExpectation) Then(err error) *StatusNotifierMock {
	e.results = &StatusNotifierMockNotifyOrderStatusResults{err}
	return e.mock
}

// NotifyOrderStatus implements order.StatusNotifier
func (mmNotifyOrderStatus *StatusNotifierMock) NotifyOrderStatus(ctx context.Context, m models.OrderID, status models.Status) (err error) {
	mm_atomic.AddUint64(&mmNotifyOrderStatus.beforeNotifyOrderStatusCounter, 1)
	defer mm_atomic.AddUint64(&mmNotifyOrderStatus.afterNotifyOrderStatusCounter, 1)

	if mmNotifyOrderStatus.inspectFuncNotifyOrderStatus != nil {
		mmNotifyOrderStatus.inspectFuncNotifyOrderStatus(ctx, m, status)
	}

	mm_params := &StatusNotifierMockNotifyOrderStatusParams{ctx, m, status}

	// Record call args
	mmNotifyOrderStatus.NotifyOrderStatusMock.mutex.Lock()
	mmNotifyOrderStatus.NotifyOrderStatusMock.callArgs = append(mmNotifyOrderStatus.NotifyOrderStatusMock.callArgs, mm_params)
	mmNotifyOrderStatus.NotifyOrderStatusMock.mutex.Unlock()

	for _, e := range mmNotifyOrderStatus.NotifyOrderStatusMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmNotifyOrderStatus.NotifyOrderStatusMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmNotifyOrderStatus.NotifyOrderStatusMock.defaultExpectation.Counter, 1)
		mm_want := mmNotifyOrderStatus.NotifyOrderStatusMock.defaultExpectation.params
		mm_got := StatusNotifierMockNotifyOrderStatusParams{ctx, m, status}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmNotifyOrderStatus.t.Errorf("StatusNotifierMock.NotifyOrderStatus got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmNotifyOrderStatus.NotifyOrderStatusMock.defaultExpectation.results
		if mm_results == nil {
			mmNotifyOrderStatus.t.Fatal("No results are set for the StatusNotifierMock.NotifyOrderStatus")
		}
		return (*mm_results).err
	}
	if mmNotifyOrderStatus.funcNotifyOrderStatus != nil {
		return mmNotifyOrderStatus.funcNotifyOrderStatus(ctx, m, status)
	}
	mmNotifyOrderStatus.t.Fatalf("Unexpected call to StatusNotifierMock.NotifyOrderStatus. %v %v %v", ctx, m, status)
	return
}

// NotifyOrderStatusAfterCounter returns a count of finished StatusNotifierMock.NotifyOrderStatus invocations
func (mmNotifyOrderStatus *StatusNotifierMock) NotifyOrderStatusAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmNotifyOrderStatus.afterNotifyOrderStatusCounter)
}

// NotifyOrderStatusBeforeCounter returns a count of StatusNotifierMock.NotifyOrderStatus invocations
func (mmNotifyOrderStatus *StatusNotifierMock) NotifyOrderStatusBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmNotifyOrderStatus.beforeNotifyOrderStatusCounter)
}

// Calls returns a list of arguments used in each call to StatusNotifierMock.NotifyOrderStatus.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmNotifyOrderStatus *mStatusNotifierMockNotifyOrderStatus) Calls() []*StatusNotifierMockNotifyOrderStatusParams {
	mmNotifyOrderStatus.mutex.RLock()

	argCopy := make([]*StatusNotifierMockNotifyOrderStatusParams, len(mmNotifyOrderStatus.callArgs))
	copy(argCopy, mmNotifyOrderStatus.callArgs)

	mmNotifyOrderStatus.mutex.RUnlock()

	return argCopy
}

// MinimockNotifyOrderStatusDone returns true if the count of the NotifyOrderStatus invocations corresponds
// the number of defined expectations
func (m *StatusNotifierMock) MinimockNotifyOrderStatusDone() bool {
	for _, e := range m.NotifyOrderStatusMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NotifyOrderStatusMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNotifyOrderStatusCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcNotifyOrderStatus != nil && mm_atomic.LoadUint64(&m.afterNotifyOrderStatusCounter) < 1 {
		return false
	}
	return true
}

// MinimockNotifyOrderStatusInspect logs each unmet expectation
func (m *StatusNotifierMock) MinimockNotifyOrderStatusInspect() {
	for _, e := range m.NotifyOrderStatusMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to StatusNotifierMock.NotifyOrderStatus with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.NotifyOrderStatusMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterNotifyOrderStatusCounter) < 1 {
		if m.NotifyOrderStatusMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to StatusNotifierMock.NotifyOrderStatus")
		} else {
			m.t.Errorf("Expected call to StatusNotifierMock.NotifyOrderStatus with params: %#v", *m.NotifyOrderStatusMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcNotifyOrderStatus != nil && mm_atomic.LoadUint64(&m.afterNotifyOrderStatusCounter) < 1 {
		m.t.Error("Expected call to StatusNotifierMock.NotifyOrderStatus")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *StatusNotifierMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockNotifyOrderStatusInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *StatusNotifierMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *StatusNotifierMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockNotifyOrderStatusDone()
}
