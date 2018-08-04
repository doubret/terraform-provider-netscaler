package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerAuditnslogpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_auditnslogpolicy,
		Read:          read_auditnslogpolicy,
		Update:        update_auditnslogpolicy,
		Delete:        delete_auditnslogpolicy,
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

func key_auditnslogpolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_auditnslogpolicy(d *schema.ResourceData) nitro.Auditnslogpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Auditnslogpolicy{
		Name:   d.Get("name").(string),
		Action: d.Get("action").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_auditnslogpolicy(d *schema.ResourceData, resource *nitro.Auditnslogpolicy) {
	d.Set("name", resource.Name)
	d.Set("action", resource.Action)
	d.Set("rule", resource.Rule)
	d.SetId(resource.Name)
}

func create_auditnslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_auditnslogpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_auditnslogpolicy(d)

	exists, err := client.ExistsAuditnslogpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuditnslogpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_auditnslogpolicy(d, resource)
	} else {
		err := client.AddAuditnslogpolicy(get_auditnslogpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAuditnslogpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_auditnslogpolicy(d, resource)
	}

	return nil
}

func read_auditnslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_auditnslogpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_auditnslogpolicy(d)

	exists, err := client.ExistsAuditnslogpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuditnslogpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_auditnslogpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_auditnslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_auditnslogpolicy")

	return nil
}

func delete_auditnslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_auditnslogpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_auditnslogpolicy(d)

	exists, err := client.ExistsAuditnslogpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAuditnslogpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
