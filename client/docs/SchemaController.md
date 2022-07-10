# SchemaController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Division** | Pointer to **string** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Representative** | Pointer to [**SchemaControllerRepresentative**](SchemaControllerRepresentative.md) |  | [optional] 

## Methods

### NewSchemaController

`func NewSchemaController() *SchemaController`

NewSchemaController instantiates a new SchemaController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaControllerWithDefaults

`func NewSchemaControllerWithDefaults() *SchemaController`

NewSchemaControllerWithDefaults instantiates a new SchemaController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SchemaController) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaController) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaController) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchemaController) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDivision

`func (o *SchemaController) GetDivision() string`

GetDivision returns the Division field if non-nil, zero value otherwise.

### GetDivisionOk

`func (o *SchemaController) GetDivisionOk() (*string, bool)`

GetDivisionOk returns a tuple with the Division field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivision

`func (o *SchemaController) SetDivision(v string)`

SetDivision sets Division field to given value.

### HasDivision

`func (o *SchemaController) HasDivision() bool`

HasDivision returns a boolean if a field has been set.

### GetAddress

`func (o *SchemaController) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SchemaController) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SchemaController) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SchemaController) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCountry

`func (o *SchemaController) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SchemaController) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SchemaController) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SchemaController) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRepresentative

`func (o *SchemaController) GetRepresentative() SchemaControllerRepresentative`

GetRepresentative returns the Representative field if non-nil, zero value otherwise.

### GetRepresentativeOk

`func (o *SchemaController) GetRepresentativeOk() (*SchemaControllerRepresentative, bool)`

GetRepresentativeOk returns a tuple with the Representative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepresentative

`func (o *SchemaController) SetRepresentative(v SchemaControllerRepresentative)`

SetRepresentative sets Representative field to given value.

### HasRepresentative

`func (o *SchemaController) HasRepresentative() bool`

HasRepresentative returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


