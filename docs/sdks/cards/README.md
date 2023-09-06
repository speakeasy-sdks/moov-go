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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    linkApplePay := shared.LinkApplePay{
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
    }
    accountID := "2e9817ee-17cb-4e61-a6b7-b95bc0ab3c20"

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

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `linkApplePay`                                             | [shared.LinkApplePay](../../models/shared/linkapplepay.md) | :heavy_check_mark:                                         | N/A                                                        |
| `accountID`                                                | *string*                                                   | :heavy_check_mark:                                         | ID of the account                                          |


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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    cardRequest := shared.CardRequest{
        BillingAddress: &shared.Address{
            AddressLine1: moov.String("123 Main Street"),
            AddressLine2: moov.String("Apt 302"),
            City: moov.String("Boulder"),
            Country: moov.String("US"),
            PostalCode: moov.String("80301"),
            StateOrProvince: moov.String("CO"),
        },
        CardCvv: moov.String("0123"),
        CardNumber: moov.String("minus"),
        CardOnFile: moov.Bool(false),
        Expiration: &shared.CardExpiration{
            Month: moov.String("01"),
            Year: moov.String("21"),
        },
        HolderName: moov.String("Jules Jackson"),
        MerchantAccountID: moov.String("4f3789fd-871f-499d-92ef-d121aa6f1e67"),
    }
    accountID := "4bdb04f1-5756-4082-968e-a19f1d170513"
    xWaitFor := shared.SchemasWaitForPaymentMethod

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
| `cardRequest`                                                                                        | [shared.CardRequest](../../models/shared/cardrequest.md)                                             | :heavy_check_mark:                                                                                   | N/A                                                                                                  |
| `accountID`                                                                                          | *string*                                                                                             | :heavy_check_mark:                                                                                   | ID of the account                                                                                    |
| `xWaitFor`                                                                                           | [*shared.SchemasWaitFor](../../models/shared/schemaswaitfor.md)                                      | :heavy_minus_sign:                                                                                   | Optional header that indicates whether to return a synchronous response or an asynchronous response. |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |


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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    accountID := "39d08086-a184-4039-8c26-071f93f5f064"

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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    createApplePaySession := shared.CreateApplePaySession{
        DisplayName: "Example Merchant",
        Domain: "checkout.classbooker.dev",
    }
    accountID := "2dac7af5-15cc-4413-aa63-aae8d67864db"

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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `createApplePaySession`                                                      | [shared.CreateApplePaySession](../../models/shared/createapplepaysession.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `accountID`                                                                  | *string*                                                                     | :heavy_check_mark:                                                           | ID of the account                                                            |


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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    accountID := "b675fd5e-60b3-475e-94f6-fbee41f33317"
    cardID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

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

**[*operations.DeleteCardResponse](../../models/operations/deletecardresponse.md), error**


## Get

Fetch a specific card associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.read` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    accountID := "fe35b60e-b1ea-4426-955b-a3c28744ed53"
    cardID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    accountID := "b88f3a8d-8f5c-40b2-b2fb-7b194a276b26"

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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    registerApplePayMerchantDomains := shared.RegisterApplePayMerchantDomains{
        DisplayName: "Example Merchant",
        Domains: []string{
            "unde",
        },
    }
    accountID := "16fe1f08-f429-44e3-a98f-447f603e8b44"

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

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `registerApplePayMerchantDomains`                                                                | [shared.RegisterApplePayMerchantDomains](../../models/shared/registerapplepaymerchantdomains.md) | :heavy_check_mark:                                                                               | N/A                                                                                              |
| `accountID`                                                                                      | *string*                                                                                         | :heavy_check_mark:                                                                               | ID of the account                                                                                |


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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    cardUpdateRequest := shared.CardUpdateRequest{
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
    }
    accountID := "5e80ca55-efd2-40e4-97e1-858b6a89fbe3"
    cardID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

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

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |                                                                      |
| `cardUpdateRequest`                                                  | [shared.CardUpdateRequest](../../models/shared/cardupdaterequest.md) | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |
| `accountID`                                                          | *string*                                                             | :heavy_check_mark:                                                   | ID of the account                                                    |                                                                      |
| `cardID`                                                             | *string*                                                             | :heavy_check_mark:                                                   | ID of the card                                                       | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                 |


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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    updateApplePayMerchantDomains := shared.UpdateApplePayMerchantDomains{
        AddDomains: []string{
            "dolorum",
        },
        RemoveDomains: []string{
            "nostrum",
        },
    }
    accountID := "aa8e4824-d0ab-4407-9088-e51862065e90"

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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `updateApplePayMerchantDomains`                                                              | [shared.UpdateApplePayMerchantDomains](../../models/shared/updateapplepaymerchantdomains.md) | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `accountID`                                                                                  | *string*                                                                                     | :heavy_check_mark:                                                                           | ID of the account                                                                            |


### Response

**[*operations.UpdateApplePayMerchantDomainsResponse](../../models/operations/updateapplepaymerchantdomainsresponse.md), error**

