# InsertionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int32** | The HTTP status code of the response | [optional] 
**Success** | Pointer to **bool** | True if all provided wiki entries have been inserted successfully into the database | [optional] 
**InsertedCount** | Pointer to **int32** | The number of inserted wiki entries on the operation | [optional] 

## Methods

### NewInsertionResponse

`func NewInsertionResponse() *InsertionResponse`

NewInsertionResponse instantiates a new InsertionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsertionResponseWithDefaults

`func NewInsertionResponseWithDefaults() *InsertionResponse`

NewInsertionResponseWithDefaults instantiates a new InsertionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *InsertionResponse) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *InsertionResponse) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *InsertionResponse) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *InsertionResponse) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetSuccess

`func (o *InsertionResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *InsertionResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *InsertionResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *InsertionResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetInsertedCount

`func (o *InsertionResponse) GetInsertedCount() int32`

GetInsertedCount returns the InsertedCount field if non-nil, zero value otherwise.

### GetInsertedCountOk

`func (o *InsertionResponse) GetInsertedCountOk() (*int32, bool)`

GetInsertedCountOk returns a tuple with the InsertedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertedCount

`func (o *InsertionResponse) SetInsertedCount(v int32)`

SetInsertedCount sets InsertedCount field to given value.

### HasInsertedCount

`func (o *InsertionResponse) HasInsertedCount() bool`

HasInsertedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


