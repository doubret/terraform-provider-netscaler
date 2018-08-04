package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerLbwlm() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbwlm,
		Read:          read_lbwlm,
		Update:        update_lbwlm,
		Delete:        delete_lbwlm,
		Schema: map[string]*schema.Schema{
			"wlmname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"katimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"lbuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func create_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbwlm")

	return nil
}

func read_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbwlm")

	return nil
}

func update_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_lbwlm")

	return nil
}

func delete_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbwlm")

	return nil
}
