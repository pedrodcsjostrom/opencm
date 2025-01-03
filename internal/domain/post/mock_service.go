// Code generated by mockery v2.43.2. DO NOT EDIT.

package post

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// AddSocialMediaPublisher provides a mock function with given fields: ctx, projectID, postID, publisherID
func (_m *MockService) AddSocialMediaPublisher(ctx context.Context, projectID string, postID string, publisherID string) error {
	ret := _m.Called(ctx, projectID, postID, publisherID)

	if len(ret) == 0 {
		panic("no return value specified for AddSocialMediaPublisher")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, projectID, postID, publisherID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddToProjectQueue provides a mock function with given fields: ctx, projectID, postID
func (_m *MockService) AddToProjectQueue(ctx context.Context, projectID string, postID string) error {
	ret := _m.Called(ctx, projectID, postID)

	if len(ret) == 0 {
		panic("no return value specified for AddToProjectQueue")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, projectID, postID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ArchivePost provides a mock function with given fields: ctx, id
func (_m *MockService) ArchivePost(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for ArchivePost")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreatePost provides a mock function with given fields: ctx, projectID, title, textContent, imageURLs, videoURLs, isIdea, scheduledAt
func (_m *MockService) CreatePost(ctx context.Context, projectID string, title string, textContent string, imageURLs []string, videoURLs []string, isIdea bool, scheduledAt time.Time) (*Post, error) {
	ret := _m.Called(ctx, projectID, title, textContent, imageURLs, videoURLs, isIdea, scheduledAt)

	if len(ret) == 0 {
		panic("no return value specified for CreatePost")
	}

	var r0 *Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, []string, []string, bool, time.Time) (*Post, error)); ok {
		return rf(ctx, projectID, title, textContent, imageURLs, videoURLs, isIdea, scheduledAt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, []string, []string, bool, time.Time) *Post); ok {
		r0 = rf(ctx, projectID, title, textContent, imageURLs, videoURLs, isIdea, scheduledAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, []string, []string, bool, time.Time) error); ok {
		r1 = rf(ctx, projectID, title, textContent, imageURLs, videoURLs, isIdea, scheduledAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePost provides a mock function with given fields: ctx, id
func (_m *MockService) DeletePost(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeletePost")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DequeuePostsToPublish provides a mock function with given fields: ctx, projectID
func (_m *MockService) DequeuePostsToPublish(ctx context.Context, projectID string) ([]*QPost, error) {
	ret := _m.Called(ctx, projectID)

	if len(ret) == 0 {
		panic("no return value specified for DequeuePostsToPublish")
	}

	var r0 []*QPost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*QPost, error)); ok {
		return rf(ctx, projectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*QPost); ok {
		r0 = rf(ctx, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*QPost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindScheduledReadyPosts provides a mock function with given fields: ctx, offset, chunkSize
func (_m *MockService) FindScheduledReadyPosts(ctx context.Context, offset int, chunkSize int) ([]*QPost, error) {
	ret := _m.Called(ctx, offset, chunkSize)

	if len(ret) == 0 {
		panic("no return value specified for FindScheduledReadyPosts")
	}

	var r0 []*QPost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, int) ([]*QPost, error)); ok {
		return rf(ctx, offset, chunkSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, int) []*QPost); ok {
		r0 = rf(ctx, offset, chunkSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*QPost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, offset, chunkSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPost provides a mock function with given fields: ctx, id
func (_m *MockService) GetPost(ctx context.Context, id string) (*Post, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetPost")
	}

	var r0 *Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*Post, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *Post); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProjectQueuedPosts provides a mock function with given fields: ctx, projectID
func (_m *MockService) GetProjectQueuedPosts(ctx context.Context, projectID string) ([]*Post, error) {
	ret := _m.Called(ctx, projectID)

	if len(ret) == 0 {
		panic("no return value specified for GetProjectQueuedPosts")
	}

	var r0 []*Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*Post, error)); ok {
		return rf(ctx, projectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*Post); ok {
		r0 = rf(ctx, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetQueuePost provides a mock function with given fields: ctx, id
func (_m *MockService) GetQueuePost(ctx context.Context, id string) (*QPost, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetQueuePost")
	}

	var r0 *QPost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*QPost, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *QPost); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*QPost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListProjectPosts provides a mock function with given fields: ctx, projectID
func (_m *MockService) ListProjectPosts(ctx context.Context, projectID string) ([]*Post, error) {
	ret := _m.Called(ctx, projectID)

	if len(ret) == 0 {
		panic("no return value specified for ListProjectPosts")
	}

	var r0 []*Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*Post, error)); ok {
		return rf(ctx, projectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*Post); ok {
		r0 = rf(ctx, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MovePostInQueue provides a mock function with given fields: ctx, projectID, currentIndex, newIndex
func (_m *MockService) MovePostInQueue(ctx context.Context, projectID string, currentIndex int, newIndex int) error {
	ret := _m.Called(ctx, projectID, currentIndex, newIndex)

	if len(ret) == 0 {
		panic("no return value specified for MovePostInQueue")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int) error); ok {
		r0 = rf(ctx, projectID, currentIndex, newIndex)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SchedulePost provides a mock function with given fields: ctx, id, scheduled_at
func (_m *MockService) SchedulePost(ctx context.Context, id string, scheduled_at time.Time) error {
	ret := _m.Called(ctx, id, scheduled_at)

	if len(ret) == 0 {
		panic("no return value specified for SchedulePost")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, time.Time) error); ok {
		r0 = rf(ctx, id, scheduled_at)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
