/*
 * Nchf_ConvergedCharging
 *
 * Charging Function Service API
 *
 * API version 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

import "time"

// PDUSessionInformation TS 132 291 Section 6.1.6.2.2.8
type PduSessionInformation struct {
	NetworkSlicingInfo                   NetworkSlicingInfo                     `json:"networkSlicingInfo" yaml:"networkSlicingInfo" bson:"networkSlicingInfo" mapstructure:"NetworkSlicingInfo"`
	PduSessionID                         uint8                                  `json:"pduSessionID" yaml:"pduSessionID" bson:"pduSessionID" mapstructure:"PduSessionID"`
	PduType                              PduSessionType                         `json:"pduType" yaml:"pduType" bson:"pduType" mapstructure:"PduType"`
	SscMode                              SscMode                                `json:"sscMode" yaml:"sscMode" bson:"sscMode" mapstructure:"SscMode"`
	HPlmnId                              PlmnId                                 `json:"hPlmnId" yaml:"hPlmnId" bson:"hPlmnId" mapstructure:"HPlmnId"`
	ServingNetworkFunctionID             ServingNFID                            `json:"servingNetworkFunctionID" yaml:"servingNetworkFunctionID" bson:"servingNetworkFunctionID" mapstructure:"ServingNetworkFunctionID"`
	ServingCNPlmnId                      PlmnId                                 `json:"servingCNPlmnId" yaml:"servingCNPlmnId" bson:"servingCNPlmnId" mapstructure:"ServingCNPlmnId"`
	RatType                              RatType                                `json:"ratType" yaml:"ratType" bson:"ratType" mapstructure:"RatType"`
	DnnId                                string                                 `json:"dnnId" yaml:"dnnId" bson:"dnnId" mapstructure:"DnnId"`
	DnnSelectionMode                     []DnnSelectionMode                     `json:"dnnSelectionMode" yaml:"dnnSelectionMode" bson:"dnnSelectionMode" mapstructure:"DnnSelectionMode"`
	ChargingCharacteristics              string                                 `json:"chargingCharacteristics" yaml:"chargingCharacteristics" bson:"chargingCharacteristics" mapstructure:"ChargingCharacteristics"`
	ChargingCharacteristicsSelectionMode []ChargingCharacteristicsSelectionMode `json:"chargingCharacteristicsSelectionMode" yaml:"chargingCharacteristicsSelectionMode" bson:"chargingCharacteristicsSelectionMode" mapstructure:"ChargingCharacteristicsSelectionMode"`
	StartTime                            *time.Time                             `json:"startTime" yaml:"startTime" bson:"startTime" mapstructure:"StartTime"`
	StopTime                             *time.Time                             `json:"stopTime" yaml:"stopTime" bson:"stopTime" mapstructure:"StopTime"`
	Var3gppPSDataOffStatus               []Var3GPPPSDataOffStatus               `json:"3gppPSDataOffStatus" yaml:"3gppPSDataOffStatus" bson:"3gppPSDataOffStatus" mapstructure:"3gppPSDataOffStatus"`
	SessionStopIndicator                 bool                                   `json:"sessionStopIndicator" yaml:"sessionStopIndicator" bson:"sessionStopIndicator" mapstructure:"SessionStopIndicator"`
	PduAddress                           PduAddress                             `json:"pduAddress" yaml:"pduAddress" bson:"pduAddress" mapstructure:"PduAddress"`
	Diagnostics                          int                                    `json:"diagnostics" yaml:"diagnostics" bson:"diagnostics" mapstructure:"Diagnostics"`
	AuthorizedQoSInformation             AuthorizedDefaultQos                   `json:"authorizedQoSInformation" yaml:"authorizedQoSInformation" bson:"authorizedQoSInformation" mapstructure:"AuthorizedQoSInformation"`
	SubscribedQoSInformation             SubscribedDefaultQos                   `json:"subscribedQoSInformation" yaml:"subscribedQoSInformation" bson:"subscribedQoSInformation" mapstructure:"SubscribedQoSInformation"`
	AuthorizedSessionAMBR                Ambr                                   `json:"authorizedSessionAMBR" yaml:"authorizedSessionAMBR" bson:"authorizedSessionAMBR" mapstructure:"AuthorizedSessionAMBR"`
	SubscribedSessionAMBR                Ambr                                   `json:"subscribedSessionAMBR" yaml:"subscribedSessionAMBR" bson:"subscribedSessionAMBR" mapstructure:"SubscribedSessionAMBR"`
}
