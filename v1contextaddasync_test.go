// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Alchemyst-ai/alchemyst-sdk-golang"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/testutil"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/option"
)

func TestV1ContextAddAsyncNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := alchemystai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Context.AddAsync.New(context.TODO(), alchemystai.V1ContextAddAsyncNewParams{
		ContextType: alchemystai.V1ContextAddAsyncNewParamsContextTypeResource,
		Documents: []alchemystai.V1ContextAddAsyncNewParamsDocument{{
			Content: alchemystai.String("Customer asked about pricing for the Scale plan."),
		}},
		Scope:  alchemystai.V1ContextAddAsyncNewParamsScopeInternal,
		Source: "support-inbox",
		Metadata: alchemystai.V1ContextAddAsyncNewParamsMetadata{
			FileName:     alchemystai.String("support_thread_TCK-1234.txt"),
			FileSize:     alchemystai.Float(2048),
			FileType:     alchemystai.String("text/plain"),
			GroupName:    []string{"support", "pricing"},
			LastModified: alchemystai.String("2025-01-10T12:34:56.000Z"),
		},
	})
	if err != nil {
		var apierr *alchemystai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContextAddAsyncCancel(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := alchemystai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Context.AddAsync.Cancel(context.TODO(), "id")
	if err != nil {
		var apierr *alchemystai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
