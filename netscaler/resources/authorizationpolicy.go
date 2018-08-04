package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerAuthorizationpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_authorizationpolicy_func,
		//                Read:          read_authorizationpolicy_func,
		//                Update:        update_authorizationpolicy_func,
		//                Delete:        delete_authorizationpolicy_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
