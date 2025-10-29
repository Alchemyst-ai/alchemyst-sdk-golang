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

func TestV1ContextMemoryDeleteWithOptionalParams(t *testing.T) {
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
	err := client.V1.Context.Memory.Delete(context.TODO(), alchemystai.V1ContextMemoryDeleteParams{
		MemoryID:       alchemystai.String("memoryId"),
		OrganizationID: alchemystai.String("organization_id"),
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

func TestV1ContextMemoryAddWithOptionalParams(t *testing.T) {
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
	err := client.V1.Context.Memory.Add(context.TODO(), alchemystai.V1ContextMemoryAddParams{
		Contents: []alchemystai.V1ContextMemoryAddParamsContent{{
			Content: alchemystai.String("content"),
		}},
		MemoryID: alchemystai.String("memoryId"),
	})
	if err != nil {
		var apierr *alchemystai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
