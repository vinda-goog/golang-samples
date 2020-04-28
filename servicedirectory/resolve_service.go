package servicedirectory

// [START servicedirectory]
import (
	"context"
	"fmt"
	"io"

	servicedirectory "cloud.google.com/go/servicedirectory/apiv1beta1"
	sdpb "google.golang.org/genproto/googleapis/cloud/servicedirectory/v1beta1"
)

func resolveService(w io.Writer, projectId string) error {
	location := "us-east4"
	namespaceId := "golang-test-namespace"
	serviceId := "golang-test-service"

	ctx := context.Background()
	// Create a lookup client.
	resolver, err := servicedirectory.NewLookupClient(ctx)
	if err != nil {
		return err
	}
	// Now Resolve the service.
	lookupRequest := &sdpb.ResolveServiceRequest{
		Name: fmt.Sprintf("projects/%s/locations/%s/namespaces/%s/services/%s", projectId, location, namespaceId, serviceId),
	}
	result, err := resolver.ResolveService(ctx, lookupRequest)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "Successfully Resolved Service %v\n", result)
	return nil
}

// [END servicedirectory]
