package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerResponderpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_responderpolicy,
		Read:          read_responderpolicy,
		Update:        update_responderpolicy,
		Delete:        delete_responderpolicy,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"appflowaction": &schema.Schema{
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
			"logaction": &schema.Schema{
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
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_responderpolicy(d *schema.ResourceData) nitro.Responderpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Responderpolicy{
		Action:        d.Get("action").(string),
		Appflowaction: d.Get("appflowaction").(string),
		Comment:       d.Get("comment").(string),
		Logaction:     d.Get("logaction").(string),
		Name:          d.Get("name").(string),
		Rule:          d.Get("rule").(string),
		Undefaction:   d.Get("undefaction").(string),
	}

	return resource
}

func set_responderpolicy(d *schema.ResourceData, resource *nitro.Responderpolicy) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("action", resource.Action)
	d.Set("appflowaction", resource.Appflowaction)
	d.Set("comment", resource.Comment)
	d.Set("logaction", resource.Logaction)
	d.Set("name", resource.Name)
	d.Set("rule", resource.Rule)
	d.Set("undefaction", resource.Undefaction)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_responderpolicy_key(d *schema.ResourceData) nitro.ResponderpolicyKey {

	key := nitro.ResponderpolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_responderpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_responderpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_responderpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsResponderpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetResponderpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_responderpolicy(d, resource)
	} else {
		err := client.AddResponderpolicy(get_responderpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetResponderpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_responderpolicy(d, resource)
	}

	return nil
}

func read_responderpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_responderpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_responderpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsResponderpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetResponderpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_responderpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_responderpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_responderpolicy")

	client := meta.(*nitro.NitroClient)

	update := nitro.ResponderpolicyUpdate{}
	unset := nitro.ResponderpolicyUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("rule") {
		updateFlag = true

		value := d.Get("rule").(string)
		update.Rule = value

		if value == "" {
			unsetFlag = true

			unset.Rule = true
		}

	}
	if d.HasChange("action") {
		updateFlag = true

		value := d.Get("action").(string)
		update.Action = value

		if value == "" {
			unsetFlag = true

			unset.Action = true
		}

	}
	if d.HasChange("undefaction") {
		updateFlag = true

		value := d.Get("undefaction").(string)
		update.Undefaction = value

		if value == "" {
			unsetFlag = true

			unset.Undefaction = true
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
	if d.HasChange("logaction") {
		updateFlag = true

		value := d.Get("logaction").(string)
		update.Logaction = value

		if value == "" {
			unsetFlag = true

			unset.Logaction = true
		}

	}
	if d.HasChange("appflowaction") {
		updateFlag = true

		value := d.Get("appflowaction").(string)
		update.Appflowaction = value

		if value == "" {
			unsetFlag = true

			unset.Appflowaction = true
		}

	}
	key := get_responderpolicy_key(d)

	if updateFlag {
		if err := client.UpdateResponderpolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetResponderpolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetResponderpolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_responderpolicy(d, resource)
	}

	return nil
}

func delete_responderpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_responderpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_responderpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsResponderpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteResponderpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
