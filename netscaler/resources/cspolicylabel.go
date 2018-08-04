package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerCspolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_cspolicylabel_func,
		//                Read:          read_cspolicylabel_func,
		//                Update:        update_cspolicylabel_func,
		//                Delete:        delete_cspolicylabel_func,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cspolicylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
