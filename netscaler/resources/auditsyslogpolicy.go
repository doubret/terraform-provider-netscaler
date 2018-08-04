package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerAuditsyslogpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_auditsyslogpolicy,
		Read:          read_auditsyslogpolicy,
		Update:        update_auditsyslogpolicy,
		Delete:        delete_auditsyslogpolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func create_auditsyslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_auditsyslogpolicy")

	return nil
}

func read_auditsyslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_auditsyslogpolicy")

	return nil
}

func update_auditsyslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_auditsyslogpolicy")

	return nil
}

func delete_auditsyslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_auditsyslogpolicy")

	return nil
}
