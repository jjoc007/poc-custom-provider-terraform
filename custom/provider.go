package custom

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/jjoc007/poc-custom-provider-terraform/custom/clients"
	"github.com/rs/zerolog/log"
)

// Provider creates the Ogmio provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name custom Server",
			},
			"location_files": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "location files",
			},
		},
		ConfigureFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"custom_jjoc_file": resourceFile(),
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	log.Debug().Msgf("[DEBUG] providerConfigure")
	log.Debug().Msgf("[DEBUG] providerConfigure %s , %s",d.Get("name").(string), d.Get("location_files").(string))
	config := &clients.ConfigCustomServer{
		Name:              d.Get("name").(string),
		LocationFiles: d.Get("location_files").(string),

	}

	return clients.NewConfig(config), nil
}
