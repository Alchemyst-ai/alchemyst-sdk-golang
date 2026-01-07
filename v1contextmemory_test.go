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

func TestV1ContextMemoryUpdate(t *testing.T) {
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
	_, err := client.V1.Context.Memory.Update(context.TODO(), alchemystai.V1ContextMemoryUpdateParams{
		Contents: []alchemystai.V1ContextMemoryUpdateParamsContent{{
			ID:        alchemystai.String("msg-1"),
			Content:   alchemystai.String("Customer asked about pricing for the Scale plan."),
			CreatedAt: alchemystai.String("2025-01-10T12:34:56.000Z"),
			Metadata: map[string]any{
				"messageId": "bar",
			},
			Role: alchemystai.String("user"),
		}, {
			ID:        alchemystai.String("msg-2"),
			Content:   alchemystai.String("Updated answer about the Scale plan pricing after discounts."),
			CreatedAt: alchemystai.String("2025-01-10T12:36:00.000Z"),
			Metadata: map[string]any{
				"messageId": "bar",
			},
			Role: alchemystai.String("assistant"),
		}},
		MemoryID: "support-thread-TCK-1234",
	})
	if err != nil {
		var apierr *alchemystai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

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
		MemoryID:       "support-thread-TCK-1234",
		OrganizationID: alchemystai.String("org_01HXYZABC"),
		ByDoc:          alchemystai.Bool(true),
		ByID:           alchemystai.Bool(false),
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
	_, err := client.V1.Context.Memory.Add(context.TODO(), alchemystai.V1ContextMemoryAddParams{
		Contents: []alchemystai.V1ContextMemoryAddParamsContent{{
			Content: "Customer asked about pricing for the Scale plan.",
			Metadata: alchemystai.V1ContextMemoryAddParamsContentMetadata{
				MessageID: alchemystai.String("messageId"),
			},
		}},
		MemoryID: "support-thread-TCK-1234",
		Metadata: alchemystai.V1ContextMemoryAddParamsMetadata{
			GroupName: []string{"string"},
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
