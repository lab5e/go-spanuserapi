# \OrgsApi

All URIs are relative to *https://api.lab5e.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAPIToken**](OrgsApi.md#AddAPIToken) | **Post** /dejavu/v1/orgs/{orgId}/teams/{teamId}/apitokens | Add API token to team
[**AddTeamMember**](OrgsApi.md#AddTeamMember) | **Post** /dejavu/v1/orgs/{orgId}/teams/{teamId}/members | Add a team member
[**CreateOrg**](OrgsApi.md#CreateOrg) | **Post** /dejavu/v1/orgs | Create new organisation
[**CreateOrgInvites**](OrgsApi.md#CreateOrgInvites) | **Post** /dejavu/v1/orgs/{orgId}/invites | Invite members into org
[**CreateOrgTeam**](OrgsApi.md#CreateOrgTeam) | **Post** /dejavu/v1/orgs/{orgId}/teams | Create a team
[**DeleteOrg**](OrgsApi.md#DeleteOrg) | **Delete** /dejavu/v1/orgs/{orgId} | Delete organisation
[**DeleteOrgTeam**](OrgsApi.md#DeleteOrgTeam) | **Delete** /dejavu/v1/orgs/{orgId}/teams/{teamId} | Delete a team
[**GetAPIToken**](OrgsApi.md#GetAPIToken) | **Get** /dejavu/v1/orgs/{orgId}/teams/{teamId}/apitokens/{tokenId} | Retrieve API token
[**GetOrg**](OrgsApi.md#GetOrg) | **Get** /dejavu/v1/orgs/{orgId} | Organisation details
[**GetOrgTeam**](OrgsApi.md#GetOrgTeam) | **Get** /dejavu/v1/orgs/{orgId}/teams/{teamId} | Retrieve team information
[**HandleInvite**](OrgsApi.md#HandleInvite) | **Patch** /dejavu/v1/orgs/invites/{inviteId} | Respond to invite
[**ListAPITokens**](OrgsApi.md#ListAPITokens) | **Get** /dejavu/v1/orgs/{orgId}/teams/apitokens | List API tokens for teams in organisation
[**ListInvites**](OrgsApi.md#ListInvites) | **Post** /dejavu/v1/orgs/invites | List invites
[**ListOrgMembers**](OrgsApi.md#ListOrgMembers) | **Get** /dejavu/v1/orgs/{orgId}/members | List members of organisation
[**ListOrgTeams**](OrgsApi.md#ListOrgTeams) | **Get** /dejavu/v1/orgs/{orgId}/teams | List teams in org
[**ListOrgs**](OrgsApi.md#ListOrgs) | **Get** /dejavu/v1/orgs | List organisations
[**RemoveAPIToken**](OrgsApi.md#RemoveAPIToken) | **Delete** /dejavu/v1/orgs/{orgId}/teams/{teamId}/apitokens/{tokenId} | Remove API token from team
[**RemoveOrgMember**](OrgsApi.md#RemoveOrgMember) | **Delete** /dejavu/v1/orgs/{orgId}/members/{memberId} | Remove member from organisation
[**RemoveTeamMember**](OrgsApi.md#RemoveTeamMember) | **Delete** /dejavu/v1/orgs/{orgId}/teams/{teamId}/members/{memberId} | Remove a member from a team
[**UpdateAPIToken**](OrgsApi.md#UpdateAPIToken) | **Patch** /dejavu/v1/orgs/{orgId}/teams/{teamId}/apitokens/{tokenId} | Update API token
[**UpdateOrg**](OrgsApi.md#UpdateOrg) | **Patch** /dejavu/v1/orgs/{orgId} | Update organisation
[**UpdateOrgMemberRole**](OrgsApi.md#UpdateOrgMemberRole) | **Patch** /dejavu/v1/orgs/{orgId}/members/{memberId} | Update org member role
[**UpdateOrgTeam**](OrgsApi.md#UpdateOrgTeam) | **Patch** /dejavu/v1/orgs/{orgId}/teams/{teamId} | Update a team
[**UpdateTeamMember**](OrgsApi.md#UpdateTeamMember) | **Patch** /dejavu/v1/orgs/{orgId}/teams/{teamId}/members | Change membership role for a team member



## AddAPIToken

> APIToken AddAPIToken(ctx, orgId, teamId).Body(body).Execute()

Add API token to team



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
    orgId := "orgId_example" // string | 
    teamId := "teamId_example" // string | 
    body := *openapiclient.NewAddAPITokenRequest() // AddAPITokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.AddAPIToken(context.Background(), orgId, teamId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.AddAPIToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAPIToken`: APIToken
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.AddAPIToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAPITokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**AddAPITokenRequest**](AddAPITokenRequest.md) |  | 

### Return type

[**APIToken**](APIToken.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddTeamMember

> OrgTeam AddTeamMember(ctx, orgId, teamId).Body(body).Execute()

Add a team member



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
    orgId := "orgId_example" // string | 
    teamId := "teamId_example" // string | 
    body := *openapiclient.NewTeamMemberRequest() // TeamMemberRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.AddTeamMember(context.Background(), orgId, teamId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.AddTeamMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTeamMember`: OrgTeam
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.AddTeamMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**TeamMemberRequest**](TeamMemberRequest.md) |  | 

### Return type

[**OrgTeam**](OrgTeam.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrg

> Organisation CreateOrg(ctx).Body(body).Execute()

Create new organisation



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
    body := *openapiclient.NewCreateOrgRequest() // CreateOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.CreateOrg(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.CreateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrg`: Organisation
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.CreateOrg`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateOrgRequest**](CreateOrgRequest.md) |  | 

### Return type

[**Organisation**](Organisation.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgInvites

> CreateOrgInvitesResponse CreateOrgInvites(ctx, orgId).Body(body).Execute()

Invite members into org



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
    orgId := "orgId_example" // string | 
    body := *openapiclient.NewCreateOrgInvitesRequest() // CreateOrgInvitesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.CreateOrgInvites(context.Background(), orgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.CreateOrgInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrgInvites`: CreateOrgInvitesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.CreateOrgInvites`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrgInvitesRequest**](CreateOrgInvitesRequest.md) |  | 

### Return type

[**CreateOrgInvitesResponse**](CreateOrgInvitesResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrgTeam

> OrgTeam CreateOrgTeam(ctx, orgId).Body(body).Execute()

Create a team



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
    orgId := "orgId_example" // string | 
    body := *openapiclient.NewCreateOrgTeamRequest() // CreateOrgTeamRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.CreateOrgTeam(context.Background(), orgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.CreateOrgTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrgTeam`: OrgTeam
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.CreateOrgTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateOrgTeamRequest**](CreateOrgTeamRequest.md) |  | 

### Return type

[**OrgTeam**](OrgTeam.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrg

> map[string]interface{} DeleteOrg(ctx, orgId).Execute()

Delete organisation



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
    orgId := "orgId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.DeleteOrg(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.DeleteOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrg`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.DeleteOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgTeam

> map[string]interface{} DeleteOrgTeam(ctx, orgId, teamId).Execute()

Delete a team



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
    orgId := "orgId_example" // string | Organisation ID
    teamId := "teamId_example" // string | Team ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.DeleteOrgTeam(context.Background(), orgId, teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.DeleteOrgTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgTeam`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.DeleteOrgTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organisation ID | 
**teamId** | **string** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAPIToken

> APIToken GetAPIToken(ctx, orgId, teamId, tokenId).Execute()

Retrieve API token



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
    orgId := "orgId_example" // string | 
    teamId := "teamId_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.GetAPIToken(context.Background(), orgId, teamId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.GetAPIToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAPIToken`: APIToken
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.GetAPIToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**teamId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAPITokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**APIToken**](APIToken.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrg

> Organisation GetOrg(ctx, orgId).Execute()

Organisation details



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
    orgId := "orgId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.GetOrg(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.GetOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrg`: Organisation
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.GetOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Organisation**](Organisation.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgTeam

> OrgTeam GetOrgTeam(ctx, orgId, teamId).Execute()

Retrieve team information



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
    orgId := "orgId_example" // string | 
    teamId := "teamId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.GetOrgTeam(context.Background(), orgId, teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.GetOrgTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgTeam`: OrgTeam
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.GetOrgTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrgTeam**](OrgTeam.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HandleInvite

> map[string]interface{} HandleInvite(ctx, inviteId).Body(body).Execute()

Respond to invite



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
    inviteId := "inviteId_example" // string | The invitie ID
    body := *openapiclient.NewHandleInviteRequest() // HandleInviteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.HandleInvite(context.Background(), inviteId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.HandleInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HandleInvite`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.HandleInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inviteId** | **string** | The invitie ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiHandleInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**HandleInviteRequest**](HandleInviteRequest.md) |  | 

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


## ListAPITokens

> ListAPITokenResponse ListAPITokens(ctx, orgId).TeamId(teamId).Execute()

List API tokens for teams in organisation



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
    orgId := "orgId_example" // string | 
    teamId := "teamId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.ListAPITokens(context.Background(), orgId).TeamId(teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.ListAPITokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAPITokens`: ListAPITokenResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.ListAPITokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAPITokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **teamId** | **string** |  | 

### Return type

[**ListAPITokenResponse**](ListAPITokenResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInvites

> ListInvitesResponse ListInvites(ctx).Body(body).Execute()

List invites



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
    resp, r, err := apiClient.OrgsApi.ListInvites(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.ListInvites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInvites`: ListInvitesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.ListInvites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInvitesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**ListInvitesResponse**](ListInvitesResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgMembers

> ListOrgMembersResponse ListOrgMembers(ctx, orgId).Execute()

List members of organisation



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
    orgId := "orgId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.ListOrgMembers(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.ListOrgMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgMembers`: ListOrgMembersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.ListOrgMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListOrgMembersResponse**](ListOrgMembersResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgTeams

> ListOrgTeamsResponse ListOrgTeams(ctx, orgId).Execute()

List teams in org



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
    orgId := "orgId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.ListOrgTeams(context.Background(), orgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.ListOrgTeams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgTeams`: ListOrgTeamsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.ListOrgTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListOrgTeamsResponse**](ListOrgTeamsResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrgs

> ListOrgsResponse ListOrgs(ctx).Execute()

List organisations



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
    resp, r, err := apiClient.OrgsApi.ListOrgs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.ListOrgs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgs`: ListOrgsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.ListOrgs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgsRequest struct via the builder pattern


### Return type

[**ListOrgsResponse**](ListOrgsResponse.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAPIToken

> map[string]interface{} RemoveAPIToken(ctx, orgId, teamId, tokenId).Execute()

Remove API token from team



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
    orgId := "orgId_example" // string | 
    teamId := "teamId_example" // string | 
    tokenId := "tokenId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.RemoveAPIToken(context.Background(), orgId, teamId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.RemoveAPIToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveAPIToken`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.RemoveAPIToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**teamId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAPITokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveOrgMember

> OrgMember RemoveOrgMember(ctx, orgId, memberId).Execute()

Remove member from organisation



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
    orgId := "orgId_example" // string | 
    memberId := "memberId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.RemoveOrgMember(context.Background(), orgId, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.RemoveOrgMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveOrgMember`: OrgMember
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.RemoveOrgMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveOrgMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OrgMember**](OrgMember.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTeamMember

> OrgTeam RemoveTeamMember(ctx, orgId, teamId, memberId).Execute()

Remove a member from a team



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
    orgId := "orgId_example" // string | 
    teamId := "teamId_example" // string | 
    memberId := "memberId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.RemoveTeamMember(context.Background(), orgId, teamId, memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.RemoveTeamMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveTeamMember`: OrgTeam
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.RemoveTeamMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**teamId** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**OrgTeam**](OrgTeam.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAPIToken

> APIToken UpdateAPIToken(ctx, orgId, teamId, tokenId).Body(body).Execute()

Update API token



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
    orgId := "orgId_example" // string | 
    teamId := "teamId_example" // string | 
    tokenId := "tokenId_example" // string | 
    body := *openapiclient.NewUpdateAPITokenRequest() // UpdateAPITokenRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.UpdateAPIToken(context.Background(), orgId, teamId, tokenId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.UpdateAPIToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAPIToken`: APIToken
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.UpdateAPIToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**teamId** | **string** |  | 
**tokenId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAPITokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**UpdateAPITokenRequest**](UpdateAPITokenRequest.md) |  | 

### Return type

[**APIToken**](APIToken.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrg

> Organisation UpdateOrg(ctx, orgId).Body(body).Execute()

Update organisation



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
    orgId := "orgId_example" // string | 
    body := *openapiclient.NewUpdateOrgRequest() // UpdateOrgRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.UpdateOrg(context.Background(), orgId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.UpdateOrg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrg`: Organisation
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.UpdateOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateOrgRequest**](UpdateOrgRequest.md) |  | 

### Return type

[**Organisation**](Organisation.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgMemberRole

> OrgMember UpdateOrgMemberRole(ctx, orgId, memberId).Body(body).Execute()

Update org member role



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
    orgId := "orgId_example" // string | 
    memberId := "memberId_example" // string | 
    body := *openapiclient.NewUpdateOrgMemberRoleRequest() // UpdateOrgMemberRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.UpdateOrgMemberRole(context.Background(), orgId, memberId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.UpdateOrgMemberRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgMemberRole`: OrgMember
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.UpdateOrgMemberRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgMemberRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateOrgMemberRoleRequest**](UpdateOrgMemberRoleRequest.md) |  | 

### Return type

[**OrgMember**](OrgMember.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgTeam

> OrgTeam UpdateOrgTeam(ctx, orgId, teamId).Body(body).Execute()

Update a team



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
    orgId := "orgId_example" // string | Organisation ID
    teamId := "teamId_example" // string | Team ID
    body := *openapiclient.NewUpdateOrgTeamRequest() // UpdateOrgTeamRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.UpdateOrgTeam(context.Background(), orgId, teamId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.UpdateOrgTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgTeam`: OrgTeam
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.UpdateOrgTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** | Organisation ID | 
**teamId** | **string** | Team ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateOrgTeamRequest**](UpdateOrgTeamRequest.md) |  | 

### Return type

[**OrgTeam**](OrgTeam.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeamMember

> OrgTeam UpdateTeamMember(ctx, orgId, teamId).Body(body).Execute()

Change membership role for a team member



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
    orgId := "orgId_example" // string | 
    teamId := "teamId_example" // string | 
    body := *openapiclient.NewTeamMemberRequest() // TeamMemberRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrgsApi.UpdateTeamMember(context.Background(), orgId, teamId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgsApi.UpdateTeamMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeamMember`: OrgTeam
    fmt.Fprintf(os.Stdout, "Response from `OrgsApi.UpdateTeamMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgId** | **string** |  | 
**teamId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**TeamMemberRequest**](TeamMemberRequest.md) |  | 

### Return type

[**OrgTeam**](OrgTeam.md)

### Authorization

[bearer](../README.md#bearer), [APIToken](../README.md#APIToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

