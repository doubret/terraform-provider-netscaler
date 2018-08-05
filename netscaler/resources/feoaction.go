package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
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
			"pageextendcache": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_feoaction(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_feoaction(d *schema.ResourceData) nitro.Feoaction {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Feoaction{
		Name:                   d.Get("name").(string),
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
		Pageextendcache:        d.Get("pageextendcache").(bool),
	}

	return resource
}

func set_feoaction(d *schema.ResourceData, resource *nitro.Feoaction) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
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
	d.Set("pageextendcache", resource.Pageextendcache)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_feoaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_feoaction")

	client := meta.(*nitro.NitroClient)

	key := key_feoaction(d)

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

	key := key_feoaction(d)

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

	return nil
}

func delete_feoaction(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_feoaction")

	client := meta.(*nitro.NitroClient)

	key := key_feoaction(d)

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
