// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alchemystaisdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Alchemyst-ai/alchemyst-sdk-golang"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/internal/testutil"
	"github.com/Alchemyst-ai/alchemyst-sdk-golang/option"
)

func TestV1ContextDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Context.Delete(context.TODO(), alchemystaisdk.V1ContextDeleteParams{
		ByDoc:          alchemystaisdk.Bool(true),
		ByID:           alchemystaisdk.Bool(true),
		OrganizationID: alchemystaisdk.String("organization_id"),
		Source:         alchemystaisdk.String("source"),
		UserID:         alchemystaisdk.String("user_id"),
	})
	if err != nil {
		var apierr *alchemystaisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContextAddWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Context.Add(context.TODO(), alchemystaisdk.V1ContextAddParams{
		ContextType: alchemystaisdk.V1ContextAddParamsContextTypeResource,
		Documents: []alchemystaisdk.V1ContextAddParamsDocument{{
			Content: alchemystaisdk.String("content"),
		}},
		Metadata: alchemystaisdk.V1ContextAddParamsMetadata{
			FileName:     alchemystaisdk.String("fileName"),
			FileSize:     alchemystaisdk.Float(0),
			FileType:     alchemystaisdk.String("fileType"),
			GroupName:    []string{"string"},
			LastModified: alchemystaisdk.String("lastModified"),
		},
		Scope:  alchemystaisdk.V1ContextAddParamsScopeInternal,
		Source: alchemystaisdk.String("source"),
	})
	if err != nil {
		var apierr *alchemystaisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ContextSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Context.Search(context.TODO(), alchemystaisdk.V1ContextSearchParams{
		MinimumSimilarityThreshold: 0.5,
		Query:                      "search query for user preferences",
		SimilarityThreshold:        0.8,
		Metadata:                   map[string]interface{}{},
		Scope:                      alchemystaisdk.V1ContextSearchParamsScopeInternal,
		UserID:                     alchemystaisdk.String("user123"),
	})
	if err != nil {
		var apierr *alchemystaisdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
