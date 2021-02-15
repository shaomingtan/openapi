/*
 * Nchf_spendinglimitcontrol
 *
 * Spending Limit Control
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import "time"

// PenPolCounterStatus 3GPP TS 29 594 Section 5.6.2.5
type PenPolCounterStatus struct {
	PolicyCounterStatus string     `json:"policyCounterStatus" yaml:"policyCounterStatus" bson:"policyCounterStatus" mapstructure:"PolicyCounterStatus"`
	ActivationTime      *time.Time `json:"activationTime" yaml:"activationTime" bson:"activationTime" mapstructure:"ActivationTime"`
}
