package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerServicegroupLbmonitorBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_servicegroup_lbmonitor_binding_func,
		//                Read:          read_servicegroup_lbmonitor_binding_func,
		//                Update:        update_servicegroup_lbmonitor_binding_func,
		//                Delete:        delete_servicegroup_lbmonitor_binding_func,
		Schema: map[string]*schema.Schema{
			"monitor_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
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
