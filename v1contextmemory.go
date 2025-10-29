// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystaisdk

import (
	"context"
	"net/http"
	"slices"

	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/apijson"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/requestconfig"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/option"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/packages/param"
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

// Deletes memory context data based on provided parameters
func (r *V1ContextMemoryService) Delete(ctx context.Context, body V1ContextMemoryDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v1/context/memory/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// This endpoint adds memory context data, fetching chat history if needed.
func (r *V1ContextMemoryService) Add(ctx context.Context, body V1ContextMemoryAddParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "api/v1/context/memory/add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type V1ContextMemoryDeleteParams struct {
	// Optional organization ID
	OrganizationID param.Opt[string] `json:"organization_id,omitzero"`
	// Optional user ID
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// The ID of the memory to delete
	MemoryID param.Opt[string] `json:"memoryId,omitzero"`
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
	// The ID of the memory
	MemoryID param.Opt[string] `json:"memoryId,omitzero"`
	// Array of content objects with additional properties allowed
	Contents []V1ContextMemoryAddParamsContent `json:"contents,omitzero"`
	paramObj
}

func (r V1ContextMemoryAddParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextMemoryAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ContextMemoryAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextMemoryAddParamsContent struct {
	Content     param.Opt[string] `json:"content,omitzero"`
	ExtraFields map[string]any    `json:"-"`
	paramObj
}

func (r V1ContextMemoryAddParamsContent) MarshalJSON() (data []byte, err error) {
	type shadow V1ContextMemoryAddParamsContent
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *V1ContextMemoryAddParamsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
