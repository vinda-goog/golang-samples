package servicedirectory

// [START servicedirectory]
import (
	"context"
	"fmt"

	servicedirectory "cloud.google.com/go/servicedirectory/apiv1beta1"
	sdpb "google.golang.org/genproto/googleapis/cloud/servicedirectory/v1beta1"
)

func deleteNamespace(projectId string) error {
	location := "us-east4"
	namespaceId := "golang-test-namespace"

	ctx := context.Background()
	// Create a registration client.
	client, err := servicedirectory.NewRegistrationClient(ctx)
	if err != nil {
		return err
	}

	// Delete a Namespace.
	deleteNsReq := &sdpb.DeleteNamespaceRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", projectId, location, namespaceId),
	}
	deleteErr := client.DeleteNamespace(ctx, deleteNsReq)
	if deleteErr != nil {
		return deleteErr
	}
	return nil
}

// [END servicedirectory]
