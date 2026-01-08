// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/apijson"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/apiquery"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/requestconfig"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/option"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/packages/param"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/packages/respjson"
)

// V1ContextAddAsyncStatusService contains methods and other services that help
// with interacting with the alchemyst_ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContextAddAsyncStatusService] method instead.
type V1ContextAddAsyncStatusService struct {
	Options []option.RequestOption
}

// NewV1ContextAddAsyncStatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1ContextAddAsyncStatusService(opts ...option.RequestOption) (r V1ContextAddAsyncStatusService) {
	r = V1ContextAddAsyncStatusService{}
	r.Options = opts
	return
}

// Returns the status and result of a context add job by job id.
func (r *V1ContextAddAsyncStatusService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *V1ContextAddAsyncStatusGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/context/add-async/%s/status", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all jobs (active, waiting, delayed, failed, completed) belonging to the
// authenticated user.
func (r *V1ContextAddAsyncStatusService) List(ctx context.Context, query V1ContextAddAsyncStatusListParams, opts ...option.RequestOption) (res *V1ContextAddAsyncStatusListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/add-async/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type V1ContextAddAsyncStatusGetResponse struct {
	JobID        string  `json:"jobId,required"`
	Status       string  `json:"status,required"`
	Success      bool    `json:"success,required"`
	AttemptsMade float64 `json:"attemptsMade"`
	FailedReason string  `json:"failedReason"`
	FinishedOn   float64 `json:"finishedOn"`
	ProcessedOn  float64 `json:"processedOn"`
	// Result of the job (if available)
	Result any `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID        respjson.Field
		Status       respjson.Field
		Success      respjson.Field
		AttemptsMade respjson.Field
		FailedReason respjson.Field
		FinishedOn   respjson.Field
		ProcessedOn  respjson.Field
		Result       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextAddAsyncStatusGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextAddAsyncStatusGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextAddAsyncStatusListResponse struct {
	Jobs    []V1ContextAddAsyncStatusListResponseJob `json:"jobs,required"`
	Success bool                                     `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Jobs        respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextAddAsyncStatusListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextAddAsyncStatusListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextAddAsyncStatusListResponseJob struct {
	AttemptsMade float64 `json:"attemptsMade,required"`
	Data         any     `json:"data,required"`
	JobID        string  `json:"jobId,required"`
	Status       string  `json:"status,required"`
	FailedReason string  `json:"failedReason"`
	FinishedOn   float64 `json:"finishedOn"`
	ProcessedOn  float64 `json:"processedOn"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttemptsMade respjson.Field
		Data         respjson.Field
		JobID        respjson.Field
		Status       respjson.Field
		FailedReason respjson.Field
		FinishedOn   respjson.Field
		ProcessedOn  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextAddAsyncStatusListResponseJob) RawJSON() string { return r.JSON.raw }
func (r *V1ContextAddAsyncStatusListResponseJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextAddAsyncStatusListParams struct {
	// Maximum number of jobs to return
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	// Number of jobs to skip before starting to collect the result set
	Offset param.Opt[string] `query:"offset,omitzero" json:"-"`
	// Type of jobs to list
	//
	// Any of "all", "active", "failed", "completed".
	Type V1ContextAddAsyncStatusListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1ContextAddAsyncStatusListParams]'s query parameters as
// `url.Values`.
func (r V1ContextAddAsyncStatusListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Type of jobs to list
type V1ContextAddAsyncStatusListParamsType string

const (
	V1ContextAddAsyncStatusListParamsTypeAll       V1ContextAddAsyncStatusListParamsType = "all"
	V1ContextAddAsyncStatusListParamsTypeActive    V1ContextAddAsyncStatusListParamsType = "active"
	V1ContextAddAsyncStatusListParamsTypeFailed    V1ContextAddAsyncStatusListParamsType = "failed"
	V1ContextAddAsyncStatusListParamsTypeCompleted V1ContextAddAsyncStatusListParamsType = "completed"
)
