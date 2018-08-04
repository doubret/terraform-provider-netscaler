package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerDbdbprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_dbdbprofile_func,
		//                Read:          read_dbdbprofile_func,
		//                Update:        update_dbdbprofile_func,
		//                Delete:        delete_dbdbprofile_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"conmultiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"enablecachingconmuxoff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"interpretquery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"kcdaccount": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"stickiness": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
