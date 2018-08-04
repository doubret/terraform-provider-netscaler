package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerPqpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_pqpolicy,
		Read:          read_pqpolicy,
		Update:        update_pqpolicy,
		Delete:        delete_pqpolicy,
		Schema: map[string]*schema.Schema{
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"polqdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"qdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_pqpolicy(d *schema.ResourceData) string {
	return d.Get("policyname").(string)
}

func get_pqpolicy(d *schema.ResourceData) nitro.Pqpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Pqpolicy{
		Policyname: d.Get("policyname").(string),
		Polqdepth:  d.Get("polqdepth").(int),
		Priority:   d.Get("priority").(int),
		Qdepth:     d.Get("qdepth").(int),
		Rule:       d.Get("rule").(string),
		Weight:     d.Get("weight").(int),
	}

	return resource
}

func set_pqpolicy(d *schema.ResourceData, resource *nitro.Pqpolicy) {
	d.Set("policyname", resource.Policyname)
	d.Set("polqdepth", resource.Polqdepth)
	d.Set("priority", resource.Priority)
	d.Set("qdepth", resource.Qdepth)
	d.Set("rule", resource.Rule)
	d.Set("weight", resource.Weight)
	d.SetId(resource.Policyname)
}

func create_pqpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_pqpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_pqpolicy(d)

	exists, err := client.ExistsPqpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPqpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_pqpolicy(d, resource)
	} else {
		err := client.AddPqpolicy(get_pqpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetPqpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_pqpolicy(d, resource)
	}

	return nil
}

func read_pqpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_pqpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_pqpolicy(d)

	exists, err := client.ExistsPqpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetPqpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_pqpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_pqpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_pqpolicy")

	return nil
}

func delete_pqpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_pqpolicy")

	client := meta.(*nitro.NitroClient)

	key := key_pqpolicy(d)

	exists, err := client.ExistsPqpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeletePqpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
