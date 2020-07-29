/*
 * ProductOffering aggregation composite interface
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Representing a Context
type Context struct {

	AdditionalAttributes []Characteristic `json:"additionalAttributes,omitempty"`

	Channel *ChannelRef `json:"channel,omitempty"`

	InstalledBase *InstalledBase `json:"installedBase,omitempty"`

	Place *PlaceRef `json:"place,omitempty"`

	Serviceability *Serviceability `json:"serviceability,omitempty"`

	ShoppingCart *ShoppingCart `json:"shoppingCart,omitempty"`
}