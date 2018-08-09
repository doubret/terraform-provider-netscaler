package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAuditnslogpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_auditnslogpolicy,
		Read:          read_auditnslogpolicy,
		Update:        update_auditnslogpolicy,
		Delete:        delete_auditnslogpolicy,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func get_auditnslogpolicy(d *schema.ResourceData) nitro.Auditnslogpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Auditnslogpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_auditnslogpolicy(d *schema.ResourceData, resource *nitro.Auditnslogpolicy) {
	var _ = strconv.Itoa

	d.Set("action", resource.Action)
	d.Set("name", resource.Name)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_auditnslogpolicy_key(d *schema.ResourceData) nitro.AuditnslogpolicyKey {

	key := nitro.AuditnslogpolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_auditnslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_auditnslogpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_auditnslogpolicy(d)
	key := resource.ToKey()

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

	resource := get_auditnslogpolicy(d)
	key := resource.ToKey()

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

	client := meta.(*nitro.NitroClient)

	update := nitro.AuditnslogpolicyUpdate{}
	unset := nitro.AuditnslogpolicyUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("rule") {
		updateFlag = true

		value := d.Get("rule").(string)
		update.Rule = value

		if value == "" {
			unsetFlag = true

			unset.Rule = true
		}

	}
	if d.HasChange("action") {
		updateFlag = true

		value := d.Get("action").(string)
		update.Action = value

		if value == "" {
			unsetFlag = true

			unset.Action = true
		}

	}
	key := get_auditnslogpolicy_key(d)

	if updateFlag {
		if err := client.UpdateAuditnslogpolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetAuditnslogpolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetAuditnslogpolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_auditnslogpolicy(d, resource)
	}

	return nil
}

func delete_auditnslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_auditnslogpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_auditnslogpolicy(d)
	key := resource.ToKey()

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
