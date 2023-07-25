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

type BackupSlotOperationResponse struct {
	HttpResponse *http.Response
	Model        *BackupItem
}

// BackupSlot ...
func (c WebAppsClient) BackupSlot(ctx context.Context, id SlotId, input BackupRequest) (result BackupSlotOperationResponse, err error) {
	req, err := c.preparerForBackupSlot(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "BackupSlot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "BackupSlot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForBackupSlot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "BackupSlot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForBackupSlot prepares the BackupSlot request.
func (c WebAppsClient) preparerForBackupSlot(ctx context.Context, id SlotId, input BackupRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/backup", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForBackupSlot handles the response to the BackupSlot request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForBackupSlot(resp *http.Response) (result BackupSlotOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}