package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerVideooptimizationpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_videooptimizationpolicylabel,
		Read:          read_videooptimizationpolicylabel,
		Update:        update_videooptimizationpolicylabel,
		Delete:        delete_videooptimizationpolicylabel,
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

func create_videooptimizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_videooptimizationpolicylabel")

	return nil
}

func read_videooptimizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_videooptimizationpolicylabel")

	return nil
}

func update_videooptimizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_videooptimizationpolicylabel")

	return nil
}

func delete_videooptimizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_videooptimizationpolicylabel")

	return nil
}
