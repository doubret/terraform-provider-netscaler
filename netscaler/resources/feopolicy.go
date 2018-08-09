package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerFeopolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_feopolicy,
		Read:          read_feopolicy,
		Update:        update_feopolicy,
		Delete:        delete_feopolicy,
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

func get_feopolicy(d *schema.ResourceData) nitro.Feopolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Feopolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_feopolicy(d *schema.ResourceData, resource *nitro.Feopolicy) {
	var _ = strconv.Itoa

	d.Set("action", resource.Action)
	d.Set("name", resource.Name)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_feopolicy_key(d *schema.ResourceData) nitro.FeopolicyKey {

	key := nitro.FeopolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_feopolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_feopolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_feopolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsFeopolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFeopolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_feopolicy(d, resource)
	} else {
		err := client.AddFeopolicy(get_feopolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetFeopolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_feopolicy(d, resource)
	}

	return nil
}

func read_feopolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_feopolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_feopolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsFeopolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFeopolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_feopolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_feopolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_feopolicy")

	client := meta.(*nitro.NitroClient)

	update := nitro.FeopolicyUpdate{}
	unset := nitro.FeopolicyUnset{}

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
	key := get_feopolicy_key(d)

	if updateFlag {
		if err := client.UpdateFeopolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetFeopolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetFeopolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_feopolicy(d, resource)
	}

	return nil
}

func delete_feopolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_feopolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_feopolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsFeopolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteFeopolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
