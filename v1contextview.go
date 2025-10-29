// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystaisdk

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/alchemyst-ai-sdk-go/internal/apijson"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/option"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/packages/respjson"
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

// Gets the context information for the authenticated user
func (r *V1ContextViewService) Get(ctx context.Context, opts ...option.RequestOption) (res *V1ContextViewGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/view"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Fetches documents view for authenticated user with optional organization context
func (r *V1ContextViewService) Docs(ctx context.Context, opts ...option.RequestOption) (res *V1ContextViewDocsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/context/view/docs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type V1ContextViewGetResponse struct {
	// List of context items
	Context []any `json:"context"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ContextViewGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ContextViewGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ContextViewDocsResponse = any
