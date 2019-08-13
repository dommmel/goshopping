// +build private

package tests

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/dommmel/goshopping/shopify"
)

var (
	client *shopify.Client
	//debug  godebug.DebugFunction
	spit   spew.ConfigState
)

func init() {
	apiKey := os.Getenv("SHOPIFY_API_KEY")
	password := os.Getenv("SHOPIFY_PASSWORD")
	shopName := os.Getenv("SHOP_NAME")
	client = shopify.NewPrivateClient(nil, apiKey, password, shopName)
	//debug = godebug.Debug("response")
	spit = spew.ConfigState{Indent: " ", DisableCapacities: true, DisablePointerAddresses: true}
}
