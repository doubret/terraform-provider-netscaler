package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func NetscalerNshttpprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_nshttpprofile,
		Read:          read_nshttpprofile,
		Update:        update_nshttpprofile,
		Delete:        delete_nshttpprofile,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"adpttimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"altsvc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"apdexcltresptimethreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"clientiphdrexpr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cmponpush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"conmultiplex": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dropextracrlf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dropextradata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dropinvalreqs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2direct": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2headertablesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2initialwindowsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2maxconcurrentstreams": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2maxframesize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2maxheaderlistsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"http2minseverconn": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"incomphdrdelay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"markconnreqinval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"markhttp09inval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxheaderlen": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxreq": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxreusepool": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"minreusepool": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"persistentetag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"reqtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"reqtimeoutaction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"reusepooltimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rtsptunnel": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"spdy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"weblog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"websocket": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_nshttpprofile(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_nshttpprofile(d *schema.ResourceData) nitro.Nshttpprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Nshttpprofile{
		Name:        d.Get("name").(string),
		Adpttimeout: d.Get("adpttimeout").(string),
		Altsvc:      d.Get("altsvc").(string),
		Apdexcltresptimethreshold: d.Get("apdexcltresptimethreshold").(int),
		Clientiphdrexpr:           d.Get("clientiphdrexpr").(string),
		Cmponpush:                 d.Get("cmponpush").(string),
		Conmultiplex:              d.Get("conmultiplex").(string),
		Dropextracrlf:             d.Get("dropextracrlf").(string),
		Dropextradata:             d.Get("dropextradata").(string),
		Dropinvalreqs:             d.Get("dropinvalreqs").(string),
		Http2:                     d.Get("http2").(string),
		Http2direct:               d.Get("http2direct").(string),
		Http2headertablesize:      d.Get("http2headertablesize").(int),
		Http2initialwindowsize:    d.Get("http2initialwindowsize").(int),
		Http2maxconcurrentstreams: d.Get("http2maxconcurrentstreams").(int),
		Http2maxframesize:         d.Get("http2maxframesize").(int),
		Http2maxheaderlistsize:    d.Get("http2maxheaderlistsize").(int),
		Http2minseverconn:         d.Get("http2minseverconn").(int),
		Incomphdrdelay:            d.Get("incomphdrdelay").(int),
		Markconnreqinval:          d.Get("markconnreqinval").(string),
		Markhttp09inval:           d.Get("markhttp09inval").(string),
		Maxheaderlen:              d.Get("maxheaderlen").(int),
		Maxreq:                    d.Get("maxreq").(int),
		Maxreusepool:              d.Get("maxreusepool").(int),
		Minreusepool:              d.Get("minreusepool").(int),
		Persistentetag:            d.Get("persistentetag").(string),
		Reqtimeout:                d.Get("reqtimeout").(int),
		Reqtimeoutaction:          d.Get("reqtimeoutaction").(string),
		Reusepooltimeout:          d.Get("reusepooltimeout").(int),
		Rtsptunnel:                d.Get("rtsptunnel").(string),
		Spdy:                      d.Get("spdy").(string),
		Weblog:                    d.Get("weblog").(string),
		Websocket:                 d.Get("websocket").(string),
	}

	return resource
}

func set_nshttpprofile(d *schema.ResourceData, resource *nitro.Nshttpprofile) {
	d.Set("name", resource.Name)
	d.Set("adpttimeout", resource.Adpttimeout)
	d.Set("altsvc", resource.Altsvc)
	d.Set("apdexcltresptimethreshold", resource.Apdexcltresptimethreshold)
	d.Set("clientiphdrexpr", resource.Clientiphdrexpr)
	d.Set("cmponpush", resource.Cmponpush)
	d.Set("conmultiplex", resource.Conmultiplex)
	d.Set("dropextracrlf", resource.Dropextracrlf)
	d.Set("dropextradata", resource.Dropextradata)
	d.Set("dropinvalreqs", resource.Dropinvalreqs)
	d.Set("http2", resource.Http2)
	d.Set("http2direct", resource.Http2direct)
	d.Set("http2headertablesize", resource.Http2headertablesize)
	d.Set("http2initialwindowsize", resource.Http2initialwindowsize)
	d.Set("http2maxconcurrentstreams", resource.Http2maxconcurrentstreams)
	d.Set("http2maxframesize", resource.Http2maxframesize)
	d.Set("http2maxheaderlistsize", resource.Http2maxheaderlistsize)
	d.Set("http2minseverconn", resource.Http2minseverconn)
	d.Set("incomphdrdelay", resource.Incomphdrdelay)
	d.Set("markconnreqinval", resource.Markconnreqinval)
	d.Set("markhttp09inval", resource.Markhttp09inval)
	d.Set("maxheaderlen", resource.Maxheaderlen)
	d.Set("maxreq", resource.Maxreq)
	d.Set("maxreusepool", resource.Maxreusepool)
	d.Set("minreusepool", resource.Minreusepool)
	d.Set("persistentetag", resource.Persistentetag)
	d.Set("reqtimeout", resource.Reqtimeout)
	d.Set("reqtimeoutaction", resource.Reqtimeoutaction)
	d.Set("reusepooltimeout", resource.Reusepooltimeout)
	d.Set("rtsptunnel", resource.Rtsptunnel)
	d.Set("spdy", resource.Spdy)
	d.Set("weblog", resource.Weblog)
	d.Set("websocket", resource.Websocket)
	d.SetId(resource.Name)
}

func create_nshttpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_nshttpprofile")

	client := meta.(*nitro.NitroClient)

	key := key_nshttpprofile(d)

	exists, err := client.ExistsNshttpprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetNshttpprofile(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_nshttpprofile(d, resource)
	} else {
		err := client.AddNshttpprofile(get_nshttpprofile(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetNshttpprofile(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_nshttpprofile(d, resource)
	}

	return nil
}

func read_nshttpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_nshttpprofile")

	client := meta.(*nitro.NitroClient)

	key := key_nshttpprofile(d)

	exists, err := client.ExistsNshttpprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetNshttpprofile(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_nshttpprofile(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_nshttpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_nshttpprofile")

	return nil
}

func delete_nshttpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_nshttpprofile")

	client := meta.(*nitro.NitroClient)

	key := key_nshttpprofile(d)

	exists, err := client.ExistsNshttpprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteNshttpprofile(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
