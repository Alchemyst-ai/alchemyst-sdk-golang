// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystaisdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/alchemyst-ai-sdk-go/internal/apijson"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/option"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/packages/respjson"
)

// V1ContextTraceService contains methods and other services that help with
// interacting with the alchemyst_ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContextTraceService] method instead.
type V1ContextTraceService struct {
	Options []option.RequestOption
}

// NewV1ContextTraceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ContextTraceService(opts ...option.RequestOption) (r V1ContextTraceService) {
	r = V1ContextTraceService{}
	r.Options = opts
	return
}

// Retrieves a list of traces for the authenticated user
func (r *V1ContextTraceService) List(ctx context.Context, opts ...option.RequestOption) (res *V1ContextTraceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/traces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a data trace for the authenticated user with the specified trace ID
func (r *V1ContextTraceService) Delete(ctx context.Context, traceID string, opts ...option.RequestOption) (res *V1ContextTraceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if traceID == "" {
		err = errors.New("missing required traceId parameter")
		return
	}
	path := fmt.Sprintf("api/v1/context/traces/%s/delete", traceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type V1ContextTraceListResponse struct {
	Traces []V1ContextTraceListResponseTrace `json:"traces"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Traces      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextTraceListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextTraceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextTraceListResponseTrace struct {
	ID        string    `json:"_id"`
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	Data      any       `json:"data"`
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	UserID    string    `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Data        respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextTraceListResponseTrace) RawJSON() string { return r.JSON.raw }
func (r *V1ContextTraceListResponseTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextTraceDeleteResponse struct {
	// The deleted trace data
	Trace any `json:"trace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Trace       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextTraceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextTraceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
