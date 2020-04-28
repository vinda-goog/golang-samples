package servicedirectory

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func TestEndpoint(t *testing.T) {
	// <test setup code>
	tc := testutil.SystemTest(t)
	createNsBuf := new(bytes.Buffer)
	createNsErr := createNamespace(createNsBuf, tc.ProjectID)
	if createNsErr != nil {
		fmt.Printf("Failed to create namespace in test setup. Skipping: %v", createNsErr)
		os.Exit(0)
	}
	createServiceBuf := new(bytes.Buffer)
	createServiceErr := createService(createServiceBuf, tc.ProjectID)
	if createServiceErr != nil {
		fmt.Printf("Failed to create service in test setup. Skipping: %v", createServiceErr)
		os.Exit(0)
	}
	t.Run("create", func(t *testing.T) {
		tc := testutil.SystemTest(t)
		buf := new(bytes.Buffer)
		err := createEndpoint(buf, tc.ProjectID)

		if err != nil {
			t.Errorf("CreateEndpoint %v", err)
		}

		got := buf.String()
		fmt.Println(got)
		if want := "services/golang-test-endpoint"; !strings.Contains(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("delete", func(t * testing.T) {
		tc := testutil.SystemTest(t)
		err := deleteEndpoint(tc.ProjectID)

		if err != nil {
			t.Errorf("DeleteEndpoint: %v", err)
		}
	})

	// <test tear-down code>
	deleteErr := deleteNamespace(tc.ProjectID)

	if deleteErr != nil {
		fmt.Println("Failed to delete namespace in test tear down.")
	}
}

