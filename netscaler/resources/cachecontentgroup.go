package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerCachecontentgroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_cachecontentgroup,
		Read:          read_cachecontentgroup,
		Update:        update_cachecontentgroup,
		Delete:        delete_cachecontentgroup,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"absexpiry": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"absexpirygmt": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"alwaysevalpolicies": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cachecontrol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"expireatlastbyte": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"flashcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"heurexpiryparam": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"hitparams": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"hitselector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ignoreparamvaluecase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ignorereloadreq": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ignorereqcachinghdrs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"insertage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"insertetag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"insertvia": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"invalparams": &schema.Schema{
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"invalrestrictedtohost": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"invalselector": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"lazydnsresolve": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"matchcookies": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxressize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"memlimit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"minhits": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"minressize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"persistha": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"pinned": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"polleverytime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"prefetch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"prefetchmaxpending": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"prefetchperiod": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"prefetchperiodmillisec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"quickabortsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"relexpiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"relexpirymillisec": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"removecookies": &schema.Schema{
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
			"weaknegrelexpiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"weakposrelexpiry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_cachecontentgroup(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_cachecontentgroup(d *schema.ResourceData) nitro.Cachecontentgroup {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Cachecontentgroup{
		Name:                   d.Get("name").(string),
		Absexpiry:              utils.Convert_set_to_string_array(d.Get("absexpiry").(*schema.Set)),
		Absexpirygmt:           utils.Convert_set_to_string_array(d.Get("absexpirygmt").(*schema.Set)),
		Alwaysevalpolicies:     d.Get("alwaysevalpolicies").(string),
		Cachecontrol:           d.Get("cachecontrol").(string),
		Expireatlastbyte:       d.Get("expireatlastbyte").(string),
		Flashcache:             d.Get("flashcache").(string),
		Heurexpiryparam:        d.Get("heurexpiryparam").(int),
		Hitparams:              utils.Convert_set_to_string_array(d.Get("hitparams").(*schema.Set)),
		Hitselector:            d.Get("hitselector").(string),
		Ignoreparamvaluecase:   d.Get("ignoreparamvaluecase").(string),
		Ignorereloadreq:        d.Get("ignorereloadreq").(string),
		Ignorereqcachinghdrs:   d.Get("ignorereqcachinghdrs").(string),
		Insertage:              d.Get("insertage").(string),
		Insertetag:             d.Get("insertetag").(string),
		Insertvia:              d.Get("insertvia").(string),
		Invalparams:            utils.Convert_set_to_string_array(d.Get("invalparams").(*schema.Set)),
		Invalrestrictedtohost:  d.Get("invalrestrictedtohost").(string),
		Invalselector:          d.Get("invalselector").(string),
		Lazydnsresolve:         d.Get("lazydnsresolve").(string),
		Matchcookies:           d.Get("matchcookies").(string),
		Maxressize:             d.Get("maxressize").(int),
		Memlimit:               d.Get("memlimit").(int),
		Minhits:                d.Get("minhits").(int),
		Minressize:             d.Get("minressize").(int),
		Persistha:              d.Get("persistha").(string),
		Pinned:                 d.Get("pinned").(string),
		Polleverytime:          d.Get("polleverytime").(string),
		Prefetch:               d.Get("prefetch").(string),
		Prefetchmaxpending:     d.Get("prefetchmaxpending").(int),
		Prefetchperiod:         d.Get("prefetchperiod").(int),
		Prefetchperiodmillisec: d.Get("prefetchperiodmillisec").(int),
		Quickabortsize:         d.Get("quickabortsize").(int),
		Relexpiry:              d.Get("relexpiry").(int),
		Relexpirymillisec:      d.Get("relexpirymillisec").(int),
		Removecookies:          d.Get("removecookies").(string),
		Type:                   d.Get("type").(string),
		Weaknegrelexpiry:       d.Get("weaknegrelexpiry").(int),
		Weakposrelexpiry:       d.Get("weakposrelexpiry").(int),
	}

	return resource
}

func set_cachecontentgroup(d *schema.ResourceData, resource *nitro.Cachecontentgroup) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("absexpiry", resource.Absexpiry)
	d.Set("absexpirygmt", resource.Absexpirygmt)
	d.Set("alwaysevalpolicies", resource.Alwaysevalpolicies)
	d.Set("cachecontrol", resource.Cachecontrol)
	d.Set("expireatlastbyte", resource.Expireatlastbyte)
	d.Set("flashcache", resource.Flashcache)
	d.Set("heurexpiryparam", resource.Heurexpiryparam)
	d.Set("hitparams", resource.Hitparams)
	d.Set("hitselector", resource.Hitselector)
	d.Set("ignoreparamvaluecase", resource.Ignoreparamvaluecase)
	d.Set("ignorereloadreq", resource.Ignorereloadreq)
	d.Set("ignorereqcachinghdrs", resource.Ignorereqcachinghdrs)
	d.Set("insertage", resource.Insertage)
	d.Set("insertetag", resource.Insertetag)
	d.Set("insertvia", resource.Insertvia)
	d.Set("invalparams", resource.Invalparams)
	d.Set("invalrestrictedtohost", resource.Invalrestrictedtohost)
	d.Set("invalselector", resource.Invalselector)
	d.Set("lazydnsresolve", resource.Lazydnsresolve)
	d.Set("matchcookies", resource.Matchcookies)
	d.Set("maxressize", resource.Maxressize)
	d.Set("memlimit", resource.Memlimit)
	d.Set("minhits", resource.Minhits)
	d.Set("minressize", resource.Minressize)
	d.Set("persistha", resource.Persistha)
	d.Set("pinned", resource.Pinned)
	d.Set("polleverytime", resource.Polleverytime)
	d.Set("prefetch", resource.Prefetch)
	d.Set("prefetchmaxpending", resource.Prefetchmaxpending)
	d.Set("prefetchperiod", resource.Prefetchperiod)
	d.Set("prefetchperiodmillisec", resource.Prefetchperiodmillisec)
	d.Set("quickabortsize", resource.Quickabortsize)
	d.Set("relexpiry", resource.Relexpiry)
	d.Set("relexpirymillisec", resource.Relexpirymillisec)
	d.Set("removecookies", resource.Removecookies)
	d.Set("type", resource.Type)
	d.Set("weaknegrelexpiry", resource.Weaknegrelexpiry)
	d.Set("weakposrelexpiry", resource.Weakposrelexpiry)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_cachecontentgroup(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cachecontentgroup")

	client := meta.(*nitro.NitroClient)

	key := key_cachecontentgroup(d)

	exists, err := client.ExistsCachecontentgroup(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCachecontentgroup(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_cachecontentgroup(d, resource)
	} else {
		err := client.AddCachecontentgroup(get_cachecontentgroup(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetCachecontentgroup(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_cachecontentgroup(d, resource)
	}

	return nil
}

func read_cachecontentgroup(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_cachecontentgroup")

	client := meta.(*nitro.NitroClient)

	key := key_cachecontentgroup(d)

	exists, err := client.ExistsCachecontentgroup(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetCachecontentgroup(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_cachecontentgroup(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_cachecontentgroup(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_cachecontentgroup")

	return nil
}

func delete_cachecontentgroup(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cachecontentgroup")

	client := meta.(*nitro.NitroClient)

	key := key_cachecontentgroup(d)

	exists, err := client.ExistsCachecontentgroup(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteCachecontentgroup(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
