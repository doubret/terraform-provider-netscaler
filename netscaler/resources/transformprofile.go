package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerTransformprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_transformprofile,
		Read:          read_transformprofile,
		Update:        update_transformprofile,
		Delete:        delete_transformprofile,
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
			"onlytransformabsurlinbody": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func create_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_transformprofile")

	return nil
}

func read_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_transformprofile")

	return nil
}

func update_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_transformprofile")

	return nil
}

func delete_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_transformprofile")

	return nil
}
