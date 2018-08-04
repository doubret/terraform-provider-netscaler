package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerPolicystringmap() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_policystringmap_func,
		//                Read:          read_policystringmap_func,
		//                Update:        update_policystringmap_func,
		//                Delete:        delete_policystringmap_func,
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
		},
	}
}
