# SchemaDataDisclosedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Purposes** | Pointer to [**[]SchemaDataDisclosedInnerPurposesInner**](SchemaDataDisclosedInnerPurposesInner.md) |  | [optional] 
**LegalBases** | Pointer to [**[]SchemaDataDisclosedInnerLegalBasesInner**](SchemaDataDisclosedInnerLegalBasesInner.md) |  | [optional] 
**LegitimateInterests** | Pointer to [**[]SchemaDataDisclosedInnerLegitimateInterestsInner**](SchemaDataDisclosedInnerLegitimateInterestsInner.md) |  | [optional] 
**Recipients** | Pointer to [**[]SchemaDataDisclosedInnerRecipientsInner**](SchemaDataDisclosedInnerRecipientsInner.md) |  | [optional] 
**Storage** | Pointer to [**[]SchemaDataDisclosedInnerStorageInner**](SchemaDataDisclosedInnerStorageInner.md) |  | [optional] 
**NonDisclosure** | Pointer to [**SchemaDataDisclosedInnerNonDisclosure**](SchemaDataDisclosedInnerNonDisclosure.md) |  | [optional] 

## Methods

### NewSchemaDataDisclosedInner

`func NewSchemaDataDisclosedInner() *SchemaDataDisclosedInner`

NewSchemaDataDisclosedInner instantiates a new SchemaDataDisclosedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaDataDisclosedInnerWithDefaults

`func NewSchemaDataDisclosedInnerWithDefaults() *SchemaDataDisclosedInner`

NewSchemaDataDisclosedInnerWithDefaults instantiates a new SchemaDataDisclosedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchemaDataDisclosedInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemaDataDisclosedInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemaDataDisclosedInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemaDataDisclosedInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategory

`func (o *SchemaDataDisclosedInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SchemaDataDisclosedInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SchemaDataDisclosedInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SchemaDataDisclosedInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetPurposes

`func (o *SchemaDataDisclosedInner) GetPurposes() []SchemaDataDisclosedInnerPurposesInner`

GetPurposes returns the Purposes field if non-nil, zero value otherwise.

### GetPurposesOk

`func (o *SchemaDataDisclosedInner) GetPurposesOk() (*[]SchemaDataDisclosedInnerPurposesInner, bool)`

GetPurposesOk returns a tuple with the Purposes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposes

`func (o *SchemaDataDisclosedInner) SetPurposes(v []SchemaDataDisclosedInnerPurposesInner)`

SetPurposes sets Purposes field to given value.

### HasPurposes

`func (o *SchemaDataDisclosedInner) HasPurposes() bool`

HasPurposes returns a boolean if a field has been set.

### GetLegalBases

`func (o *SchemaDataDisclosedInner) GetLegalBases() []SchemaDataDisclosedInnerLegalBasesInner`

GetLegalBases returns the LegalBases field if non-nil, zero value otherwise.

### GetLegalBasesOk

`func (o *SchemaDataDisclosedInner) GetLegalBasesOk() (*[]SchemaDataDisclosedInnerLegalBasesInner, bool)`

GetLegalBasesOk returns a tuple with the LegalBases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegalBases

`func (o *SchemaDataDisclosedInner) SetLegalBases(v []SchemaDataDisclosedInnerLegalBasesInner)`

SetLegalBases sets LegalBases field to given value.

### HasLegalBases

`func (o *SchemaDataDisclosedInner) HasLegalBases() bool`

HasLegalBases returns a boolean if a field has been set.

### GetLegitimateInterests

`func (o *SchemaDataDisclosedInner) GetLegitimateInterests() []SchemaDataDisclosedInnerLegitimateInterestsInner`

GetLegitimateInterests returns the LegitimateInterests field if non-nil, zero value otherwise.

### GetLegitimateInterestsOk

`func (o *SchemaDataDisclosedInner) GetLegitimateInterestsOk() (*[]SchemaDataDisclosedInnerLegitimateInterestsInner, bool)`

GetLegitimateInterestsOk returns a tuple with the LegitimateInterests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegitimateInterests

`func (o *SchemaDataDisclosedInner) SetLegitimateInterests(v []SchemaDataDisclosedInnerLegitimateInterestsInner)`

SetLegitimateInterests sets LegitimateInterests field to given value.

### HasLegitimateInterests

`func (o *SchemaDataDisclosedInner) HasLegitimateInterests() bool`

HasLegitimateInterests returns a boolean if a field has been set.

### GetRecipients

`func (o *SchemaDataDisclosedInner) GetRecipients() []SchemaDataDisclosedInnerRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *SchemaDataDisclosedInner) GetRecipientsOk() (*[]SchemaDataDisclosedInnerRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *SchemaDataDisclosedInner) SetRecipients(v []SchemaDataDisclosedInnerRecipientsInner)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *SchemaDataDisclosedInner) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetStorage

`func (o *SchemaDataDisclosedInner) GetStorage() []SchemaDataDisclosedInnerStorageInner`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *SchemaDataDisclosedInner) GetStorageOk() (*[]SchemaDataDisclosedInnerStorageInner, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *SchemaDataDisclosedInner) SetStorage(v []SchemaDataDisclosedInnerStorageInner)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *SchemaDataDisclosedInner) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetNonDisclosure

`func (o *SchemaDataDisclosedInner) GetNonDisclosure() SchemaDataDisclosedInnerNonDisclosure`

GetNonDisclosure returns the NonDisclosure field if non-nil, zero value otherwise.

### GetNonDisclosureOk

`func (o *SchemaDataDisclosedInner) GetNonDisclosureOk() (*SchemaDataDisclosedInnerNonDisclosure, bool)`

GetNonDisclosureOk returns a tuple with the NonDisclosure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDisclosure

`func (o *SchemaDataDisclosedInner) SetNonDisclosure(v SchemaDataDisclosedInnerNonDisclosure)`

SetNonDisclosure sets NonDisclosure field to given value.

### HasNonDisclosure

`func (o *SchemaDataDisclosedInner) HasNonDisclosure() bool`

HasNonDisclosure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


