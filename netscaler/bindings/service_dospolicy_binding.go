package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerServiceDospolicyBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_service_dospolicy_binding,
		Read:          read_service_dospolicy_binding,
		Update:        nil,
		Delete:        delete_service_dospolicy_binding,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_service_dospolicy_binding(d *schema.ResourceData) nitro.ServiceDospolicyBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.ServiceDospolicyBinding{
		Name:       d.Get("name").(string),
		Policyname: d.Get("policyname").(string),
	}

	return resource
}

func set_service_dospolicy_binding(d *schema.ResourceData, resource *nitro.ServiceDospolicyBinding) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("name", resource.Name)
	d.Set("policyname", resource.Policyname)

	var key []string

	key = append(key, resource.Name)
	key = append(key, resource.Policyname)
	d.SetId(strings.Join(key, "-"))
}

func get_service_dospolicy_binding_key(d *schema.ResourceData) nitro.ServiceDospolicyBindingKey {

	key := nitro.ServiceDospolicyBindingKey{
		d.Get("name").(string),
		d.Get("policyname").(string),
	}
	return key
}

func create_service_dospolicy_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_service_dospolicy_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_service_dospolicy_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsServiceDospolicyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServiceDospolicyBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_service_dospolicy_binding(d, resource)
	} else {
		err := client.AddServiceDospolicyBinding(get_service_dospolicy_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetServiceDospolicyBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_service_dospolicy_binding(d, resource)
	}

	return nil
}

func read_service_dospolicy_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_service_dospolicy_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_service_dospolicy_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsServiceDospolicyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServiceDospolicyBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_service_dospolicy_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_service_dospolicy_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_service_dospolicy_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_service_dospolicy_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsServiceDospolicyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteServiceDospolicyBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
