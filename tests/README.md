# goshopping tests #

This directory contains test that will exercise against the live Shopify API.
It is recommended that these tests only be run using a dedicated test account.

Run tests using:
    
<!--     SHOPIFY_ACCESS_TOKEN=XXX go test -v -tags=oauth

and/or  -->

    SHOPIFY_API_KEY=XXX SHOPIFY_PASSWORD=XXX SHOP_NAME=xxx go test -v -tags=private