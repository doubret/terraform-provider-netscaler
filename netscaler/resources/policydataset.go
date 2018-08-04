package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerPolicydataset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policydataset,
		Read:          read_policydataset,
		Update:        update_policydataset,
		Delete:        delete_policydataset,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"indextype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func create_policydataset(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policydataset")

	return nil
}

func read_policydataset(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policydataset")

	return nil
}

func update_policydataset(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_policydataset")

	return nil
}

func delete_policydataset(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policydataset")

	return nil
}
