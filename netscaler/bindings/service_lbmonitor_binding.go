package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerServiceLbmonitorBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_service_lbmonitor_binding,
		Read:          read_service_lbmonitor_binding,
		Update:        nil,
		Delete:        delete_service_lbmonitor_binding,
		Schema: map[string]*schema.Schema{
			"monitor_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"monstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"passive": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func create_service_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_service_lbmonitor_binding")

	return nil
}

func read_service_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_service_lbmonitor_binding")

	return nil
}

func delete_service_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_service_lbmonitor_binding")

	return nil
}
