package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerCspolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_cspolicylabel,
		Read:          read_cspolicylabel,
		Update:        update_cspolicylabel,
		Delete:        delete_cspolicylabel,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cspolicylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func create_cspolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cspolicylabel")

	return nil
}

func read_cspolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_cspolicylabel")

	return nil
}

func update_cspolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_cspolicylabel")

	return nil
}

func delete_cspolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cspolicylabel")

	return nil
}
