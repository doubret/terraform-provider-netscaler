package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerTmtrafficpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_tmtrafficpolicy_func,
		//                Read:          read_tmtrafficpolicy_func,
		//                Update:        update_tmtrafficpolicy_func,
		//                Delete:        delete_tmtrafficpolicy_func,
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
