package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func get_cachecontentgroup(d *schema.ResourceData) nitro.Cachecontentgroup {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Cachecontentgroup{
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
		Name:                   d.Get("name").(string),
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
	var _ = strconv.FormatBool

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
	d.Set("name", resource.Name)
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

func get_cachecontentgroup_key(d *schema.ResourceData) nitro.CachecontentgroupKey {

	key := nitro.CachecontentgroupKey{
		d.Get("name").(string),
	}
	return key
}

func create_cachecontentgroup(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_cachecontentgroup")

	client := meta.(*nitro.NitroClient)

	resource := get_cachecontentgroup(d)
	key := resource.ToKey()

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

	resource := get_cachecontentgroup(d)
	key := resource.ToKey()

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

	client := meta.(*nitro.NitroClient)

	update := nitro.CachecontentgroupUpdate{}
	unset := nitro.CachecontentgroupUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("weakposrelexpiry") {
		updateFlag = true

		value := d.Get("weakposrelexpiry").(int)
		update.Weakposrelexpiry = value

		if value == 0 {
			unsetFlag = true

			unset.Weakposrelexpiry = true
		}

	}
	if d.HasChange("heurexpiryparam") {
		updateFlag = true

		value := d.Get("heurexpiryparam").(int)
		update.Heurexpiryparam = value

		if value == 0 {
			unsetFlag = true

			unset.Heurexpiryparam = true
		}

	}
	if d.HasChange("relexpiry") {
		updateFlag = true

		value := d.Get("relexpiry").(int)
		update.Relexpiry = value

		if value == 0 {
			unsetFlag = true

			unset.Relexpiry = true
		}

	}
	if d.HasChange("relexpirymillisec") {
		updateFlag = true

		value := d.Get("relexpirymillisec").(int)
		update.Relexpirymillisec = value

		if value == 0 {
			unsetFlag = true

			unset.Relexpirymillisec = true
		}

	}
	if d.HasChange("absexpiry") {
		updateFlag = true

		value := utils.Convert_set_to_string_array(d.Get("absexpiry").(*schema.Set))
		update.Absexpiry = value

	}
	if d.HasChange("absexpirygmt") {
		updateFlag = true

		value := utils.Convert_set_to_string_array(d.Get("absexpirygmt").(*schema.Set))
		update.Absexpirygmt = value

	}
	if d.HasChange("weaknegrelexpiry") {
		updateFlag = true

		value := d.Get("weaknegrelexpiry").(int)
		update.Weaknegrelexpiry = value

		if value == 0 {
			unsetFlag = true

			unset.Weaknegrelexpiry = true
		}

	}
	if d.HasChange("hitparams") {
		updateFlag = true

		value := utils.Convert_set_to_string_array(d.Get("hitparams").(*schema.Set))
		update.Hitparams = value

	}
	if d.HasChange("invalparams") {
		updateFlag = true

		value := utils.Convert_set_to_string_array(d.Get("invalparams").(*schema.Set))
		update.Invalparams = value

	}
	if d.HasChange("ignoreparamvaluecase") {
		updateFlag = true

		value := d.Get("ignoreparamvaluecase").(string)
		update.Ignoreparamvaluecase = value

		if value == "" {
			unsetFlag = true

			unset.Ignoreparamvaluecase = true
		}

	}
	if d.HasChange("matchcookies") {
		updateFlag = true

		value := d.Get("matchcookies").(string)
		update.Matchcookies = value

		if value == "" {
			unsetFlag = true

			unset.Matchcookies = true
		}

	}
	if d.HasChange("invalrestrictedtohost") {
		updateFlag = true

		value := d.Get("invalrestrictedtohost").(string)
		update.Invalrestrictedtohost = value

		if value == "" {
			unsetFlag = true

			unset.Invalrestrictedtohost = true
		}

	}
	if d.HasChange("polleverytime") {
		updateFlag = true

		value := d.Get("polleverytime").(string)
		update.Polleverytime = value

		if value == "" {
			unsetFlag = true

			unset.Polleverytime = true
		}

	}
	if d.HasChange("ignorereloadreq") {
		updateFlag = true

		value := d.Get("ignorereloadreq").(string)
		update.Ignorereloadreq = value

		if value == "" {
			unsetFlag = true

			unset.Ignorereloadreq = true
		}

	}
	if d.HasChange("removecookies") {
		updateFlag = true

		value := d.Get("removecookies").(string)
		update.Removecookies = value

		if value == "" {
			unsetFlag = true

			unset.Removecookies = true
		}

	}
	if d.HasChange("prefetch") {
		updateFlag = true

		value := d.Get("prefetch").(string)
		update.Prefetch = value

		if value == "" {
			unsetFlag = true

			unset.Prefetch = true
		}

	}
	if d.HasChange("prefetchperiod") {
		updateFlag = true

		value := d.Get("prefetchperiod").(int)
		update.Prefetchperiod = value

		if value == 0 {
			unsetFlag = true

			unset.Prefetchperiod = true
		}

	}
	if d.HasChange("prefetchperiodmillisec") {
		updateFlag = true

		value := d.Get("prefetchperiodmillisec").(int)
		update.Prefetchperiodmillisec = value

		if value == 0 {
			unsetFlag = true

			unset.Prefetchperiodmillisec = true
		}

	}
	if d.HasChange("prefetchmaxpending") {
		updateFlag = true

		value := d.Get("prefetchmaxpending").(int)
		update.Prefetchmaxpending = value

		if value == 0 {
			unsetFlag = true

			unset.Prefetchmaxpending = true
		}

	}
	if d.HasChange("flashcache") {
		updateFlag = true

		value := d.Get("flashcache").(string)
		update.Flashcache = value

		if value == "" {
			unsetFlag = true

			unset.Flashcache = true
		}

	}
	if d.HasChange("expireatlastbyte") {
		updateFlag = true

		value := d.Get("expireatlastbyte").(string)
		update.Expireatlastbyte = value

		if value == "" {
			unsetFlag = true

			unset.Expireatlastbyte = true
		}

	}
	if d.HasChange("insertvia") {
		updateFlag = true

		value := d.Get("insertvia").(string)
		update.Insertvia = value

		if value == "" {
			unsetFlag = true

			unset.Insertvia = true
		}

	}
	if d.HasChange("insertage") {
		updateFlag = true

		value := d.Get("insertage").(string)
		update.Insertage = value

		if value == "" {
			unsetFlag = true

			unset.Insertage = true
		}

	}
	if d.HasChange("insertetag") {
		updateFlag = true

		value := d.Get("insertetag").(string)
		update.Insertetag = value

		if value == "" {
			unsetFlag = true

			unset.Insertetag = true
		}

	}
	if d.HasChange("cachecontrol") {
		updateFlag = true

		value := d.Get("cachecontrol").(string)
		update.Cachecontrol = value

		if value == "" {
			unsetFlag = true

			unset.Cachecontrol = true
		}

	}
	if d.HasChange("quickabortsize") {
		updateFlag = true

		value := d.Get("quickabortsize").(int)
		update.Quickabortsize = value

		if value == 0 {
			unsetFlag = true

			unset.Quickabortsize = true
		}

	}
	if d.HasChange("minressize") {
		updateFlag = true

		value := d.Get("minressize").(int)
		update.Minressize = value

		if value == 0 {
			unsetFlag = true

			unset.Minressize = true
		}

	}
	if d.HasChange("maxressize") {
		updateFlag = true

		value := d.Get("maxressize").(int)
		update.Maxressize = value

		if value == 0 {
			unsetFlag = true

			unset.Maxressize = true
		}

	}
	if d.HasChange("memlimit") {
		updateFlag = true

		value := d.Get("memlimit").(int)
		update.Memlimit = value

		if value == 0 {
			unsetFlag = true

			unset.Memlimit = true
		}

	}
	if d.HasChange("ignorereqcachinghdrs") {
		updateFlag = true

		value := d.Get("ignorereqcachinghdrs").(string)
		update.Ignorereqcachinghdrs = value

		if value == "" {
			unsetFlag = true

			unset.Ignorereqcachinghdrs = true
		}

	}
	if d.HasChange("minhits") {
		updateFlag = true

		value := d.Get("minhits").(int)
		update.Minhits = value

		if value == 0 {
			unsetFlag = true

			unset.Minhits = true
		}

	}
	if d.HasChange("alwaysevalpolicies") {
		updateFlag = true

		value := d.Get("alwaysevalpolicies").(string)
		update.Alwaysevalpolicies = value

		if value == "" {
			unsetFlag = true

			unset.Alwaysevalpolicies = true
		}

	}
	if d.HasChange("persistha") {
		updateFlag = true

		value := d.Get("persistha").(string)
		update.Persistha = value

		if value == "" {
			unsetFlag = true

			unset.Persistha = true
		}

	}
	if d.HasChange("pinned") {
		updateFlag = true

		value := d.Get("pinned").(string)
		update.Pinned = value

		if value == "" {
			unsetFlag = true

			unset.Pinned = true
		}

	}
	if d.HasChange("lazydnsresolve") {
		updateFlag = true

		value := d.Get("lazydnsresolve").(string)
		update.Lazydnsresolve = value

		if value == "" {
			unsetFlag = true

			unset.Lazydnsresolve = true
		}

	}
	if d.HasChange("hitselector") {
		updateFlag = true

		value := d.Get("hitselector").(string)
		update.Hitselector = value

		if value == "" {
			unsetFlag = true

			unset.Hitselector = true
		}

	}
	if d.HasChange("invalselector") {
		updateFlag = true

		value := d.Get("invalselector").(string)
		update.Invalselector = value

		if value == "" {
			unsetFlag = true

			unset.Invalselector = true
		}

	}
	key := get_cachecontentgroup_key(d)

	if updateFlag {
		if err := client.UpdateCachecontentgroup(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetCachecontentgroup(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetCachecontentgroup(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_cachecontentgroup(d, resource)
	}

	return nil
}

func delete_cachecontentgroup(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_cachecontentgroup")

	client := meta.(*nitro.NitroClient)

	resource := get_cachecontentgroup(d)
	key := resource.ToKey()

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
