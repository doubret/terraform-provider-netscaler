package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerCachepolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_cachepolicy,
		Read:          read_cachepolicy,
		Update:        update_cachepolicy,
		Delete:        delete_cachepolicy,
		Schema: map[string]*schema.Schema{
			"policyname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"invalgroups": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"invalobjects": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"storeingroup": &schema.Schema{
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

func key_cachepolicy(d *schema.ResourceData) string {
	return d.Get("policyname").(string)
}

func get_cachepolicy(d *schema.ResourceData) nitro.Cachepolicy {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Cachepolicy{
		Policyname:   d.Get("policyname").(string),
		Action:       d.Get("action").(string),
		Invalgroups:  utils.Convert_set_to_string_array(d.Get("invalgroups").(*schema.Set)),
		Invalobjects: utils.Convert_set_to_string_array(d.Get("invalobjects").(*schema.Set)),
		Rule:         d.Get("rule").(string),
		Storeingroup: d.Get("storeingroup").(string),
		Undefaction:  d.Get("undefaction").(string),
	}

	return resource
}

func set_cachepolicy(d *schema.ResourceData, resource *nitro.Cachepolicy) {
	d.Set("policyname", resource.Policyname)
	d.Set("action", resource.Action)
	d.Set("invalgroups", resource.Invalgroups)
	d.Set("invalobjects", resource.Invalobjects)
	d.Set("rule", resource.Rule)
	d.Set("storeingroup", resource.Storeingroup)
	d.Set("undefaction", resource.Undefaction)
	d.SetId(resource.Policyname)
}

func create_cachepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cachepolicy")

	client := meta.(*nitro.NitroClient)

	key := key_cachepolicy(d)

	exists, err := client.ExistsCachepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCachepolicy(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_cachepolicy(d, resource)
	} else {
		err := client.AddCachepolicy(get_cachepolicy(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCachepolicy(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_cachepolicy(d, resource)
	}

	return nil
}

func read_cachepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_cachepolicy")

	client := meta.(*nitro.NitroClient)

	key := key_cachepolicy(d)

	exists, err := client.ExistsCachepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCachepolicy(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_cachepolicy(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_cachepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_cachepolicy")

	return nil
}

func delete_cachepolicy(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cachepolicy")

	client := meta.(*nitro.NitroClient)

	key := key_cachepolicy(d)

	exists, err := client.ExistsCachepolicy(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCachepolicy(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
