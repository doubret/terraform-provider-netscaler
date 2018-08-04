package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerServicegroupServicegroupmemberBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_servicegroup_servicegroupmember_binding,
		Read:          read_servicegroup_servicegroupmember_binding,
		Update:        nil,
		Delete:        delete_servicegroup_servicegroupmember_binding,
		Schema: map[string]*schema.Schema{
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"servername": &schema.Schema{
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

func create_servicegroup_servicegroupmember_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_servicegroup_servicegroupmember_binding")

	return nil
}

func read_servicegroup_servicegroupmember_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_servicegroup_servicegroupmember_binding")

	return nil
}

func delete_servicegroup_servicegroupmember_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_servicegroup_servicegroupmember_binding")

	return nil
}
