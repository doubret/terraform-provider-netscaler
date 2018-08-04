package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerLbwlm() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_lbwlm_func,
		//                Read:          read_lbwlm_func,
		//                Update:        update_lbwlm_func,
		//                Delete:        delete_lbwlm_func,
		Schema: map[string]*schema.Schema{
			"wlmname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"katimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"lbuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
