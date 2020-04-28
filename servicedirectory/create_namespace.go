package servicedirectory

// [START servicedirectory]
import (
	"context"
	"fmt"
	"io"

	servicedirectory "cloud.google.com/go/servicedirectory/apiv1beta1"
	sdpb "google.golang.org/genproto/googleapis/cloud/servicedirectory/v1beta1"
)

func createNamespace(w io.Writer, projectId string) error {
	location := "us-east4"
	namespaceId := "golang-test-namespace"

	ctx := context.Background()
	// Create a registration client.
	client, err := servicedirectory.NewRegistrationClient(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("create request %s, %s, %s", projectId, location, namespaceId)
	// Create a Namespace.
	createNsReq := &sdpb.CreateNamespaceRequest{
		Parent:      fmt.Sprintf("projects/%s/locations/%s", projectId, location),
		NamespaceId: namespaceId,
	}
	resp, err := client.CreateNamespace(ctx, createNsReq)
	fmt.Printf("create request response %v error: %v", resp, err)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "servicedirectory.CreateNamespace result: %s\n", resp.Name)
	return nil
}

// [END servicedirectory]
