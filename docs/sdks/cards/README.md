# Cards

## Overview

You can link credit or debit cards to Moov accounts. You can use a card as a source for making transfers, which charges the card. To link a card to a Moov account and avoid some of the burden of PCI compliance, use the [card link Moov Drop](https://docs.moov.io/moovjs/drops/card-link). You cannot add a card via the Dashboard. If you're linking a card via API, you must provide Moov with a copy of your PCI attestation of compliance. When testing cards, use the designated [card numbers for test mode](https://docs.moov.io/guides/set-up-your-account/test-mode/#cards). You must contact Moov before going live in production with cards. Read our guide on [cards](https://docs.moov.io/guides/sources/cards/) for more information.

### Available Operations

* [LinkApplePayToken](#linkapplepaytoken) - Link Apple Pay token
* [LinkCard](#linkcard) - Link card
* [ListCards](#listcards) - List cards
* [CreateApplePaySession](#createapplepaysession) - Create Apple Pay session
* [Delete](#delete) - Disable card
* [Get](#get) - Get card
* [ListApplePayDomains](#listapplepaydomains) - Get Apple Pay domains
* [RegisterApplePayDomain](#registerapplepaydomain) - Register Apple Pay domains
* [Update](#update) - Update card
* [UpdateApplePayDomains](#updateapplepaydomains) - Update Apple Pay domains

## LinkApplePayToken

Connect an Apple Pay token to the specified account.
The `token` data is defined by Apple Pay and should be passed through from Apple Pay's response unmodified.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Cards.LinkApplePayToken(ctx, operations.PostLinkApplePayTokenRequest{
        LinkApplePay: shared.LinkApplePay{
            Token: shared.LinkApplePayToken{
                PaymentData: shared.LinkApplePayTokenPaymentData{
                    Data: "3+f4oOTwPa6f1UZ6tG...CE=",
                    Header: shared.LinkApplePayTokenPaymentDataHeader{
                        EphemeralPublicKey: "MFkwEK...Md==",
                        PublicKeyHash: "l0CnXdMv...D1I=",
                        TransactionID: "32b...4f3",
                    },
                    Signature: "MIAGCSqGSIb3DQ.AAAA==",
                    Version: "EC_v1",
                },
                PaymentMethod: shared.LinkApplePayTokenPaymentMethod{
                    DisplayName: "Visa 1234",
                    Network: "Visa",
                    Type: "debit",
                },
                TransactionIdentifier: "32b...4f3",
            },
        },
        AccountID: "3c7e0bc7-178e-4479-af2a-70c688282aa4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkedApplePayPaymentMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.PostLinkApplePayTokenRequest](../../models/operations/postlinkapplepaytokenrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.PostLinkApplePayTokenResponse](../../models/operations/postlinkapplepaytokenresponse.md), error**


## LinkCard

Link a card to an existing Moov account. Only use this endpoint if you have provided Moov with a copy of your PCI attestation of compliance. 
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Cards.LinkCard(ctx, operations.PostLinkCardRequest{
        CardRequest: shared.CardRequest{
            BillingAddress: &shared.Address{
                AddressLine1: moov.String("123 Main Street"),
                AddressLine2: moov.String("Apt 302"),
                City: moov.String("Boulder"),
                Country: moov.String("US"),
                PostalCode: moov.String("80301"),
                StateOrProvince: moov.String("CO"),
            },
            CardCvv: moov.String("0123"),
            CardNumber: moov.String("atque"),
            CardOnFile: moov.Bool(false),
            Expiration: &shared.CardExpiration{
                Month: moov.String("01"),
                Year: moov.String("21"),
            },
            HolderName: moov.String("Jules Jackson"),
            MerchantAccountID: moov.String("2562f222-e981-47ee-97cb-e61e6b7b95bc"),
        },
        XWaitFor: shared.SchemasWaitForPaymentMethod.ToPointer(),
        AccountID: "0ab3c20c-4f37-489f-9871-f99dd2efd121",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Card != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PostLinkCardRequest](../../models/operations/postlinkcardrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `opts`                                                                           | [][operations.Option](../../models/operations/option.md)                         | :heavy_minus_sign:                                                               | The options for this request.                                                    |


### Response

**[*operations.PostLinkCardResponse](../../models/operations/postlinkcardresponse.md), error**


## ListCards

List all the cards associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.read` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Cards.ListCards(ctx, operations.GetListCardsRequest{
        AccountID: "aa6f1e67-4bdb-404f-9575-6082d68ea19f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Cards != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetListCardsRequest](../../models/operations/getlistcardsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.GetListCardsResponse](../../models/operations/getlistcardsresponse.md), error**


## CreateApplePaySession

Create a session with Apple Pay to facilitate a payment.
A successful response from this endpoint should be passed through to Apple Pay unchanged.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.write` scope.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Cards.CreateApplePaySession(ctx, operations.PostApplePaySessionRequest{
        CreateApplePaySession: shared.CreateApplePaySession{
            DisplayName: "Example Merchant",
            Domain: "checkout.classbooker.dev",
        },
        AccountID: "1d170513-39d0-4808-aa18-40394c26071f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplePaySession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PostApplePaySessionRequest](../../models/operations/postapplepaysessionrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PostApplePaySessionResponse](../../models/operations/postapplepaysessionresponse.md), error**


## Delete

Disables a card associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Cards.Delete(ctx, operations.DeleteCardRequest{
        AccountID: "93f5f064-2dac-47af-915c-c413aa63aae8",
        CardID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.DeleteCardRequest](../../models/operations/deletecardrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.DeleteCardResponse](../../models/operations/deletecardresponse.md), error**


## Get

Fetch a specific card associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.read` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Cards.Get(ctx, operations.GetCardRequest{
        AccountID: "d67864db-b675-4fd5-a60b-375ed4f6fbee",
        CardID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Card != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [operations.GetCardRequest](../../models/operations/getcardrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.GetCardResponse](../../models/operations/getcardresponse.md), error**


## ListApplePayDomains

Get domains registered with Apple Pay.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.read` scope.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Cards.ListApplePayDomains(ctx, operations.GetApplePayMerchantDomainsRequest{
        AccountID: "41f33317-fe35-4b60-ab1e-a426555ba3c2",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplePayMerchantDomains != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetApplePayMerchantDomainsRequest](../../models/operations/getapplepaymerchantdomainsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetApplePayMerchantDomainsResponse](../../models/operations/getapplepaymerchantdomainsresponse.md), error**


## RegisterApplePayDomain

Add domains to be registered with Apple Pay.
<br><br> Any domains that will be used to accept payments must first be [verified](https://docs.moov.io/guides/money-movement/cards/apple-pay/#step-1-register-your-domains) with Apple.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.write` scope.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Cards.RegisterApplePayDomain(ctx, operations.PostApplePayMerchantDomainsRequest{
        RegisterApplePayMerchantDomains: shared.RegisterApplePayMerchantDomains{
            DisplayName: "Example Merchant",
            Domains: []string{
                "in",
                "dolore",
                "aliquam",
            },
        },
        AccountID: "ed53b88f-3a8d-48f5-80b2-f2fb7b194a27",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplePayMerchantDomains != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.PostApplePayMerchantDomainsRequest](../../models/operations/postapplepaymerchantdomainsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.PostApplePayMerchantDomainsResponse](../../models/operations/postapplepaymerchantdomainsresponse.md), error**


## Update

Update a Linked Card and/or resubmit it for verification. If a value is provided for CVV, 
a new verification ($0 authorization) will be submitted for the card. Updating the expiration date or 
address will update the information stored on file for the card but will not be verified. 
Only use this endpoint if you have provided Moov with a copy of your PCI attestation of compliance. 
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Cards.Update(ctx, operations.UpdateCardRequest{
        CardUpdateRequest: shared.CardUpdateRequest{
            BillingAddress: &shared.UpdateAddress{
                AddressLine1: moov.String("123 Main Street"),
                AddressLine2: moov.String("Apt 302"),
                City: moov.String("Boulder"),
                Country: moov.String("US"),
                PostalCode: moov.String("80301"),
                StateOrProvince: moov.String("CO"),
            },
            CardCvv: moov.String("123"),
            CardOnFile: moov.Bool(false),
            Expiration: &shared.UpdateCardExpiration{
                Month: moov.String("01"),
                Year: moov.String("21"),
            },
        },
        AccountID: "6b26916f-e1f0-48f4-a94e-3698f447f603",
        CardID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Card != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.UpdateCardRequest](../../models/operations/updatecardrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.UpdateCardResponse](../../models/operations/updatecardresponse.md), error**


## UpdateApplePayDomains

Add or remove domains to be registered with Apple Pay.
<br><br> Any domains that will be used to accept payments must first be [verified](https://docs.moov.io/guides/money-movement/cards/apple-pay/#step-1-register-your-domains) with Apple.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.write` scope.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Cards.UpdateApplePayDomains(ctx, operations.UpdateApplePayMerchantDomainsRequest{
        UpdateApplePayMerchantDomains: shared.UpdateApplePayMerchantDomains{
            AddDomains: []string{
                "praesentium",
                "facilis",
                "quaerat",
                "incidunt",
            },
            RemoveDomains: []string{
                "debitis",
                "rem",
            },
        },
        AccountID: "0ca55efd-20e4-457e-9858-b6a89fbe3a5a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.UpdateApplePayMerchantDomainsRequest](../../models/operations/updateapplepaymerchantdomainsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.UpdateApplePayMerchantDomainsResponse](../../models/operations/updateapplepaymerchantdomainsresponse.md), error**

