package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerSpilloverpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_spilloverpolicy,
		Read:          read_spilloverpolicy,
		Update:        update_spilloverpolicy,
		Delete:        delete_spilloverpolicy,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"comment": &schema.Schema{
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

func get_spilloverpolicy(d *schema.ResourceData) nitro.Spilloverpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Spilloverpolicy{
		Action:  d.Get("action").(string),
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Rule:    d.Get("rule").(string),
	}

	return resource
}

func set_spilloverpolicy(d *schema.ResourceData, resource *nitro.Spilloverpolicy) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("action", resource.Action)
	d.Set("comment", resource.Comment)
	d.Set("name", resource.Name)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_spilloverpolicy_key(d *schema.ResourceData) nitro.SpilloverpolicyKey {

	key := nitro.SpilloverpolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_spilloverpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_spilloverpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_spilloverpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsSpilloverpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSpilloverpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_spilloverpolicy(d, resource)
	} else {
		err := client.AddSpilloverpolicy(get_spilloverpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetSpilloverpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_spilloverpolicy(d, resource)
	}

	return nil
}

func read_spilloverpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_spilloverpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_spilloverpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsSpilloverpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSpilloverpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_spilloverpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_spilloverpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_spilloverpolicy")

	client := meta.(*nitro.NitroClient)

	update := nitro.SpilloverpolicyUpdate{}
	unset := nitro.SpilloverpolicyUnset{}

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
	if d.HasChange("comment") {
		updateFlag = true

		value := d.Get("comment").(string)
		update.Comment = value

		if value == "" {
			unsetFlag = true

			unset.Comment = true
		}

	}
	key := get_spilloverpolicy_key(d)

	if updateFlag {
		if err := client.UpdateSpilloverpolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetSpilloverpolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetSpilloverpolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_spilloverpolicy(d, resource)
	}

	return nil
}

func delete_spilloverpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_spilloverpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_spilloverpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsSpilloverpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteSpilloverpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
