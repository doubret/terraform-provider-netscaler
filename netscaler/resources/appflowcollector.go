package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerAppflowcollector() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_appflowcollector,
		Read:          read_appflowcollector,
		Update:        update_appflowcollector,
		Delete:        delete_appflowcollector,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"transport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func create_appflowcollector(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appflowcollector")

	return nil
}

func read_appflowcollector(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_appflowcollector")

	return nil
}

func update_appflowcollector(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_appflowcollector")

	return nil
}

func delete_appflowcollector(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appflowcollector")

	return nil
}
