package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerLbvserverServiceBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbvserver_service_binding,
		Read:          read_lbvserver_service_binding,
		Update:        nil,
		Delete:        delete_lbvserver_service_binding,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_lbvserver_service_binding(d *schema.ResourceData) nitro.LbvserverServiceBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.LbvserverServiceBinding{
		Name:        d.Get("name").(string),
		Servicename: d.Get("servicename").(string),
		Weight:      d.Get("weight").(int),
	}

	return resource
}

func set_lbvserver_service_binding(d *schema.ResourceData, resource *nitro.LbvserverServiceBinding) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("servicename", resource.Servicename)
	d.Set("weight", resource.Weight)

	var key []string

	key = append(key, resource.Name)
	key = append(key, resource.Servicename)
	d.SetId(strings.Join(key, "-"))
}

func get_lbvserver_service_binding_key(d *schema.ResourceData) nitro.LbvserverServiceBindingKey {

	key := nitro.LbvserverServiceBindingKey{
		d.Get("name").(string),
		d.Get("servicename").(string),
	}
	return key
}

func create_lbvserver_service_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbvserver_service_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_lbvserver_service_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbvserverServiceBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbvserverServiceBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbvserver_service_binding(d, resource)
	} else {
		err := client.AddLbvserverServiceBinding(get_lbvserver_service_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbvserverServiceBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbvserver_service_binding(d, resource)
	}

	return nil
}

func read_lbvserver_service_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbvserver_service_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_lbvserver_service_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbvserverServiceBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbvserverServiceBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbvserver_service_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_lbvserver_service_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbvserver_service_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_lbvserver_service_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbvserverServiceBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbvserverServiceBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
