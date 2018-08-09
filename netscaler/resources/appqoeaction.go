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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func get_appqoeaction(d *schema.ResourceData) nitro.Appqoeaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Appqoeaction{
		Altcontentpath:    d.Get("altcontentpath").(string),
		Altcontentsvcname: d.Get("altcontentsvcname").(string),
		Customfile:        d.Get("customfile").(string),
		Delay:             d.Get("delay").(int),
		Dosaction:         d.Get("dosaction").(string),
		Dostrigexpression: d.Get("dostrigexpression").(string),
		Maxconn:           d.Get("maxconn").(int),
		Name:              d.Get("name").(string),
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

	d.Set("altcontentpath", resource.Altcontentpath)
	d.Set("altcontentsvcname", resource.Altcontentsvcname)
	d.Set("customfile", resource.Customfile)
	d.Set("delay", resource.Delay)
	d.Set("dosaction", resource.Dosaction)
	d.Set("dostrigexpression", resource.Dostrigexpression)
	d.Set("maxconn", resource.Maxconn)
	d.Set("name", resource.Name)
	d.Set("polqdepth", resource.Polqdepth)
	d.Set("priority", resource.Priority)
	d.Set("priqdepth", resource.Priqdepth)
	d.Set("respondwith", resource.Respondwith)
	d.Set("tcpprofile", resource.Tcpprofile)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_appqoeaction_key(d *schema.ResourceData) nitro.AppqoeactionKey {

	key := nitro.AppqoeactionKey{
		d.Get("name").(string),
	}
	return key
}

func create_appqoeaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_appqoeaction")

	client := meta.(*nitro.NitroClient)

	resource := get_appqoeaction(d)
	key := resource.ToKey()

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

	resource := get_appqoeaction(d)
	key := resource.ToKey()

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

	update := nitro.AppqoeactionUpdate{}
	unset := nitro.AppqoeactionUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("priority") {
		updateFlag = true

		value := d.Get("priority").(string)
		update.Priority = value

		if value == "" {
			unsetFlag = true

			unset.Priority = true
		}

	}
	if d.HasChange("altcontentsvcname") {
		updateFlag = true

		value := d.Get("altcontentsvcname").(string)
		update.Altcontentsvcname = value

		if value == "" {
			unsetFlag = true

			unset.Altcontentsvcname = true
		}

	}
	if d.HasChange("altcontentpath") {
		updateFlag = true

		value := d.Get("altcontentpath").(string)
		update.Altcontentpath = value

		if value == "" {
			unsetFlag = true

			unset.Altcontentpath = true
		}

	}
	if d.HasChange("polqdepth") {
		updateFlag = true

		value := d.Get("polqdepth").(int)
		update.Polqdepth = value

		if value == 0 {
			unsetFlag = true

			unset.Polqdepth = true
		}

	}
	if d.HasChange("priqdepth") {
		updateFlag = true

		value := d.Get("priqdepth").(int)
		update.Priqdepth = value

		if value == 0 {
			unsetFlag = true

			unset.Priqdepth = true
		}

	}
	if d.HasChange("maxconn") {
		updateFlag = true

		value := d.Get("maxconn").(int)
		update.Maxconn = value

		if value == 0 {
			unsetFlag = true

			unset.Maxconn = true
		}

	}
	if d.HasChange("delay") {
		updateFlag = true

		value := d.Get("delay").(int)
		update.Delay = value

		if value == 0 {
			unsetFlag = true

			unset.Delay = true
		}

	}
	if d.HasChange("dostrigexpression") {
		updateFlag = true

		value := d.Get("dostrigexpression").(string)
		update.Dostrigexpression = value

		if value == "" {
			unsetFlag = true

			unset.Dostrigexpression = true
		}

	}
	if d.HasChange("dosaction") {
		updateFlag = true

		value := d.Get("dosaction").(string)
		update.Dosaction = value

		if value == "" {
			unsetFlag = true

			unset.Dosaction = true
		}

	}
	if d.HasChange("tcpprofile") {
		updateFlag = true

		value := d.Get("tcpprofile").(string)
		update.Tcpprofile = value

		if value == "" {
			unsetFlag = true

			unset.Tcpprofile = true
		}

	}
	key := get_appqoeaction_key(d)

	if updateFlag {
		if err := client.UpdateAppqoeaction(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetAppqoeaction(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetAppqoeaction(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_appqoeaction(d, resource)
	}

	return nil
}

func delete_appqoeaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_appqoeaction")

	client := meta.(*nitro.NitroClient)

	resource := get_appqoeaction(d)
	key := resource.ToKey()

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
