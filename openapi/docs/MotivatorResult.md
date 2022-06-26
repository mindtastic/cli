# MotivatorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MotivatorId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Results** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMotivatorResult

`func NewMotivatorResult() *MotivatorResult`

NewMotivatorResult instantiates a new MotivatorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotivatorResultWithDefaults

`func NewMotivatorResultWithDefaults() *MotivatorResult`

NewMotivatorResultWithDefaults instantiates a new MotivatorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMotivatorId

`func (o *MotivatorResult) GetMotivatorId() string`

GetMotivatorId returns the MotivatorId field if non-nil, zero value otherwise.

### GetMotivatorIdOk

`func (o *MotivatorResult) GetMotivatorIdOk() (*string, bool)`

GetMotivatorIdOk returns a tuple with the MotivatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotivatorId

`func (o *MotivatorResult) SetMotivatorId(v string)`

SetMotivatorId sets MotivatorId field to given value.

### HasMotivatorId

`func (o *MotivatorResult) HasMotivatorId() bool`

HasMotivatorId returns a boolean if a field has been set.

### GetTimestamp

`func (o *MotivatorResult) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MotivatorResult) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MotivatorResult) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MotivatorResult) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetResults

`func (o *MotivatorResult) GetResults() []string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MotivatorResult) GetResultsOk() (*[]string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MotivatorResult) SetResults(v []string)`

SetResults sets Results field to given value.

### HasResults

`func (o *MotivatorResult) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


