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
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/packages/respjson"
)

// V1OrgContextService contains methods and other services that help with
// interacting with the alchemyst_ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OrgContextService] method instead.
type V1OrgContextService struct {
	Options []option.RequestOption
}

// NewV1OrgContextService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1OrgContextService(opts ...option.RequestOption) (r V1OrgContextService) {
	r = V1OrgContextService{}
	r.Options = opts
	return
}

// View organization context
func (r *V1OrgContextService) View(ctx context.Context, body V1OrgContextViewParams, opts ...option.RequestOption) (res *V1OrgContextViewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/org/context/view"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type V1OrgContextViewResponse struct {
	Contexts any `json:"contexts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contexts    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1OrgContextViewResponse) RawJSON() string { return r.JSON.raw }
func (r *V1OrgContextViewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1OrgContextViewParams struct {
	UserIDs []string `json:"userIds,omitzero,required"`
	paramObj
}

func (r V1OrgContextViewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1OrgContextViewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1OrgContextViewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
