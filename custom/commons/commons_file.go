package commons

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

//nolint:dupl Similar Methods between entities
var FileSchema = map[string]*schema.Schema{
	"name": {
		Type:        schema.TypeString,
		Description: "The file name",
		ForceNew:    true,
		Required:    true,
	},
	"content": {
		Type:        schema.TypeString,
		Description: "The ciontent file",
		ForceNew:    true,
		Required:    true,
	},
}