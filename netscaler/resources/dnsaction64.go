package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerDnsaction64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_dnsaction64,
		Read:          read_dnsaction64,
		Update:        update_dnsaction64,
		Delete:        delete_dnsaction64,
		Schema: map[string]*schema.Schema{
			"actionname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"excluderule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mappedrule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func create_dnsaction64(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dnsaction64")

	return nil
}

func read_dnsaction64(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_dnsaction64")

	return nil
}

func update_dnsaction64(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_dnsaction64")

	return nil
}

func delete_dnsaction64(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dnsaction64")

	return nil
}
