package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerFilterpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_filterpolicy_func,
		//                Read:          read_filterpolicy_func,
		//                Update:        update_filterpolicy_func,
		//                Delete:        delete_filterpolicy_func,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"reqaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
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
