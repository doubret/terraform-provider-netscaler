package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCmpaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_cmpaction,
		Read:          read_cmpaction,
		Update:        update_cmpaction,
		Delete:        delete_cmpaction,
		Schema: map[string]*schema.Schema{
			"addvaryheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cmptype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"deltatype": &schema.Schema{
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
			"varyheadervalue": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_cmpaction(d *schema.ResourceData) nitro.Cmpaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Cmpaction{
		Addvaryheader:   d.Get("addvaryheader").(string),
		Cmptype:         d.Get("cmptype").(string),
		Deltatype:       d.Get("deltatype").(string),
		Name:            d.Get("name").(string),
		Varyheadervalue: d.Get("varyheadervalue").(string),
	}

	return resource
}

func set_cmpaction(d *schema.ResourceData, resource *nitro.Cmpaction) {
	var _ = strconv.Itoa

	d.Set("addvaryheader", resource.Addvaryheader)
	d.Set("cmptype", resource.Cmptype)
	d.Set("deltatype", resource.Deltatype)
	d.Set("name", resource.Name)
	d.Set("varyheadervalue", resource.Varyheadervalue)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_cmpaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cmpaction")

	client := meta.(*nitro.NitroClient)

	resource := get_cmpaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsCmpaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCmpaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_cmpaction(d, resource)
	} else {
		err := client.AddCmpaction(get_cmpaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCmpaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_cmpaction(d, resource)
	}

	return nil
}

func read_cmpaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_cmpaction")

	client := meta.(*nitro.NitroClient)

	resource := get_cmpaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsCmpaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCmpaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_cmpaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_cmpaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_cmpaction")

	// TODO
	// client := meta.(*nitro.NitroClient)

	// err := client.UpdateCmpaction(get_cmpaction(d))

	// if err != nil {
	//       return err
	// }

	return nil
}

func delete_cmpaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cmpaction")

	client := meta.(*nitro.NitroClient)

	resource := get_cmpaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsCmpaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCmpaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
