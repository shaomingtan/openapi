/*
 * Nchf_ConvergedCharging
 *
 * Charging Function Service API
 *
 * API version 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// MultipleUnitUsage 3GPP TS 32 291 Table 6.1.6.2.1.5
type MultipleUnitUsage struct {
	RatingGroup       uint32              `json:"ratingGroup" yaml:"ratingGroup" bson:"ratingGroup" mapstructure:"RatingGroup"`
	RequestedUnit     RequestedUnit       `json:"requestedUnit,omitempty" yaml:"requestedUnit" bson:"requestedUnit" mapstructure:"RequestedUnit"`
	UsedUnitContainer []UsedUnitContainer `json:"usedUnitContainer" yaml:"usedUnitContainer" bson:"usedUnitContainer" mapstructure:"UsedUnitContainer"`
}
