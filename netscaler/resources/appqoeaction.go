package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerAppqoeaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_appqoeaction,
		Read:          read_appqoeaction,
		Update:        update_appqoeaction,
		Delete:        delete_appqoeaction,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"altcontentpath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"altcontentsvcname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"customfile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dosaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dostrigexpression": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxconn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"polqdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"priqdepth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"respondwith": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"tcpprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_appqoeaction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_appqoeaction(d *schema.ResourceData) nitro.Appqoeaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appqoeaction{
		Name:              d.Get("name").(string),
		Altcontentpath:    d.Get("altcontentpath").(string),
		Altcontentsvcname: d.Get("altcontentsvcname").(string),
		Customfile:        d.Get("customfile").(string),
		Delay:             d.Get("delay").(int),
		Dosaction:         d.Get("dosaction").(string),
		Dostrigexpression: d.Get("dostrigexpression").(string),
		Maxconn:           d.Get("maxconn").(int),
		Polqdepth:         d.Get("polqdepth").(int),
		Priority:          d.Get("priority").(string),
		Priqdepth:         d.Get("priqdepth").(int),
		Respondwith:       d.Get("respondwith").(string),
		Tcpprofile:        d.Get("tcpprofile").(string),
	}

	return resource
}

func set_appqoeaction(d *schema.ResourceData, resource *nitro.Appqoeaction) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("altcontentpath", resource.Altcontentpath)
	d.Set("altcontentsvcname", resource.Altcontentsvcname)
	d.Set("customfile", resource.Customfile)
	d.Set("delay", resource.Delay)
	d.Set("dosaction", resource.Dosaction)
	d.Set("dostrigexpression", resource.Dostrigexpression)
	d.Set("maxconn", resource.Maxconn)
	d.Set("polqdepth", resource.Polqdepth)
	d.Set("priority", resource.Priority)
	d.Set("priqdepth", resource.Priqdepth)
	d.Set("respondwith", resource.Respondwith)
	d.Set("tcpprofile", resource.Tcpprofile)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_appqoeaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appqoeaction")

	client := meta.(*nitro.NitroClient)

	key := key_appqoeaction(d)

	exists, err := client.ExistsAppqoeaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppqoeaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_appqoeaction(d, resource)
	} else {
		err := client.AddAppqoeaction(get_appqoeaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetAppqoeaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_appqoeaction(d, resource)
	}

	return nil
}

func read_appqoeaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_appqoeaction")

	client := meta.(*nitro.NitroClient)

	key := key_appqoeaction(d)

	exists, err := client.ExistsAppqoeaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetAppqoeaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_appqoeaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_appqoeaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_appqoeaction")

	client := meta.(*nitro.NitroClient)

	err := client.UpdateAppqoeaction(get_appqoeaction(d))

	if err != nil {
		return err
	}

	return nil
}

func delete_appqoeaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appqoeaction")

	client := meta.(*nitro.NitroClient)

	key := key_appqoeaction(d)

	exists, err := client.ExistsAppqoeaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteAppqoeaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
