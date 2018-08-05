package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerLbvserverServicegroupBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbvserver_servicegroup_binding,
		Read:          read_lbvserver_servicegroup_binding,
		Update:        nil,
		Delete:        delete_lbvserver_servicegroup_binding,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
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

func key_lbvserver_servicegroup_binding(d *schema.ResourceData) nitro.LbvserverServicegroupBindingKey {
	key := nitro.LbvserverServicegroupBindingKey{
		Name:             d.Get("name").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
	}

	return key
}

func get_lbvserver_servicegroup_binding(d *schema.ResourceData) nitro.LbvserverServicegroupBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.LbvserverServicegroupBinding{
		Name:             d.Get("name").(string),
		Servicegroupname: d.Get("servicegroupname").(string),
		Weight:           d.Get("weight").(int),
	}

	return resource
}

func set_lbvserver_servicegroup_binding(d *schema.ResourceData, resource *nitro.LbvserverServicegroupBinding) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("servicegroupname", resource.Servicegroupname)
	d.Set("weight", resource.Weight)
	var key []string

	key = append(key, resource.Name)
	key = append(key, resource.Servicegroupname)
	d.SetId(strings.Join(key, "-"))
}

func create_lbvserver_servicegroup_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbvserver_servicegroup_binding")

	client := meta.(*nitro.NitroClient)

	key := key_lbvserver_servicegroup_binding(d)

	exists, err := client.ExistsLbvserverServicegroupBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbvserverServicegroupBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbvserver_servicegroup_binding(d, resource)
	} else {
		err := client.AddLbvserverServicegroupBinding(get_lbvserver_servicegroup_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbvserverServicegroupBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbvserver_servicegroup_binding(d, resource)
	}

	return nil
}

func read_lbvserver_servicegroup_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbvserver_servicegroup_binding")

	client := meta.(*nitro.NitroClient)

	key := key_lbvserver_servicegroup_binding(d)

	exists, err := client.ExistsLbvserverServicegroupBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbvserverServicegroupBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbvserver_servicegroup_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_lbvserver_servicegroup_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbvserver_servicegroup_binding")

	client := meta.(*nitro.NitroClient)

	key := key_lbvserver_servicegroup_binding(d)

	exists, err := client.ExistsLbvserverServicegroupBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbvserverServicegroupBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
