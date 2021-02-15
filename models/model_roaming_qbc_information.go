/*
 * Nchf_ConvergedCharging
 *
 * Charging Function Service API
 *
 * API version 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// RoamingQBCInformation TS 132 291 Section 6.1.6.2.2.13
type RoamingQBCInformation struct {
	MultipleQFIcontainer   []MultipleQFIcontainer `json:"multipleQFIcontainer" yaml:"multipleQFIcontainer" bson:"multipleQFIcontainer" mapstructure:"MultipleQFIcontainer"`
	UPFID                  string                 `json:"uPFID" yaml:"uPFID" bson:"uPFID" mapstructure:"UPFID"`
	RoamingChargingProfile RoamingChargingProfile `json:"roamingChargingProfile" yaml:"roamingChargingProfile" bson:"roamingChargingProfile" mapstructure:"RoamingChargingProfile"`
}