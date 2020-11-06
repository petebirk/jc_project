# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/pbirk007/JCProject/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestHash**](DevelopersApi.md#RequestHash) | **Post** /hash | creates a new hash request
[**RetrieveHash**](DevelopersApi.md#RetrieveHash) | **Get** /hash/{requestId} | retrieves the password hash

# **RequestHash**
> int32 RequestHash(ctx, optional)
creates a new hash request

Pass in password field in request body.  Queue the password hash request, but don't hash it for 5 seconds.  Return a unique request ID that can be used with /hash/${requestid} for obtaining the base64 encoded hash value. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DevelopersApiRequestHashOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DevelopersApiRequestHashOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **password** | **optional.**|  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveHash**
> string RetrieveHash(ctx, requestId)
retrieves the password hash

Submit the requestId in the API path and the password hash associated with this requestId will be returned.   If the password hash has not been created yet, it will return a 202 Accepted response until the password hash is ready to return. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **requestId** | **string**| The value returned in the original /hash request to retrieve the password hash.  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

