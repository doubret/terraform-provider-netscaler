package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerPolicystringmapPatternBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policystringmap_pattern_binding,
		Read:          read_policystringmap_pattern_binding,
		Update:        nil,
		Delete:        delete_policystringmap_pattern_binding,
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_policystringmap_pattern_binding(d *schema.ResourceData) nitro.PolicystringmapPatternBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.PolicystringmapPatternBinding{
		Key:   d.Get("key").(string),
		Name:  d.Get("name").(string),
		Value: d.Get("value").(string),
	}

	return resource
}

func set_policystringmap_pattern_binding(d *schema.ResourceData, resource *nitro.PolicystringmapPatternBinding) {
	var _ = strconv.Itoa

	d.Set("key", resource.Key)
	d.Set("name", resource.Name)
	d.Set("value", resource.Value)

	var key []string

	key = append(key, resource.Name)
	key = append(key, resource.Key)
	d.SetId(strings.Join(key, "-"))
}

func get_policystringmap_pattern_binding_key(d *schema.ResourceData) nitro.PolicystringmapPatternBindingKey {

	key := nitro.PolicystringmapPatternBindingKey{
		d.Get("name").(string),
		d.Get("key").(string),
	}
	return key
}

func create_policystringmap_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policystringmap_pattern_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_policystringmap_pattern_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicystringmapPatternBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicystringmapPatternBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_policystringmap_pattern_binding(d, resource)
	} else {
		err := client.AddPolicystringmapPatternBinding(get_policystringmap_pattern_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetPolicystringmapPatternBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_policystringmap_pattern_binding(d, resource)
	}

	return nil
}

func read_policystringmap_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policystringmap_pattern_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_policystringmap_pattern_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicystringmapPatternBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicystringmapPatternBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_policystringmap_pattern_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_policystringmap_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policystringmap_pattern_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_policystringmap_pattern_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicystringmapPatternBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeletePolicystringmapPatternBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
