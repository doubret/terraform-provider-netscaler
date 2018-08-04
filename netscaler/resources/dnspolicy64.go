package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerDnspolicy64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_dnspolicy64,
		Read:          read_dnspolicy64,
		Update:        update_dnspolicy64,
		Delete:        delete_dnspolicy64,
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

func create_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dnspolicy64")

	return nil
}

func read_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_dnspolicy64")

	return nil
}

func update_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_dnspolicy64")

	return nil
}

func delete_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dnspolicy64")

	return nil
}
