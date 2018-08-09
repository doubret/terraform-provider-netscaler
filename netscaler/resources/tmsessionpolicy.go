package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerTmsessionpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_tmsessionpolicy,
		Read:          read_tmsessionpolicy,
		Update:        update_tmsessionpolicy,
		Delete:        delete_tmsessionpolicy,
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

func get_tmsessionpolicy(d *schema.ResourceData) nitro.Tmsessionpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Tmsessionpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_tmsessionpolicy(d *schema.ResourceData, resource *nitro.Tmsessionpolicy) {
	var _ = strconv.Itoa

	d.Set("action", resource.Action)
	d.Set("name", resource.Name)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_tmsessionpolicy_key(d *schema.ResourceData) nitro.TmsessionpolicyKey {

	key := nitro.TmsessionpolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_tmsessionpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_tmsessionpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsTmsessionpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTmsessionpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_tmsessionpolicy(d, resource)
	} else {
		err := client.AddTmsessionpolicy(get_tmsessionpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetTmsessionpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_tmsessionpolicy(d, resource)
	}

	return nil
}

func read_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_tmsessionpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_tmsessionpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsTmsessionpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTmsessionpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_tmsessionpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_tmsessionpolicy")

	client := meta.(*nitro.NitroClient)

	update := nitro.TmsessionpolicyUpdate{}
	unset := nitro.TmsessionpolicyUnset{}

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
	key := get_tmsessionpolicy_key(d)

	if updateFlag {
		if err := client.UpdateTmsessionpolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetTmsessionpolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetTmsessionpolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_tmsessionpolicy(d, resource)
	}

	return nil
}

func delete_tmsessionpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_tmsessionpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_tmsessionpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsTmsessionpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteTmsessionpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
