package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerAppqoepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_appqoepolicy_func,
		//                Read:          read_appqoepolicy_func,
		//                Update:        update_appqoepolicy_func,
		//                Delete:        delete_appqoepolicy_func,
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
