package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerAppflowpolicylabelAppflowpolicyBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_appflowpolicylabel_appflowpolicy_binding_func,
		//                Read:          read_appflowpolicylabel_appflowpolicy_binding_func,
		//                Update:        update_appflowpolicylabel_appflowpolicy_binding_func,
		//                Delete:        delete_appflowpolicylabel_appflowpolicy_binding_func,
		Schema: map[string]*schema.Schema{
			"gotopriorityexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"invoke": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"invoke_labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"labeltype": &schema.Schema{
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
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
