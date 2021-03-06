/*
 * Nchf_ConvergedCharging
 *
 * Charging Function Service API
 *
 * API version 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// PDUAddress TS 132 291 Section 6.1.6.2.2.11
type PduAddress struct {
	PduIPv4Address           string `json:"pduIPv4Address" yaml:"pduIPv4Address" bson:"pduIPv4Address" mapstructure:"PduIPv4Address"`
	PduIPv6AddresswithPrefix string `json:"pduIPv6AddresswithPrefix" yaml:"pduIPv6AddresswithPrefix" bson:"pduIPv6AddresswithPrefix" mapstructure:"PduIPv6AddresswithPrefix"`
	PduAddressprefixlength   int64  `json:"pduAddressprefixlength" yaml:"pduAddressprefixlength" bson:"pduAddressprefixlength" mapstructure:"PduAddressprefixlength"`
	IPv4dynamicAddressFlag   bool   `json:"iPv4dynamicAddressFlag" yaml:"iPv4dynamicAddressFlag" bson:"iPv4dynamicAddressFlag" mapstructure:"IPv4dynamicAddressFlag"`
	IPv6dynamicPrefixFlag    bool   `json:"iPv6dynamicPrefixFlag" yaml:"iPv6dynamicPrefixFlag" bson:"iPv6dynamicPrefixFlag" mapstructure:"IPv6dynamicPrefixFlag"`
}
