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

func TestV1ContextDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Context.Delete(context.TODO(), alchemystai.V1ContextDeleteParams{
		ByDoc:          alchemystai.Bool(true),
		ByID:           alchemystai.Bool(true),
		OrganizationID: alchemystai.String("organization_id"),
		Source:         alchemystai.String("source"),
		UserID:         alchemystai.String("user_id"),
	})
	if err != nil {
		var apierr *alchemystai.Error
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
	client := alchemystai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Context.Add(context.TODO(), alchemystai.V1ContextAddParams{
		ContextType: alchemystai.V1ContextAddParamsContextTypeResource,
		Documents: []alchemystai.V1ContextAddParamsDocument{{
			Content: alchemystai.String("content"),
		}},
		Metadata: alchemystai.V1ContextAddParamsMetadata{
			FileName:     alchemystai.String("fileName"),
			FileSize:     alchemystai.Float(0),
			FileType:     alchemystai.String("fileType"),
			GroupName:    []string{"string"},
			LastModified: alchemystai.String("lastModified"),
		},
		Scope:  alchemystai.V1ContextAddParamsScopeInternal,
		Source: alchemystai.String("source"),
	})
	if err != nil {
		var apierr *alchemystai.Error
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
	client := alchemystai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Context.Search(context.TODO(), alchemystai.V1ContextSearchParams{
		MinimumSimilarityThreshold: 0.5,
		Query:                      "search query for user preferences",
		SimilarityThreshold:        0.8,
		Metadata:                   map[string]interface{}{},
		Scope:                      alchemystai.V1ContextSearchParamsScopeInternal,
		UserID:                     alchemystai.String("user123"),
	})
	if err != nil {
		var apierr *alchemystai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
