package shopify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/google/go-querystring/query"
)


type Client struct {
	client *http.Client // HTTP client used to communicate with the API.

	rateLimit *rateLimiter
	// Base URL for API requests
	BaseURL *url.URL

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Shopify API.
	Products   *ProductsService
	Metafields *MetafieldsService
}

type service struct {
	client *Client
}

type resourceCount struct {
	Count int `json:"count,omitempty"`
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
type ListOptions struct {
	// For paginated result sets, page of results to retrieve.
	Page int `url:"page,omitempty"`

	// For paginated result sets, the number of results to include per page.
	Limit int `url:"limit,omitempty"`
}

func NewPrivateClient(httpClient *http.Client, apiKey string, password string, shopName string) *Client {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: time.Second * 20}
	}

	apiUrl := fmt.Sprintf("https://%s:%s@%s.myshopify.com/admin/", apiKey, password, shopName)
	baseURL, _ := url.Parse(apiUrl)
	c := &Client{
		rateLimit: rateLimitFor(shopName),
		client:    httpClient,
		BaseURL:   baseURL,
	}
	c.common.client = c
	c.Products = (*ProductsService)(&c.common)
	c.Metafields = (*MetafieldsService)(&c.common)
	return c
}

// addOptionsWithDefaults adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
// It also adds a url parameter "limit" of value "250"

func addOptionsWithDefaults(s string, opt interface{}) (string, error) {

	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}
	// Add Default
	qs.Add("limit", "250")
	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)
	// debug("url: %s", u)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	ctx, req = withContext(ctx, req)

	c.rateLimit.Wait()
	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()
	err = CheckResponse(resp)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}

	return resp, err
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return fmt.Errorf("Status returned: %d", r.StatusCode)
}
