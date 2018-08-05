package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerServicegroupLbmonitorBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_servicegroup_lbmonitor_binding,
		Read:          read_servicegroup_lbmonitor_binding,
		Update:        nil,
		Delete:        delete_servicegroup_lbmonitor_binding,
		Schema: map[string]*schema.Schema{
			"monitor_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func key_servicegroup_lbmonitor_binding(d *schema.ResourceData) nitro.ServicegroupLbmonitorBindingKey {
	key := nitro.ServicegroupLbmonitorBindingKey{
		Servicegroupname: d.Get("servicegroupname").(string),
		Monitor_name:     d.Get("monitor_name").(string),
	}

	return key
}

func get_servicegroup_lbmonitor_binding(d *schema.ResourceData) nitro.ServicegroupLbmonitorBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.ServicegroupLbmonitorBinding{
		Monitor_name:     d.Get("monitor_name").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
		Weight:           d.Get("weight").(int),
	}

	return resource
}

func set_servicegroup_lbmonitor_binding(d *schema.ResourceData, resource *nitro.ServicegroupLbmonitorBinding) {
	var _ = strconv.Itoa

	d.Set("monitor_name", resource.Monitor_name)
	d.Set("servicegroupname", resource.Servicegroupname)
	d.Set("weight", resource.Weight)
	var key []string

	key = append(key, resource.Servicegroupname)
	key = append(key, resource.Monitor_name)
	d.SetId(strings.Join(key, "-"))
}

func create_servicegroup_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_servicegroup_lbmonitor_binding")

	client := meta.(*nitro.NitroClient)

	key := key_servicegroup_lbmonitor_binding(d)

	exists, err := client.ExistsServicegroupLbmonitorBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServicegroupLbmonitorBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_servicegroup_lbmonitor_binding(d, resource)
	} else {
		err := client.AddServicegroupLbmonitorBinding(get_servicegroup_lbmonitor_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetServicegroupLbmonitorBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_servicegroup_lbmonitor_binding(d, resource)
	}

	return nil
}

func read_servicegroup_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_servicegroup_lbmonitor_binding")

	client := meta.(*nitro.NitroClient)

	key := key_servicegroup_lbmonitor_binding(d)

	exists, err := client.ExistsServicegroupLbmonitorBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServicegroupLbmonitorBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_servicegroup_lbmonitor_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_servicegroup_lbmonitor_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_servicegroup_lbmonitor_binding")

	client := meta.(*nitro.NitroClient)

	key := key_servicegroup_lbmonitor_binding(d)

	exists, err := client.ExistsServicegroupLbmonitorBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteServicegroupLbmonitorBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
