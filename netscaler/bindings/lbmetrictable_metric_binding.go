package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerLbmetrictableMetricBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_lbmetrictable_metric_binding_func,
		//                Read:          read_lbmetrictable_metric_binding_func,
		//                Update:        update_lbmetrictable_metric_binding_func,
		//                Delete:        delete_lbmetrictable_metric_binding_func,
		Schema: map[string]*schema.Schema{
			"metric": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metrictable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"snmpoid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
