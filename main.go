package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/nokia/eda/apps/terraform-provider-bootstrap/internal/provider"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "github.com/nokia-eda/bootstrap-v1",
	}

	err := providerserver.Serve(context.Background(), provider.New("v1"), opts)
	if err != nil {
		log.Fatal(err.Error())
	}
}
