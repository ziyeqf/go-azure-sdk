package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMigrateMySqlStatusSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *MigrateMySqlStatus
}

// GetMigrateMySqlStatusSlot ...
func (c WebAppsClient) GetMigrateMySqlStatusSlot(ctx context.Context, id SlotId) (result GetMigrateMySqlStatusSlotOperationResponse, err error) {
	req, err := c.preparerForGetMigrateMySqlStatusSlot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMigrateMySqlStatusSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMigrateMySqlStatusSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetMigrateMySqlStatusSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetMigrateMySqlStatusSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetMigrateMySqlStatusSlot prepares the GetMigrateMySqlStatusSlot request.
func (c WebAppsClient) preparerForGetMigrateMySqlStatusSlot(ctx context.Context, id SlotId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/migrateMySql/status", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetMigrateMySqlStatusSlot handles the response to the GetMigrateMySqlStatusSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetMigrateMySqlStatusSlot(resp *http.Response) (result GetMigrateMySqlStatusSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}