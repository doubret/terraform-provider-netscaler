package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAppqoepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_appqoepolicy,
		Read:          read_appqoepolicy,
		Update:        update_appqoepolicy,
		Delete:        delete_appqoepolicy,
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

func get_appqoepolicy(d *schema.ResourceData) nitro.Appqoepolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appqoepolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_appqoepolicy(d *schema.ResourceData, resource *nitro.Appqoepolicy) {
	var _ = strconv.Itoa

	d.Set("action", resource.Action)
	d.Set("name", resource.Name)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_appqoepolicy_key(d *schema.ResourceData) nitro.AppqoepolicyKey {

	key := nitro.AppqoepolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_appqoepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appqoepolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_appqoepolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppqoepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppqoepolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_appqoepolicy(d, resource)
	} else {
		err := client.AddAppqoepolicy(get_appqoepolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAppqoepolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_appqoepolicy(d, resource)
	}

	return nil
}

func read_appqoepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_appqoepolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_appqoepolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppqoepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppqoepolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_appqoepolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_appqoepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_appqoepolicy")

	client := meta.(*nitro.NitroClient)

	update := nitro.AppqoepolicyUpdate{}
	unset := nitro.AppqoepolicyUnset{}

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
	key := get_appqoepolicy_key(d)

	if updateFlag {
		if err := client.UpdateAppqoepolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetAppqoepolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetAppqoepolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_appqoepolicy(d, resource)
	}

	return nil
}

func delete_appqoepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appqoepolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_appqoepolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppqoepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAppqoepolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
