package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerTmsessionpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_tmsessionpolicy_func,
		//                Read:          read_tmsessionpolicy_func,
		//                Update:        update_tmsessionpolicy_func,
		//                Delete:        delete_tmsessionpolicy_func,
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
