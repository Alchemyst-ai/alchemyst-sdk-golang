// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystai

import (
	"context"
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

// V1ContextViewService contains methods and other services that help with
// interacting with the alchemyst_ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ContextViewService] method instead.
type V1ContextViewService struct {
	Options []option.RequestOption
}

// NewV1ContextViewService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ContextViewService(opts ...option.RequestOption) (r V1ContextViewService) {
	r = V1ContextViewService{}
	r.Options = opts
	return
}

// Gets the context information for the authenticated user.
func (r *V1ContextViewService) Get(ctx context.Context, query V1ContextViewGetParams, opts ...option.RequestOption) (res *V1ContextViewGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/view"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Fetches documents view for authenticated user with optional organization
// context.
func (r *V1ContextViewService) Docs(ctx context.Context, query V1ContextViewDocsParams, opts ...option.RequestOption) (res *V1ContextViewDocsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/view/docs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type V1ContextViewGetResponse struct {
	// List of context items
	Contexts []V1ContextViewGetResponseContext `json:"contexts,required"`
	Success  bool                              `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contexts    respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextViewGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextViewGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextViewGetResponseContext struct {
	// The content of the context item
	Content string `json:"content"`
	// Additional metadata for the context
	Metadata V1ContextViewGetResponseContextMetadata `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextViewGetResponseContext) RawJSON() string { return r.JSON.raw }
func (r *V1ContextViewGetResponseContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional metadata for the context
type V1ContextViewGetResponseContextMetadata struct {
	FileName     string   `json:"fileName"`
	FileSize     float64  `json:"fileSize"`
	FileType     string   `json:"fileType"`
	GroupName    []string `json:"groupName"`
	LastModified string   `json:"lastModified"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileName     respjson.Field
		FileSize     respjson.Field
		FileType     respjson.Field
		GroupName    respjson.Field
		LastModified respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextViewGetResponseContextMetadata) RawJSON() string { return r.JSON.raw }
func (r *V1ContextViewGetResponseContextMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextViewDocsResponse struct {
	Documents []V1ContextViewDocsResponseDocument `json:"documents,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Documents   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextViewDocsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextViewDocsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextViewDocsResponseDocument struct {
	// Name of the file
	FileName string `json:"fileName,required"`
	// Size of the file in bytes
	FileSize float64 `json:"fileSize,required"`
	// Type/MIME of the file
	FileType string `json:"fileType,required"`
	// Array of group names to which the file belongs
	GroupName []string `json:"groupName,required"`
	// Last modified timestamp (ISO format)
	LastModified string `json:"lastModified,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileName     respjson.Field
		FileSize     respjson.Field
		FileType     respjson.Field
		GroupName    respjson.Field
		LastModified respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextViewDocsResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *V1ContextViewDocsResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextViewGetParams struct {
	// Name of the file to retrieve context for
	FileName param.Opt[string] `query:"file_name,omitzero" json:"-"`
	// Magic key for context retrieval
	MagicKey param.Opt[string] `query:"magic_key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1ContextViewGetParams]'s query parameters as `url.Values`.
func (r V1ContextViewGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ContextViewDocsParams struct {
	// Optional magic key for special access or filtering
	MagicKey param.Opt[string] `query:"magic_key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1ContextViewDocsParams]'s query parameters as
// `url.Values`.
func (r V1ContextViewDocsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
