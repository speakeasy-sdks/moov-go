# Cards
(*Cards*)

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
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    linkApplePay := shared.LinkApplePay{
        Token: shared.Token{
            PaymentData: shared.PaymentData{
                Data: "3+f4oOTwPa6f1UZ6tG...CE=",
                Header: shared.Header{
                    EphemeralPublicKey: "MFkwEK...Md==",
                    PublicKeyHash: "l0CnXdMv...D1I=",
                    TransactionID: "32b...4f3",
                },
                Signature: "MIAGCSqGSIb3DQ.AAAA==",
                Version: "EC_v1",
            },
            PaymentMethod: shared.LinkApplePayPaymentMethod{
                DisplayName: "Visa 1234",
                Network: "Visa",
                Type: "debit",
            },
            TransactionIdentifier: "32b...4f3",
        },
    }

    var accountID string = "1f061848-80f9-474a-a77b-17cb1fda3296"

    ctx := context.Background()
    res, err := s.Cards.LinkApplePayToken(ctx, linkApplePay, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.LinkedApplePayPaymentMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `linkApplePay`                                                    | [shared.LinkApplePay](../../../pkg/models/shared/linkapplepay.md) | :heavy_check_mark:                                                | N/A                                                               |
| `accountID`                                                       | *string*                                                          | :heavy_check_mark:                                                | ID of the account                                                 |


### Response

**[*operations.PostLinkApplePayTokenResponse](../../pkg/models/operations/postlinkapplepaytokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## LinkCard

Link a card to an existing Moov account. Only use this endpoint if you have provided Moov with a copy of your PCI attestation of compliance. 
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    cardRequest := shared.CardRequest{
        BillingAddress: &shared.Address{
            AddressLine1: moovgo.String("123 Main Street"),
            AddressLine2: moovgo.String("Apt 302"),
            City: moovgo.String("Boulder"),
            Country: moovgo.String("US"),
            PostalCode: moovgo.String("80301"),
            StateOrProvince: moovgo.String("CO"),
        },
        CardCvv: moovgo.String("0123"),
        Expiration: &shared.CardExpiration{
            Month: moovgo.String("01"),
            Year: moovgo.String("21"),
        },
        HolderName: moovgo.String("Jules Jackson"),
    }

    var accountID string = "8cfd9cf0-8cf1-4dc7-b48b-a0e013b33b1b"

    var xWaitFor *shared.SchemasWaitFor = shared.SchemasWaitForPaymentMethod

    ctx := context.Background()
    res, err := s.Cards.LinkCard(ctx, cardRequest, accountID, xWaitFor)
    if err != nil {
        log.Fatal(err)
    }

    if res.Card != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `cardRequest`                                                                                        | [shared.CardRequest](../../../pkg/models/shared/cardrequest.md)                                      | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `accountID`                                                                                          | *string*                                                                                             | :heavy_check_mark:                                                                                   | ID of the account                                                                                    |
| `xWaitFor`                                                                                           | [*shared.SchemasWaitFor](../../../pkg/models/shared/schemaswaitfor.md)                               | :heavy_minus_sign:                                                                                   | Optional header that indicates whether to return a synchronous response or an asynchronous response. |
| `opts`                                                                                               | [][operations.Option](../../pkg/models/operations/option.md)                                         | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


### Response

**[*operations.PostLinkCardResponse](../../pkg/models/operations/postlinkcardresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListCards

List all the cards associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.read` scope.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "3540bb23-258b-486f-9192-4d7f6c3c94e4"

    ctx := context.Background()
    res, err := s.Cards.ListCards(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Cards != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |


### Response

**[*operations.GetListCardsResponse](../../pkg/models/operations/getlistcardsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## CreateApplePaySession

Create a session with Apple Pay to facilitate a payment.
A successful response from this endpoint should be passed through to Apple Pay unchanged.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.write` scope.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    createApplePaySession := shared.CreateApplePaySession{
        DisplayName: "Example Merchant",
        Domain: "checkout.classbooker.dev",
    }

    var accountID string = "94160d72-b710-419d-89b1-719f72571386"

    ctx := context.Background()
    res, err := s.Cards.CreateApplePaySession(ctx, createApplePaySession, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplePaySession != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                           | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `ctx`                                                                               | [context.Context](https://pkg.go.dev/context#Context)                               | :heavy_check_mark:                                                                  | The context to use for the request.                                                 |
| `createApplePaySession`                                                             | [shared.CreateApplePaySession](../../../pkg/models/shared/createapplepaysession.md) | :heavy_check_mark:                                                                  | N/A                                                                                 |
| `accountID`                                                                         | *string*                                                                            | :heavy_check_mark:                                                                  | ID of the account                                                                   |


### Response

**[*operations.PostApplePaySessionResponse](../../pkg/models/operations/postapplepaysessionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Delete

Disables a card associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "8db863f6-ef9b-413a-8a70-cb816b33de6b"

    var cardID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.Cards.Delete(ctx, accountID, cardID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `cardID`                                              | *string*                                              | :heavy_check_mark:                                    | ID of the card                                        | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


### Response

**[*operations.DeleteCardResponse](../../pkg/models/operations/deletecardresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Get

Fetch a specific card associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.read` scope.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "b18d8d81-fd7b-4764-a31e-475cb1f36591"

    var cardID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.Cards.Get(ctx, accountID, cardID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Card != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `cardID`                                              | *string*                                              | :heavy_check_mark:                                    | ID of the card                                        | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


### Response

**[*operations.GetCardResponse](../../pkg/models/operations/getcardresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListApplePayDomains

Get domains registered with Apple Pay.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.read` scope.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "4611290f-802a-4b57-a6c6-be763a3142ab"

    ctx := context.Background()
    res, err := s.Cards.ListApplePayDomains(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplePayMerchantDomains != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |


### Response

**[*operations.GetApplePayMerchantDomainsResponse](../../pkg/models/operations/getapplepaymerchantdomainsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## RegisterApplePayDomain

Add domains to be registered with Apple Pay.
<br><br> Any domains that will be used to accept payments must first be [verified](https://docs.moov.io/guides/money-movement/cards/apple-pay/#step-1-register-your-domains) with Apple.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.write` scope.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    registerApplePayMerchantDomains := shared.RegisterApplePayMerchantDomains{
        DisplayName: "Example Merchant",
        Domains: []string{
            "checkout.classbooker.dev",
        },
    }

    var accountID string = "a7f34d07-3cbe-4fbd-9d16-8b2a9900d9ad"

    ctx := context.Background()
    res, err := s.Cards.RegisterApplePayDomain(ctx, registerApplePayMerchantDomains, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.ApplePayMerchantDomains != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                               | Type                                                                                                    | Required                                                                                                | Description                                                                                             |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                   | [context.Context](https://pkg.go.dev/context#Context)                                                   | :heavy_check_mark:                                                                                      | The context to use for the request.                                                                     |
| `registerApplePayMerchantDomains`                                                                       | [shared.RegisterApplePayMerchantDomains](../../../pkg/models/shared/registerapplepaymerchantdomains.md) | :heavy_check_mark:                                                                                      | N/A                                                                                                     |
| `accountID`                                                                                             | *string*                                                                                                | :heavy_check_mark:                                                                                      | ID of the account                                                                                       |


### Response

**[*operations.PostApplePayMerchantDomainsResponse](../../pkg/models/operations/postapplepaymerchantdomainsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    cardUpdateRequest := shared.CardUpdateRequest{
        BillingAddress: &shared.UpdateAddress{
            AddressLine1: moovgo.String("123 Main Street"),
            AddressLine2: moovgo.String("Apt 302"),
            City: moovgo.String("Boulder"),
            Country: moovgo.String("US"),
            PostalCode: moovgo.String("80301"),
            StateOrProvince: moovgo.String("CO"),
        },
        CardCvv: moovgo.String("123"),
        Expiration: &shared.UpdateCardExpiration{
            Month: moovgo.String("01"),
            Year: moovgo.String("21"),
        },
    }

    var accountID string = "d0905bf4-aa77-4f20-8e77-54c352acfe54"

    var cardID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.Cards.Update(ctx, cardUpdateRequest, accountID, cardID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Card != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                   | Type                                                                        | Required                                                                    | Description                                                                 | Example                                                                     |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ctx`                                                                       | [context.Context](https://pkg.go.dev/context#Context)                       | :heavy_check_mark:                                                          | The context to use for the request.                                         |                                                                             |
| `cardUpdateRequest`                                                         | [shared.CardUpdateRequest](../../../pkg/models/shared/cardupdaterequest.md) | :heavy_check_mark:                                                          | N/A                                                                         |                                                                             |
| `accountID`                                                                 | *string*                                                                    | :heavy_check_mark:                                                          | ID of the account                                                           |                                                                             |
| `cardID`                                                                    | *string*                                                                    | :heavy_check_mark:                                                          | ID of the card                                                              | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                        |


### Response

**[*operations.UpdateCardResponse](../../pkg/models/operations/updatecardresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateApplePayDomains

Add or remove domains to be registered with Apple Pay.
<br><br> Any domains that will be used to accept payments must first be [verified](https://docs.moov.io/guides/money-movement/cards/apple-pay/#step-1-register-your-domains) with Apple.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.write` scope.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    updateApplePayMerchantDomains := shared.UpdateApplePayMerchantDomains{
        AddDomains: []string{
            "pay.classbooker.dev",
        },
        RemoveDomains: []string{
            "checkout.classbooker.dev",
        },
    }

    var accountID string = "8073d21a-7f13-415c-bf5b-aa0e8c6d7b6b"

    ctx := context.Background()
    res, err := s.Cards.UpdateApplePayDomains(ctx, updateApplePayMerchantDomains, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                           | Type                                                                                                | Required                                                                                            | Description                                                                                         |
| --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                               | :heavy_check_mark:                                                                                  | The context to use for the request.                                                                 |
| `updateApplePayMerchantDomains`                                                                     | [shared.UpdateApplePayMerchantDomains](../../../pkg/models/shared/updateapplepaymerchantdomains.md) | :heavy_check_mark:                                                                                  | N/A                                                                                                 |
| `accountID`                                                                                         | *string*                                                                                            | :heavy_check_mark:                                                                                  | ID of the account                                                                                   |


### Response

**[*operations.UpdateApplePayMerchantDomainsResponse](../../pkg/models/operations/updateapplepaymerchantdomainsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
