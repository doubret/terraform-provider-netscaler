package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCmppolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_cmppolicylabel,
		Read:          read_cmppolicylabel,
		Update:        nil,
		Delete:        delete_cmppolicylabel,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_cmppolicylabel(d *schema.ResourceData) nitro.Cmppolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Cmppolicylabel{
		Labelname: d.Get("labelname").(string),
		Type:      d.Get("type").(string),
	}

	return resource
}

func set_cmppolicylabel(d *schema.ResourceData, resource *nitro.Cmppolicylabel) {
	var _ = strconv.Itoa

	d.Set("labelname", resource.Labelname)
	d.Set("type", resource.Type)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func create_cmppolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cmppolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_cmppolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsCmppolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCmppolicylabel(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_cmppolicylabel(d, resource)
	} else {
		err := client.AddCmppolicylabel(get_cmppolicylabel(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCmppolicylabel(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_cmppolicylabel(d, resource)
	}

	return nil
}

func read_cmppolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_cmppolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_cmppolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsCmppolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCmppolicylabel(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_cmppolicylabel(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_cmppolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cmppolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_cmppolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsCmppolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCmppolicylabel(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
