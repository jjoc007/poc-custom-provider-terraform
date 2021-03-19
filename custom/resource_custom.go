package custom

import (
	"github.com/jjoc007/poc-custom-provider-terraform/custom/commons"
	"github.com/rs/zerolog/log"

	"github.com/jjoc007/poc-custom-provider-terraform/custom/clients"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Schema: commons.FileSchema,
		Create: resourceCreateFile,
		Read:   resourceReadFile,
		Delete: resourceDeleteFile,
	}
}

func resourceCreateFile(d *schema.ResourceData, m interface{}) (err error) {
	fileClient := &clients.FileClient{Config: m.(clients.Config)}

	file := &clients.File{
		Name:    d.Get("name").(string),
		Content: d.Get("content").(string),
	}

	if _, err := fileClient.CreateFile(file); err != nil {
		log.Err(err).Msgf("[DEBUG] custom::create file name  %s",file.Name)
		return err
	}

	d.SetId(file.Name)

	return resourceReadFile(d, m)
}

//nolint:dupl Similar Methods between entities
func resourceReadFile(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	log.Debug().Msgf("[DEBUG] custom::read - file name %s ",name)
	fileClient := &clients.FileClient{Config: m.(clients.Config)}

	file, err := fileClient.GetFile(name)
	if err != nil {
		log.Err(err).Msgf("[DEBUG] custom::create file name  %s",file.Name)
		return err
	}

	d.Set("name", file.Name)
	d.Set("content", file.Content)

	log.Debug().Msgf("[DEBUG] custom::read - file name %s exist",name)
	return nil
}

func resourceDeleteFile(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	log.Debug().Msgf("[DEBUG] custom::delete - file name %s ",name)
	fileClient := &clients.FileClient{Config: m.(clients.Config)}
	result, err := fileClient.DeleteFile(name)
	if err != nil || !result {
		log.Err(err).Msgf("custom::delete file name  %s not exist",name)
		return err
	}

	log.Debug().Msgf("custom::delete - file name %s deleted",name)
	return nil
}
