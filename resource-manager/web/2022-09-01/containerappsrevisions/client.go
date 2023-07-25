package containerappsrevisions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsRevisionsClient struct {
	Client *resourcemanager.Client
}

func NewContainerAppsRevisionsClientWithBaseURI(api environments.Api) (*ContainerAppsRevisionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "containerappsrevisions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContainerAppsRevisionsClient: %+v", err)
	}

	return &ContainerAppsRevisionsClient{
		Client: client,
	}, nil
}
