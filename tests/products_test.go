package tests

import (
	"context"
	"testing"
)

func TestListProduct(t *testing.T) {
	ctx := context.Background()
	//opt := &shopify.ProductListOptions{Fields: []string{"id"}}
	products, _, err := client.Products.List(ctx, nil)
	if err != nil {
		t.Fatalf("Products.List() returned error: %v", err)
	}
	if len(products) == 0 {
		t.Errorf("Products.List() returned no events")
	}
	//debug(spit.Sdump(products))

}

func TestAutoPagingListProduct(t *testing.T) {
	ctx := context.Background()
	//opt := &shopify.ProductListOptions{Fields: []string{"id"}}
	products, err := client.Products.AutoPagingList(ctx, nil)
	if err != nil {
		t.Fatalf("Products.List() returned error: %v", err)
	}
	if len(products) == 0 {
		t.Errorf("Products.List() returned no events")
	}
	//debug(spit.Sdump(products))

}
