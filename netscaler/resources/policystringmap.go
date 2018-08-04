package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerPolicystringmap() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policystringmap,
		Read:          read_policystringmap,
		Update:        update_policystringmap,
		Delete:        delete_policystringmap,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func create_policystringmap(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policystringmap")

	return nil
}

func read_policystringmap(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policystringmap")

	return nil
}

func update_policystringmap(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_policystringmap")

	return nil
}

func delete_policystringmap(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policystringmap")

	return nil
}
