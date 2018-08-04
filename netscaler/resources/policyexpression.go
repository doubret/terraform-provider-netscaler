package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerPolicyexpression() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_policyexpression_func,
		//                Read:          read_policyexpression_func,
		//                Update:        update_policyexpression_func,
		//                Delete:        delete_policyexpression_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"clientsecuritymessage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
