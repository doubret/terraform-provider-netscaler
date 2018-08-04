package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerAppflowpolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_appflowpolicylabel,
		Read:          read_appflowpolicylabel,
		Update:        update_appflowpolicylabel,
		Delete:        delete_appflowpolicylabel,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"policylabeltype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func key_appflowpolicylabel(d *schema.ResourceData) string {
	return d.Get("labelname").(string)
}

func get_appflowpolicylabel(d *schema.ResourceData) nitro.Appflowpolicylabel {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appflowpolicylabel{
		Labelname:       d.Get("labelname").(string),
		Policylabeltype: d.Get("policylabeltype").(string),
	}

	return resource
}

func set_appflowpolicylabel(d *schema.ResourceData, resource *nitro.Appflowpolicylabel) {
	d.Set("labelname", resource.Labelname)
	d.Set("policylabeltype", resource.Policylabeltype)
	d.SetId(resource.Labelname)
}

func create_appflowpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appflowpolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_appflowpolicylabel(d)

	exists, err := client.ExistsAppflowpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowpolicylabel(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_appflowpolicylabel(d, resource)
	} else {
		err := client.AddAppflowpolicylabel(get_appflowpolicylabel(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAppflowpolicylabel(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_appflowpolicylabel(d, resource)
	}

	return nil
}

func read_appflowpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_appflowpolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_appflowpolicylabel(d)

	exists, err := client.ExistsAppflowpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowpolicylabel(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_appflowpolicylabel(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_appflowpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_appflowpolicylabel")

	return nil
}

func delete_appflowpolicylabel(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appflowpolicylabel")

	client := meta.(*nitro.NitroClient)

	key := key_appflowpolicylabel(d)

	exists, err := client.ExistsAppflowpolicylabel(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAppflowpolicylabel(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
