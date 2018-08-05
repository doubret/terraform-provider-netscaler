package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
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

func key_rewriteaction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_rewriteaction(d *schema.ResourceData) nitro.Rewriteaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Rewriteaction{
		Name:              d.Get("name").(string),
		Bypasssafetycheck: d.Get("bypasssafetycheck").(string),
		Comment:           d.Get("comment").(string),
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

	d.Set("name", resource.Name)
	d.Set("bypasssafetycheck", resource.Bypasssafetycheck)
	d.Set("comment", resource.Comment)
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

func create_rewriteaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_rewriteaction")

	client := meta.(*nitro.NitroClient)

	key := key_rewriteaction(d)

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

	key := key_rewriteaction(d)

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

	return nil
}

func delete_rewriteaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_rewriteaction")

	client := meta.(*nitro.NitroClient)

	key := key_rewriteaction(d)

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
