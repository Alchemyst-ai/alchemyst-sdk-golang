// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystaisdk

import (
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/option"
)

// V1Service contains methods and other services that help with interacting with
// the alchemyst_ai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	Options []option.RequestOption
	Context V1ContextService
	Org     V1OrgService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r V1Service) {
	r = V1Service{}
	r.Options = opts
	r.Context = NewV1ContextService(opts...)
	r.Org = NewV1OrgService(opts...)
	return
}
