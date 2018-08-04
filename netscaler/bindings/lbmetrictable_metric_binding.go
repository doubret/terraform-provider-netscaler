package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerLbmetrictableMetricBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbmetrictable_metric_binding,
		Read:          read_lbmetrictable_metric_binding,
		Update:        nil,
		Delete:        delete_lbmetrictable_metric_binding,
		Schema: map[string]*schema.Schema{
			"metric": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"metrictable": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func create_lbmetrictable_metric_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbmetrictable_metric_binding")

	return nil
}

func read_lbmetrictable_metric_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbmetrictable_metric_binding")

	return nil
}

func delete_lbmetrictable_metric_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbmetrictable_metric_binding")

	return nil
}
