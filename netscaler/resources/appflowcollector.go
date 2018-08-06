package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAppflowcollector() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_appflowcollector,
		Read:          read_appflowcollector,
		Update:        update_appflowcollector,
		Delete:        delete_appflowcollector,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_appflowcollector(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_appflowcollector(d *schema.ResourceData) nitro.Appflowcollector {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appflowcollector{
		Name:       d.Get("name").(string),
		Ipaddress:  d.Get("ipaddress").(string),
		Netprofile: d.Get("netprofile").(string),
		Port:       d.Get("port").(int),
	}

	return resource
}

func set_appflowcollector(d *schema.ResourceData, resource *nitro.Appflowcollector) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("ipaddress", resource.Ipaddress)
	d.Set("netprofile", resource.Netprofile)
	d.Set("port", resource.Port)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_appflowcollector(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appflowcollector")

	client := meta.(*nitro.NitroClient)

	key := key_appflowcollector(d)

	exists, err := client.ExistsAppflowcollector(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowcollector(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_appflowcollector(d, resource)
	} else {
		err := client.AddAppflowcollector(get_appflowcollector(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAppflowcollector(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_appflowcollector(d, resource)
	}

	return nil
}

func read_appflowcollector(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_appflowcollector")

	client := meta.(*nitro.NitroClient)

	key := key_appflowcollector(d)

	exists, err := client.ExistsAppflowcollector(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowcollector(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_appflowcollector(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_appflowcollector(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_appflowcollector")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateAppflowcollector(get_appflowcollector(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_appflowcollector(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appflowcollector")

	client := meta.(*nitro.NitroClient)

	key := key_appflowcollector(d)

	exists, err := client.ExistsAppflowcollector(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAppflowcollector(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
