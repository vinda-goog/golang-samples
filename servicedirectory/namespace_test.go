package servicedirectory

import (
	"bytes"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func TestCreateNamespace(t *testing.T) {
	tc := testutil.SystemTest(t)
	buf := new(bytes.Buffer)
	err := createNamespace(buf, tc.ProjectID)

	got := buf.String()

	if err != nil {
		t.Errorf("CreateNamespace: %v", err)
		return
	}
	if want := "namespaces/golang-test-namespace"; !strings.Contains(got, want) {
		t.Errorf("got %q, want %q", got, want)
		return
	}
}

//func TestDeleteNamespace(t *testing.T) {
//	tc := testutil.SystemTest(t)
//	err := deleteNamespace(tc.ProjectID)
//
//	if err != nil {
//		t.Errorf("DeleteNamespace: %v", err)
//		return
//	}
//}
