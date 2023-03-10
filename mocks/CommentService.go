// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	comment "social-media-app/features/comment"

	mock "github.com/stretchr/testify/mock"
)

// CommentService is an autogenerated mock type for the CommentService type
type CommentService struct {
	mock.Mock
}

// Add provides a mock function with given fields: newComment, PostID, token
func (_m *CommentService) Add(newComment comment.Core, PostID uint, token interface{}) (comment.Core, error) {
	ret := _m.Called(newComment, PostID, token)

	var r0 comment.Core
	if rf, ok := ret.Get(0).(func(comment.Core, uint, interface{}) comment.Core); ok {
		r0 = rf(newComment, PostID, token)
	} else {
		r0 = ret.Get(0).(comment.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(comment.Core, uint, interface{}) error); ok {
		r1 = rf(newComment, PostID, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: commentID, token
func (_m *CommentService) Delete(commentID uint, token interface{}) error {
	ret := _m.Called(commentID, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, interface{}) error); ok {
		r0 = rf(commentID, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListComments provides a mock function with given fields: PostID
func (_m *CommentService) ListComments(PostID uint) ([]comment.Core, error) {
	ret := _m.Called(PostID)

	var r0 []comment.Core
	if rf, ok := ret.Get(0).(func(uint) []comment.Core); ok {
		r0 = rf(PostID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]comment.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(PostID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: newComment, commentID, token
func (_m *CommentService) Update(newComment comment.Core, commentID uint, token interface{}) (comment.Core, error) {
	ret := _m.Called(newComment, commentID, token)

	var r0 comment.Core
	if rf, ok := ret.Get(0).(func(comment.Core, uint, interface{}) comment.Core); ok {
		r0 = rf(newComment, commentID, token)
	} else {
		r0 = ret.Get(0).(comment.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(comment.Core, uint, interface{}) error); ok {
		r1 = rf(newComment, commentID, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCommentService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentService creates a new instance of CommentService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentService(t mockConstructorTestingTNewCommentService) *CommentService {
	mock := &CommentService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
