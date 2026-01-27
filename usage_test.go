// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystai_test

import (
	"context"
	"os"
	"testing"

	"github.com/Alchemyst-ai/alchemyst-sdk-golang"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/testutil"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/option"
)

func TestUsage(t *testing.T) {
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
	t.Skip("Prism tests are disabled")
	response, err := client.V1.Context.Add(context.TODO(), alchemystai.V1ContextAddParams{
		ContextType: alchemystai.V1ContextAddParamsContextTypeResource,
		Documents: []alchemystai.V1ContextAddParamsDocument{{
			Content: alchemystai.String("The content of the document"),
		}},
		Scope:  alchemystai.V1ContextAddParamsScopeInternal,
		Source: "platform.api.context.add",
		Metadata: alchemystai.V1ContextAddParamsMetadata{
			FileName:     alchemystai.String("notes.txt"),
			FileType:     alchemystai.String("text/plain"),
			LastModified: alchemystai.String("2025-10-01T18:42:40.419Z"),
			FileSize:     alchemystai.Float(1024),
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.ContextID)
}
