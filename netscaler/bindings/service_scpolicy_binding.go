package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerServiceScpolicyBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_service_scpolicy_binding_func,
		//                Read:          read_service_scpolicy_binding_func,
		//                Update:        update_service_scpolicy_binding_func,
		//                Delete:        delete_service_scpolicy_binding_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
