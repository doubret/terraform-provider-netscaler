package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerResponderaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_responderaction,
		Read:          read_responderaction,
		Update:        update_responderaction,
		Delete:        delete_responderaction,
		Schema: map[string]*schema.Schema{
			"bypasssafetycheck": &schema.Schema{
				Type:     schema.TypeString,
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
			"htmlpage": &schema.Schema{
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
			"reasonphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"responsestatuscode": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"target": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
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

func get_responderaction(d *schema.ResourceData) nitro.Responderaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Responderaction{
		Bypasssafetycheck:  d.Get("bypasssafetycheck").(string),
		Comment:            d.Get("comment").(string),
		Htmlpage:           d.Get("htmlpage").(string),
		Name:               d.Get("name").(string),
		Reasonphrase:       d.Get("reasonphrase").(string),
		Responsestatuscode: d.Get("responsestatuscode").(int),
		Target:             d.Get("target").(string),
		Type:               d.Get("type").(string),
	}

	return resource
}

func set_responderaction(d *schema.ResourceData, resource *nitro.Responderaction) {
	var _ = strconv.Itoa

	d.Set("bypasssafetycheck", resource.Bypasssafetycheck)
	d.Set("comment", resource.Comment)
	d.Set("htmlpage", resource.Htmlpage)
	d.Set("name", resource.Name)
	d.Set("reasonphrase", resource.Reasonphrase)
	d.Set("responsestatuscode", resource.Responsestatuscode)
	d.Set("target", resource.Target)
	d.Set("type", resource.Type)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_responderaction_key(d *schema.ResourceData) nitro.ResponderactionKey {

	key := nitro.ResponderactionKey{
		d.Get("name").(string),
	}
	return key
}

func create_responderaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_responderaction")

	client := meta.(*nitro.NitroClient)

	resource := get_responderaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsResponderaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetResponderaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_responderaction(d, resource)
	} else {
		err := client.AddResponderaction(get_responderaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetResponderaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_responderaction(d, resource)
	}

	return nil
}

func read_responderaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_responderaction")

	client := meta.(*nitro.NitroClient)

	resource := get_responderaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsResponderaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetResponderaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_responderaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_responderaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_responderaction")

	client := meta.(*nitro.NitroClient)

	update := nitro.ResponderactionUpdate{}
	unset := nitro.ResponderactionUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("target") {
		updateFlag = true

		value := d.Get("target").(string)
		update.Target = value

		if value == "" {
			unsetFlag = true

			unset.Target = true
		}

	}
	if d.HasChange("bypasssafetycheck") {
		updateFlag = true

		value := d.Get("bypasssafetycheck").(string)
		update.Bypasssafetycheck = value

		if value == "" {
			unsetFlag = true

			unset.Bypasssafetycheck = true
		}

	}
	if d.HasChange("htmlpage") {
		updateFlag = true

		value := d.Get("htmlpage").(string)
		update.Htmlpage = value

		if value == "" {
			unsetFlag = true

			unset.Htmlpage = true
		}

	}
	if d.HasChange("responsestatuscode") {
		updateFlag = true

		value := d.Get("responsestatuscode").(int)
		update.Responsestatuscode = value

		if value == 0 {
			unsetFlag = true

			unset.Responsestatuscode = true
		}

	}
	if d.HasChange("reasonphrase") {
		updateFlag = true

		value := d.Get("reasonphrase").(string)
		update.Reasonphrase = value

		if value == "" {
			unsetFlag = true

			unset.Reasonphrase = true
		}

	}
	if d.HasChange("comment") {
		updateFlag = true

		value := d.Get("comment").(string)
		update.Comment = value

		if value == "" {
			unsetFlag = true

			unset.Comment = true
		}

	}
	key := get_responderaction_key(d)

	if updateFlag {
		if err := client.UpdateResponderaction(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetResponderaction(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetResponderaction(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_responderaction(d, resource)
	}

	return nil
}

func delete_responderaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_responderaction")

	client := meta.(*nitro.NitroClient)

	resource := get_responderaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsResponderaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteResponderaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
