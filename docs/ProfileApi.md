# \ProfileApi

All URIs are relative to *https://api.lab5e.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NewMFAToken**](ProfileApi.md#NewMFAToken) | **Post** /dejavu/v1/mfa/new | Assign MFA token
[**Profile**](ProfileApi.md#Profile) | **Get** /dejavu/v1/profile | Retrieve profile
[**RemoveMFAToken**](ProfileApi.md#RemoveMFAToken) | **Post** /dejavu/v1/mfa/remove | Remove MFA token
[**UpdatePassword**](ProfileApi.md#UpdatePassword) | **Post** /dejavu/v1/profile/password | Change password
[**UpdateProfile**](ProfileApi.md#UpdateProfile) | **Post** /dejavu/v1/profile | Update profile
[**VerifyNewMFAToken**](ProfileApi.md#VerifyNewMFAToken) | **Post** /dejavu/v1/mfa/verify | Verify MFA token
[**VerifyPhone**](ProfileApi.md#VerifyPhone) | **Post** /dejavu/v1/phone/verify | Verify phone number



## NewMFAToken

> NewMFATokenResponse NewMFAToken(ctx).Body(body).Execute()

Assign MFA token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanuserapi"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.NewMFAToken(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.NewMFAToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NewMFAToken`: NewMFATokenResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.NewMFAToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNewMFATokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**NewMFATokenResponse**](NewMFATokenResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Profile

> ProfileResponse Profile(ctx).Execute()

Retrieve profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanuserapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.Profile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.Profile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Profile`: ProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.Profile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiProfileRequest struct via the builder pattern


### Return type

[**ProfileResponse**](ProfileResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMFAToken

> RemoveMFATokenResponse RemoveMFAToken(ctx).Body(body).Execute()

Remove MFA token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanuserapi"
)

func main() {
    body := *openapiclient.NewRemoveMFATokenRequest() // RemoveMFATokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.RemoveMFAToken(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.RemoveMFAToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveMFAToken`: RemoveMFATokenResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.RemoveMFAToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMFATokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RemoveMFATokenRequest**](RemoveMFATokenRequest.md) |  | 

### Return type

[**RemoveMFATokenResponse**](RemoveMFATokenResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePassword

> UpdatePasswordResponse UpdatePassword(ctx).Body(body).Execute()

Change password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanuserapi"
)

func main() {
    body := *openapiclient.NewUpdatePasswordRequest() // UpdatePasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.UpdatePassword(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.UpdatePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePassword`: UpdatePasswordResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.UpdatePassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdatePasswordRequest**](UpdatePasswordRequest.md) |  | 

### Return type

[**UpdatePasswordResponse**](UpdatePasswordResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProfile

> UpdateProfileResponse UpdateProfile(ctx).Body(body).Execute()

Update profile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanuserapi"
)

func main() {
    body := *openapiclient.NewUpdateProfileRequest() // UpdateProfileRequest | Update the user profile. The existing password is required in all cases. If an MFA token is assigned to the account it is required to update the profile information.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.UpdateProfile(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.UpdateProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfile`: UpdateProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UpdateProfileRequest**](UpdateProfileRequest.md) | Update the user profile. The existing password is required in all cases. If an MFA token is assigned to the account it is required to update the profile information. | 

### Return type

[**UpdateProfileResponse**](UpdateProfileResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyNewMFAToken

> VerifyNewMFATokenResponse VerifyNewMFAToken(ctx).Body(body).Execute()

Verify MFA token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanuserapi"
)

func main() {
    body := *openapiclient.NewVerifyNewMFATokenRequest() // VerifyNewMFATokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.VerifyNewMFAToken(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.VerifyNewMFAToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyNewMFAToken`: VerifyNewMFATokenResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.VerifyNewMFAToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyNewMFATokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VerifyNewMFATokenRequest**](VerifyNewMFATokenRequest.md) |  | 

### Return type

[**VerifyNewMFATokenResponse**](VerifyNewMFATokenResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPhone

> VerifyPhoneResponse VerifyPhone(ctx).Body(body).Execute()

Verify phone number



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/lab5e/go-spanuserapi"
)

func main() {
    body := *openapiclient.NewVerifyPhoneRequest() // VerifyPhoneRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProfileApi.VerifyPhone(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileApi.VerifyPhone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyPhone`: VerifyPhoneResponse
    fmt.Fprintf(os.Stdout, "Response from `ProfileApi.VerifyPhone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPhoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VerifyPhoneRequest**](VerifyPhoneRequest.md) |  | 

### Return type

[**VerifyPhoneResponse**](VerifyPhoneResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

