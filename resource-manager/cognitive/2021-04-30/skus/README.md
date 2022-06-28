
## `github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2021-04-30/skus` Documentation

The `skus` SDK allows for interaction with the Azure Resource Manager Service `cognitive` (API Version `2021-04-30`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/cognitive/2021-04-30/skus"
```


### Client Initialization

```go
client := skus.NewSkusClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
if err != nil {
	// handle the error
}
```


### Example Usage: `SkusClient.ResourceSkusList`

```go
ctx := context.TODO()
id := skus.NewSubscriptionID()
// alternatively `client.ResourceSkusList(ctx, id)` can be used to do batched pagination
items, err := client.ResourceSkusListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```