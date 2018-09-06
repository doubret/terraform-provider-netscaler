package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerNetprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_netprofile,
		Read:          read_netprofile,
		Update:        update_netprofile,
		Delete:        delete_netprofile,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"overridelsn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"srcip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"srcippersistency": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_netprofile(d *schema.ResourceData) nitro.Netprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Netprofile{
		Name:             d.Get("name").(string),
		Overridelsn:      d.Get("overridelsn").(string),
		Srcip:            d.Get("srcip").(string),
		Srcippersistency: d.Get("srcippersistency").(string),
		Td:               d.Get("td").(int),
	}

	return resource
}

func set_netprofile(d *schema.ResourceData, resource *nitro.Netprofile) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("name", resource.Name)
	d.Set("overridelsn", resource.Overridelsn)
	d.Set("srcip", resource.Srcip)
	d.Set("srcippersistency", resource.Srcippersistency)
	d.Set("td", resource.Td)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_netprofile_key(d *schema.ResourceData) nitro.NetprofileKey {

	key := nitro.NetprofileKey{
		d.Get("name").(string),
	}
	return key
}

func create_netprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_netprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_netprofile(d)
	key := resource.ToKey()

	exists, err := client.ExistsNetprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetNetprofile(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_netprofile(d, resource)
	} else {
		err := client.AddNetprofile(get_netprofile(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetNetprofile(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_netprofile(d, resource)
	}

	return nil
}

func read_netprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_netprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_netprofile(d)
	key := resource.ToKey()

	exists, err := client.ExistsNetprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetNetprofile(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_netprofile(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_netprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_netprofile")

	client := meta.(*nitro.NitroClient)

	update := nitro.NetprofileUpdate{}
	unset := nitro.NetprofileUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("srcip") {
		updateFlag = true

		value := d.Get("srcip").(string)
		update.Srcip = value

		if value == "" {
			unsetFlag = true

			unset.Srcip = true
		}

	}
	if d.HasChange("srcippersistency") {
		updateFlag = true

		value := d.Get("srcippersistency").(string)
		update.Srcippersistency = value

		if value == "" {
			unsetFlag = true

			unset.Srcippersistency = true
		}

	}
	if d.HasChange("overridelsn") {
		updateFlag = true

		value := d.Get("overridelsn").(string)
		update.Overridelsn = value

		if value == "" {
			unsetFlag = true

			unset.Overridelsn = true
		}

	}
	key := get_netprofile_key(d)

	if updateFlag {
		if err := client.UpdateNetprofile(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetNetprofile(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetNetprofile(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_netprofile(d, resource)
	}

	return nil
}

func delete_netprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_netprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_netprofile(d)
	key := resource.ToKey()

	exists, err := client.ExistsNetprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteNetprofile(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
