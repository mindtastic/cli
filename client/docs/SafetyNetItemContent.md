# SafetyNetItemContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Type** | **string** | Which kind of item is this (i.a., influences symbol in UI)? | 
**Strategies** | **[]string** | How can this item help the user to feel better? An array of strings describing the strategies. | 
**Feedback** | Pointer to [**[]SafetyNetItemFeedback**](SafetyNetItemFeedback.md) |  | [optional] 

## Methods

### NewSafetyNetItemContent

`func NewSafetyNetItemContent(name string, type_ string, strategies []string, ) *SafetyNetItemContent`

NewSafetyNetItemContent instantiates a new SafetyNetItemContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSafetyNetItemContentWithDefaults

`func NewSafetyNetItemContentWithDefaults() *SafetyNetItemContent`

NewSafetyNetItemContentWithDefaults instantiates a new SafetyNetItemContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SafetyNetItemContent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SafetyNetItemContent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SafetyNetItemContent) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SafetyNetItemContent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SafetyNetItemContent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SafetyNetItemContent) SetType(v string)`

SetType sets Type field to given value.


### GetStrategies

`func (o *SafetyNetItemContent) GetStrategies() []string`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *SafetyNetItemContent) GetStrategiesOk() (*[]string, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *SafetyNetItemContent) SetStrategies(v []string)`

SetStrategies sets Strategies field to given value.


### GetFeedback

`func (o *SafetyNetItemContent) GetFeedback() []SafetyNetItemFeedback`

GetFeedback returns the Feedback field if non-nil, zero value otherwise.

### GetFeedbackOk

`func (o *SafetyNetItemContent) GetFeedbackOk() (*[]SafetyNetItemFeedback, bool)`

GetFeedbackOk returns a tuple with the Feedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedback

`func (o *SafetyNetItemContent) SetFeedback(v []SafetyNetItemFeedback)`

SetFeedback sets Feedback field to given value.

### HasFeedback

`func (o *SafetyNetItemContent) HasFeedback() bool`

HasFeedback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


