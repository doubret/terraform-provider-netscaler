package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerLbmetrictable() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbmetrictable,
		Read:          read_lbmetrictable,
		Update:        update_lbmetrictable,
		Delete:        delete_lbmetrictable,
		Schema: map[string]*schema.Schema{
			"metrictable": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func create_lbmetrictable(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbmetrictable")

	return nil
}

func read_lbmetrictable(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbmetrictable")

	return nil
}

func update_lbmetrictable(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_lbmetrictable")

	return nil
}

func delete_lbmetrictable(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbmetrictable")

	return nil
}
