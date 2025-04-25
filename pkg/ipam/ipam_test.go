package ipam

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/containernetworking/cni/pkg/invoke"
	"github.com/containernetworking/cni/pkg/types"
	current "github.com/containernetworking/cni/pkg/types/100"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockDelegate is a mock implementation of the Delegate interface
type MockDelegate struct {
	mock.Mock
}

func (m *MockDelegate) DelegateAdd(ctx context.Context, plugin string, netconf []byte, opts ...invoke.Option) (types.Result, error) {
	args := m.Called(ctx, plugin, netconf, opts)
	return args.Get(0).(types.Result), args.Error(1)
}

func (m *MockDelegate) DelegateCheck(ctx context.Context, plugin string, netconf []byte, opts ...invoke.Option) error {
	args := m.Called(ctx, plugin, netconf, opts)
	return args.Error(0)
}

func (m *MockDelegate) DelegateDel(ctx context.Context, plugin string, netconf []byte, opts ...invoke.Option) error {
	args := m.Called(ctx, plugin, netconf, opts)
	return args.Error(0)
}

func (m *MockDelegate) DelegateStatus(ctx context.Context, plugin string, netconf []byte, opts ...invoke.Option) error {
	args := m.Called(ctx, plugin, netconf, opts)
	return args.Error(0)
}

func TestExecAdd_Success(t *testing.T) {
	mockDelegate := new(MockDelegate)
	mockDelegate.On("DelegateAdd", mock.Anything, "test-plugin", mock.Anything, mock.Anything).Return(&current.Result{}, nil)

	invoke.DefaultDelegate = mockDelegate
	result, err := ExecAdd("test-plugin", []byte("{}"))

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestExecAdd_Failure(t *testing.T) {
	mockDelegate := new(MockDelegate)
	mockDelegate.On("DelegateAdd", mock.Anything, "test-plugin", mock.Anything, mock.Anything).Return(nil, errors.New("delegate failed"))

	invoke.DefaultDelegate = mockDelegate
	result, err := ExecAdd("test-plugin", []byte("{}"))

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestConfigureIface_NoInterfaces(t *testing.T) {
	err := ConfigureIface("eth0", &current.Result{})
	assert.Error(t, err)
	assert.Equal(t, "no interfaces to configure", err.Error())
}