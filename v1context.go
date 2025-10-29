// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystaisdk

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/alchemyst-ai-sdk-go/internal/apijson"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/option"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/packages/param"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/packages/respjson"
)

// V1ContextService contains methods and other services that help with interacting
// with the alchemyst_ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContextService] method instead.
type V1ContextService struct {
	Options []option.RequestOption
	Traces  V1ContextTraceService
	View    V1ContextViewService
	Memory  V1ContextMemoryService
}

// NewV1ContextService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ContextService(opts ...option.RequestOption) (r V1ContextService) {
	r = V1ContextService{}
	r.Options = opts
	r.Traces = NewV1ContextTraceService(opts...)
	r.View = NewV1ContextViewService(opts...)
	r.Memory = NewV1ContextMemoryService(opts...)
	return
}

// Deletes context data based on provided parameters
func (r *V1ContextService) Delete(ctx context.Context, body V1ContextDeleteParams, opts ...option.RequestOption) (res *V1ContextDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint accepts context data and sends it to a context processor for
// further handling. It returns a success or error response depending on the result
// from the context processor.
func (r *V1ContextService) Add(ctx context.Context, body V1ContextAddParams, opts ...option.RequestOption) (res *V1ContextAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This endpoint sends a search request to the context processor to retrieve
// relevant context data based on the provided query.
func (r *V1ContextService) Search(ctx context.Context, body V1ContextSearchParams, opts ...option.RequestOption) (res *V1ContextSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1ContextDeleteResponse = any

type V1ContextAddResponse = any

type V1ContextSearchResponse struct {
	Contexts []V1ContextSearchResponseContext `json:"contexts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contexts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextSearchResponseContext struct {
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	Metadata  any       `json:"metadata"`
	Score     float64   `json:"score"`
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		CreatedAt   respjson.Field
		Metadata    respjson.Field
		Score       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextSearchResponseContext) RawJSON() string { return r.JSON.raw }
func (r *V1ContextSearchResponseContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextDeleteParams struct {
	// Flag to delete by document
	ByDoc param.Opt[bool] `json:"by_doc,omitzero"`
	// Flag to delete by ID
	ByID param.Opt[bool] `json:"by_id,omitzero"`
	// Optional organization ID
	OrganizationID param.Opt[string] `json:"organization_id,omitzero"`
	// Optional user ID
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Source identifier for the context
	Source param.Opt[string] `json:"source,omitzero"`
	paramObj
}

func (r V1ContextDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContextDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextAddParams struct {
	// The source of the context data
	Source param.Opt[string] `json:"source,omitzero"`
	// Type of context being added
	//
	// Any of "resource", "conversation", "instruction".
	ContextType V1ContextAddParamsContextType `json:"context_type,omitzero"`
	// Array of documents with content and additional metadata
	Documents []V1ContextAddParamsDocument `json:"documents,omitzero"`
	// Additional metadata for the context
	Metadata V1ContextAddParamsMetadata `json:"metadata,omitzero"`
	// Scope of the context
	//
	// Any of "internal", "external".
	Scope V1ContextAddParamsScope `json:"scope,omitzero"`
	paramObj
}

func (r V1ContextAddParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContextAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of context being added
type V1ContextAddParamsContextType string

const (
	V1ContextAddParamsContextTypeResource     V1ContextAddParamsContextType = "resource"
	V1ContextAddParamsContextTypeConversation V1ContextAddParamsContextType = "conversation"
	V1ContextAddParamsContextTypeInstruction  V1ContextAddParamsContextType = "instruction"
)

type V1ContextAddParamsDocument struct {
	// The content of the document
	Content     param.Opt[string] `json:"content,omitzero"`
	ExtraFields map[string]string `json:"-"`
	paramObj
}

func (r V1ContextAddParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextAddParamsDocument
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *V1ContextAddParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional metadata for the context
type V1ContextAddParamsMetadata struct {
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

func (r V1ContextAddParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextAddParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContextAddParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scope of the context
type V1ContextAddParamsScope string

const (
	V1ContextAddParamsScopeInternal V1ContextAddParamsScope = "internal"
	V1ContextAddParamsScopeExternal V1ContextAddParamsScope = "external"
)

type V1ContextSearchParams struct {
	// Minimum similarity threshold
	MinimumSimilarityThreshold float64 `json:"minimum_similarity_threshold,required"`
	// The search query used to search for context data
	Query string `json:"query,required"`
	// Maximum similarity threshold (must be >= minimum_similarity_threshold)
	SimilarityThreshold float64 `json:"similarity_threshold,required"`
	// The ID of the user making the request
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Additional metadata for the search
	Metadata any `json:"metadata,omitzero"`
	// Search scope
	//
	// Any of "internal", "external".
	Scope V1ContextSearchParamsScope `json:"scope,omitzero"`
	paramObj
}

func (r V1ContextSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContextSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search scope
type V1ContextSearchParamsScope string

const (
	V1ContextSearchParamsScopeInternal V1ContextSearchParamsScope = "internal"
	V1ContextSearchParamsScopeExternal V1ContextSearchParamsScope = "external"
)
