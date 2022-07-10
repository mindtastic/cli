# MotivatorInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Fields** | Pointer to [**map[string]MotivatorInputElement**](MotivatorInputElement.md) |  | [optional] 

## Methods

### NewMotivatorInput

`func NewMotivatorInput() *MotivatorInput`

NewMotivatorInput instantiates a new MotivatorInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMotivatorInputWithDefaults

`func NewMotivatorInputWithDefaults() *MotivatorInput`

NewMotivatorInputWithDefaults instantiates a new MotivatorInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MotivatorInput) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MotivatorInput) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MotivatorInput) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MotivatorInput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFields

`func (o *MotivatorInput) GetFields() map[string]MotivatorInputElement`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *MotivatorInput) GetFieldsOk() (*map[string]MotivatorInputElement, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *MotivatorInput) SetFields(v map[string]MotivatorInputElement)`

SetFields sets Fields field to given value.

### HasFields

`func (o *MotivatorInput) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


