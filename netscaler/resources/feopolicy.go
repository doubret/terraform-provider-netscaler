package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerFeopolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_feopolicy,
		Read:          read_feopolicy,
		Update:        update_feopolicy,
		Delete:        delete_feopolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func create_feopolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_feopolicy")

	return nil
}

func read_feopolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_feopolicy")

	return nil
}

func update_feopolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_feopolicy")

	return nil
}

func delete_feopolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_feopolicy")

	return nil
}
