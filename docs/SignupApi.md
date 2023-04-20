# \SignupApi

All URIs are relative to *https://api.lab5e.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndUserLicenseAgreement**](SignupApi.md#EndUserLicenseAgreement) | **Get** /dejavu/v1/eula | EULA text
[**GetGithubProfile**](SignupApi.md#GetGithubProfile) | **Get** /dejavu/v1/ghprofile/{githubToken} | Retrieve GitHub profile
[**Signup**](SignupApi.md#Signup) | **Post** /dejavu/v1/signup | Sign up new user
[**UpdateGithubProfile**](SignupApi.md#UpdateGithubProfile) | **Post** /dejavu/v1/ghprofile/{githubToken} | Update GitHub profile
[**VerifyEmail**](SignupApi.md#VerifyEmail) | **Post** /dejavu/v1/verifyemail | Verify email
[**VerifyGithubProfile**](SignupApi.md#VerifyGithubProfile) | **Post** /dejavu/v1/ghprofile/{emailToken}/verify | Verify GitHub profile



## EndUserLicenseAgreement

> EndUserLicenseAgreementResponse EndUserLicenseAgreement(ctx).Format(format).Execute()

EULA text



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
    format := "format_example" // string |  (optional) (default to "FORMAT_UNSPECIFIED")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignupApi.EndUserLicenseAgreement(context.Background()).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupApi.EndUserLicenseAgreement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EndUserLicenseAgreement`: EndUserLicenseAgreementResponse
    fmt.Fprintf(os.Stdout, "Response from `SignupApi.EndUserLicenseAgreement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEndUserLicenseAgreementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | [default to &quot;FORMAT_UNSPECIFIED&quot;]

### Return type

[**EndUserLicenseAgreementResponse**](EndUserLicenseAgreementResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGithubProfile

> GetGithubProfileResponse GetGithubProfile(ctx, githubToken).Execute()

Retrieve GitHub profile



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
    githubToken := "githubToken_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignupApi.GetGithubProfile(context.Background(), githubToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupApi.GetGithubProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGithubProfile`: GetGithubProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `SignupApi.GetGithubProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGithubProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetGithubProfileResponse**](GetGithubProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Signup

> SignupResponse Signup(ctx).Body(body).Execute()

Sign up new user



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
    body := *openapiclient.NewSignupRequest() // SignupRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignupApi.Signup(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupApi.Signup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Signup`: SignupResponse
    fmt.Fprintf(os.Stdout, "Response from `SignupApi.Signup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SignupRequest**](SignupRequest.md) |  | 

### Return type

[**SignupResponse**](SignupResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGithubProfile

> map[string]interface{} UpdateGithubProfile(ctx, githubToken).Body(body).Execute()

Update GitHub profile



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
    githubToken := "githubToken_example" // string | 
    body := *openapiclient.NewUpdateGithubProfileRequest() // UpdateGithubProfileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignupApi.UpdateGithubProfile(context.Background(), githubToken).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupApi.UpdateGithubProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGithubProfile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `SignupApi.UpdateGithubProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGithubProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateGithubProfileRequest**](UpdateGithubProfileRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyEmail

> VerifyEmailResponse VerifyEmail(ctx).Body(body).Execute()

Verify email



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
    body := *openapiclient.NewVerifyEmailRequest() // VerifyEmailRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignupApi.VerifyEmail(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupApi.VerifyEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyEmail`: VerifyEmailResponse
    fmt.Fprintf(os.Stdout, "Response from `SignupApi.VerifyEmail`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VerifyEmailRequest**](VerifyEmailRequest.md) |  | 

### Return type

[**VerifyEmailResponse**](VerifyEmailResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyGithubProfile

> VerifyGithubProfileResponse VerifyGithubProfile(ctx, emailToken).Body(body).Execute()

Verify GitHub profile



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
    emailToken := "emailToken_example" // string | 
    body := *openapiclient.NewVerifyGithubProfileRequest() // VerifyGithubProfileRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SignupApi.VerifyGithubProfile(context.Background(), emailToken).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SignupApi.VerifyGithubProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyGithubProfile`: VerifyGithubProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `SignupApi.VerifyGithubProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailToken** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyGithubProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VerifyGithubProfileRequest**](VerifyGithubProfileRequest.md) |  | 

### Return type

[**VerifyGithubProfileResponse**](VerifyGithubProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

