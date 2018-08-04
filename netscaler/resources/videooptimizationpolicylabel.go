package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerVideooptimizationpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_videooptimizationpolicylabel_func,
		//                Read:          read_videooptimizationpolicylabel_func,
		//                Update:        update_videooptimizationpolicylabel_func,
		//                Delete:        delete_videooptimizationpolicylabel_func,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"policylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
