package servicedirectory

// [START servicedirectory]
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

// [END servicedirectory]
