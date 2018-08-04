package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerPolicydatasetValueBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policydataset_value_binding,
		Read:          read_policydataset_value_binding,
		Update:        nil,
		Delete:        delete_policydataset_value_binding,
		Schema: map[string]*schema.Schema{
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
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func create_policydataset_value_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policydataset_value_binding")

	return nil
}

func read_policydataset_value_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policydataset_value_binding")

	return nil
}

func delete_policydataset_value_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policydataset_value_binding")

	return nil
}
