package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerFilterpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_filterpolicy,
		Read:          read_filterpolicy,
		Update:        update_filterpolicy,
		Delete:        delete_filterpolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"reqaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"resaction": &schema.Schema{
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

func get_filterpolicy(d *schema.ResourceData) nitro.Filterpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Filterpolicy{
		Name:      d.Get("name").(string),
		Reqaction: d.Get("reqaction").(string),
		Resaction: d.Get("resaction").(string),
		Rule:      d.Get("rule").(string),
	}

	return resource
}

func set_filterpolicy(d *schema.ResourceData, resource *nitro.Filterpolicy) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("name", resource.Name)
	d.Set("reqaction", resource.Reqaction)
	d.Set("resaction", resource.Resaction)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_filterpolicy_key(d *schema.ResourceData) nitro.FilterpolicyKey {

	key := nitro.FilterpolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_filterpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_filterpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_filterpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsFilterpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFilterpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_filterpolicy(d, resource)
	} else {
		err := client.AddFilterpolicy(get_filterpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetFilterpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_filterpolicy(d, resource)
	}

	return nil
}

func read_filterpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_filterpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_filterpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsFilterpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFilterpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_filterpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_filterpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_filterpolicy")

	client := meta.(*nitro.NitroClient)

	update := nitro.FilterpolicyUpdate{}
	unset := nitro.FilterpolicyUnset{}

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
	if d.HasChange("reqaction") {
		updateFlag = true

		value := d.Get("reqaction").(string)
		update.Reqaction = value

		if value == "" {
			unsetFlag = true

			unset.Reqaction = true
		}

	}
	if d.HasChange("resaction") {
		updateFlag = true

		value := d.Get("resaction").(string)
		update.Resaction = value

		if value == "" {
			unsetFlag = true

			unset.Resaction = true
		}

	}
	key := get_filterpolicy_key(d)

	if updateFlag {
		if err := client.UpdateFilterpolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetFilterpolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetFilterpolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_filterpolicy(d, resource)
	}

	return nil
}

func delete_filterpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_filterpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_filterpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsFilterpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteFilterpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
