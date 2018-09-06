package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerTransformprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_transformprofile,
		Read:          read_transformprofile,
		Update:        update_transformprofile,
		Delete:        delete_transformprofile,
		Schema: map[string]*schema.Schema{
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
			"onlytransformabsurlinbody": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_transformprofile(d *schema.ResourceData) nitro.Transformprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Transformprofile{
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Onlytransformabsurlinbody: d.Get("onlytransformabsurlinbody").(string),
		Type: d.Get("type").(string),
	}

	return resource
}

func set_transformprofile(d *schema.ResourceData, resource *nitro.Transformprofile) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("comment", resource.Comment)
	d.Set("name", resource.Name)
	d.Set("onlytransformabsurlinbody", resource.Onlytransformabsurlinbody)
	d.Set("type", resource.Type)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_transformprofile_key(d *schema.ResourceData) nitro.TransformprofileKey {

	key := nitro.TransformprofileKey{
		d.Get("name").(string),
	}
	return key
}

func create_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_transformprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_transformprofile(d)
	key := resource.ToKey()

	exists, err := client.ExistsTransformprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTransformprofile(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_transformprofile(d, resource)
	} else {
		err := client.AddTransformprofile(get_transformprofile(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetTransformprofile(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_transformprofile(d, resource)
	}

	return nil
}

func read_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_transformprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_transformprofile(d)
	key := resource.ToKey()

	exists, err := client.ExistsTransformprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetTransformprofile(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_transformprofile(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_transformprofile")

	client := meta.(*nitro.NitroClient)

	update := nitro.TransformprofileUpdate{}
	unset := nitro.TransformprofileUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("type") {
		updateFlag = true

		value := d.Get("type").(string)
		update.Type = value

		if value == "" {
			unsetFlag = true

			unset.Type = true
		}

	}
	if d.HasChange("onlytransformabsurlinbody") {
		updateFlag = true

		value := d.Get("onlytransformabsurlinbody").(string)
		update.Onlytransformabsurlinbody = value

		if value == "" {
			unsetFlag = true

			unset.Onlytransformabsurlinbody = true
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
	key := get_transformprofile_key(d)

	if updateFlag {
		if err := client.UpdateTransformprofile(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetTransformprofile(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetTransformprofile(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_transformprofile(d, resource)
	}

	return nil
}

func delete_transformprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_transformprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_transformprofile(d)
	key := resource.ToKey()

	exists, err := client.ExistsTransformprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteTransformprofile(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
