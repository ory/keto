# \SyntaxApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckOplSyntax**](SyntaxApi.md#CheckOplSyntax) | **Post** /opl/syntax/check | Check the syntax of an OPL file



## CheckOplSyntax

> PostCheckOplSyntaxResponse CheckOplSyntax(ctx).Body(body).Execute()

Check the syntax of an OPL file



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := "body_example" // string | the OPL content to check

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntaxApi.CheckOplSyntax(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntaxApi.CheckOplSyntax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckOplSyntax`: PostCheckOplSyntaxResponse
    fmt.Fprintf(os.Stdout, "Response from `SyntaxApi.CheckOplSyntax`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckOplSyntaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | the OPL content to check | 

### Return type

[**PostCheckOplSyntaxResponse**](PostCheckOplSyntaxResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

