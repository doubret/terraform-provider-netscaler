package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerPolicydataset() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_policydataset,
		Read:          read_policydataset,
		Update:        nil,
		Delete:        delete_policydataset,
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
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func get_policydataset(d *schema.ResourceData) nitro.Policydataset {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Policydataset{
		Comment:   d.Get("comment").(string),
		Indextype: d.Get("indextype").(string),
		Name:      d.Get("name").(string),
		Type:      d.Get("type").(string),
	}

	return resource
}

func set_policydataset(d *schema.ResourceData, resource *nitro.Policydataset) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("comment", resource.Comment)
	d.Set("indextype", resource.Indextype)
	d.Set("name", resource.Name)
	d.Set("type", resource.Type)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_policydataset_key(d *schema.ResourceData) nitro.PolicydatasetKey {

	key := nitro.PolicydatasetKey{
		d.Get("name").(string),
	}
	return key
}

func create_policydataset(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_policydataset")

	client := meta.(*nitro.NitroClient)

	resource := get_policydataset(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicydataset(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicydataset(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_policydataset(d, resource)
	} else {
		err := client.AddPolicydataset(get_policydataset(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetPolicydataset(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_policydataset(d, resource)
	}

	return nil
}

func read_policydataset(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_policydataset")

	client := meta.(*nitro.NitroClient)

	resource := get_policydataset(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicydataset(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPolicydataset(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_policydataset(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func delete_policydataset(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_policydataset")

	client := meta.(*nitro.NitroClient)

	resource := get_policydataset(d)
	key := resource.ToKey()

	exists, err := client.ExistsPolicydataset(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeletePolicydataset(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
