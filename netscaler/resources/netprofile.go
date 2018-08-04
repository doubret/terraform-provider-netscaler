package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerNetprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_netprofile_func,
		//                Read:          read_netprofile_func,
		//                Update:        update_netprofile_func,
		//                Delete:        delete_netprofile_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"overridelsn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"srcip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"srcippersistency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
