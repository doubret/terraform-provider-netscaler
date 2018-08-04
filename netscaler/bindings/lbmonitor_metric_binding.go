package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerLbmonitorMetricBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_lbmonitor_metric_binding_func,
		//                Read:          read_lbmonitor_metric_binding_func,
		//                Update:        update_lbmonitor_metric_binding_func,
		//                Delete:        delete_lbmonitor_metric_binding_func,
		Schema: map[string]*schema.Schema{
			"metric": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metricthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metricweight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"monitorname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
