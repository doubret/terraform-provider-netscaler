package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerPolicypatsetPatternBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policypatset_pattern_binding,
		Read:          read_policypatset_pattern_binding,
		Update:        nil,
		Delete:        delete_policypatset_pattern_binding,
		Schema: map[string]*schema.Schema{
			"charset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"string": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func create_policypatset_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policypatset_pattern_binding")

	return nil
}

func read_policypatset_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policypatset_pattern_binding")

	return nil
}

func delete_policypatset_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policypatset_pattern_binding")

	return nil
}
