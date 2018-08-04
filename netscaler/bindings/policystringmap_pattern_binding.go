package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerPolicystringmapPatternBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policystringmap_pattern_binding,
		Read:          read_policystringmap_pattern_binding,
		Update:        nil,
		Delete:        delete_policystringmap_pattern_binding,
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func create_policystringmap_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policystringmap_pattern_binding")

	return nil
}

func read_policystringmap_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policystringmap_pattern_binding")

	return nil
}

func delete_policystringmap_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policystringmap_pattern_binding")

	return nil
}
