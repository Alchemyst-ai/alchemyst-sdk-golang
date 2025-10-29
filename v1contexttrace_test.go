// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystaisdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/alchemyst-ai-sdk-go"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/internal/testutil"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/option"
)

func TestV1ContextTraceList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := alchemystaisdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Context.Traces.List(context.TODO())
	if err != nil {
		var apierr *alchemystaisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContextTraceDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := alchemystaisdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Context.Traces.Delete(context.TODO(), "traceId")
	if err != nil {
		var apierr *alchemystaisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
