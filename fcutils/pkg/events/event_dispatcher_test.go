package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handle          TestEventHandler
	handle2         TestEventHandler
	handle3         TestEventHandler
	eventDispatcher *EventDispatcher
}

func (s *EventDispatcherTestSuite) SetupTest() {
	s.eventDispatcher = NewEventDispatcher()
	s.handle = TestEventHandler{ID: 1}
	s.handle2 = TestEventHandler{ID: 2}
	s.handle3 = TestEventHandler{ID: 3}
	s.event = TestEvent{
		Name:    "test",
		Payload: "test",
	}
	s.event2 = TestEvent{
		Name:    "test2",
		Payload: "test2",
	}
}
func (s *EventDispatcherTestSuite) TestEventDispatchRegister() {
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handle)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handle2)
	s.Nil(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	assert.Equal(s.T(), &s.handle, s.eventDispatcher.handlers[s.event.GetName()][0])
	assert.Equal(s.T(), &s.handle2, s.eventDispatcher.handlers[s.event.GetName()][1])

}

func (s *EventDispatcherTestSuite) TestEventDispatchRegisterWithSameHandler() {
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handle)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))

	err = s.eventDispatcher.Register(s.event.GetName(), &s.handle)
	s.Equal(ErrHandlerAlreadyRegistered, err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
}

func (s *EventDispatcherTestSuite) TestEventDispatchClear() {

	// event 1
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handle)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
	err = s.eventDispatcher.Register(s.event.GetName(), &s.handle2)
	s.Nil(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	// event 2
	err = s.eventDispatcher.Register(s.event2.GetName(), &s.handle3)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event2.GetName()]))

	s.eventDispatcher.Clear()
	s.Equal(0, len(s.eventDispatcher.handlers))
}

func (s *EventDispatcherTestSuite) TestEventDispatchHas() {
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handle)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
	err = s.eventDispatcher.Register(s.event.GetName(), &s.handle2)
	s.Nil(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	assert.True(s.T(), s.eventDispatcher.Has(s.event.GetName(), &s.handle))
	assert.True(s.T(), s.eventDispatcher.Has(s.event.GetName(), &s.handle2))
	assert.False(s.T(), s.eventDispatcher.Has("test3", &s.handle))
	assert.False(s.T(), s.eventDispatcher.Has(s.event.GetName(), &s.handle3))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (s *EventDispatcherTestSuite) TestEventDispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &s.event)

	eh2 := &MockHandler{}
	eh2.On("Handle", &s.event)

	s.eventDispatcher.Register(s.event.GetName(), eh)
	s.eventDispatcher.Register(s.event.GetName(), eh2)

	s.eventDispatcher.Dispatch(&s.event)
	eh.AssertExpectations(s.T())
	eh2.AssertExpectations(s.T())
	eh.AssertNumberOfCalls(s.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(s.T(), "Handle", 1)

}

func (s *EventDispatcherTestSuite) TestEventDispatchRemove() {

	// event 1
	err := s.eventDispatcher.Register(s.event.GetName(), &s.handle)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
	err = s.eventDispatcher.Register(s.event.GetName(), &s.handle2)
	s.Nil(err)
	s.Equal(2, len(s.eventDispatcher.handlers[s.event.GetName()]))

	// event 2
	err = s.eventDispatcher.Register(s.event2.GetName(), &s.handle3)
	s.Nil(err)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event2.GetName()]))

	s.eventDispatcher.Remove(s.event.GetName(), &s.handle)
	s.Equal(1, len(s.eventDispatcher.handlers[s.event.GetName()]))
	assert.Equal(s.T(), &s.handle2, s.eventDispatcher.handlers[s.event.GetName()][0])

	s.eventDispatcher.Remove(s.event.GetName(), &s.handle2)
	s.Equal(0, len(s.eventDispatcher.handlers[s.event.GetName()]))

	s.eventDispatcher.Remove(s.event.GetName(), &s.handle3)
	s.Equal(0, len(s.eventDispatcher.handlers[s.event.GetName()]))

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
