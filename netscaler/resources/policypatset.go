package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerPolicypatset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policypatset,
		Read:          read_policypatset,
		Update:        nil,
		Delete:        delete_policypatset,
		Schema: map[string]*schema.Schema{
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"indextype": &schema.Schema{
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
		},
	}
}

func get_policypatset(d *schema.ResourceData) nitro.Policypatset {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Policypatset{
		Comment:   d.Get("comment").(string),
		Indextype: d.Get("indextype").(string),
		Name:      d.Get("name").(string),
	}

	return resource
}

func set_policypatset(d *schema.ResourceData, resource *nitro.Policypatset) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("comment", resource.Comment)
	d.Set("indextype", resource.Indextype)
	d.Set("name", resource.Name)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_policypatset_key(d *schema.ResourceData) nitro.PolicypatsetKey {

	key := nitro.PolicypatsetKey{
		d.Get("name").(string),
	}
	return key
}

func create_policypatset(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policypatset")

	client := meta.(*nitro.NitroClient)

	resource := get_policypatset(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicypatset(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicypatset(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_policypatset(d, resource)
	} else {
		err := client.AddPolicypatset(get_policypatset(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetPolicypatset(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_policypatset(d, resource)
	}

	return nil
}

func read_policypatset(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policypatset")

	client := meta.(*nitro.NitroClient)

	resource := get_policypatset(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicypatset(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicypatset(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_policypatset(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_policypatset(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policypatset")

	client := meta.(*nitro.NitroClient)

	resource := get_policypatset(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicypatset(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeletePolicypatset(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
