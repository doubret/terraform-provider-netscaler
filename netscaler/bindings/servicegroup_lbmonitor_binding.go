package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerServicegroupLbmonitorBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_servicegroup_lbmonitor_binding,
		Read:          read_servicegroup_lbmonitor_binding,
		Update:        nil,
		Delete:        delete_servicegroup_lbmonitor_binding,
		Schema: map[string]*schema.Schema{
			"monitor_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func create_servicegroup_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_servicegroup_lbmonitor_binding")

	return nil
}

func read_servicegroup_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_servicegroup_lbmonitor_binding")

	return nil
}

func delete_servicegroup_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_servicegroup_lbmonitor_binding")

	return nil
}
