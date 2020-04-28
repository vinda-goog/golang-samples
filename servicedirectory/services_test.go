package servicedirectory

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func TestMain(m *testing.M) {
	tc, ok := testutil.ContextMain(m)
	if !ok {
		fmt.Println("Could not set up tests. Set GOLANG_SAMPLES_PROJECT_ID? Skipping.")
		os.Exit(0)
	}
	createBuf := new(bytes.Buffer)
	err := createNamespace(createBuf, tc.ProjectID)

	if err != nil {
		fmt.Printf("Failed to create namespace in test setup. Skipping: %v", err)
		os.Exit(0)
	}
	exitCode := m.Run()

	//deleteErr := deleteNamespace(tc.ProjectID)
	//if deleteErr != nil {
	//	fmt.Println("Failed to delete namespace in test tear down.")
	//}
	os.Exit(exitCode)
}

func TestCreateService(t *testing.T) {
	tc := testutil.SystemTest(t)
	buf := new(bytes.Buffer)
	err := createService(buf, tc.ProjectID)

	if err != nil {
		t.Errorf("CreateService: %v", err)
	}

	got := buf.String()
	fmt.Println(got)
	if want := "services/golang-test-service"; !strings.Contains(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
func TestResolveService(t *testing.T) {
	tc := testutil.SystemTest(t)
	buf := new(bytes.Buffer)
	err := resolveService(buf, tc.ProjectID)

	if err != nil {
		t.Errorf("CreateService: %v", err)
	}

	got := buf.String()
	fmt.Println(got)
	if want := "services/golang-test-service"; !strings.Contains(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestDeleteService(t *testing.T) {
	tc := testutil.SystemTest(t)
	err := deleteService(tc.ProjectID)

	if err != nil {
		t.Errorf("DeleteService: %v", err)
	}
}
