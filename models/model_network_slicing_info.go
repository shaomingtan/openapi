/*
 * Nchf_ConvergedCharging
 *
 * Charging Function Service API
 *
 * API version 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// NetworkSlicingInfo TS 132 291 Section 6.1.6.2.2.10
type NetworkSlicingInfo struct {
	SNSSAI Snssai `json:"sNSSAI" yaml:"sNSSAI" bson:"sNSSAI" mapstructure:"SNSSAI"`
}
