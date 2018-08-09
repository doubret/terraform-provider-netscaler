package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerServiceLbmonitorBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_service_lbmonitor_binding,
		Read:          read_service_lbmonitor_binding,
		Update:        nil,
		Delete:        delete_service_lbmonitor_binding,
		Schema: map[string]*schema.Schema{
			"monitor_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"monstate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"passive": &schema.Schema{
				Type:     schema.TypeBool,
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

func get_service_lbmonitor_binding(d *schema.ResourceData) nitro.ServiceLbmonitorBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.ServiceLbmonitorBinding{
		Monitor_name: d.Get("monitor_name").(string),
		Monstate:     d.Get("monstate").(string),
		Name:         d.Get("name").(string),
		Passive:      d.Get("passive").(bool),
		Weight:       d.Get("weight").(int),
	}

	return resource
}

func set_service_lbmonitor_binding(d *schema.ResourceData, resource *nitro.ServiceLbmonitorBinding) {
	var _ = strconv.Itoa

	d.Set("monitor_name", resource.Monitor_name)
	d.Set("monstate", resource.Monstate)
	d.Set("name", resource.Name)
	d.Set("passive", resource.Passive)
	d.Set("weight", resource.Weight)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_service_lbmonitor_binding_key(d *schema.ResourceData) nitro.ServiceLbmonitorBindingKey {

	key := nitro.ServiceLbmonitorBindingKey{
		d.Get("name").(string),
	}
	return key
}

func create_service_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_service_lbmonitor_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_service_lbmonitor_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsServiceLbmonitorBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServiceLbmonitorBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_service_lbmonitor_binding(d, resource)
	} else {
		err := client.AddServiceLbmonitorBinding(get_service_lbmonitor_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetServiceLbmonitorBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_service_lbmonitor_binding(d, resource)
	}

	return nil
}

func read_service_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_service_lbmonitor_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_service_lbmonitor_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsServiceLbmonitorBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServiceLbmonitorBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_service_lbmonitor_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_service_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_service_lbmonitor_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_service_lbmonitor_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsServiceLbmonitorBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteServiceLbmonitorBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
