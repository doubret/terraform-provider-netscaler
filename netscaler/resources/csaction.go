package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerCsaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_csaction,
		Read:          read_csaction,
		Update:        update_csaction,
		Delete:        delete_csaction,
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
			"targetlbvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"targetvserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"targetvserverexpr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_csaction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_csaction(d *schema.ResourceData) nitro.Csaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Csaction{
		Name:              d.Get("name").(string),
		Comment:           d.Get("comment").(string),
		Targetlbvserver:   d.Get("targetlbvserver").(string),
		Targetvserver:     d.Get("targetvserver").(string),
		Targetvserverexpr: d.Get("targetvserverexpr").(string),
	}

	return resource
}

func set_csaction(d *schema.ResourceData, resource *nitro.Csaction) {
	d.Set("name", resource.Name)
	d.Set("comment", resource.Comment)
	d.Set("targetlbvserver", resource.Targetlbvserver)
	d.Set("targetvserver", resource.Targetvserver)
	d.Set("targetvserverexpr", resource.Targetvserverexpr)
	d.SetId(resource.Name)
}

func create_csaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_csaction")

	client := meta.(*nitro.NitroClient)

	key := key_csaction(d)

	exists, err := client.ExistsCsaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCsaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_csaction(d, resource)
	} else {
		err := client.AddCsaction(get_csaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCsaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_csaction(d, resource)
	}

	return nil
}

func read_csaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_csaction")

	client := meta.(*nitro.NitroClient)

	key := key_csaction(d)

	exists, err := client.ExistsCsaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCsaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_csaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_csaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_csaction")

	return nil
}

func delete_csaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_csaction")

	client := meta.(*nitro.NitroClient)

	key := key_csaction(d)

	exists, err := client.ExistsCsaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCsaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
