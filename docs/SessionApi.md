# \SessionApi

All URIs are relative to *https://api.lab5e.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListSessions**](SessionApi.md#ListSessions) | **Post** /dejavu/v1/sessions | List sessions
[**Login**](SessionApi.md#Login) | **Post** /dejavu/v1/login | Sign in to the service
[**Logout**](SessionApi.md#Logout) | **Post** /dejavu/v1/logout | Sign out
[**RecoverPassword**](SessionApi.md#RecoverPassword) | **Post** /dejavu/v1/recover | Recover/reset password



## ListSessions

> ListSessionsResponse ListSessions(ctx).Body(body).Execute()

List sessions



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
    resp, r, err := apiClient.SessionApi.ListSessions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionApi.ListSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessions`: ListSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionApi.ListSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**ListSessionsResponse**](ListSessionsResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Login

> LoginResponse Login(ctx).Body(body).Execute()

Sign in to the service



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
    body := *openapiclient.NewLoginRequest() // LoginRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionApi.Login(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionApi.Login``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Login`: LoginResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionApi.Login`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LoginRequest**](LoginRequest.md) |  | 

### Return type

[**LoginResponse**](LoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Logout

> map[string]interface{} Logout(ctx).Body(body).Execute()

Sign out



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
    resp, r, err := apiClient.SessionApi.Logout(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionApi.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Logout`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SessionApi.Logout`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoverPassword

> RecoverPasswordResponse RecoverPassword(ctx).Body(body).Execute()

Recover/reset password



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
    body := *openapiclient.NewRecoverPasswordRequest() // RecoverPasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SessionApi.RecoverPassword(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SessionApi.RecoverPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoverPassword`: RecoverPasswordResponse
    fmt.Fprintf(os.Stdout, "Response from `SessionApi.RecoverPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecoverPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RecoverPasswordRequest**](RecoverPasswordRequest.md) |  | 

### Return type

[**RecoverPasswordResponse**](RecoverPasswordResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

