package shopify

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ProductsService service

type ProductForUpdateContainer struct {
	Product *Product `json:"product"`
}

type ProductList struct {
	Products []*Product `json:"products"`
}
type Product struct {
	BodyHtml       *string      `json:"body_html,omitempty"`
	CreatedAt      *time.Time   `json:"created_at,omitempty"`
	Handle         *string      `json:"handle,omitempty"`
	Id             *int         `json:"id,omitempty"`
	ProductType    *string      `json:"product_type,omitempty"`
	PublishedAt    *time.Time   `json:"published_at,omitempty"`
	PublishedScope *string      `json:"published_scope,omitempty"`
	TemplateSuffix *string      `json:"template_suffix,omitempty"`
	Title          *string      `json:"title,omitempty"`
	UpdatedAt      *time.Time   `json:"updated_at,omitempty"`
	Vendor         *string      `json:"vendor,omitempty"`
	Tags           *string      `json:"tags,omitempty"`
	Variants       []*Variant   `json:"variants,omitempty"`
	Options        []*Option    `json:"options,omitempty"`
	Images         []*Image     `json:"images,omitempty"`
	Metafields     []*Metafield `json:"metafields,omitempty"`
}

type Option struct {
	Id        *int    `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	Position  *int    `json:"position,omitempty"`
	ProductId *int    `json:"product_id,omitempty"`
}

type Variant struct {
	Barcode              *string    `json:"barcode,omitempty"`
	CompareAtPrice       *string    `json:"compare_at_price,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	FulfillmentService   *string    `json:"fulfillment_service,omitempty"`
	Grams                *int       `json:"grams,omitempty"`
	Id                   *int       `json:"id,omitempty"`
	InventoryManagement  *string    `json:"inventory_management,omitempty"`
	InventoryPolicy      *string    `json:"inventory_policy,omitempty"`
	Option1              *string    `json:"option1,omitempty"`
	Option2              *string    `json:"option2,omitempty"`
	Option3              *string    `json:"option3,omitempty"`
	Position             *int       `json:"position,omitempty"`
	Price                *string    `json:"price,omitempty"`
	ProductId            *int       `json:"product_id,omitempty"`
	RequiresShipping     *bool      `json:"requires_shipping,omitempty"`
	Sku                  *string    `json:"sku,omitempty"`
	Taxable              *bool      `json:"taxable,omitempty"`
	Title                *string    `json:"title,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	InventoryQuantity    *int       `json:"inventory_quantity,omitempty"`
	OldInventoryQuantity *int       `json:"old_inventory_quantity,omitempty"`
	ImageId              *int       `json:"image_id,omitempty"`
}

type Image struct {
	ID         *int       `json:"id,omitempty"`
	ProductID  *int       `json:"product_id,omitempty"`
	Position   *int       `json:"position,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	Width      *int       `json:"width,omitempty"`
	Height     *int       `json:"height,omitempty"`
	Src        *string    `json:"src,omitempty"`
	VariantIds []*int     `json:"variant_ids,omitempty"`
}

type ProductListOptions struct {
	Ids          []int     `url:"ids,comma,omitempty"`
	SinceId      int       `url:"since_id,omitempty"`
	Title        string    `url:"title,omitempty"`
	Vendor       string    `url:"vendor,omitempty"`
	Handle       string    `url:"handle,omitempty"`
	ProductType  string    `url:"product_type,omitempty"`
	CollectionId int       `url:"collection_id,omitempty"`
	Fields       []string  `url:"fields,comma,omitempty"`
	CreatedAtMin time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax time.Time `url:"updated_at_max,omitempty"`
}

func (p *ProductsService) List(ctx context.Context, opt *ProductListOptions) ([]*Product, *http.Response, error) {
	u := "products.json"
	u, err := addOptionsWithDefaults(u, opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var productList ProductList
	resp, err := p.client.Do(ctx, req, &productList)
	if err != nil {
		return nil, resp, err
	}

	return productList.Products, resp, nil
}

func (p *ProductsService) Edit(ctx context.Context, product *Product) (*http.Response, error) {
	innerProduct := &ProductForUpdateContainer{Product: product}
	u := fmt.Sprintf("products/%d.json", *product.Id)
	req, err := p.client.NewRequest("PUT", u, innerProduct)

	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
