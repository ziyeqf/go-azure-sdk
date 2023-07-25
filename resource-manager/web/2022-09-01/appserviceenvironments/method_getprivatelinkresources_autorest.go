package appserviceenvironments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivateLinkResourcesOperationResponse struct {
	HttpResponse *http.Response
	Model        *PrivateLinkResourcesWrapper
}

// GetPrivateLinkResources ...
func (c AppServiceEnvironmentsClient) GetPrivateLinkResources(ctx context.Context, id HostingEnvironmentId) (result GetPrivateLinkResourcesOperationResponse, err error) {
	req, err := c.preparerForGetPrivateLinkResources(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetPrivateLinkResources", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetPrivateLinkResources", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetPrivateLinkResources(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appserviceenvironments.AppServiceEnvironmentsClient", "GetPrivateLinkResources", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetPrivateLinkResources prepares the GetPrivateLinkResources request.
func (c AppServiceEnvironmentsClient) preparerForGetPrivateLinkResources(ctx context.Context, id HostingEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/privateLinkResources", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetPrivateLinkResources handles the response to the GetPrivateLinkResources request. The method always
// closes the http.Response Body.
func (c AppServiceEnvironmentsClient) responderForGetPrivateLinkResources(resp *http.Response) (result GetPrivateLinkResourcesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}