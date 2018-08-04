package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerFeopolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_feopolicy_func,
		//                Read:          read_feopolicy_func,
		//                Update:        update_feopolicy_func,
		//                Delete:        delete_feopolicy_func,
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
