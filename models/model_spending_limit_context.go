/*
 * Nchf_spendinglimitcontrol
 *
 * Spending Limit Control
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// SpendingLimitContext 3GPP TS 29 594 Section 5.6.2.2
type SpendingLimitContext struct {
	Supi              string   `json:"supi" yaml:"supi" bson:"supi" mapstructure:"Supi"`
	Gpsi              string   `json:"gpsi,omitempty" yaml:"gpsi" bson:"gpsi" mapstructure:"Gpsi"`
	PolicyCounterIds  []string `json:"policyCounterIds" yaml:"policyCounterIds" bson:"policyCounterIds" mapstructure:"PolicyCounterIds"`
	NotifUri          string   `json:"notifUri" yaml:"notifUri" bson:"notifUri" mapstructure:"NotifUri"`
	SupportedFeatures string   `json:"supportedFeatures,omitempty" yaml:"supportedFeatures" bson:"supportedFeatures" mapstructure:"SupportedFeatures"`
}
