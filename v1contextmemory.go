// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystai

import (
	"context"
	"net/http"
	"slices"

	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/apijson"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/requestconfig"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/option"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/packages/param"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/packages/respjson"
)

// V1ContextMemoryService contains methods and other services that help with
// interacting with the alchemyst_ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContextMemoryService] method instead.
type V1ContextMemoryService struct {
	Options []option.RequestOption
}

// NewV1ContextMemoryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ContextMemoryService(opts ...option.RequestOption) (r V1ContextMemoryService) {
	r = V1ContextMemoryService{}
	r.Options = opts
	return
}

// This endpoint updates memory context data.
func (r *V1ContextMemoryService) Update(ctx context.Context, body V1ContextMemoryUpdateParams, opts ...option.RequestOption) (res *V1ContextMemoryUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/memory/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes memory context data based on provided parameters.
func (r *V1ContextMemoryService) Delete(ctx context.Context, body V1ContextMemoryDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "api/v1/context/memory/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// This endpoint adds memory (chat history) as context.
func (r *V1ContextMemoryService) Add(ctx context.Context, body V1ContextMemoryAddParams, opts ...option.RequestOption) (res *V1ContextMemoryAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/memory/add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1ContextMemoryUpdateResponse struct {
	MemoryID       string  `json:"memory_id,required"`
	Success        bool    `json:"success,required"`
	UpdatedEntries float64 `json:"updated_entries,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MemoryID       respjson.Field
		Success        respjson.Field
		UpdatedEntries respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextMemoryUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextMemoryUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextMemoryAddResponse struct {
	ContextID          string  `json:"context_id,required"`
	Success            bool    `json:"success,required"`
	ProcessedDocuments float64 `json:"processed_documents"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContextID          respjson.Field
		Success            respjson.Field
		ProcessedDocuments respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextMemoryAddResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextMemoryAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextMemoryUpdateParams struct {
	// Array of updated content objects
	Contents []V1ContextMemoryUpdateParamsContent `json:"contents,omitzero,required"`
	// The ID of the memory to update
	MemoryID string `json:"memoryId,required"`
	paramObj
}

func (r V1ContextMemoryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextMemoryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContextMemoryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextMemoryUpdateParamsContent struct {
	// Unique ID for the message
	ID param.Opt[string] `json:"id,omitzero"`
	// The content of the memory entry
	Content param.Opt[string] `json:"content,omitzero"`
	// Creation timestamp
	CreatedAt param.Opt[string] `json:"createdAt,omitzero"`
	// Role of the message (e.g., user, assistant)
	Role param.Opt[string] `json:"role,omitzero"`
	// Additional metadata for the memory entry
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r V1ContextMemoryUpdateParamsContent) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextMemoryUpdateParamsContent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContextMemoryUpdateParamsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextMemoryDeleteParams struct {
	// Organization ID
	OrganizationID param.Opt[string] `json:"organization_id,omitzero,required"`
	// The ID of the memory to delete
	MemoryID string `json:"memoryId,required"`
	// Delete by document flag
	ByDoc param.Opt[bool] `json:"by_doc,omitzero"`
	// Delete by ID flag
	ByID param.Opt[bool] `json:"by_id,omitzero"`
	// Optional user ID
	UserID param.Opt[string] `json:"user_id,omitzero"`
	paramObj
}

func (r V1ContextMemoryDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextMemoryDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContextMemoryDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextMemoryAddParams struct {
	// Array of content objects. Each object must contain at least the 'content' field.
	// Additional properties are allowed.
	Contents []V1ContextMemoryAddParamsContent `json:"contents,omitzero,required"`
	// The ID of the memory
	MemoryID string `json:"memoryId,required"`
	// Optional metadata for the memory context. Defaults to ["default"]
	Metadata V1ContextMemoryAddParamsMetadata `json:"metadata,omitzero"`
	paramObj
}

func (r V1ContextMemoryAddParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextMemoryAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContextMemoryAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Content is required.
type V1ContextMemoryAddParamsContent struct {
	// The content of the memory message
	Content string `json:"content,required"`
	// Additional metadata for the message (optional)
	Metadata    V1ContextMemoryAddParamsContentMetadata `json:"metadata,omitzero"`
	ExtraFields map[string]any                          `json:"-"`
	paramObj
}

func (r V1ContextMemoryAddParamsContent) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextMemoryAddParamsContent
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *V1ContextMemoryAddParamsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional metadata for the message (optional)
type V1ContextMemoryAddParamsContentMetadata struct {
	// Unique message ID
	MessageID   param.Opt[string] `json:"messageId,omitzero"`
	ExtraFields map[string]any    `json:"-"`
	paramObj
}

func (r V1ContextMemoryAddParamsContentMetadata) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextMemoryAddParamsContentMetadata
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *V1ContextMemoryAddParamsContentMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional metadata for the memory context. Defaults to ["default"]
type V1ContextMemoryAddParamsMetadata struct {
	// Optional group names for the memory context. Defaults to ["default"] if not
	// provided.
	GroupName   []string       `json:"groupName,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r V1ContextMemoryAddParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextMemoryAddParamsMetadata
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *V1ContextMemoryAddParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
