package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerNstcpprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_nstcpprofile,
		Read:          read_nstcpprofile,
		Update:        update_nstcpprofile,
		Delete:        delete_nstcpprofile,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ackaggregation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ackonpush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"buffersize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"burstratecontrol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"delayedack": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dropestconnontimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"drophalfclosedconnontimeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dsack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dupackthresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dynamicreceivebuffering": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ecn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"establishclientconn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"fack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"flavor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"frto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"hystart": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"initialcwnd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ka": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"kaconnidletime": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"kamaxprobes": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"kaprobeinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"kaprobeupdatelastactivity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxburst": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxcwnd": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxpktpermss": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"minrto": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mptcp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mptcpdropdataonpreestsf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mptcpfastopen": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mptcpsessiontimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"mss": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"nagle": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"oooqsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"pktperretx": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rateqmax": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rstmaxack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rstwindowattenuate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sack": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sendbuffsize": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"slowstartincr": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"spoofsyndrop": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"syncookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"tcpfastopen": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"tcpmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"tcprate": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"tcpsegoffload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"ws": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"wsval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_nstcpprofile(d *schema.ResourceData) string {
	return d.Get("name").(string)
}

func get_nstcpprofile(d *schema.ResourceData) nitro.Nstcpprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Nstcpprofile{
		Name:                        d.Get("name").(string),
		Ackaggregation:              d.Get("ackaggregation").(string),
		Ackonpush:                   d.Get("ackonpush").(string),
		Buffersize:                  d.Get("buffersize").(int),
		Burstratecontrol:            d.Get("burstratecontrol").(string),
		Delayedack:                  d.Get("delayedack").(int),
		Dropestconnontimeout:        d.Get("dropestconnontimeout").(string),
		Drophalfclosedconnontimeout: d.Get("drophalfclosedconnontimeout").(string),
		Dsack:                   d.Get("dsack").(string),
		Dupackthresh:            d.Get("dupackthresh").(int),
		Dynamicreceivebuffering: d.Get("dynamicreceivebuffering").(string),
		Ecn:                       d.Get("ecn").(string),
		Establishclientconn:       d.Get("establishclientconn").(string),
		Fack:                      d.Get("fack").(string),
		Flavor:                    d.Get("flavor").(string),
		Frto:                      d.Get("frto").(string),
		Hystart:                   d.Get("hystart").(string),
		Initialcwnd:               d.Get("initialcwnd").(int),
		Ka:                        d.Get("ka").(string),
		Kaconnidletime:            d.Get("kaconnidletime").(int),
		Kamaxprobes:               d.Get("kamaxprobes").(int),
		Kaprobeinterval:           d.Get("kaprobeinterval").(int),
		Kaprobeupdatelastactivity: d.Get("kaprobeupdatelastactivity").(string),
		Maxburst:                  d.Get("maxburst").(int),
		Maxcwnd:                   d.Get("maxcwnd").(int),
		Maxpktpermss:              d.Get("maxpktpermss").(int),
		Minrto:                    d.Get("minrto").(int),
		Mptcp:                     d.Get("mptcp").(string),
		Mptcpdropdataonpreestsf: d.Get("mptcpdropdataonpreestsf").(string),
		Mptcpfastopen:           d.Get("mptcpfastopen").(string),
		Mptcpsessiontimeout:     d.Get("mptcpsessiontimeout").(int),
		Mss:                     d.Get("mss").(int),
		Nagle:                   d.Get("nagle").(string),
		Oooqsize:                d.Get("oooqsize").(int),
		Pktperretx:              d.Get("pktperretx").(int),
		Rateqmax:                d.Get("rateqmax").(int),
		Rstmaxack:               d.Get("rstmaxack").(string),
		Rstwindowattenuate:      d.Get("rstwindowattenuate").(string),
		Sack:                    d.Get("sack").(string),
		Sendbuffsize:            d.Get("sendbuffsize").(int),
		Slowstartincr:           d.Get("slowstartincr").(int),
		Spoofsyndrop:            d.Get("spoofsyndrop").(string),
		Syncookie:               d.Get("syncookie").(string),
		Tcpfastopen:             d.Get("tcpfastopen").(string),
		Tcpmode:                 d.Get("tcpmode").(string),
		Tcprate:                 d.Get("tcprate").(int),
		Tcpsegoffload:           d.Get("tcpsegoffload").(string),
		Timestamp:               d.Get("timestamp").(string),
		Ws:                      d.Get("ws").(string),
		Wsval:                   d.Get("wsval").(int),
	}

	return resource
}

func set_nstcpprofile(d *schema.ResourceData, resource *nitro.Nstcpprofile) {
	var _ = strconv.Itoa

	d.Set("name", resource.Name)
	d.Set("ackaggregation", resource.Ackaggregation)
	d.Set("ackonpush", resource.Ackonpush)
	d.Set("buffersize", resource.Buffersize)
	d.Set("burstratecontrol", resource.Burstratecontrol)
	d.Set("delayedack", resource.Delayedack)
	d.Set("dropestconnontimeout", resource.Dropestconnontimeout)
	d.Set("drophalfclosedconnontimeout", resource.Drophalfclosedconnontimeout)
	d.Set("dsack", resource.Dsack)
	d.Set("dupackthresh", resource.Dupackthresh)
	d.Set("dynamicreceivebuffering", resource.Dynamicreceivebuffering)
	d.Set("ecn", resource.Ecn)
	d.Set("establishclientconn", resource.Establishclientconn)
	d.Set("fack", resource.Fack)
	d.Set("flavor", resource.Flavor)
	d.Set("frto", resource.Frto)
	d.Set("hystart", resource.Hystart)
	d.Set("initialcwnd", resource.Initialcwnd)
	d.Set("ka", resource.Ka)
	d.Set("kaconnidletime", resource.Kaconnidletime)
	d.Set("kamaxprobes", resource.Kamaxprobes)
	d.Set("kaprobeinterval", resource.Kaprobeinterval)
	d.Set("kaprobeupdatelastactivity", resource.Kaprobeupdatelastactivity)
	d.Set("maxburst", resource.Maxburst)
	d.Set("maxcwnd", resource.Maxcwnd)
	d.Set("maxpktpermss", resource.Maxpktpermss)
	d.Set("minrto", resource.Minrto)
	d.Set("mptcp", resource.Mptcp)
	d.Set("mptcpdropdataonpreestsf", resource.Mptcpdropdataonpreestsf)
	d.Set("mptcpfastopen", resource.Mptcpfastopen)
	d.Set("mptcpsessiontimeout", resource.Mptcpsessiontimeout)
	d.Set("mss", resource.Mss)
	d.Set("nagle", resource.Nagle)
	d.Set("oooqsize", resource.Oooqsize)
	d.Set("pktperretx", resource.Pktperretx)
	d.Set("rateqmax", resource.Rateqmax)
	d.Set("rstmaxack", resource.Rstmaxack)
	d.Set("rstwindowattenuate", resource.Rstwindowattenuate)
	d.Set("sack", resource.Sack)
	d.Set("sendbuffsize", resource.Sendbuffsize)
	d.Set("slowstartincr", resource.Slowstartincr)
	d.Set("spoofsyndrop", resource.Spoofsyndrop)
	d.Set("syncookie", resource.Syncookie)
	d.Set("tcpfastopen", resource.Tcpfastopen)
	d.Set("tcpmode", resource.Tcpmode)
	d.Set("tcprate", resource.Tcprate)
	d.Set("tcpsegoffload", resource.Tcpsegoffload)
	d.Set("timestamp", resource.Timestamp)
	d.Set("ws", resource.Ws)
	d.Set("wsval", resource.Wsval)

	var key []string

	key = append(key, resource.Name)
	d.SetId(strings.Join(key, "-"))
}

func create_nstcpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_nstcpprofile")

	client := meta.(*nitro.NitroClient)

	key := key_nstcpprofile(d)

	exists, err := client.ExistsNstcpprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetNstcpprofile(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_nstcpprofile(d, resource)
	} else {
		err := client.AddNstcpprofile(get_nstcpprofile(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetNstcpprofile(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_nstcpprofile(d, resource)
	}

	return nil
}

func read_nstcpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_nstcpprofile")

	client := meta.(*nitro.NitroClient)

	key := key_nstcpprofile(d)

	exists, err := client.ExistsNstcpprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetNstcpprofile(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_nstcpprofile(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_nstcpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_nstcpprofile")

	return nil
}

func delete_nstcpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_nstcpprofile")

	client := meta.(*nitro.NitroClient)

	key := key_nstcpprofile(d)

	exists, err := client.ExistsNstcpprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteNstcpprofile(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
