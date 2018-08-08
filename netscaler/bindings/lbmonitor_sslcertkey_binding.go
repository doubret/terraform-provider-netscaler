package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerLbmonitorSslcertkeyBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbmonitor_sslcertkey_binding,
		Read:          read_lbmonitor_sslcertkey_binding,
		Update:        nil,
		Delete:        delete_lbmonitor_sslcertkey_binding,
		Schema: map[string]*schema.Schema{
			"ca": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certkeyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"crlcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"monitorname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ocspcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_lbmonitor_sslcertkey_binding(d *schema.ResourceData) nitro.LbmonitorSslcertkeyBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.LbmonitorSslcertkeyBinding{
		Ca:          d.Get("ca").(bool),
		Certkeyname: d.Get("certkeyname").(string),
		Crlcheck:    d.Get("crlcheck").(string),
		Monitorname: d.Get("monitorname").(string),
		Ocspcheck:   d.Get("ocspcheck").(string),
	}

	return resource
}

func set_lbmonitor_sslcertkey_binding(d *schema.ResourceData, resource *nitro.LbmonitorSslcertkeyBinding) {
	var _ = strconv.Itoa

	d.Set("ca", resource.Ca)
	d.Set("certkeyname", resource.Certkeyname)
	d.Set("crlcheck", resource.Crlcheck)
	d.Set("monitorname", resource.Monitorname)
	d.Set("ocspcheck", resource.Ocspcheck)

	var key []string

	key = append(key, resource.Monitorname)
	key = append(key, resource.Certkeyname)
	d.SetId(strings.Join(key, "-"))
}

func create_lbmonitor_sslcertkey_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbmonitor_sslcertkey_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_lbmonitor_sslcertkey_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbmonitorSslcertkeyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbmonitorSslcertkeyBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbmonitor_sslcertkey_binding(d, resource)
	} else {
		err := client.AddLbmonitorSslcertkeyBinding(get_lbmonitor_sslcertkey_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbmonitorSslcertkeyBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbmonitor_sslcertkey_binding(d, resource)
	}

	return nil
}

func read_lbmonitor_sslcertkey_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbmonitor_sslcertkey_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_lbmonitor_sslcertkey_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbmonitorSslcertkeyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbmonitorSslcertkeyBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbmonitor_sslcertkey_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_lbmonitor_sslcertkey_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbmonitor_sslcertkey_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_lbmonitor_sslcertkey_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsLbmonitorSslcertkeyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbmonitorSslcertkeyBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
