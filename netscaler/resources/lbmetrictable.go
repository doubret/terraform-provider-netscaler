package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerLbmetrictable() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_lbmetrictable_func,
		//                Read:          read_lbmetrictable_func,
		//                Update:        update_lbmetrictable_func,
		//                Delete:        delete_lbmetrictable_func,
		Schema: map[string]*schema.Schema{
			"metrictable": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
