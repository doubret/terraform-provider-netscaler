package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
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

func get_authorizationpolicy(d *schema.ResourceData) nitro.Authorizationpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Authorizationpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_authorizationpolicy(d *schema.ResourceData, resource *nitro.Authorizationpolicy) {
	var _ = strconv.Itoa

	d.Set("action", resource.Action)
	d.Set("name", resource.Name)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_authorizationpolicy_key(d *schema.ResourceData) nitro.AuthorizationpolicyKey {

	key := nitro.AuthorizationpolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_authorizationpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_authorizationpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_authorizationpolicy(d)
	key := resource.ToKey()

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

	resource := get_authorizationpolicy(d)
	key := resource.ToKey()

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

	update := nitro.AuthorizationpolicyUpdate{}
	unset := nitro.AuthorizationpolicyUnset{}

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
	key := get_authorizationpolicy_key(d)

	if updateFlag {
		if err := client.UpdateAuthorizationpolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetAuthorizationpolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetAuthorizationpolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_authorizationpolicy(d, resource)
	}

	return nil
}

func delete_authorizationpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_authorizationpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_authorizationpolicy(d)
	key := resource.ToKey()

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
