package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerLbwlm() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbwlm,
		Read:          read_lbwlm,
		Update:        update_lbwlm,
		Delete:        delete_lbwlm,
		Schema: map[string]*schema.Schema{
			"ipaddress": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"katimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"lbuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"wlmname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_lbwlm(d *schema.ResourceData) nitro.Lbwlm {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Lbwlm{
		Ipaddress: d.Get("ipaddress").(string),
		Katimeout: d.Get("katimeout").(int),
		Lbuid:     d.Get("lbuid").(string),
		Port:      d.Get("port").(int),
		Wlmname:   d.Get("wlmname").(string),
	}

	return resource
}

func set_lbwlm(d *schema.ResourceData, resource *nitro.Lbwlm) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("ipaddress", resource.Ipaddress)
	d.Set("katimeout", resource.Katimeout)
	d.Set("lbuid", resource.Lbuid)
	d.Set("port", resource.Port)
	d.Set("wlmname", resource.Wlmname)

	var key []string

	key = append(key, resource.Wlmname)
	d.SetId(strings.Join(key, "-"))
}

func get_lbwlm_key(d *schema.ResourceData) nitro.LbwlmKey {

	key := nitro.LbwlmKey{
		d.Get("wlmname").(string),
	}
	return key
}

func create_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbwlm")

	client := meta.(*nitro.NitroClient)

	resource := get_lbwlm(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbwlm(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbwlm(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbwlm(d, resource)
	} else {
		err := client.AddLbwlm(get_lbwlm(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbwlm(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbwlm(d, resource)
	}

	return nil
}

func read_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbwlm")

	client := meta.(*nitro.NitroClient)

	resource := get_lbwlm(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbwlm(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbwlm(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbwlm(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_lbwlm")

	client := meta.(*nitro.NitroClient)

	update := nitro.LbwlmUpdate{}
	unset := nitro.LbwlmUnset{}

	updateFlag := false
	unsetFlag := false

	update.Wlmname = d.Get("wlmname").(string)
	unset.Wlmname = d.Get("wlmname").(string)

	if d.HasChange("katimeout") {
		updateFlag = true

		value := d.Get("katimeout").(int)
		update.Katimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Katimeout = true
		}

	}
	key := get_lbwlm_key(d)

	if updateFlag {
		if err := client.UpdateLbwlm(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetLbwlm(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetLbwlm(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_lbwlm(d, resource)
	}

	return nil
}

func delete_lbwlm(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbwlm")

	client := meta.(*nitro.NitroClient)

	resource := get_lbwlm(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbwlm(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbwlm(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
