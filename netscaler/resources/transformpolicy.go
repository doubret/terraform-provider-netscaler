package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerTransformpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_transformpolicy,
		Read:          read_transformpolicy,
		Update:        update_transformpolicy,
		Delete:        delete_transformpolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"logaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"profilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_transformpolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_transformpolicy(d *schema.ResourceData) nitro.Transformpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Transformpolicy{
		Name:        d.Get("name").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Profilename: d.Get("profilename").(string),
		Rule:        d.Get("rule").(string),
	}

	return resource
}

func set_transformpolicy(d *schema.ResourceData, resource *nitro.Transformpolicy) {
	d.Set("name", resource.Name)
	d.Set("comment", resource.Comment)
	d.Set("logaction", resource.Logaction)
	d.Set("profilename", resource.Profilename)
	d.Set("rule", resource.Rule)
	d.SetId(resource.Name)
}

func create_transformpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_transformpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_transformpolicy(d)

	exists, err := client.ExistsTransformpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTransformpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_transformpolicy(d, resource)
	} else {
		err := client.AddTransformpolicy(get_transformpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetTransformpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_transformpolicy(d, resource)
	}

	return nil
}

func read_transformpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_transformpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_transformpolicy(d)

	exists, err := client.ExistsTransformpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTransformpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_transformpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_transformpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_transformpolicy")

	return nil
}

func delete_transformpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_transformpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_transformpolicy(d)

	exists, err := client.ExistsTransformpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteTransformpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
