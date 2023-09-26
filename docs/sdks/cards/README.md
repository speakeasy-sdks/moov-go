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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
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
    accountID := "4f157560-82d6-48ea-99f1-d17051339d08"

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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
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
        CardNumber: moovgo.String("aut"),
        CardOnFile: moovgo.Bool(false),
        Expiration: &shared.CardExpiration{
            Month: moovgo.String("01"),
            Year: moovgo.String("21"),
        },
        HolderName: moovgo.String("Jules Jackson"),
        MerchantAccountID: moovgo.String("86a18403-94c2-4607-9f93-f5f0642dac7a"),
    }
    accountID := "f515cc41-3aa6-43aa-a8d6-7864dbb675fd"
    xWaitFor := "nemo"

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
| `xWaitFor`                                                                                           | **string*                                                                                            | :heavy_minus_sign:                                                                                   | Optional header that indicates whether to return a synchronous response or an asynchronous response. |
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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    accountID := "e60b375e-d4f6-4fbe-a41f-33317fe35b60"

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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
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
    accountID := "eb1ea426-555b-4a3c-a874-4ed53b88f3a8"

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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    accountID := "d8f5c0b2-f2fb-47b1-94a2-76b26916fe1f"
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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    accountID := "08f4294e-3698-4f44-bf60-3e8b445e80ca"
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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    accountID := "55efd20e-457e-4185-8b6a-89fbe3a5aa8e"

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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
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
            "tempora",
        },
    }
    accountID := "824d0ab4-0750-488e-9186-2065e904f3b1"

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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
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
        CardOnFile: moovgo.Bool(false),
        Expiration: &shared.UpdateCardExpiration{
            Month: moovgo.String("01"),
            Year: moovgo.String("21"),
        },
    }
    accountID := "194b8abf-603a-479f-9dfe-0ab7da8a50ce"
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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    updateApplePayMerchantDomains := shared.UpdateApplePayMerchantDomains{
        AddDomains: []string{
            "quasi",
        },
        RemoveDomains: []string{
            "atque",
        },
    }
    accountID := "7f86bc17-3d68-49ee-a952-6f8d986e881e"

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

