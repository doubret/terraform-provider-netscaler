package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_policystringmap(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_policystringmap(d *schema.ResourceData) nitro.Policystringmap {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Policystringmap{
		Name:    d.Get("name").(string),
		Comment: d.Get("comment").(string),
	}

	return resource
}

func set_policystringmap(d *schema.ResourceData, resource *nitro.Policystringmap) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("comment", resource.Comment)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_policystringmap(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policystringmap")

	client := meta.(*nitro.NitroClient)

	key := key_policystringmap(d)

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

	key := key_policystringmap(d)

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

	err := client.UpdatePolicystringmap(get_policystringmap(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_policystringmap(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policystringmap")

	client := meta.(*nitro.NitroClient)

	key := key_policystringmap(d)

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
