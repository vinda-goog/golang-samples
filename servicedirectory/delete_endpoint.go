package servicedirectory

// [START servicedirectory]
import (
	"context"
	"fmt"

	servicedirectory "cloud.google.com/go/servicedirectory/apiv1beta1"
	sdpb "google.golang.org/genproto/googleapis/cloud/servicedirectory/v1beta1"
)

func deleteEndpoint(projectId string) error {
	location := "us-east4"
	namespaceId := "golang-test-namespace"
	serviceId := "golang-test-service"
	endpointId := "golang-test-endpoint"

	ctx := context.Background()
	// Create a registration client.
	client, err := servicedirectory.NewRegistrationClient(ctx)
	if err != nil {
		return err
	}

	deleteEndpointReq := &sdpb.DeleteEndpointRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s/endpoints/%s", projectId, location, namespaceId, serviceId, endpointId),
	}
	deleteErr := client.DeleteEndpoint(ctx, deleteEndpointReq)
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}

// [END servicedirectory]
