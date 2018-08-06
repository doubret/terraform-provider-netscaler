package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAuthorizationpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_authorizationpolicy,
		Read:          read_authorizationpolicy,
		Update:        update_authorizationpolicy,
		Delete:        delete_authorizationpolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func key_authorizationpolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_authorizationpolicy(d *schema.ResourceData) nitro.Authorizationpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Authorizationpolicy{
		Name:   d.Get("name").(string),
		Action: d.Get("action").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_authorizationpolicy(d *schema.ResourceData, resource *nitro.Authorizationpolicy) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("action", resource.Action)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_authorizationpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_authorizationpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_authorizationpolicy(d)

	exists, err := client.ExistsAuthorizationpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuthorizationpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_authorizationpolicy(d, resource)
	} else {
		err := client.AddAuthorizationpolicy(get_authorizationpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAuthorizationpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_authorizationpolicy(d, resource)
	}

	return nil
}

func read_authorizationpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_authorizationpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_authorizationpolicy(d)

	exists, err := client.ExistsAuthorizationpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuthorizationpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_authorizationpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_authorizationpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_authorizationpolicy")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateAuthorizationpolicy(get_authorizationpolicy(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_authorizationpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_authorizationpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_authorizationpolicy(d)

	exists, err := client.ExistsAuthorizationpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAuthorizationpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
