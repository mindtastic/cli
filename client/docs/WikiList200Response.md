# WikiList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryCount** | **int32** |  | 
**Entries** | [**[]WikiEntryHeader**](WikiEntryHeader.md) |  | 

## Methods

### NewWikiList200Response

`func NewWikiList200Response(entryCount int32, entries []WikiEntryHeader, ) *WikiList200Response`

NewWikiList200Response instantiates a new WikiList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWikiList200ResponseWithDefaults

`func NewWikiList200ResponseWithDefaults() *WikiList200Response`

NewWikiList200ResponseWithDefaults instantiates a new WikiList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryCount

`func (o *WikiList200Response) GetEntryCount() int32`

GetEntryCount returns the EntryCount field if non-nil, zero value otherwise.

### GetEntryCountOk

`func (o *WikiList200Response) GetEntryCountOk() (*int32, bool)`

GetEntryCountOk returns a tuple with the EntryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryCount

`func (o *WikiList200Response) SetEntryCount(v int32)`

SetEntryCount sets EntryCount field to given value.


### GetEntries

`func (o *WikiList200Response) GetEntries() []WikiEntryHeader`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *WikiList200Response) GetEntriesOk() (*[]WikiEntryHeader, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *WikiList200Response) SetEntries(v []WikiEntryHeader)`

SetEntries sets Entries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


