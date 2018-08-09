package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAppflowaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_appflowaction,
		Read:          read_appflowaction,
		Update:        update_appflowaction,
		Delete:        delete_appflowaction,
		Schema: map[string]*schema.Schema{
			"clientsidemeasurements": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"collectors": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"pagetracking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"securityinsight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"webinsight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_appflowaction(d *schema.ResourceData) nitro.Appflowaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appflowaction{
		Clientsidemeasurements: d.Get("clientsidemeasurements").(string),
		Collectors:             utils.Convert_set_to_string_array(d.Get("collectors").(*schema.Set)),
		Comment:                d.Get("comment").(string),
		Name:                   d.Get("name").(string),
		Pagetracking:           d.Get("pagetracking").(string),
		Securityinsight:        d.Get("securityinsight").(string),
		Webinsight:             d.Get("webinsight").(string),
	}

	return resource
}

func set_appflowaction(d *schema.ResourceData, resource *nitro.Appflowaction) {
	var _ = strconv.Itoa

	d.Set("clientsidemeasurements", resource.Clientsidemeasurements)
	d.Set("collectors", resource.Collectors)
	d.Set("comment", resource.Comment)
	d.Set("name", resource.Name)
	d.Set("pagetracking", resource.Pagetracking)
	d.Set("securityinsight", resource.Securityinsight)
	d.Set("webinsight", resource.Webinsight)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_appflowaction_key(d *schema.ResourceData) nitro.AppflowactionKey {

	key := nitro.AppflowactionKey{
		d.Get("name").(string),
	}
	return key
}

func create_appflowaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appflowaction")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_appflowaction(d, resource)
	} else {
		err := client.AddAppflowaction(get_appflowaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAppflowaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_appflowaction(d, resource)
	}

	return nil
}

func read_appflowaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_appflowaction")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppflowaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_appflowaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_appflowaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_appflowaction")

	client := meta.(*nitro.NitroClient)

	update := nitro.AppflowactionUpdate{}
	unset := nitro.AppflowactionUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("collectors") {
		updateFlag = true

		value := utils.Convert_set_to_string_array(d.Get("collectors").(*schema.Set))
		update.Collectors = value

	}
	if d.HasChange("clientsidemeasurements") {
		updateFlag = true

		value := d.Get("clientsidemeasurements").(string)
		update.Clientsidemeasurements = value

		if value == "" {
			unsetFlag = true

			unset.Clientsidemeasurements = true
		}

	}
	if d.HasChange("pagetracking") {
		updateFlag = true

		value := d.Get("pagetracking").(string)
		update.Pagetracking = value

		if value == "" {
			unsetFlag = true

			unset.Pagetracking = true
		}

	}
	if d.HasChange("webinsight") {
		updateFlag = true

		value := d.Get("webinsight").(string)
		update.Webinsight = value

		if value == "" {
			unsetFlag = true

			unset.Webinsight = true
		}

	}
	if d.HasChange("securityinsight") {
		updateFlag = true

		value := d.Get("securityinsight").(string)
		update.Securityinsight = value

		if value == "" {
			unsetFlag = true

			unset.Securityinsight = true
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
	key := get_appflowaction_key(d)

	if updateFlag {
		if err := client.UpdateAppflowaction(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetAppflowaction(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetAppflowaction(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_appflowaction(d, resource)
	}

	return nil
}

func delete_appflowaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appflowaction")

	client := meta.(*nitro.NitroClient)

	resource := get_appflowaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsAppflowaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAppflowaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
