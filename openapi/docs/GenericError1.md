# GenericError1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** | The HTTP status code of the reponse | [optional] 
**Error** | Pointer to **string** | Error message | [optional] 

## Methods

### NewGenericError1

`func NewGenericError1() *GenericError1`

NewGenericError1 instantiates a new GenericError1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericError1WithDefaults

`func NewGenericError1WithDefaults() *GenericError1`

NewGenericError1WithDefaults instantiates a new GenericError1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *GenericError1) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *GenericError1) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *GenericError1) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *GenericError1) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *GenericError1) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GenericError1) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GenericError1) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GenericError1) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


