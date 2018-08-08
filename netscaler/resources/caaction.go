package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCaaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_caaction,
		Read:          read_caaction,
		Update:        update_caaction,
		Delete:        delete_caaction,
		Schema: map[string]*schema.Schema{
			"accumressize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"lbvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
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
				ForceNew: false,
			},
		},
	}
}

func get_caaction(d *schema.ResourceData) nitro.Caaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Caaction{
		Accumressize: d.Get("accumressize").(int),
		Comment:      d.Get("comment").(string),
		Lbvserver:    d.Get("lbvserver").(string),
		Name:         d.Get("name").(string),
		Type:         d.Get("type").(string),
	}

	return resource
}

func set_caaction(d *schema.ResourceData, resource *nitro.Caaction) {
	var _ = strconv.Itoa

	d.Set("accumressize", resource.Accumressize)
	d.Set("comment", resource.Comment)
	d.Set("lbvserver", resource.Lbvserver)
	d.Set("name", resource.Name)
	d.Set("type", resource.Type)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_caaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_caaction")

	client := meta.(*nitro.NitroClient)

	resource := get_caaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsCaaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCaaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_caaction(d, resource)
	} else {
		err := client.AddCaaction(get_caaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCaaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_caaction(d, resource)
	}

	return nil
}

func read_caaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_caaction")

	client := meta.(*nitro.NitroClient)

	resource := get_caaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsCaaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCaaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_caaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_caaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_caaction")

	// TODO
	// client := meta.(*nitro.NitroClient)

	// err := client.UpdateCaaction(get_caaction(d))

	// if err != nil {
	//       return err
	// }

	return nil
}

func delete_caaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_caaction")

	client := meta.(*nitro.NitroClient)

	resource := get_caaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsCaaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCaaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
