package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerSslvserverSslcertkeyBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_sslvserver_sslcertkey_binding,
		Read:          read_sslvserver_sslcertkey_binding,
		Update:        nil,
		Delete:        delete_sslvserver_sslcertkey_binding,
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
			"ocspcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"skipcaname": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"snicert": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"vservername": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_sslvserver_sslcertkey_binding(d *schema.ResourceData) nitro.SslvserverSslcertkeyBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.SslvserverSslcertkeyBinding{
		Ca:          d.Get("ca").(bool),
		Certkeyname: d.Get("certkeyname").(string),
		Crlcheck:    d.Get("crlcheck").(string),
		Ocspcheck:   d.Get("ocspcheck").(string),
		Skipcaname:  d.Get("skipcaname").(bool),
		Snicert:     d.Get("snicert").(bool),
		Vservername: d.Get("vservername").(string),
	}

	return resource
}

func set_sslvserver_sslcertkey_binding(d *schema.ResourceData, resource *nitro.SslvserverSslcertkeyBinding) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("ca", resource.Ca)
	d.Set("certkeyname", resource.Certkeyname)
	d.Set("crlcheck", resource.Crlcheck)
	d.Set("ocspcheck", resource.Ocspcheck)
	d.Set("skipcaname", resource.Skipcaname)
	d.Set("snicert", resource.Snicert)
	d.Set("vservername", resource.Vservername)

	var key []string

	key = append(key, resource.Vservername)
	key = append(key, resource.Certkeyname)
	key = append(key, strconv.FormatBool(resource.Ca))
	d.SetId(strings.Join(key, "-"))
}

func get_sslvserver_sslcertkey_binding_key(d *schema.ResourceData) nitro.SslvserverSslcertkeyBindingKey {

	key := nitro.SslvserverSslcertkeyBindingKey{
		d.Get("vservername").(string),
		d.Get("certkeyname").(string),
		d.Get("ca").(bool),
	}
	return key
}

func create_sslvserver_sslcertkey_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_sslvserver_sslcertkey_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_sslvserver_sslcertkey_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsSslvserverSslcertkeyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSslvserverSslcertkeyBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_sslvserver_sslcertkey_binding(d, resource)
	} else {
		err := client.AddSslvserverSslcertkeyBinding(get_sslvserver_sslcertkey_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetSslvserverSslcertkeyBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_sslvserver_sslcertkey_binding(d, resource)
	}

	return nil
}

func read_sslvserver_sslcertkey_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_sslvserver_sslcertkey_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_sslvserver_sslcertkey_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsSslvserverSslcertkeyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSslvserverSslcertkeyBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_sslvserver_sslcertkey_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_sslvserver_sslcertkey_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_sslvserver_sslcertkey_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_sslvserver_sslcertkey_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsSslvserverSslcertkeyBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteSslvserverSslcertkeyBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
