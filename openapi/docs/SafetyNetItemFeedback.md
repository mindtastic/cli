# SafetyNetItemFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItHelped** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSafetyNetItemFeedback

`func NewSafetyNetItemFeedback() *SafetyNetItemFeedback`

NewSafetyNetItemFeedback instantiates a new SafetyNetItemFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafetyNetItemFeedbackWithDefaults

`func NewSafetyNetItemFeedbackWithDefaults() *SafetyNetItemFeedback`

NewSafetyNetItemFeedbackWithDefaults instantiates a new SafetyNetItemFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItHelped

`func (o *SafetyNetItemFeedback) GetItHelped() bool`

GetItHelped returns the ItHelped field if non-nil, zero value otherwise.

### GetItHelpedOk

`func (o *SafetyNetItemFeedback) GetItHelpedOk() (*bool, bool)`

GetItHelpedOk returns a tuple with the ItHelped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItHelped

`func (o *SafetyNetItemFeedback) SetItHelped(v bool)`

SetItHelped sets ItHelped field to given value.

### HasItHelped

`func (o *SafetyNetItemFeedback) HasItHelped() bool`

HasItHelped returns a boolean if a field has been set.

### GetComment

`func (o *SafetyNetItemFeedback) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SafetyNetItemFeedback) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SafetyNetItemFeedback) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SafetyNetItemFeedback) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetTimestamp

`func (o *SafetyNetItemFeedback) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SafetyNetItemFeedback) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SafetyNetItemFeedback) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SafetyNetItemFeedback) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


