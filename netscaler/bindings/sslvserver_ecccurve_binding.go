package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerSslvserverEcccurveBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_sslvserver_ecccurve_binding,
		Read:          read_sslvserver_ecccurve_binding,
		Update:        nil,
		Delete:        delete_sslvserver_ecccurve_binding,
		Schema: map[string]*schema.Schema{
			"ecccurvename": &schema.Schema{
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

func get_sslvserver_ecccurve_binding(d *schema.ResourceData) nitro.SslvserverEcccurveBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.SslvserverEcccurveBinding{
		Ecccurvename: d.Get("ecccurvename").(string),
		Vservername:  d.Get("vservername").(string),
	}

	return resource
}

func set_sslvserver_ecccurve_binding(d *schema.ResourceData, resource *nitro.SslvserverEcccurveBinding) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("ecccurvename", resource.Ecccurvename)
	d.Set("vservername", resource.Vservername)

	var key []string

	key = append(key, resource.Vservername)
	key = append(key, resource.Ecccurvename)
	d.SetId(strings.Join(key, "-"))
}

func get_sslvserver_ecccurve_binding_key(d *schema.ResourceData) nitro.SslvserverEcccurveBindingKey {

	key := nitro.SslvserverEcccurveBindingKey{
		d.Get("vservername").(string),
		d.Get("ecccurvename").(string),
	}
	return key
}

func create_sslvserver_ecccurve_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_sslvserver_ecccurve_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_sslvserver_ecccurve_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsSslvserverEcccurveBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSslvserverEcccurveBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_sslvserver_ecccurve_binding(d, resource)
	} else {
		err := client.AddSslvserverEcccurveBinding(get_sslvserver_ecccurve_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetSslvserverEcccurveBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_sslvserver_ecccurve_binding(d, resource)
	}

	return nil
}

func read_sslvserver_ecccurve_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_sslvserver_ecccurve_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_sslvserver_ecccurve_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsSslvserverEcccurveBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetSslvserverEcccurveBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_sslvserver_ecccurve_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_sslvserver_ecccurve_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_sslvserver_ecccurve_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_sslvserver_ecccurve_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsSslvserverEcccurveBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteSslvserverEcccurveBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
