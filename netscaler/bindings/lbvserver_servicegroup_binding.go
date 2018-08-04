package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerLbvserverServicegroupBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_lbvserver_servicegroup_binding_func,
		//                Read:          read_lbvserver_servicegroup_binding_func,
		//                Update:        update_lbvserver_servicegroup_binding_func,
		//                Delete:        delete_lbvserver_servicegroup_binding_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
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
