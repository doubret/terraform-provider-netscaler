package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerDnspolicy64() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_dnspolicy64,
		Read:          read_dnspolicy64,
		Update:        update_dnspolicy64,
		Delete:        delete_dnspolicy64,
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

func get_dnspolicy64(d *schema.ResourceData) nitro.Dnspolicy64 {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Dnspolicy64{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_dnspolicy64(d *schema.ResourceData, resource *nitro.Dnspolicy64) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("action", resource.Action)
	d.Set("name", resource.Name)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_dnspolicy64_key(d *schema.ResourceData) nitro.Dnspolicy64Key {

	key := nitro.Dnspolicy64Key{
		d.Get("name").(string),
	}
	return key
}

func create_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dnspolicy64")

	client := meta.(*nitro.NitroClient)

	resource := get_dnspolicy64(d)
	key := resource.ToKey()

	exists, err := client.ExistsDnspolicy64(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDnspolicy64(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_dnspolicy64(d, resource)
	} else {
		err := client.AddDnspolicy64(get_dnspolicy64(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetDnspolicy64(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_dnspolicy64(d, resource)
	}

	return nil
}

func read_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_dnspolicy64")

	client := meta.(*nitro.NitroClient)

	resource := get_dnspolicy64(d)
	key := resource.ToKey()

	exists, err := client.ExistsDnspolicy64(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDnspolicy64(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_dnspolicy64(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_dnspolicy64")

	client := meta.(*nitro.NitroClient)

	update := nitro.Dnspolicy64Update{}
	unset := nitro.Dnspolicy64Unset{}

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
	key := get_dnspolicy64_key(d)

	if updateFlag {
		if err := client.UpdateDnspolicy64(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetDnspolicy64(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetDnspolicy64(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_dnspolicy64(d, resource)
	}

	return nil
}

func delete_dnspolicy64(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dnspolicy64")

	client := meta.(*nitro.NitroClient)

	resource := get_dnspolicy64(d)
	key := resource.ToKey()

	exists, err := client.ExistsDnspolicy64(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteDnspolicy64(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
