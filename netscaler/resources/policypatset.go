package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerPolicypatset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_policypatset_func,
		//                Read:          read_policypatset_func,
		//                Update:        update_policypatset_func,
		//                Delete:        delete_policypatset_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"indextype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
