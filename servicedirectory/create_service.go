package servicedirectory

// [START servicedirectory]
import (
	"context"
	"fmt"
	"io"

	servicedirectory "cloud.google.com/go/servicedirectory/apiv1beta1"
	sdpb "google.golang.org/genproto/googleapis/cloud/servicedirectory/v1beta1"
)

func createService(w io.Writer, projectId string) error {
	location := "us-east4"
	namespaceId := "golang-test-namespace"
	serviceId := "golang-test-service"

	ctx := context.Background()
	// Create a registration client.
	client, err := servicedirectory.NewRegistrationClient(ctx)
	if err != nil {
		return err
	}
	// Create a Service.
	createServiceReq := &sdpb.CreateServiceRequest{
		Parent:    fmt.Sprintf("projects/%s/locations/%s/namespaces/%s", projectId, location, namespaceId),
		ServiceId: serviceId,
		Service: &sdpb.Service{
			Metadata: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
	}
	service, err := client.CreateService(ctx, createServiceReq)
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "servicedirectory.Createservice result %s\n", service.Name)
	return nil
}

// [END servicedirectory]
