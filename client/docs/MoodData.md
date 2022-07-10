# MoodData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoodDay** | **time.Time** | The date of the day the mood entry should created for. It expects an [&#x60;ISO 8601&#x60; / &#x60;RFC 3399&#x60; timestamp](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6). Sending only the date part without any &#x60;time&#x60; component is also allowed, but not recommended, as this could change in future versions of the API. | 
**MoodType** | **string** |  | 
**MoodDescr** | Pointer to **string** |  | [optional] 

## Methods

### NewMoodData

`func NewMoodData(moodDay time.Time, moodType string, ) *MoodData`

NewMoodData instantiates a new MoodData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoodDataWithDefaults

`func NewMoodDataWithDefaults() *MoodData`

NewMoodDataWithDefaults instantiates a new MoodData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoodDay

`func (o *MoodData) GetMoodDay() time.Time`

GetMoodDay returns the MoodDay field if non-nil, zero value otherwise.

### GetMoodDayOk

`func (o *MoodData) GetMoodDayOk() (*time.Time, bool)`

GetMoodDayOk returns a tuple with the MoodDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoodDay

`func (o *MoodData) SetMoodDay(v time.Time)`

SetMoodDay sets MoodDay field to given value.


### GetMoodType

`func (o *MoodData) GetMoodType() string`

GetMoodType returns the MoodType field if non-nil, zero value otherwise.

### GetMoodTypeOk

`func (o *MoodData) GetMoodTypeOk() (*string, bool)`

GetMoodTypeOk returns a tuple with the MoodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoodType

`func (o *MoodData) SetMoodType(v string)`

SetMoodType sets MoodType field to given value.


### GetMoodDescr

`func (o *MoodData) GetMoodDescr() string`

GetMoodDescr returns the MoodDescr field if non-nil, zero value otherwise.

### GetMoodDescrOk

`func (o *MoodData) GetMoodDescrOk() (*string, bool)`

GetMoodDescrOk returns a tuple with the MoodDescr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoodDescr

`func (o *MoodData) SetMoodDescr(v string)`

SetMoodDescr sets MoodDescr field to given value.

### HasMoodDescr

`func (o *MoodData) HasMoodDescr() bool`

HasMoodDescr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


