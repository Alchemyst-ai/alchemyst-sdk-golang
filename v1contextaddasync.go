// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/apijson"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/requestconfig"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/option"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/packages/param"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/packages/respjson"
)

// V1ContextAddAsyncService contains methods and other services that help with
// interacting with the alchemyst_ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContextAddAsyncService] method instead.
type V1ContextAddAsyncService struct {
	Options []option.RequestOption
	Status  V1ContextAddAsyncStatusService
}

// NewV1ContextAddAsyncService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1ContextAddAsyncService(opts ...option.RequestOption) (r V1ContextAddAsyncService) {
	r = V1ContextAddAsyncService{}
	r.Options = opts
	r.Status = NewV1ContextAddAsyncStatusService(opts...)
	return
}

// This endpoint accepts context data and queues it for asynchronous processing by
// the context processor. It returns a success or error response depending on the
// queuing result.
func (r *V1ContextAddAsyncService) New(ctx context.Context, body V1ContextAddAsyncNewParams, opts ...option.RequestOption) (res *V1ContextAddAsyncNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/add-async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Attempts to cancel a context add job by job id.
//
//   - If the job is already completed or failed, returns 404.
//   - If the job is currently running ("active"), returns 409 and cannot be
//     cancelled.
//   - Only jobs in "waiting" or "delayed" state can be cancelled.
func (r *V1ContextAddAsyncService) Cancel(ctx context.Context, id string, opts ...option.RequestOption) (res *V1ContextAddAsyncCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("api/v1/context/add-async/%s/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type V1ContextAddAsyncNewResponse struct {
	JobID  string `json:"jobId,required"`
	Queued bool   `json:"queued,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		Queued      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextAddAsyncNewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextAddAsyncNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextAddAsyncCancelResponse struct {
	JobID   string `json:"jobId,required"`
	Message string `json:"message,required"`
	Status  string `json:"status,required"`
	Success bool   `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		Message     respjson.Field
		Status      respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextAddAsyncCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextAddAsyncCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextAddAsyncNewParams struct {
	// Type of context being added
	//
	// Any of "resource", "conversation", "instruction".
	ContextType V1ContextAddAsyncNewParamsContextType `json:"context_type,omitzero,required"`
	// Array of documents with content and additional metadata
	Documents []V1ContextAddAsyncNewParamsDocument `json:"documents,omitzero,required"`
	// Scope of the context
	//
	// Any of "internal", "external".
	Scope V1ContextAddAsyncNewParamsScope `json:"scope,omitzero,required"`
	// The source of the context data
	Source string `json:"source,required"`
	// Additional metadata for the context
	Metadata V1ContextAddAsyncNewParamsMetadata `json:"metadata,omitzero"`
	paramObj
}

func (r V1ContextAddAsyncNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextAddAsyncNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContextAddAsyncNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of context being added
type V1ContextAddAsyncNewParamsContextType string

const (
	V1ContextAddAsyncNewParamsContextTypeResource     V1ContextAddAsyncNewParamsContextType = "resource"
	V1ContextAddAsyncNewParamsContextTypeConversation V1ContextAddAsyncNewParamsContextType = "conversation"
	V1ContextAddAsyncNewParamsContextTypeInstruction  V1ContextAddAsyncNewParamsContextType = "instruction"
)

type V1ContextAddAsyncNewParamsDocument struct {
	// The content of the document
	Content     param.Opt[string] `json:"content,omitzero"`
	ExtraFields map[string]string `json:"-"`
	paramObj
}

func (r V1ContextAddAsyncNewParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextAddAsyncNewParamsDocument
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *V1ContextAddAsyncNewParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scope of the context
type V1ContextAddAsyncNewParamsScope string

const (
	V1ContextAddAsyncNewParamsScopeInternal V1ContextAddAsyncNewParamsScope = "internal"
	V1ContextAddAsyncNewParamsScopeExternal V1ContextAddAsyncNewParamsScope = "external"
)

// Additional metadata for the context
type V1ContextAddAsyncNewParamsMetadata struct {
	// Name of the file
	FileName param.Opt[string] `json:"fileName,omitzero"`
	// Size of the file in bytes
	FileSize param.Opt[float64] `json:"fileSize,omitzero"`
	// Type/MIME of the file
	FileType param.Opt[string] `json:"fileType,omitzero"`
	// Last modified timestamp
	LastModified param.Opt[string] `json:"lastModified,omitzero"`
	// Array of Group Name to which the file belongs to
	GroupName []string `json:"groupName,omitzero"`
	paramObj
}

func (r V1ContextAddAsyncNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextAddAsyncNewParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContextAddAsyncNewParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
