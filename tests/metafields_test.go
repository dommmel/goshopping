package tests

import (
	"context"
	"testing"

	"github.com/dommmel/goshopping/shopify"
)

func TestListMetafields(t *testing.T) {
	ctx := context.Background()
	opt := &shopify.MetafieldListOptions{Fields: []string{"id", "key", "value"}}
	metafields, _, err := client.Metafields.ListByProduct(ctx, 117563710, opt)
	if err != nil {
		t.Fatalf("Metafields.List() returned error: %v", err)
	}
	if len(metafields) == 0 {
		t.Errorf("Metafields.List() returned no events")
	}
	//debug(spit.Sdump(metafields))

}

// func TestUpdateMetafields(t *testing.T) {
// 	ctx := context.Background()
// 	opt := &shopify.MetafieldListOptions{Fields: []string{"id", "key", "value", "namespace"}}
// 	metafields, _, err := client.Metafields.ListByProduct(ctx, 117563710, opt)

// 	_, reserr := client.Products.UpdateByProduct(ctx, 117563710, metafields)

// 	if reserr != nil {
// 		t.Fatalf("Metafields.UpdateByProduct() returned error: %v", err)
// 	}

// }
