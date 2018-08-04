package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerServiceLbmonitorBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_service_lbmonitor_binding_func,
		//                Read:          read_service_lbmonitor_binding_func,
		//                Update:        update_service_lbmonitor_binding_func,
		//                Delete:        delete_service_lbmonitor_binding_func,
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
				Optional: true,
				Computed: true,
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
