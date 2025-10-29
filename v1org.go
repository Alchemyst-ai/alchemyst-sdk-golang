// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystai

import (
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/option"
)

// V1OrgService contains methods and other services that help with interacting with
// the alchemyst_ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1OrgService] method instead.
type V1OrgService struct {
	Options []option.RequestOption
	Context V1OrgContextService
}

// NewV1OrgService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1OrgService(opts ...option.RequestOption) (r V1OrgService) {
	r = V1OrgService{}
	r.Options = opts
	r.Context = NewV1OrgContextService(opts...)
	return
}
