package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAuditsyslogpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_auditsyslogpolicy,
		Read:          read_auditsyslogpolicy,
		Update:        update_auditsyslogpolicy,
		Delete:        delete_auditsyslogpolicy,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
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
			"rule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_auditsyslogpolicy(d *schema.ResourceData) nitro.Auditsyslogpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Auditsyslogpolicy{
		Action: d.Get("action").(string),
		Name:   d.Get("name").(string),
		Rule:   d.Get("rule").(string),
	}

	return resource
}

func set_auditsyslogpolicy(d *schema.ResourceData, resource *nitro.Auditsyslogpolicy) {
	var _ = strconv.Itoa

	d.Set("action", resource.Action)
	d.Set("name", resource.Name)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_auditsyslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_auditsyslogpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_auditsyslogpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsAuditsyslogpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuditsyslogpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_auditsyslogpolicy(d, resource)
	} else {
		err := client.AddAuditsyslogpolicy(get_auditsyslogpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAuditsyslogpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_auditsyslogpolicy(d, resource)
	}

	return nil
}

func read_auditsyslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_auditsyslogpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_auditsyslogpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsAuditsyslogpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAuditsyslogpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_auditsyslogpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_auditsyslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_auditsyslogpolicy")

	// TODO
	// client := meta.(*nitro.NitroClient)

	// err := client.UpdateAuditsyslogpolicy(get_auditsyslogpolicy(d))

	// if err != nil {
	//       return err
	// }

	return nil
}

func delete_auditsyslogpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_auditsyslogpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_auditsyslogpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsAuditsyslogpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAuditsyslogpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
