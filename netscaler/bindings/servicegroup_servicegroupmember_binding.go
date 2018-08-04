package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerServicegroupServicegroupmemberBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_servicegroup_servicegroupmember_binding_func,
		//                Read:          read_servicegroup_servicegroupmember_binding_func,
		//                Update:        update_servicegroup_servicegroupmember_binding_func,
		//                Delete:        delete_servicegroup_servicegroupmember_binding_func,
		Schema: map[string]*schema.Schema{
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servername": &schema.Schema{
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
