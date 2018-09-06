package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerVideooptimizationaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_videooptimizationaction,
		Read:          read_videooptimizationaction,
		Update:        update_videooptimizationaction,
		Delete:        delete_videooptimizationaction,
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
			"rate": &schema.Schema{
				Type:     schema.TypeInt,
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

func get_videooptimizationaction(d *schema.ResourceData) nitro.Videooptimizationaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Videooptimizationaction{
		Comment: d.Get("comment").(string),
		Name:    d.Get("name").(string),
		Rate:    d.Get("rate").(int),
		Type:    d.Get("type").(string),
	}

	return resource
}

func set_videooptimizationaction(d *schema.ResourceData, resource *nitro.Videooptimizationaction) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("comment", resource.Comment)
	d.Set("name", resource.Name)
	d.Set("rate", resource.Rate)
	d.Set("type", resource.Type)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_videooptimizationaction_key(d *schema.ResourceData) nitro.VideooptimizationactionKey {

	key := nitro.VideooptimizationactionKey{
		d.Get("name").(string),
	}
	return key
}

func create_videooptimizationaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_videooptimizationaction")

	client := meta.(*nitro.NitroClient)

	resource := get_videooptimizationaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsVideooptimizationaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetVideooptimizationaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_videooptimizationaction(d, resource)
	} else {
		err := client.AddVideooptimizationaction(get_videooptimizationaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetVideooptimizationaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_videooptimizationaction(d, resource)
	}

	return nil
}

func read_videooptimizationaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_videooptimizationaction")

	client := meta.(*nitro.NitroClient)

	resource := get_videooptimizationaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsVideooptimizationaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetVideooptimizationaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_videooptimizationaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_videooptimizationaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_videooptimizationaction")

	client := meta.(*nitro.NitroClient)

	update := nitro.VideooptimizationactionUpdate{}
	unset := nitro.VideooptimizationactionUnset{}

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
	if d.HasChange("rate") {
		updateFlag = true

		value := d.Get("rate").(int)
		update.Rate = value

		if value == 0 {
			unsetFlag = true

			unset.Rate = true
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
	key := get_videooptimizationaction_key(d)

	if updateFlag {
		if err := client.UpdateVideooptimizationaction(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetVideooptimizationaction(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetVideooptimizationaction(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_videooptimizationaction(d, resource)
	}

	return nil
}

func delete_videooptimizationaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_videooptimizationaction")

	client := meta.(*nitro.NitroClient)

	resource := get_videooptimizationaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsVideooptimizationaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteVideooptimizationaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
