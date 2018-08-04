package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerCspolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_cspolicy_func,
		//                Read:          read_cspolicy_func,
		//                Update:        update_cspolicy_func,
		//                Delete:        delete_cspolicy_func,
		Schema: map[string]*schema.Schema{
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"logaction": &schema.Schema{
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
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}
