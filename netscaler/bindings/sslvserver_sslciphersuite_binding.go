package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerSslvserverSslciphersuiteBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_sslvserver_sslciphersuite_binding,
		Read:          read_sslvserver_sslciphersuite_binding,
		Update:        nil,
		Delete:        delete_sslvserver_sslciphersuite_binding,
		Schema: map[string]*schema.Schema{
			"ciphername": &schema.Schema{
				Type:     schema.TypeString,
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

func get_sslvserver_sslciphersuite_binding(d *schema.ResourceData) nitro.SslvserverSslciphersuiteBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.SslvserverSslciphersuiteBinding{
		Ciphername:  d.Get("ciphername").(string),
		Vservername: d.Get("vservername").(string),
	}

	return resource
}

func set_sslvserver_sslciphersuite_binding(d *schema.ResourceData, resource *nitro.SslvserverSslciphersuiteBinding) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("ciphername", resource.Ciphername)
	d.Set("vservername", resource.Vservername)

	var key []string

	key = append(key, resource.Vservername)
	key = append(key, resource.Ciphername)
	d.SetId(strings.Join(key, "-"))
}

func get_sslvserver_sslciphersuite_binding_key(d *schema.ResourceData) nitro.SslvserverSslciphersuiteBindingKey {

	key := nitro.SslvserverSslciphersuiteBindingKey{
		d.Get("vservername").(string),
		d.Get("ciphername").(string),
	}
	return key
}

func create_sslvserver_sslciphersuite_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_sslvserver_sslciphersuite_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_sslvserver_sslciphersuite_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsSslvserverSslciphersuiteBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSslvserverSslciphersuiteBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_sslvserver_sslciphersuite_binding(d, resource)
	} else {
		err := client.AddSslvserverSslciphersuiteBinding(get_sslvserver_sslciphersuite_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetSslvserverSslciphersuiteBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_sslvserver_sslciphersuite_binding(d, resource)
	}

	return nil
}

func read_sslvserver_sslciphersuite_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_sslvserver_sslciphersuite_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_sslvserver_sslciphersuite_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsSslvserverSslciphersuiteBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSslvserverSslciphersuiteBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_sslvserver_sslciphersuite_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_sslvserver_sslciphersuite_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_sslvserver_sslciphersuite_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_sslvserver_sslciphersuite_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsSslvserverSslciphersuiteBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteSslvserverSslciphersuiteBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
