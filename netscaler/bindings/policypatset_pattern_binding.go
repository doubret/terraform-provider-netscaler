package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerPolicypatsetPatternBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_policypatset_pattern_binding_func,
		//                Read:          read_policypatset_pattern_binding_func,
		//                Update:        update_policypatset_pattern_binding_func,
		//                Delete:        delete_policypatset_pattern_binding_func,
		Schema: map[string]*schema.Schema{
			"charset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
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
			"string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
