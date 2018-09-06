package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerRewriteaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_rewriteaction,
		Read:          read_rewriteaction,
		Update:        update_rewriteaction,
		Delete:        delete_rewriteaction,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"refinesearch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"search": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"stringbuilderexpr": &schema.Schema{
				Type:     schema.TypeString,
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

func get_rewriteaction(d *schema.ResourceData) nitro.Rewriteaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Rewriteaction{
		Bypasssafetycheck: d.Get("bypasssafetycheck").(string),
		Comment:           d.Get("comment").(string),
		Name:              d.Get("name").(string),
		Pattern:           d.Get("pattern").(string),
		Refinesearch:      d.Get("refinesearch").(string),
		Search:            d.Get("search").(string),
		Stringbuilderexpr: d.Get("stringbuilderexpr").(string),
		Target:            d.Get("target").(string),
		Type:              d.Get("type").(string),
	}

	return resource
}

func set_rewriteaction(d *schema.ResourceData, resource *nitro.Rewriteaction) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("bypasssafetycheck", resource.Bypasssafetycheck)
	d.Set("comment", resource.Comment)
	d.Set("name", resource.Name)
	d.Set("pattern", resource.Pattern)
	d.Set("refinesearch", resource.Refinesearch)
	d.Set("search", resource.Search)
	d.Set("stringbuilderexpr", resource.Stringbuilderexpr)
	d.Set("target", resource.Target)
	d.Set("type", resource.Type)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_rewriteaction_key(d *schema.ResourceData) nitro.RewriteactionKey {

	key := nitro.RewriteactionKey{
		d.Get("name").(string),
	}
	return key
}

func create_rewriteaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_rewriteaction")

	client := meta.(*nitro.NitroClient)

	resource := get_rewriteaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsRewriteaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetRewriteaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_rewriteaction(d, resource)
	} else {
		err := client.AddRewriteaction(get_rewriteaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetRewriteaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_rewriteaction(d, resource)
	}

	return nil
}

func read_rewriteaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_rewriteaction")

	client := meta.(*nitro.NitroClient)

	resource := get_rewriteaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsRewriteaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetRewriteaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_rewriteaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_rewriteaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_rewriteaction")

	client := meta.(*nitro.NitroClient)

	update := nitro.RewriteactionUpdate{}
	unset := nitro.RewriteactionUnset{}

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
	if d.HasChange("stringbuilderexpr") {
		updateFlag = true

		value := d.Get("stringbuilderexpr").(string)
		update.Stringbuilderexpr = value

		if value == "" {
			unsetFlag = true

			unset.Stringbuilderexpr = true
		}

	}
	if d.HasChange("pattern") {
		updateFlag = true

		value := d.Get("pattern").(string)
		update.Pattern = value

		if value == "" {
			unsetFlag = true

			unset.Pattern = true
		}

	}
	if d.HasChange("search") {
		updateFlag = true

		value := d.Get("search").(string)
		update.Search = value

		if value == "" {
			unsetFlag = true

			unset.Search = true
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
	if d.HasChange("refinesearch") {
		updateFlag = true

		value := d.Get("refinesearch").(string)
		update.Refinesearch = value

		if value == "" {
			unsetFlag = true

			unset.Refinesearch = true
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
	key := get_rewriteaction_key(d)

	if updateFlag {
		if err := client.UpdateRewriteaction(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetRewriteaction(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetRewriteaction(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_rewriteaction(d, resource)
	}

	return nil
}

func delete_rewriteaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_rewriteaction")

	client := meta.(*nitro.NitroClient)

	resource := get_rewriteaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsRewriteaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteRewriteaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
