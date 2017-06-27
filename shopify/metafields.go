package shopify

import (
	"context"

	"fmt"
	"net/http"
	"time"
)

type MetafieldsService service

type MetafieldList struct {
	Metafields []*Metafield `json:"metafields"`
}

type Metafield struct {
	CreatedAt     *time.Time `json:"created_at,omitempty"`
	Description   *string    `json:"description,omitempty"`
	Id            *int       `json:"id,omitempty"`
	Key           *string    `json:"key"`
	Namespace     *string    `json:"namespace,omitempty"`
	OwnerId       *int       `json:"owner_id,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	Value         *string    `json:"value"`
	ValueType     *string    `json:"value_type,omitempty"`
	OwnerResource *string    `json:"owner_resource,omitempty"`
}

type ProductForUpdateContainer struct {
	Product ProductForUpdate `json:"product"`
}

type ProductForUpdate struct {
	Id         int          `json:"id"`
	Metafields []*Metafield `json:"metafields"`
}

type MetafieldListOptions struct {
	SinceId      int       `url:"since_id,omitempty"`
	CreatedAtMin time.Time `url:"created_at_min,omitempty"`
	CreatedAtMax time.Time `url:"created_at_max,omitempty"`
	UpdatedAtMin time.Time `url:"updated_at_min,omitempty"`
	UpdatedAtMax time.Time `url:"updated_at_max,omitempty"`
	Namespace    string    `url:"namespace,omitempty"`
	Key          string    `url:"key,omitempty"`
	ValueType    string    `url:"value_type,omitempty"`
	Fields       []string  `url:"fields,comma,omitempty"`
}

func (p *MetafieldsService) ListByProduct(ctx context.Context, productId int, opt *MetafieldListOptions) ([]*Metafield, *http.Response, error) {
	u := fmt.Sprintf("products/%d/metafields.json", productId)
	u, err := addOptionsWithDefaults(u, opt)
	if err != nil {
		return nil, nil, err
	}
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var metafieldList MetafieldList
	resp, err := p.client.Do(ctx, req, &metafieldList)
	if err != nil {
		return nil, resp, err
	}

	return metafieldList.Metafields, resp, nil
}

func (p *MetafieldsService) UpdateByProduct(ctx context.Context, productId int, metafields []*Metafield) (*http.Response, error) {
	u := fmt.Sprintf("products/%d.json", productId)

	product := ProductForUpdateContainer{
		Product: ProductForUpdate{productId, metafields},
	}

	req, err := p.client.NewRequest("PUT", u, product)

	if err != nil {
		return nil, err
	}

	resp, err := p.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
