package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerDospolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_dospolicy,
		Read:          read_dospolicy,
		Update:        update_dospolicy,
		Delete:        delete_dospolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cltdetectrate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"qdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_dospolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_dospolicy(d *schema.ResourceData) nitro.Dospolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Dospolicy{
		Name:          d.Get("name").(string),
		Cltdetectrate: d.Get("cltdetectrate").(int),
		Qdepth:        d.Get("qdepth").(int),
	}

	return resource
}

func set_dospolicy(d *schema.ResourceData, resource *nitro.Dospolicy) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("cltdetectrate", resource.Cltdetectrate)
	d.Set("qdepth", resource.Qdepth)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dospolicy")

	client := meta.(*nitro.NitroClient)

	key := key_dospolicy(d)

	exists, err := client.ExistsDospolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDospolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_dospolicy(d, resource)
	} else {
		err := client.AddDospolicy(get_dospolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetDospolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_dospolicy(d, resource)
	}

	return nil
}

func read_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_dospolicy")

	client := meta.(*nitro.NitroClient)

	key := key_dospolicy(d)

	exists, err := client.ExistsDospolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDospolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_dospolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_dospolicy")

	return nil
}

func delete_dospolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dospolicy")

	client := meta.(*nitro.NitroClient)

	key := key_dospolicy(d)

	exists, err := client.ExistsDospolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteDospolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
