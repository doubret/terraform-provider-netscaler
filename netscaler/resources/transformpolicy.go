package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerTransformpolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_transformpolicy,
		Read:          read_transformpolicy,
		Update:        update_transformpolicy,
		Delete:        delete_transformpolicy,
		Schema: map[string]*schema.Schema{
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

func get_transformpolicy(d *schema.ResourceData) nitro.Transformpolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Transformpolicy{
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Name:        d.Get("name").(string),
		Profilename: d.Get("profilename").(string),
		Rule:        d.Get("rule").(string),
	}

	return resource
}

func set_transformpolicy(d *schema.ResourceData, resource *nitro.Transformpolicy) {
	var _ = strconv.Itoa

	d.Set("comment", resource.Comment)
	d.Set("logaction", resource.Logaction)
	d.Set("name", resource.Name)
	d.Set("profilename", resource.Profilename)
	d.Set("rule", resource.Rule)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_transformpolicy_key(d *schema.ResourceData) nitro.TransformpolicyKey {

	key := nitro.TransformpolicyKey{
		d.Get("name").(string),
	}
	return key
}

func create_transformpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_transformpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_transformpolicy(d)
	key := resource.ToKey()

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

	resource := get_transformpolicy(d)
	key := resource.ToKey()

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

	client := meta.(*nitro.NitroClient)

	update := nitro.TransformpolicyUpdate{}
	unset := nitro.TransformpolicyUnset{}

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
	if d.HasChange("profilename") {
		updateFlag = true

		value := d.Get("profilename").(string)
		update.Profilename = value

		if value == "" {
			unsetFlag = true

			unset.Profilename = true
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
	key := get_transformpolicy_key(d)

	if updateFlag {
		if err := client.UpdateTransformpolicy(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetTransformpolicy(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetTransformpolicy(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_transformpolicy(d, resource)
	}

	return nil
}

func delete_transformpolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_transformpolicy")

	client := meta.(*nitro.NitroClient)

	resource := get_transformpolicy(d)
	key := resource.ToKey()

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
