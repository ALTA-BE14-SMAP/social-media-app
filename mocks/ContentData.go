// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	content "social-media-app/features/content"

	mock "github.com/stretchr/testify/mock"
)

// ContentData is an autogenerated mock type for the ContentData type
type ContentData struct {
	mock.Mock
}

// Add provides a mock function with given fields: newContent, id
func (_m *ContentData) Add(newContent content.CoreContent, id uint) (content.CoreContent, error) {
	ret := _m.Called(newContent, id)

	var r0 content.CoreContent
	if rf, ok := ret.Get(0).(func(content.CoreContent, uint) content.CoreContent); ok {
		r0 = rf(newContent, id)
	} else {
		r0 = ret.Get(0).(content.CoreContent)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(content.CoreContent, uint) error); ok {
		r1 = rf(newContent, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: userId, contentId
func (_m *ContentData) Delete(userId uint, contentId uint) error {
	ret := _m.Called(userId, contentId)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint, uint) error); ok {
		r0 = rf(userId, contentId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *ContentData) GetAll() ([]content.CoreContent, error) {
	ret := _m.Called()

	var r0 []content.CoreContent
	if rf, ok := ret.Get(0).(func() []content.CoreContent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]content.CoreContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: userId, tes
func (_m *ContentData) GetById(userId uint, tes uint) ([]content.CoreContent, error) {
	ret := _m.Called(userId, tes)

	var r0 []content.CoreContent
	if rf, ok := ret.Get(0).(func(uint, uint) []content.CoreContent); ok {
		r0 = rf(userId, tes)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]content.CoreContent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint) error); ok {
		r1 = rf(userId, tes)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userId, contentId, updatedData
func (_m *ContentData) Update(userId uint, contentId uint, updatedData content.CoreContent) (content.CoreContent, error) {
	ret := _m.Called(userId, contentId, updatedData)

	var r0 content.CoreContent
	if rf, ok := ret.Get(0).(func(uint, uint, content.CoreContent) content.CoreContent); ok {
		r0 = rf(userId, contentId, updatedData)
	} else {
		r0 = ret.Get(0).(content.CoreContent)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint, uint, content.CoreContent) error); ok {
		r1 = rf(userId, contentId, updatedData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewContentData interface {
	mock.TestingT
	Cleanup(func())
}

// NewContentData creates a new instance of ContentData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContentData(t mockConstructorTestingTNewContentData) *ContentData {
	mock := &ContentData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
