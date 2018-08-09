package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerFeoaction() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_feoaction,
		Read:          read_feoaction,
		Update:        update_feoaction,
		Delete:        delete_feoaction,
		Schema: map[string]*schema.Schema{
			"cachemaxage": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"clientsidemeasurements": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"convertimporttolink": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"csscombine": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cssimginline": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cssinline": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cssminify": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cssmovetohead": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dnsshards": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"domainsharding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"htmlminify": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"imggiftopng": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"imginline": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"imglazyload": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"imgshrinktoattrib": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"imgtojpegxr": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"imgtowebp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"jpgoptimize": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"jsinline": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"jsminify": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"jsmovetoend": &schema.Schema{
				Type:     schema.TypeBool,
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
			"pageextendcache": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_feoaction(d *schema.ResourceData) nitro.Feoaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Feoaction{
		Cachemaxage:            d.Get("cachemaxage").(int),
		Clientsidemeasurements: d.Get("clientsidemeasurements").(bool),
		Convertimporttolink:    d.Get("convertimporttolink").(bool),
		Csscombine:             d.Get("csscombine").(bool),
		Cssimginline:           d.Get("cssimginline").(bool),
		Cssinline:              d.Get("cssinline").(bool),
		Cssminify:              d.Get("cssminify").(bool),
		Cssmovetohead:          d.Get("cssmovetohead").(bool),
		Dnsshards:              utils.Convert_set_to_string_array(d.Get("dnsshards").(*schema.Set)),
		Domainsharding:         d.Get("domainsharding").(string),
		Htmlminify:             d.Get("htmlminify").(bool),
		Imggiftopng:            d.Get("imggiftopng").(bool),
		Imginline:              d.Get("imginline").(bool),
		Imglazyload:            d.Get("imglazyload").(bool),
		Imgshrinktoattrib:      d.Get("imgshrinktoattrib").(bool),
		Imgtojpegxr:            d.Get("imgtojpegxr").(bool),
		Imgtowebp:              d.Get("imgtowebp").(bool),
		Jpgoptimize:            d.Get("jpgoptimize").(bool),
		Jsinline:               d.Get("jsinline").(bool),
		Jsminify:               d.Get("jsminify").(bool),
		Jsmovetoend:            d.Get("jsmovetoend").(bool),
		Name:                   d.Get("name").(string),
		Pageextendcache:        d.Get("pageextendcache").(bool),
	}

	return resource
}

func set_feoaction(d *schema.ResourceData, resource *nitro.Feoaction) {
	var _ = strconv.Itoa

	d.Set("cachemaxage", resource.Cachemaxage)
	d.Set("clientsidemeasurements", resource.Clientsidemeasurements)
	d.Set("convertimporttolink", resource.Convertimporttolink)
	d.Set("csscombine", resource.Csscombine)
	d.Set("cssimginline", resource.Cssimginline)
	d.Set("cssinline", resource.Cssinline)
	d.Set("cssminify", resource.Cssminify)
	d.Set("cssmovetohead", resource.Cssmovetohead)
	d.Set("dnsshards", resource.Dnsshards)
	d.Set("domainsharding", resource.Domainsharding)
	d.Set("htmlminify", resource.Htmlminify)
	d.Set("imggiftopng", resource.Imggiftopng)
	d.Set("imginline", resource.Imginline)
	d.Set("imglazyload", resource.Imglazyload)
	d.Set("imgshrinktoattrib", resource.Imgshrinktoattrib)
	d.Set("imgtojpegxr", resource.Imgtojpegxr)
	d.Set("imgtowebp", resource.Imgtowebp)
	d.Set("jpgoptimize", resource.Jpgoptimize)
	d.Set("jsinline", resource.Jsinline)
	d.Set("jsminify", resource.Jsminify)
	d.Set("jsmovetoend", resource.Jsmovetoend)
	d.Set("name", resource.Name)
	d.Set("pageextendcache", resource.Pageextendcache)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func get_feoaction_key(d *schema.ResourceData) nitro.FeoactionKey {

	key := nitro.FeoactionKey{
		d.Get("name").(string),
	}
	return key
}

func create_feoaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_feoaction")

	client := meta.(*nitro.NitroClient)

	resource := get_feoaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsFeoaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFeoaction(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_feoaction(d, resource)
	} else {
		err := client.AddFeoaction(get_feoaction(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetFeoaction(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_feoaction(d, resource)
	}

	return nil
}

func read_feoaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_feoaction")

	client := meta.(*nitro.NitroClient)

	resource := get_feoaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsFeoaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetFeoaction(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_feoaction(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_feoaction(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_feoaction")

	client := meta.(*nitro.NitroClient)

	update := nitro.FeoactionUpdate{}
	unset := nitro.FeoactionUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("pageextendcache") {
		updateFlag = true

		value := d.Get("pageextendcache").(bool)
		update.Pageextendcache = value

	}
	if d.HasChange("cachemaxage") {
		updateFlag = true

		value := d.Get("cachemaxage").(int)
		update.Cachemaxage = value

		if value == 0 {
			unsetFlag = true

			unset.Cachemaxage = true
		}

	}
	if d.HasChange("imgshrinktoattrib") {
		updateFlag = true

		value := d.Get("imgshrinktoattrib").(bool)
		update.Imgshrinktoattrib = value

	}
	if d.HasChange("imggiftopng") {
		updateFlag = true

		value := d.Get("imggiftopng").(bool)
		update.Imggiftopng = value

	}
	if d.HasChange("imgtowebp") {
		updateFlag = true

		value := d.Get("imgtowebp").(bool)
		update.Imgtowebp = value

	}
	if d.HasChange("imgtojpegxr") {
		updateFlag = true

		value := d.Get("imgtojpegxr").(bool)
		update.Imgtojpegxr = value

	}
	if d.HasChange("imginline") {
		updateFlag = true

		value := d.Get("imginline").(bool)
		update.Imginline = value

	}
	if d.HasChange("cssimginline") {
		updateFlag = true

		value := d.Get("cssimginline").(bool)
		update.Cssimginline = value

	}
	if d.HasChange("jpgoptimize") {
		updateFlag = true

		value := d.Get("jpgoptimize").(bool)
		update.Jpgoptimize = value

	}
	if d.HasChange("imglazyload") {
		updateFlag = true

		value := d.Get("imglazyload").(bool)
		update.Imglazyload = value

	}
	if d.HasChange("cssminify") {
		updateFlag = true

		value := d.Get("cssminify").(bool)
		update.Cssminify = value

	}
	if d.HasChange("cssinline") {
		updateFlag = true

		value := d.Get("cssinline").(bool)
		update.Cssinline = value

	}
	if d.HasChange("csscombine") {
		updateFlag = true

		value := d.Get("csscombine").(bool)
		update.Csscombine = value

	}
	if d.HasChange("convertimporttolink") {
		updateFlag = true

		value := d.Get("convertimporttolink").(bool)
		update.Convertimporttolink = value

	}
	if d.HasChange("jsminify") {
		updateFlag = true

		value := d.Get("jsminify").(bool)
		update.Jsminify = value

	}
	if d.HasChange("jsinline") {
		updateFlag = true

		value := d.Get("jsinline").(bool)
		update.Jsinline = value

	}
	if d.HasChange("htmlminify") {
		updateFlag = true

		value := d.Get("htmlminify").(bool)
		update.Htmlminify = value

	}
	if d.HasChange("cssmovetohead") {
		updateFlag = true

		value := d.Get("cssmovetohead").(bool)
		update.Cssmovetohead = value

	}
	if d.HasChange("jsmovetoend") {
		updateFlag = true

		value := d.Get("jsmovetoend").(bool)
		update.Jsmovetoend = value

	}
	if d.HasChange("domainsharding") {
		updateFlag = true

		value := d.Get("domainsharding").(string)
		update.Domainsharding = value

		if value == "" {
			unsetFlag = true

			unset.Domainsharding = true
		}

	}
	if d.HasChange("dnsshards") {
		updateFlag = true

		value := utils.Convert_set_to_string_array(d.Get("dnsshards").(*schema.Set))
		update.Dnsshards = value

	}
	if d.HasChange("clientsidemeasurements") {
		updateFlag = true

		value := d.Get("clientsidemeasurements").(bool)
		update.Clientsidemeasurements = value

	}
	key := get_feoaction_key(d)

	if updateFlag {
		if err := client.UpdateFeoaction(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetFeoaction(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetFeoaction(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_feoaction(d, resource)
	}

	return nil
}

func delete_feoaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_feoaction")

	client := meta.(*nitro.NitroClient)

	resource := get_feoaction(d)
	key := resource.ToKey()

	exists, err := client.ExistsFeoaction(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteFeoaction(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
