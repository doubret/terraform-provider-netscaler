package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerCmppolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_cmppolicylabel,
		Read:          read_cmppolicylabel,
		Update:        update_cmppolicylabel,
		Delete:        delete_cmppolicylabel,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func create_cmppolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cmppolicylabel")

	return nil
}

func read_cmppolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_cmppolicylabel")

	return nil
}

func update_cmppolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_cmppolicylabel")

	return nil
}

func delete_cmppolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cmppolicylabel")

	return nil
}
