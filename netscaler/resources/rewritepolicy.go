package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerRewritepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_rewritepolicy,
		Read:          read_rewritepolicy,
		Update:        update_rewritepolicy,
		Delete:        delete_rewritepolicy,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
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
			"logaction": &schema.Schema{
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
			"undefaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_rewritepolicy(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_rewritepolicy(d *schema.ResourceData) nitro.Rewritepolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Rewritepolicy{
		Name:        d.Get("name").(string),
		Action:      d.Get("action").(string),
		Comment:     d.Get("comment").(string),
		Logaction:   d.Get("logaction").(string),
		Rule:        d.Get("rule").(string),
		Undefaction: d.Get("undefaction").(string),
	}

	return resource
}

func set_rewritepolicy(d *schema.ResourceData, resource *nitro.Rewritepolicy) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("action", resource.Action)
	d.Set("comment", resource.Comment)
	d.Set("logaction", resource.Logaction)
	d.Set("rule", resource.Rule)
	d.Set("undefaction", resource.Undefaction)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_rewritepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_rewritepolicy")

	client := meta.(*nitro.NitroClient)

	key := key_rewritepolicy(d)

	exists, err := client.ExistsRewritepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetRewritepolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_rewritepolicy(d, resource)
	} else {
		err := client.AddRewritepolicy(get_rewritepolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetRewritepolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_rewritepolicy(d, resource)
	}

	return nil
}

func read_rewritepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_rewritepolicy")

	client := meta.(*nitro.NitroClient)

	key := key_rewritepolicy(d)

	exists, err := client.ExistsRewritepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetRewritepolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_rewritepolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_rewritepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_rewritepolicy")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateRewritepolicy(get_rewritepolicy(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_rewritepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_rewritepolicy")

	client := meta.(*nitro.NitroClient)

	key := key_rewritepolicy(d)

	exists, err := client.ExistsRewritepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteRewritepolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
