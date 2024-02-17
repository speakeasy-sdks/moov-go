# Accounts
(*Accounts*)

## Overview

[Accounts](https://docs.moov.io/guides/accounts/) represent a legal entity (either a business or an individual) in Moov. You can create an account for yourself or set up accounts for others, requesting different [capabilities](https://docs.moov.io/guides/accounts/capabilities/) depending on what you need to be able to do with that account. You can retrieve an account to get details on the business or individual account holder, such as an email address or employer identification number (EIN). 

Based on the type of account and its requested capabilities, we have certain [verification requirements](https://docs.moov.io/guides/accounts/identity-verification/). To see what capabilities that account has, you can use the [GET capability endpoint](https://docs.moov.io/api/#operation/getCapability). 

When you sign up for the Moov Dashboard, you will have a **facilitator account** which can be used to facilitate money movement between other accounts. A facilitator account will not show up in your list of accounts and cannot be created via API. To update your facilitator account information, use the settings page of the Moov Dashboard. 


### Available Operations

* [AssignCountry](#assigncountry) - Assign Account Countries
* [Create](#create) - Create account
* [Get](#get) - Get account
* [GetTosToken](#gettostoken) - Get terms of service token
* [List](#list) - List accounts
* [ListCountries](#listcountries) - Get Account Countries
* [Update](#update) - Patch account

## AssignCountry

Assign the countries of operation for an account. This endpoint will always overwrite the previously assigned values. <br><br> To update the account countries, you'll need to specify the `/accounts/{accountID}/profile.write` scope.

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
            AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )


    countries := shared.Countries{
        Countries: []string{
            "United States",
        },
    }

    var accountID string = "f51a6841-150c-4bc7-8c84-e981f74cfa3f"

    ctx := context.Background()
    res, err := s.Accounts.AssignCountry(ctx, countries, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Countries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `countries`                                              | [shared.Countries](../../pkg/models/shared/countries.md) | :heavy_check_mark:                                       | N/A                                                      |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | ID of the account                                        |


### Response

**[*operations.PutAccountCountriesResponse](../../pkg/models/operations/putaccountcountriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Create

You can create accounts for your users by passing the required information to Moov. <br><br> Note that `mode` field is only required when creating a facilitator account. All non-facilitator account creation requests will ignore the mode field provided and be set to the calling facilitator's mode. <br><br> If you are creating an account with the business type "llc", "partnership", or "privateCorporation", you will need to also provide [business representatives](https://docs.moov.io/api/#tag/Representatives) after creating the account for verification purposes. Once you've added your business owners as representatives, you'll then need to [patch your Moov account](https://docs.moov.io/api/#operation/patchAccount) to indicate that ownership information is complete. Read more on our [business verification requirements here](https://docs.moov.io/guides/accounts/business-verification/). <br><br> When creating an account, you will need to specify the `/accounts.write` scope.

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
            AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Create(ctx, shared.CreateAccountRequest{
        AccountType: shared.AccountTypeBusiness,
        ForeignID: moovgo.String("4528aba-b9a1-11eb-8529-0242ac13003"),
        Mode: shared.ModeProduction.ToPointer(),
        Profile: shared.CreateProfile{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.CreateAccountRequest](../../pkg/models/shared/createaccountrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateAccountResponse](../../pkg/models/operations/createaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Retrieves details for the account with the specified ID. <br><br> To get an account, you will need to specify the `/accounts/{accountID}/profile.read` scope.

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
            AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )


    var accountID string = "b18d8d81-fd7b-4764-a31e-475cb1f36591"

    ctx := context.Background()
    res, err := s.Accounts.Get(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
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

**[*operations.GetAccountResponse](../../pkg/models/operations/getaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTosToken

Generates a non-expiring token that can then be used to accept Moov's terms of service. This token can only be generated via API. Any Moov account requesting the `collect-funds`, `send-funds`, `wallet`, or `card-issuing` capabilities must accept Moov's terms of service, then have the generated terms of service token patched to the account. Read more on our [quick start guide](https://docs.moov.io/guides/quick-start/#request-capabilities).

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
            AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.GetTosToken(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.TermsOfServiceToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetTermsOfServiceTokenResponse](../../pkg/models/operations/gettermsofservicetokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

List or search accounts to which the caller is connected.<br><br>
All supported query parameters are optional. If none are provided
the response will include all connected accounts. Pagination is
supported via the `skip` and `count` query parameters.<br><br>
Searching by name and email will overlap and return results based on relevance.
<br><br> To list connected accounts, you must specify the `/accounts.read` scope when retrieving the authentication token.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.List(ctx, operations.ListAccountsRequest{
        ForeignID: moovgo.String("4528aba-b9a1-11eb-8529-0242ac13003"),
        Type: shared.AccountTypeBusiness.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListAccountsRequest](../../pkg/models/operations/listaccountsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListAccountsResponse](../../pkg/models/operations/listaccountsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListCountries

Retrieve the specified countries of operation for an account. <br><br> To get the list of countries, you'll need to specify the `/accounts/{accountID}/profile.read` scope.

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
            AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )


    var accountID string = "5694ddc2-a16c-425b-bd39-0a53b9fefa9b"

    ctx := context.Background()
    res, err := s.Accounts.ListCountries(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Countries != nil {
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

**[*operations.GetAccountCountriesResponse](../../pkg/models/operations/getaccountcountriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

To patch an account, you must specify the `/accounts/{accountID}/profile.write` scope and provide the changed information.  

When **can** profile data be updated:  
  + For unverified accounts, all profile data can be edited.
  + During the verification process, missing or incomplete profile data can be edited.
  + Verified accounts can only add missing profile data.  

  When **can't** profile data be updated:  
  + Verified accounts cannot change any existing profile data.  

If you need to update information in a locked state, please contact Moov support.


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
            AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )


    patchAccountRequest := shared.PatchAccountRequest{
        ForeignID: moovgo.String("4528aba-b9a1-11eb-8529-0242ac13003"),
    }

    var accountID string = "d0905bf4-aa77-4f20-8e77-54c352acfe54"

    ctx := context.Background()
    res, err := s.Accounts.Update(ctx, patchAccountRequest, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `patchAccountRequest`                                                        | [shared.PatchAccountRequest](../../pkg/models/shared/patchaccountrequest.md) | :heavy_check_mark:                                                           | N/A                                                                          |
| `accountID`                                                                  | *string*                                                                     | :heavy_check_mark:                                                           | ID of the account                                                            |


### Response

**[*operations.PatchAccountResponse](../../pkg/models/operations/patchaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
