package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerPolicypatset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policypatset,
		Read:          read_policypatset,
		Update:        update_policypatset,
		Delete:        delete_policypatset,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"indextype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func create_policypatset(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policypatset")

	return nil
}

func read_policypatset(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policypatset")

	return nil
}

func update_policypatset(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_policypatset")

	return nil
}

func delete_policypatset(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policypatset")

	return nil
}
