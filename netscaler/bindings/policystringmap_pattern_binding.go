package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerPolicystringmapPatternBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_policystringmap_pattern_binding_func,
		//                Read:          read_policystringmap_pattern_binding_func,
		//                Update:        update_policystringmap_pattern_binding_func,
		//                Delete:        delete_policystringmap_pattern_binding_func,
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
