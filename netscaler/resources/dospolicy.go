package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerDospolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_dospolicy,
		Read:          read_dospolicy,
		Update:        update_dospolicy,
		Delete:        delete_dospolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cltdetectrate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"qdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func create_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dospolicy")

	return nil
}

func read_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_dospolicy")

	return nil
}

func update_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_dospolicy")

	return nil
}

func delete_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dospolicy")

	return nil
}
