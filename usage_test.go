// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystaisdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/alchemyst-ai-sdk-go"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/internal/testutil"
	"github.com/stainless-sdks/alchemyst-ai-sdk-go/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.V1.Context.Add(context.TODO(), alchemystaisdk.V1ContextAddParams{
		ContextType: alchemystaisdk.V1ContextAddParamsContextTypeResource,
		Documents: []alchemystaisdk.V1ContextAddParamsDocument{{
			Content: alchemystaisdk.String("The content of the document"),
		}},
		Metadata: alchemystaisdk.V1ContextAddParamsMetadata{
			FileName:     alchemystaisdk.String("notes.txt"),
			FileType:     alchemystaisdk.String("text/plain"),
			LastModified: alchemystaisdk.String("2025-10-01T18:42:40.419Z"),
			FileSize:     alchemystaisdk.Float(1024),
		},
		Scope:  alchemystaisdk.V1ContextAddParamsScopeInternal,
		Source: alchemystaisdk.String("platform.api.context.add"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response)
}
