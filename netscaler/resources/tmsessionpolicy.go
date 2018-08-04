package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerTmsessionpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_tmsessionpolicy,
		Read:          read_tmsessionpolicy,
		Update:        update_tmsessionpolicy,
		Delete:        delete_tmsessionpolicy,
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

func create_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_tmsessionpolicy")

	return nil
}

func read_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_tmsessionpolicy")

	return nil
}

func update_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_tmsessionpolicy")

	return nil
}

func delete_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_tmsessionpolicy")

	return nil
}
