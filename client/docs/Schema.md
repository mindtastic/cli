# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**SchemaMeta**](SchemaMeta.md) |  | [optional] 
**Controller** | Pointer to [**SchemaController**](SchemaController.md) |  | [optional] 
**DataProtectionOfficer** | Pointer to [**SchemaDataProtectionOfficer**](SchemaDataProtectionOfficer.md) |  | [optional] 
**DataDisclosed** | Pointer to [**[]SchemaDataDisclosedInner**](SchemaDataDisclosedInner.md) |  | [optional] 
**ThirdCountryTransfers** | Pointer to [**[]SchemaThirdCountryTransfersInner**](SchemaThirdCountryTransfersInner.md) |  | [optional] 
**AccessAndDataPortability** | Pointer to [**SchemaAccessAndDataPortability**](SchemaAccessAndDataPortability.md) |  | [optional] 
**Sources** | Pointer to [**[]SchemaSourcesInner**](SchemaSourcesInner.md) |  | [optional] 
**RightToInformation** | Pointer to [**SchemaRightToInformation**](SchemaRightToInformation.md) |  | [optional] 
**RightToRectificationOrDeletion** | Pointer to [**SchemaRightToInformation**](SchemaRightToInformation.md) |  | [optional] 
**RightToDataPortability** | Pointer to [**SchemaRightToInformation**](SchemaRightToInformation.md) |  | [optional] 
**RightToWithdrawConsent** | Pointer to [**SchemaRightToInformation**](SchemaRightToInformation.md) |  | [optional] 
**RightToComplain** | Pointer to [**SchemaRightToComplain**](SchemaRightToComplain.md) |  | [optional] 
**AutomatedDecisionMaking** | Pointer to [**SchemaAutomatedDecisionMaking**](SchemaAutomatedDecisionMaking.md) |  | [optional] 
**ChangesOfPurpose** | Pointer to [**[]SchemaChangesOfPurposeInner**](SchemaChangesOfPurposeInner.md) |  | [optional] 

## Methods

### NewSchema

`func NewSchema() *Schema`

NewSchema instantiates a new Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithDefaults

`func NewSchemaWithDefaults() *Schema`

NewSchemaWithDefaults instantiates a new Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *Schema) GetMeta() SchemaMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Schema) GetMetaOk() (*SchemaMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Schema) SetMeta(v SchemaMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Schema) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetController

`func (o *Schema) GetController() SchemaController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *Schema) GetControllerOk() (*SchemaController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *Schema) SetController(v SchemaController)`

SetController sets Controller field to given value.

### HasController

`func (o *Schema) HasController() bool`

HasController returns a boolean if a field has been set.

### GetDataProtectionOfficer

`func (o *Schema) GetDataProtectionOfficer() SchemaDataProtectionOfficer`

GetDataProtectionOfficer returns the DataProtectionOfficer field if non-nil, zero value otherwise.

### GetDataProtectionOfficerOk

`func (o *Schema) GetDataProtectionOfficerOk() (*SchemaDataProtectionOfficer, bool)`

GetDataProtectionOfficerOk returns a tuple with the DataProtectionOfficer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataProtectionOfficer

`func (o *Schema) SetDataProtectionOfficer(v SchemaDataProtectionOfficer)`

SetDataProtectionOfficer sets DataProtectionOfficer field to given value.

### HasDataProtectionOfficer

`func (o *Schema) HasDataProtectionOfficer() bool`

HasDataProtectionOfficer returns a boolean if a field has been set.

### GetDataDisclosed

`func (o *Schema) GetDataDisclosed() []SchemaDataDisclosedInner`

GetDataDisclosed returns the DataDisclosed field if non-nil, zero value otherwise.

### GetDataDisclosedOk

`func (o *Schema) GetDataDisclosedOk() (*[]SchemaDataDisclosedInner, bool)`

GetDataDisclosedOk returns a tuple with the DataDisclosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDisclosed

`func (o *Schema) SetDataDisclosed(v []SchemaDataDisclosedInner)`

SetDataDisclosed sets DataDisclosed field to given value.

### HasDataDisclosed

`func (o *Schema) HasDataDisclosed() bool`

HasDataDisclosed returns a boolean if a field has been set.

### GetThirdCountryTransfers

`func (o *Schema) GetThirdCountryTransfers() []SchemaThirdCountryTransfersInner`

GetThirdCountryTransfers returns the ThirdCountryTransfers field if non-nil, zero value otherwise.

### GetThirdCountryTransfersOk

`func (o *Schema) GetThirdCountryTransfersOk() (*[]SchemaThirdCountryTransfersInner, bool)`

GetThirdCountryTransfersOk returns a tuple with the ThirdCountryTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdCountryTransfers

`func (o *Schema) SetThirdCountryTransfers(v []SchemaThirdCountryTransfersInner)`

SetThirdCountryTransfers sets ThirdCountryTransfers field to given value.

### HasThirdCountryTransfers

`func (o *Schema) HasThirdCountryTransfers() bool`

HasThirdCountryTransfers returns a boolean if a field has been set.

### GetAccessAndDataPortability

`func (o *Schema) GetAccessAndDataPortability() SchemaAccessAndDataPortability`

GetAccessAndDataPortability returns the AccessAndDataPortability field if non-nil, zero value otherwise.

### GetAccessAndDataPortabilityOk

`func (o *Schema) GetAccessAndDataPortabilityOk() (*SchemaAccessAndDataPortability, bool)`

GetAccessAndDataPortabilityOk returns a tuple with the AccessAndDataPortability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessAndDataPortability

`func (o *Schema) SetAccessAndDataPortability(v SchemaAccessAndDataPortability)`

SetAccessAndDataPortability sets AccessAndDataPortability field to given value.

### HasAccessAndDataPortability

`func (o *Schema) HasAccessAndDataPortability() bool`

HasAccessAndDataPortability returns a boolean if a field has been set.

### GetSources

`func (o *Schema) GetSources() []SchemaSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Schema) GetSourcesOk() (*[]SchemaSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Schema) SetSources(v []SchemaSourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *Schema) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetRightToInformation

`func (o *Schema) GetRightToInformation() SchemaRightToInformation`

GetRightToInformation returns the RightToInformation field if non-nil, zero value otherwise.

### GetRightToInformationOk

`func (o *Schema) GetRightToInformationOk() (*SchemaRightToInformation, bool)`

GetRightToInformationOk returns a tuple with the RightToInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightToInformation

`func (o *Schema) SetRightToInformation(v SchemaRightToInformation)`

SetRightToInformation sets RightToInformation field to given value.

### HasRightToInformation

`func (o *Schema) HasRightToInformation() bool`

HasRightToInformation returns a boolean if a field has been set.

### GetRightToRectificationOrDeletion

`func (o *Schema) GetRightToRectificationOrDeletion() SchemaRightToInformation`

GetRightToRectificationOrDeletion returns the RightToRectificationOrDeletion field if non-nil, zero value otherwise.

### GetRightToRectificationOrDeletionOk

`func (o *Schema) GetRightToRectificationOrDeletionOk() (*SchemaRightToInformation, bool)`

GetRightToRectificationOrDeletionOk returns a tuple with the RightToRectificationOrDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightToRectificationOrDeletion

`func (o *Schema) SetRightToRectificationOrDeletion(v SchemaRightToInformation)`

SetRightToRectificationOrDeletion sets RightToRectificationOrDeletion field to given value.

### HasRightToRectificationOrDeletion

`func (o *Schema) HasRightToRectificationOrDeletion() bool`

HasRightToRectificationOrDeletion returns a boolean if a field has been set.

### GetRightToDataPortability

`func (o *Schema) GetRightToDataPortability() SchemaRightToInformation`

GetRightToDataPortability returns the RightToDataPortability field if non-nil, zero value otherwise.

### GetRightToDataPortabilityOk

`func (o *Schema) GetRightToDataPortabilityOk() (*SchemaRightToInformation, bool)`

GetRightToDataPortabilityOk returns a tuple with the RightToDataPortability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightToDataPortability

`func (o *Schema) SetRightToDataPortability(v SchemaRightToInformation)`

SetRightToDataPortability sets RightToDataPortability field to given value.

### HasRightToDataPortability

`func (o *Schema) HasRightToDataPortability() bool`

HasRightToDataPortability returns a boolean if a field has been set.

### GetRightToWithdrawConsent

`func (o *Schema) GetRightToWithdrawConsent() SchemaRightToInformation`

GetRightToWithdrawConsent returns the RightToWithdrawConsent field if non-nil, zero value otherwise.

### GetRightToWithdrawConsentOk

`func (o *Schema) GetRightToWithdrawConsentOk() (*SchemaRightToInformation, bool)`

GetRightToWithdrawConsentOk returns a tuple with the RightToWithdrawConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightToWithdrawConsent

`func (o *Schema) SetRightToWithdrawConsent(v SchemaRightToInformation)`

SetRightToWithdrawConsent sets RightToWithdrawConsent field to given value.

### HasRightToWithdrawConsent

`func (o *Schema) HasRightToWithdrawConsent() bool`

HasRightToWithdrawConsent returns a boolean if a field has been set.

### GetRightToComplain

`func (o *Schema) GetRightToComplain() SchemaRightToComplain`

GetRightToComplain returns the RightToComplain field if non-nil, zero value otherwise.

### GetRightToComplainOk

`func (o *Schema) GetRightToComplainOk() (*SchemaRightToComplain, bool)`

GetRightToComplainOk returns a tuple with the RightToComplain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightToComplain

`func (o *Schema) SetRightToComplain(v SchemaRightToComplain)`

SetRightToComplain sets RightToComplain field to given value.

### HasRightToComplain

`func (o *Schema) HasRightToComplain() bool`

HasRightToComplain returns a boolean if a field has been set.

### GetAutomatedDecisionMaking

`func (o *Schema) GetAutomatedDecisionMaking() SchemaAutomatedDecisionMaking`

GetAutomatedDecisionMaking returns the AutomatedDecisionMaking field if non-nil, zero value otherwise.

### GetAutomatedDecisionMakingOk

`func (o *Schema) GetAutomatedDecisionMakingOk() (*SchemaAutomatedDecisionMaking, bool)`

GetAutomatedDecisionMakingOk returns a tuple with the AutomatedDecisionMaking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedDecisionMaking

`func (o *Schema) SetAutomatedDecisionMaking(v SchemaAutomatedDecisionMaking)`

SetAutomatedDecisionMaking sets AutomatedDecisionMaking field to given value.

### HasAutomatedDecisionMaking

`func (o *Schema) HasAutomatedDecisionMaking() bool`

HasAutomatedDecisionMaking returns a boolean if a field has been set.

### GetChangesOfPurpose

`func (o *Schema) GetChangesOfPurpose() []SchemaChangesOfPurposeInner`

GetChangesOfPurpose returns the ChangesOfPurpose field if non-nil, zero value otherwise.

### GetChangesOfPurposeOk

`func (o *Schema) GetChangesOfPurposeOk() (*[]SchemaChangesOfPurposeInner, bool)`

GetChangesOfPurposeOk returns a tuple with the ChangesOfPurpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangesOfPurpose

`func (o *Schema) SetChangesOfPurpose(v []SchemaChangesOfPurposeInner)`

SetChangesOfPurpose sets ChangesOfPurpose field to given value.

### HasChangesOfPurpose

`func (o *Schema) HasChangesOfPurpose() bool`

HasChangesOfPurpose returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


