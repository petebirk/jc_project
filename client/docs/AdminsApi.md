# {{classname}}

All URIs are relative to *https://virtserver.swaggerhub.com/pbirk007/JCProject/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestStatistics**](AdminsApi.md#RequestStatistics) | **Get** /stats | Return statisics on hash requests.
[**Shutdown**](AdminsApi.md#Shutdown) | **Post** /shutdown | Shuts down the system after all requests processed.

# **RequestStatistics**
> Statistics RequestStatistics(ctx, )
Return statisics on hash requests.

This will return a JSON response that contains the hash request statistics including total number of hash requests and average time in microseconds. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Statistics**](Statistics.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Shutdown**
> Shutdown(ctx, )
Shuts down the system after all requests processed.

This will reject any new requests from coming in immediately and will wait for pending requests to finish before exiting. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

