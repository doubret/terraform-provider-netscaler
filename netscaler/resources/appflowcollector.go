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
			"ipaddress": &schema.Schema{
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

func get_appflowcollector(d *schema.ResourceData) nitro.Appflowcollector {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appflowcollector{
		Ipaddress:  d.Get("ipaddress").(string),
		Name:       d.Get("name").(string),
		Netprofile: d.Get("netprofile").(string),
		Port:       d.Get("port").(int),
	}

	return resource
}

func set_appflowcollector(d *schema.ResourceData, resource *nitro.Appflowcollector) {
	var _ = strconv.Itoa

	d.Set("ipaddress", resource.Ipaddress)
	d.Set("name", resource.Name)
	d.Set("netprofile", resource.Netprofile)
	d.Set("port", resource.Port)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_appflowcollector_key(d *schema.ResourceData) nitro.AppflowcollectorKey {

	key := nitro.AppflowcollectorKey{
		d.Get("name").(string),
	}
	return key
}

func create_appflowcollector(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appflowcollector")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowcollector(d)
	key := resource.ToKey()

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

	resource := get_appflowcollector(d)
	key := resource.ToKey()

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

	update := nitro.AppflowcollectorUpdate{}
	unset := nitro.AppflowcollectorUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("ipaddress") {
		updateFlag = true

		value := d.Get("ipaddress").(string)
		update.Ipaddress = value

		if value == "" {
			unsetFlag = true

			unset.Ipaddress = true
		}

	}
	if d.HasChange("port") {
		updateFlag = true

		value := d.Get("port").(int)
		update.Port = value

		if value == 0 {
			unsetFlag = true

			unset.Port = true
		}

	}
	if d.HasChange("netprofile") {
		updateFlag = true

		value := d.Get("netprofile").(string)
		update.Netprofile = value

		if value == "" {
			unsetFlag = true

			unset.Netprofile = true
		}

	}
	key := get_appflowcollector_key(d)

	if updateFlag {
		if err := client.UpdateAppflowcollector(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetAppflowcollector(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetAppflowcollector(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_appflowcollector(d, resource)
	}

	return nil
}

func delete_appflowcollector(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appflowcollector")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowcollector(d)
	key := resource.ToKey()

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
