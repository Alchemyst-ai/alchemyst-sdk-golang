// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/apijson"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/apiquery"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/requestconfig"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/option"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/packages/param"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/packages/respjson"
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

// Returns paginated traces for the authenticated user within their organization.
func (r *V1ContextTraceService) List(ctx context.Context, query V1ContextTraceListParams, opts ...option.RequestOption) (res *V1ContextTraceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/traces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a data trace for the authenticated user with the specified trace ID.
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
	Pagination V1ContextTraceListResponsePagination `json:"pagination,required"`
	Traces     []V1ContextTraceListResponseTrace    `json:"traces,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pagination  respjson.Field
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

type V1ContextTraceListResponsePagination struct {
	HasNextPage bool  `json:"hasNextPage,required"`
	HasPrevPage bool  `json:"hasPrevPage,required"`
	Limit       int64 `json:"limit,required"`
	Page        int64 `json:"page,required"`
	Total       int64 `json:"total,required"`
	TotalPages  int64 `json:"totalPages,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNextPage respjson.Field
		HasPrevPage respjson.Field
		Limit       respjson.Field
		Page        respjson.Field
		Total       respjson.Field
		TotalPages  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextTraceListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *V1ContextTraceListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextTraceListResponseTrace struct {
	ID             string    `json:"_id,required"`
	CreatedAt      time.Time `json:"createdAt,required" format:"date-time"`
	Data           any       `json:"data,required"`
	OrganizationID string    `json:"organizationId,required"`
	Type           string    `json:"type,required"`
	UpdatedAt      time.Time `json:"updatedAt,required" format:"date-time"`
	UserID         string    `json:"userId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		OrganizationID respjson.Field
		Type           respjson.Field
		UpdatedAt      respjson.Field
		UserID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextTraceListResponseTrace) RawJSON() string { return r.JSON.raw }
func (r *V1ContextTraceListResponseTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextTraceDeleteResponse struct {
	// The deleted trace data
	Trace V1ContextTraceDeleteResponseTrace `json:"trace,required"`
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

// The deleted trace data
type V1ContextTraceDeleteResponseTrace struct {
	ID             string                                `json:"_id"`
	CreatedAt      time.Time                             `json:"createdAt" format:"date-time"`
	Data           V1ContextTraceDeleteResponseTraceData `json:"data"`
	OrganizationID string                                `json:"organizationId"`
	Type           string                                `json:"type"`
	UpdatedAt      time.Time                             `json:"updatedAt" format:"date-time"`
	UserID         string                                `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Data           respjson.Field
		OrganizationID respjson.Field
		Type           respjson.Field
		UpdatedAt      respjson.Field
		UserID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextTraceDeleteResponseTrace) RawJSON() string { return r.JSON.raw }
func (r *V1ContextTraceDeleteResponseTrace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextTraceDeleteResponseTraceData struct {
	FileName string `json:"fileName"`
	Query    string `json:"query"`
	Source   string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileName    respjson.Field
		Query       respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextTraceDeleteResponseTraceData) RawJSON() string { return r.JSON.raw }
func (r *V1ContextTraceDeleteResponseTraceData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextTraceListParams struct {
	// Number of traces per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page number for pagination
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1ContextTraceListParams]'s query parameters as
// `url.Values`.
func (r V1ContextTraceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
