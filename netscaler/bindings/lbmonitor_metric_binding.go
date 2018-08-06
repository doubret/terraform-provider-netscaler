package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerLbmonitorMetricBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_lbmonitor_metric_binding,
		Read:          read_lbmonitor_metric_binding,
		Update:        nil,
		Delete:        delete_lbmonitor_metric_binding,
		Schema: map[string]*schema.Schema{
			"metric": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"metricthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"metricweight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"monitorname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func key_lbmonitor_metric_binding(d *schema.ResourceData) nitro.LbmonitorMetricBindingKey {
	key := nitro.LbmonitorMetricBindingKey{
		Monitorname: d.Get("monitorname").(string),
		Metric:      d.Get("metric").(string),
	}

	return key
}

func get_lbmonitor_metric_binding(d *schema.ResourceData) nitro.LbmonitorMetricBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.LbmonitorMetricBinding{
		Metric:          d.Get("metric").(string),
		Metricthreshold: d.Get("metricthreshold").(int),
		Metricweight:    d.Get("metricweight").(int),
		Monitorname:     d.Get("monitorname").(string),
	}

	return resource
}

func set_lbmonitor_metric_binding(d *schema.ResourceData, resource *nitro.LbmonitorMetricBinding) {
	var _ = strconv.Itoa

	d.Set("metric", resource.Metric)
	d.Set("metricthreshold", resource.Metricthreshold)
	d.Set("metricweight", resource.Metricweight)
	d.Set("monitorname", resource.Monitorname)
	var key []string

	key = append(key, resource.Monitorname)
	key = append(key, resource.Metric)
	d.SetId(strings.Join(key, "-"))
}

func create_lbmonitor_metric_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_lbmonitor_metric_binding")

	client := meta.(*nitro.NitroClient)

	key := key_lbmonitor_metric_binding(d)

	exists, err := client.ExistsLbmonitorMetricBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbmonitorMetricBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_lbmonitor_metric_binding(d, resource)
	} else {
		err := client.AddLbmonitorMetricBinding(get_lbmonitor_metric_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetLbmonitorMetricBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_lbmonitor_metric_binding(d, resource)
	}

	return nil
}

func read_lbmonitor_metric_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_lbmonitor_metric_binding")

	client := meta.(*nitro.NitroClient)

	key := key_lbmonitor_metric_binding(d)

	exists, err := client.ExistsLbmonitorMetricBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetLbmonitorMetricBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_lbmonitor_metric_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_lbmonitor_metric_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_lbmonitor_metric_binding")

	client := meta.(*nitro.NitroClient)

	key := key_lbmonitor_metric_binding(d)

	exists, err := client.ExistsLbmonitorMetricBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteLbmonitorMetricBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
