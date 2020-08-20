// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	openapi "github.com/tendermint/cosmos-rosetta-gateway/cosmos/launchpad/client/tendermint/generated"
)

// TendermintInfoAPI is an autogenerated mock type for the TendermintInfoAPI type
type TendermintInfoAPI struct {
	mock.Mock
}

// Block provides a mock function with given fields: ctx, localVarOptionals
func (_m *TendermintInfoAPI) Block(ctx context.Context, localVarOptionals *openapi.BlockOpts) (openapi.BlockResponse, *http.Response, error) {
	ret := _m.Called(ctx, localVarOptionals)

	var r0 openapi.BlockResponse
	if rf, ok := ret.Get(0).(func(context.Context, *openapi.BlockOpts) openapi.BlockResponse); ok {
		r0 = rf(ctx, localVarOptionals)
	} else {
		r0 = ret.Get(0).(openapi.BlockResponse)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, *openapi.BlockOpts) *http.Response); ok {
		r1 = rf(ctx, localVarOptionals)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *openapi.BlockOpts) error); ok {
		r2 = rf(ctx, localVarOptionals)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// BlockByHash provides a mock function with given fields: ctx, hash
func (_m *TendermintInfoAPI) BlockByHash(ctx context.Context, hash string) (openapi.BlockResponse, *http.Response, error) {
	ret := _m.Called(ctx, hash)

	var r0 openapi.BlockResponse
	if rf, ok := ret.Get(0).(func(context.Context, string) openapi.BlockResponse); ok {
		r0 = rf(ctx, hash)
	} else {
		r0 = ret.Get(0).(openapi.BlockResponse)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string) *http.Response); ok {
		r1 = rf(ctx, hash)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, hash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NetInfo provides a mock function with given fields: ctx
func (_m *TendermintInfoAPI) NetInfo(ctx context.Context) (openapi.NetInfoResponse, *http.Response, error) {
	ret := _m.Called(ctx)

	var r0 openapi.NetInfoResponse
	if rf, ok := ret.Get(0).(func(context.Context) openapi.NetInfoResponse); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(openapi.NetInfoResponse)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context) *http.Response); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UnconfirmedTxs provides a mock function with given fields: ctx, localVarOptionals
func (_m *TendermintInfoAPI) UnconfirmedTxs(ctx context.Context, localVarOptionals *openapi.UnconfirmedTxsOpts) (openapi.UnconfirmedTransactionsResponse, *http.Response, error) {
	ret := _m.Called(ctx, localVarOptionals)

	var r0 openapi.UnconfirmedTransactionsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *openapi.UnconfirmedTxsOpts) openapi.UnconfirmedTransactionsResponse); ok {
		r0 = rf(ctx, localVarOptionals)
	} else {
		r0 = ret.Get(0).(openapi.UnconfirmedTransactionsResponse)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, *openapi.UnconfirmedTxsOpts) *http.Response); ok {
		r1 = rf(ctx, localVarOptionals)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *openapi.UnconfirmedTxsOpts) error); ok {
		r2 = rf(ctx, localVarOptionals)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TxSearch provides a mock function with given fields: ctx, query, localVarOptionals
func (_m *TendermintInfoAPI) TxSearch(ctx context.Context, query string, localVarOptionals *openapi.TxSearchOpts) (openapi.TxSearchResponse, *http.Response, error) {
	ret := _m.Called(ctx, query, localVarOptionals)

	var r0 openapi.TxSearchResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, *openapi.TxSearchOpts) openapi.TxSearchResponse); ok {
		r0 = rf(ctx, query, localVarOptionals)
	} else {
		r0 = ret.Get(0).(openapi.TxSearchResponse)
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(context.Context, string, *openapi.TxSearchOpts) *http.Response); ok {
		r1 = rf(ctx, query, localVarOptionals)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, *openapi.TxSearchOpts) error); ok {
		r2 = rf(ctx, query, localVarOptionals)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
