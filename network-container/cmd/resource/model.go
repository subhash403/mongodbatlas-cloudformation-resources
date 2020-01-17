package resource

import "github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"

/*
This file is autogenerated, do not edit;
changes will be undone by the next 'generate' command.

Updates to this type are made my editing the schema file
and executing the 'generate' command
*/

// Model is autogenerated from the json schema
type Model struct {
	ProjectId           *encoding.String `json:"ProjectId,omitempty"`
	ContainerId         *encoding.String `json:"ContainerId,omitempty"`
	RegionName          *encoding.String `json:"RegionName,omitempty"`
	Region              *encoding.String `json:"Region,omitempty"`
	Provisioned         *encoding.Bool   `json:"Provisioned,omitempty"`
	ProviderName        *encoding.String `json:"ProviderName,omitempty"`
	VpcId               *encoding.String `json:"VpcId,omitempty"`
	AtlasCidrBlock      *encoding.String `json:"AtlasCidrBlock,omitempty"`
	AzureSubscriptionId *encoding.String `json:"AzureSubscriptionId,omitempty"`
	VnetName            *encoding.String `json:"VnetName,omitempty"`
	Id                  *encoding.String `json:"Id,omitempty"`
	GcpProjectId        *encoding.String `json:"GcpProjectId,omitempty"`
	NetworkName         *encoding.String `json:"NetworkName,omitempty"`
	PageNum             *encoding.Int    `json:"PageNum,omitempty"`
	ItemsPerPage        *encoding.Int    `json:"ItemsPerPage,omitempty"`
	ApiKeys             ApiKeyDefinition `json:"ApiKeys,omitempty"`
}

// ApiKeyDefinition is autogenerated from the json schema
type ApiKeyDefinition struct {
	PublicKey  *encoding.String
	PrivateKey *encoding.String
}
