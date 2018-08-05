package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerLbmetrictableMetricBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbmetrictable_metric_binding,
		Read:          read_lbmetrictable_metric_binding,
		Update:        nil,
		Delete:        delete_lbmetrictable_metric_binding,
		Schema: map[string]*schema.Schema{
			"metric": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"metrictable": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"snmpoid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func key_lbmetrictable_metric_binding(d *schema.ResourceData) nitro.LbmetrictableMetricBindingKey {
	key := nitro.LbmetrictableMetricBindingKey{
		Metrictable: d.Get("metrictable").(string),
		Metric:      d.Get("metric").(string),
	}

	return key
}

func get_lbmetrictable_metric_binding(d *schema.ResourceData) nitro.LbmetrictableMetricBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.LbmetrictableMetricBinding{
		Metric:      d.Get("metric").(string),
		Metrictable: d.Get("metrictable").(string),
		Snmpoid:     d.Get("snmpoid").(string),
	}

	return resource
}

func set_lbmetrictable_metric_binding(d *schema.ResourceData, resource *nitro.LbmetrictableMetricBinding) {
	var _ = strconv.Itoa

	d.Set("metric", resource.Metric)
	d.Set("metrictable", resource.Metrictable)
	d.Set("snmpoid", resource.Snmpoid)
	var key []string

	key = append(key, resource.Metrictable)
	key = append(key, resource.Metric)
	d.SetId(strings.Join(key, "-"))
}

func create_lbmetrictable_metric_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbmetrictable_metric_binding")

	client := meta.(*nitro.NitroClient)

	key := key_lbmetrictable_metric_binding(d)

	exists, err := client.ExistsLbmetrictableMetricBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbmetrictableMetricBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbmetrictable_metric_binding(d, resource)
	} else {
		err := client.AddLbmetrictableMetricBinding(get_lbmetrictable_metric_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbmetrictableMetricBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbmetrictable_metric_binding(d, resource)
	}

	return nil
}

func read_lbmetrictable_metric_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbmetrictable_metric_binding")

	client := meta.(*nitro.NitroClient)

	key := key_lbmetrictable_metric_binding(d)

	exists, err := client.ExistsLbmetrictableMetricBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbmetrictableMetricBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbmetrictable_metric_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_lbmetrictable_metric_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbmetrictable_metric_binding")

	client := meta.(*nitro.NitroClient)

	key := key_lbmetrictable_metric_binding(d)

	exists, err := client.ExistsLbmetrictableMetricBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbmetrictableMetricBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
