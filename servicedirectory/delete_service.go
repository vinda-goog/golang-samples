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

// [START servicedirectory_delete_service]
import (
	"context"
	"fmt"

	servicedirectory "cloud.google.com/go/servicedirectory/apiv1beta1"
	sdpb "google.golang.org/genproto/googleapis/cloud/servicedirectory/v1beta1"
)

func deleteService(projectId string) error {
	location := "us-east4"
	namespaceId := "golang-test-namespace"
	serviceId := "golang-test-service"

	ctx := context.Background()
	// Create a registration client.
	client, err := servicedirectory.NewRegistrationClient(ctx)
	if err != nil {
		return err
	}

	deleteServiceReq := &sdpb.DeleteServiceRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", projectId, location, namespaceId, serviceId),
	}
	deleteErr := client.DeleteService(ctx, deleteServiceReq)
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}

// [END servicedirectory_delete_service]
