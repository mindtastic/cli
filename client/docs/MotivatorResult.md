# MotivatorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **time.Time** |  | 
**Values** | [**map[string]MotivatorResultEntry**](MotivatorResultEntry.md) |  | 
**Feedback** | Pointer to [**MotivatorFeedback**](MotivatorFeedback.md) |  | [optional] 

## Methods

### NewMotivatorResult

`func NewMotivatorResult(timestamp time.Time, values map[string]MotivatorResultEntry, ) *MotivatorResult`

NewMotivatorResult instantiates a new MotivatorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotivatorResultWithDefaults

`func NewMotivatorResultWithDefaults() *MotivatorResult`

NewMotivatorResultWithDefaults instantiates a new MotivatorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetValues

`func (o *MotivatorResult) GetValues() map[string]MotivatorResultEntry`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MotivatorResult) GetValuesOk() (*map[string]MotivatorResultEntry, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MotivatorResult) SetValues(v map[string]MotivatorResultEntry)`

SetValues sets Values field to given value.


### GetFeedback

`func (o *MotivatorResult) GetFeedback() MotivatorFeedback`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *MotivatorResult) GetFeedbackOk() (*MotivatorFeedback, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *MotivatorResult) SetFeedback(v MotivatorFeedback)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *MotivatorResult) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


