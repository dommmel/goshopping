package shopify

import "testing"

func TestNewPrivateClient(t *testing.T) {
	c := NewPrivateClient(nil, "apiKey", "password", "shopName")

	if got, want := c.BaseURL.String(), "https://apiKey:password@shopName.myshopify.com/admin/"; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}

}
