package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAppflowpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_appflowpolicy,
		Read:          read_appflowpolicy,
		Update:        update_appflowpolicy,
		Delete:        delete_appflowpolicy,
		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
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

func get_appflowpolicy(d *schema.ResourceData) nitro.Appflowpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appflowpolicy{
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Name:        d.Get("name").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	return resource
}

func set_appflowpolicy(d *schema.ResourceData, resource *nitro.Appflowpolicy) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("action", resource.Action)
	d.Set("comment", resource.Comment)
	d.Set("name", resource.Name)
	d.Set("rule", resource.Rule)
	d.Set("undefaction", resource.Undefaction)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_appflowpolicy_key(d *schema.ResourceData) nitro.AppflowpolicyKey {

	key := nitro.AppflowpolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_appflowpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appflowpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowpolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_appflowpolicy(d, resource)
	} else {
		err := client.AddAppflowpolicy(get_appflowpolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAppflowpolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_appflowpolicy(d, resource)
	}

	return nil
}

func read_appflowpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_appflowpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowpolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_appflowpolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_appflowpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_appflowpolicy")

	client := meta.(*nitro.NitroClient)

	update := nitro.AppflowpolicyUpdate{}
	unset := nitro.AppflowpolicyUnset{}

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
	key := get_appflowpolicy_key(d)

	if updateFlag {
		if err := client.UpdateAppflowpolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetAppflowpolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetAppflowpolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_appflowpolicy(d, resource)
	}

	return nil
}

func delete_appflowpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appflowpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowpolicy(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowpolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAppflowpolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
