package bindings

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerPolicypatsetPatternBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policypatset_pattern_binding,
		Read:          read_policypatset_pattern_binding,
		Update:        nil,
		Delete:        delete_policypatset_pattern_binding,
		Schema: map[string]*schema.Schema{
			"charset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"index": &schema.Schema{
				Type:     schema.TypeInt,
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
			"string": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_policypatset_pattern_binding(d *schema.ResourceData) nitro.PolicypatsetPatternBinding {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.PolicypatsetPatternBinding{
		Charset: d.Get("charset").(string),
		Index:   d.Get("index").(int),
		Name:    d.Get("name").(string),
		String:  d.Get("string").(string),
	}

	return resource
}

func set_policypatset_pattern_binding(d *schema.ResourceData, resource *nitro.PolicypatsetPatternBinding) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("charset", resource.Charset)
	d.Set("index", resource.Index)
	d.Set("name", resource.Name)
	d.Set("string", resource.String)

	var key []string

	key = append(key, resource.Name)
	key = append(key, resource.String)
	d.SetId(strings.Join(key, "-"))
}

func get_policypatset_pattern_binding_key(d *schema.ResourceData) nitro.PolicypatsetPatternBindingKey {

	key := nitro.PolicypatsetPatternBindingKey{
		d.Get("name").(string),
		d.Get("string").(string),
	}
	return key
}

func create_policypatset_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policypatset_pattern_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_policypatset_pattern_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicypatsetPatternBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicypatsetPatternBinding(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_policypatset_pattern_binding(d, resource)
	} else {
		err := client.AddPolicypatsetPatternBinding(get_policypatset_pattern_binding(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetPolicypatsetPatternBinding(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_policypatset_pattern_binding(d, resource)
	}

	return nil
}

func read_policypatset_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policypatset_pattern_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_policypatset_pattern_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicypatsetPatternBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicypatsetPatternBinding(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_policypatset_pattern_binding(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_policypatset_pattern_binding(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policypatset_pattern_binding")

	client := meta.(*nitro.NitroClient)

	resource := get_policypatset_pattern_binding(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicypatsetPatternBinding(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeletePolicypatsetPatternBinding(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
