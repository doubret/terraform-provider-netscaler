package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerSpilloveraction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_spilloveraction,
		Read:          read_spilloveraction,
		Update:        update_spilloveraction,
		Delete:        delete_spilloveraction,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func create_spilloveraction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_spilloveraction")

	return nil
}

func read_spilloveraction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_spilloveraction")

	return nil
}

func update_spilloveraction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_spilloveraction")

	return nil
}

func delete_spilloveraction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_spilloveraction")

	return nil
}
