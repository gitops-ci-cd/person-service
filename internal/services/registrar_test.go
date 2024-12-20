package services

import (
	"testing"

	"google.golang.org/grpc"
)

func TestRegister(t *testing.T) {
	// Create a new gRPC server
	server := grpc.NewServer()

	// Call the Register function
	Register(server, nil)

	// Verify that the services have been registered
	serviceInfo := server.GetServiceInfo()
	if _, ok := serviceInfo["com.acme.schema.v1.PersonService"]; !ok {
		t.Errorf("PersonService not registered")
	}

	// Verify that reflection has been registered
	if _, ok := serviceInfo["grpc.reflection.v1alpha.ServerReflection"]; !ok {
		t.Errorf("Reflection service not registered")
	}
}
