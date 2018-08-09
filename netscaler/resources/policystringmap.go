package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerPolicystringmap() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policystringmap,
		Read:          read_policystringmap,
		Update:        update_policystringmap,
		Delete:        delete_policystringmap,
		Schema: map[string]*schema.Schema{
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
		},
	}
}

func get_policystringmap(d *schema.ResourceData) nitro.Policystringmap {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Policystringmap{
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
	}

	return resource
}

func set_policystringmap(d *schema.ResourceData, resource *nitro.Policystringmap) {
	var _ = strconv.Itoa

	d.Set("comment", resource.Comment)
	d.Set("name", resource.Name)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_policystringmap_key(d *schema.ResourceData) nitro.PolicystringmapKey {

	key := nitro.PolicystringmapKey{
		d.Get("name").(string),
	}
	return key
}

func create_policystringmap(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policystringmap")

	client := meta.(*nitro.NitroClient)

	resource := get_policystringmap(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicystringmap(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicystringmap(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_policystringmap(d, resource)
	} else {
		err := client.AddPolicystringmap(get_policystringmap(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetPolicystringmap(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_policystringmap(d, resource)
	}

	return nil
}

func read_policystringmap(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policystringmap")

	client := meta.(*nitro.NitroClient)

	resource := get_policystringmap(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicystringmap(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicystringmap(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_policystringmap(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_policystringmap(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_policystringmap")

	client := meta.(*nitro.NitroClient)

	update := nitro.PolicystringmapUpdate{}
	unset := nitro.PolicystringmapUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("comment") {
		updateFlag = true

		value := d.Get("comment").(string)
		update.Comment = value

		if value == "" {
			unsetFlag = true

			unset.Comment = true
		}

	}
	key := get_policystringmap_key(d)

	if updateFlag {
		if err := client.UpdatePolicystringmap(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetPolicystringmap(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetPolicystringmap(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_policystringmap(d, resource)
	}

	return nil
}

func delete_policystringmap(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policystringmap")

	client := meta.(*nitro.NitroClient)

	resource := get_policystringmap(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicystringmap(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeletePolicystringmap(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
