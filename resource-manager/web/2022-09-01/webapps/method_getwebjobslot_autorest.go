package webapps

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetWebJobSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *WebJob
}

// GetWebJobSlot ...
func (c WebAppsClient) GetWebJobSlot(ctx context.Context, id SlotWebJobId) (result GetWebJobSlotOperationResponse, err error) {
	req, err := c.preparerForGetWebJobSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWebJobSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWebJobSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetWebJobSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetWebJobSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetWebJobSlot prepares the GetWebJobSlot request.
func (c WebAppsClient) preparerForGetWebJobSlot(ctx context.Context, id SlotWebJobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetWebJobSlot handles the response to the GetWebJobSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetWebJobSlot(resp *http.Response) (result GetWebJobSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}