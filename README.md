# goshopping #

goshopping is a Go client library for accessing the [Shopify API](https://help.shopify.com/api/reference).

## Usage ##

```go
import "github.com/dommmel/goshopping/shopify"
```

Construct a new Shopify client, then use it to access different parts of the Shopify API. For example for a [private app](https://help.shopify.com/api/getting-started/authentication/private-authentication)

```go
client := shopify.NewPrivateClient(nil, "apiKey", "password", "shopName")
```

