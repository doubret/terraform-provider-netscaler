package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerPolicyexpression() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policyexpression,
		Read:          read_policyexpression,
		Update:        update_policyexpression,
		Delete:        delete_policyexpression,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"clientsecuritymessage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func create_policyexpression(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policyexpression")

	return nil
}

func read_policyexpression(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policyexpression")

	return nil
}

func update_policyexpression(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_policyexpression")

	return nil
}

func delete_policyexpression(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policyexpression")

	return nil
}
