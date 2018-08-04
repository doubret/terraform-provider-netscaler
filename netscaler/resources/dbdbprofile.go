package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerDbdbprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_dbdbprofile,
		Read:          read_dbdbprofile,
		Update:        update_dbdbprofile,
		Delete:        delete_dbdbprofile,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"conmultiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"enablecachingconmuxoff": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"interpretquery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"kcdaccount": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"stickiness": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func create_dbdbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dbdbprofile")

	return nil
}

func read_dbdbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_dbdbprofile")

	return nil
}

func update_dbdbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_dbdbprofile")

	return nil
}

func delete_dbdbprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dbdbprofile")

	return nil
}
