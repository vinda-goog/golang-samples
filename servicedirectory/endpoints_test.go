// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package servicedirectory

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
)

func TestEndpoint(t *testing.T) {
	tc := testutil.SystemTest(t)
	createNsErr := createNamespace(ioutil.Discard, tc.ProjectID)
	if createNsErr != nil {
		fmt.Printf("Failed to create namespace in test setup. Skipping: %v", createNsErr)
		os.Exit(0)
	}
	createServiceErr := createService(ioutil.Discard, tc.ProjectID)
	if createServiceErr != nil {
		fmt.Printf("Failed to create service in test setup. Skipping: %v", createServiceErr)
		os.Exit(0)
	}
	t.Run("create", func(t *testing.T) {
		tc := testutil.SystemTest(t)
		buf := new(bytes.Buffer)
		if err := createEndpoint(buf, tc.ProjectID); err != nil {
			t.Errorf("CreateEndpoint %v", err)
		}

		got := buf.String()
		fmt.Println(got)
		if want := "endpoints/golang-test-endpoint"; !strings.Contains(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("delete", func(t *testing.T) {
		tc := testutil.SystemTest(t)
		if err := deleteEndpoint(tc.ProjectID); err != nil {
			t.Errorf("DeleteEndpoint: %v", err)
		}
	})

	deleteErr := deleteNamespace(tc.ProjectID)

	if deleteErr != nil {
		t.Errorf("Failed to delete namespace in test tear down: %v", deleteErr)
	}
}
