package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerCmppolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_cmppolicy_func,
		//                Read:          read_cmppolicy_func,
		//                Update:        update_cmppolicy_func,
		//                Delete:        delete_cmppolicy_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"resaction": &schema.Schema{
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
