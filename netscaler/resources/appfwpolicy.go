package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerAppfwpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_appfwpolicy_func,
		//                Read:          read_appfwpolicy_func,
		//                Update:        update_appfwpolicy_func,
		//                Delete:        delete_appfwpolicy_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
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
			"profilename": &schema.Schema{
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
