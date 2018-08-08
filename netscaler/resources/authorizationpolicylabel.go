package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAuthorizationpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_authorizationpolicylabel,
		Read:          read_authorizationpolicylabel,
		Update:        nil,
		Delete:        delete_authorizationpolicylabel,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_authorizationpolicylabel(d *schema.ResourceData) nitro.Authorizationpolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Authorizationpolicylabel{
		Labelname: d.Get("labelname").(string),
	}

	return resource
}

func set_authorizationpolicylabel(d *schema.ResourceData, resource *nitro.Authorizationpolicylabel) {
	var _ = strconv.Itoa

	d.Set("labelname", resource.Labelname)

	var key []string

	key = append(key, resource.Labelname)
	d.SetId(strings.Join(key, "-"))
}

func create_authorizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_authorizationpolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_authorizationpolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsAuthorizationpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuthorizationpolicylabel(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_authorizationpolicylabel(d, resource)
	} else {
		err := client.AddAuthorizationpolicylabel(get_authorizationpolicylabel(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAuthorizationpolicylabel(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_authorizationpolicylabel(d, resource)
	}

	return nil
}

func read_authorizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_authorizationpolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_authorizationpolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsAuthorizationpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuthorizationpolicylabel(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_authorizationpolicylabel(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_authorizationpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_authorizationpolicylabel")

	client := meta.(*nitro.NitroClient)

	resource := get_authorizationpolicylabel(d)
	key := resource.ToKey()

	exists, err := client.ExistsAuthorizationpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAuthorizationpolicylabel(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
