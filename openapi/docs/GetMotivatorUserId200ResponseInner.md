# GetMotivatorUserId200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Headline** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**TextContents** | Pointer to **[]string** |  | [optional] 
**MediaContents** | Pointer to [**[]MotivatorMediaContentsInner**](MotivatorMediaContentsInner.md) |  | [optional] 
**UserInputForm** | Pointer to **string** | May describe a form for input by the user (format to be defined, e.g., JSON Schema). | [optional] 
**MotivatorId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Results** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetMotivatorUserId200ResponseInner

`func NewGetMotivatorUserId200ResponseInner() *GetMotivatorUserId200ResponseInner`

NewGetMotivatorUserId200ResponseInner instantiates a new GetMotivatorUserId200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMotivatorUserId200ResponseInnerWithDefaults

`func NewGetMotivatorUserId200ResponseInnerWithDefaults() *GetMotivatorUserId200ResponseInner`

NewGetMotivatorUserId200ResponseInnerWithDefaults instantiates a new GetMotivatorUserId200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMotivatorUserId200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMotivatorUserId200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMotivatorUserId200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetMotivatorUserId200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetMotivatorUserId200ResponseInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMotivatorUserId200ResponseInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMotivatorUserId200ResponseInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMotivatorUserId200ResponseInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHeadline

`func (o *GetMotivatorUserId200ResponseInner) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *GetMotivatorUserId200ResponseInner) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *GetMotivatorUserId200ResponseInner) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *GetMotivatorUserId200ResponseInner) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetDescription

`func (o *GetMotivatorUserId200ResponseInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetMotivatorUserId200ResponseInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetMotivatorUserId200ResponseInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetMotivatorUserId200ResponseInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTextContents

`func (o *GetMotivatorUserId200ResponseInner) GetTextContents() []string`

GetTextContents returns the TextContents field if non-nil, zero value otherwise.

### GetTextContentsOk

`func (o *GetMotivatorUserId200ResponseInner) GetTextContentsOk() (*[]string, bool)`

GetTextContentsOk returns a tuple with the TextContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextContents

`func (o *GetMotivatorUserId200ResponseInner) SetTextContents(v []string)`

SetTextContents sets TextContents field to given value.

### HasTextContents

`func (o *GetMotivatorUserId200ResponseInner) HasTextContents() bool`

HasTextContents returns a boolean if a field has been set.

### GetMediaContents

`func (o *GetMotivatorUserId200ResponseInner) GetMediaContents() []MotivatorMediaContentsInner`

GetMediaContents returns the MediaContents field if non-nil, zero value otherwise.

### GetMediaContentsOk

`func (o *GetMotivatorUserId200ResponseInner) GetMediaContentsOk() (*[]MotivatorMediaContentsInner, bool)`

GetMediaContentsOk returns a tuple with the MediaContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaContents

`func (o *GetMotivatorUserId200ResponseInner) SetMediaContents(v []MotivatorMediaContentsInner)`

SetMediaContents sets MediaContents field to given value.

### HasMediaContents

`func (o *GetMotivatorUserId200ResponseInner) HasMediaContents() bool`

HasMediaContents returns a boolean if a field has been set.

### GetUserInputForm

`func (o *GetMotivatorUserId200ResponseInner) GetUserInputForm() string`

GetUserInputForm returns the UserInputForm field if non-nil, zero value otherwise.

### GetUserInputFormOk

`func (o *GetMotivatorUserId200ResponseInner) GetUserInputFormOk() (*string, bool)`

GetUserInputFormOk returns a tuple with the UserInputForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInputForm

`func (o *GetMotivatorUserId200ResponseInner) SetUserInputForm(v string)`

SetUserInputForm sets UserInputForm field to given value.

### HasUserInputForm

`func (o *GetMotivatorUserId200ResponseInner) HasUserInputForm() bool`

HasUserInputForm returns a boolean if a field has been set.

### GetMotivatorId

`func (o *GetMotivatorUserId200ResponseInner) GetMotivatorId() string`

GetMotivatorId returns the MotivatorId field if non-nil, zero value otherwise.

### GetMotivatorIdOk

`func (o *GetMotivatorUserId200ResponseInner) GetMotivatorIdOk() (*string, bool)`

GetMotivatorIdOk returns a tuple with the MotivatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMotivatorId

`func (o *GetMotivatorUserId200ResponseInner) SetMotivatorId(v string)`

SetMotivatorId sets MotivatorId field to given value.

### HasMotivatorId

`func (o *GetMotivatorUserId200ResponseInner) HasMotivatorId() bool`

HasMotivatorId returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetMotivatorUserId200ResponseInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetMotivatorUserId200ResponseInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetMotivatorUserId200ResponseInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetMotivatorUserId200ResponseInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetResults

`func (o *GetMotivatorUserId200ResponseInner) GetResults() []string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GetMotivatorUserId200ResponseInner) GetResultsOk() (*[]string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GetMotivatorUserId200ResponseInner) SetResults(v []string)`

SetResults sets Results field to given value.

### HasResults

`func (o *GetMotivatorUserId200ResponseInner) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


