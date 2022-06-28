
## `github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/eventhubsclusters` Documentation

The `eventhubsclusters` SDK allows for interaction with the Azure Resource Manager Service `eventhub` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/eventhub/2021-11-01/eventhubsclusters"
```


### Client Initialization

```go
client := eventhubsclusters.NewEventHubsClustersClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `EventHubsClustersClient.ClustersCreateOrUpdate`

```go
ctx := context.TODO()
id := eventhubsclusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

payload := eventhubsclusters.Cluster{
	// ...
}

future, err := client.ClustersCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `EventHubsClustersClient.ClustersDelete`

```go
ctx := context.TODO()
id := eventhubsclusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")
future, err := client.ClustersDelete(ctx, id)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```


### Example Usage: `EventHubsClustersClient.ClustersGet`

```go
ctx := context.TODO()
id := eventhubsclusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")
read, err := client.ClustersGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EventHubsClustersClient.ClustersListByResourceGroup`

```go
ctx := context.TODO()
id := eventhubsclusters.NewResourceGroupID()
// alternatively `client.ClustersListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ClustersListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EventHubsClustersClient.ClustersListBySubscription`

```go
ctx := context.TODO()
id := eventhubsclusters.NewSubscriptionID()
// alternatively `client.ClustersListBySubscription(ctx, id)` can be used to do batched pagination
items, err := client.ClustersListBySubscriptionComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `EventHubsClustersClient.ClustersUpdate`

```go
ctx := context.TODO()
id := eventhubsclusters.NewClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

payload := eventhubsclusters.Cluster{
	// ...
}

future, err := client.ClustersUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if err := future.Poller.PollUntilDone(); err != nil {
	// handle the error
}
```