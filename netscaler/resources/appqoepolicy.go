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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

func key_appqoepolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_appqoepolicy(d *schema.ResourceData) nitro.Appqoepolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appqoepolicy{
		Name:   d.Get("name").(string),
		Action: d.Get("action").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_appqoepolicy(d *schema.ResourceData, resource *nitro.Appqoepolicy) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("action", resource.Action)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_appqoepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appqoepolicy")

	client := meta.(*nitro.NitroClient)

	key := key_appqoepolicy(d)

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

	key := key_appqoepolicy(d)

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

	err := client.UpdateAppqoepolicy(get_appqoepolicy(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_appqoepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appqoepolicy")

	client := meta.(*nitro.NitroClient)

	key := key_appqoepolicy(d)

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
