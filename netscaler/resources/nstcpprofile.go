package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func get_nstcpprofile(d *schema.ResourceData) nitro.Nstcpprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Nstcpprofile{
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
		Name:                    d.Get("name").(string),
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
	d.Set("name", resource.Name)
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

func get_nstcpprofile_key(d *schema.ResourceData) nitro.NstcpprofileKey {

	key := nitro.NstcpprofileKey{
		d.Get("name").(string),
	}
	return key
}

func create_nstcpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_nstcpprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_nstcpprofile(d)
	key := resource.ToKey()

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

	resource := get_nstcpprofile(d)
	key := resource.ToKey()

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

	client := meta.(*nitro.NitroClient)

	update := nitro.NstcpprofileUpdate{}
	unset := nitro.NstcpprofileUnset{}

	updateFlag := false
	unsetFlag := false

	update.Name = d.Get("name").(string)
	unset.Name = d.Get("name").(string)

	if d.HasChange("ws") {
		updateFlag = true

		value := d.Get("ws").(string)
		update.Ws = value

		if value == "" {
			unsetFlag = true

			unset.Ws = true
		}

	}
	if d.HasChange("sack") {
		updateFlag = true

		value := d.Get("sack").(string)
		update.Sack = value

		if value == "" {
			unsetFlag = true

			unset.Sack = true
		}

	}
	if d.HasChange("wsval") {
		updateFlag = true

		value := d.Get("wsval").(int)
		update.Wsval = value

		if value == 0 {
			unsetFlag = true

			unset.Wsval = true
		}

	}
	if d.HasChange("nagle") {
		updateFlag = true

		value := d.Get("nagle").(string)
		update.Nagle = value

		if value == "" {
			unsetFlag = true

			unset.Nagle = true
		}

	}
	if d.HasChange("ackonpush") {
		updateFlag = true

		value := d.Get("ackonpush").(string)
		update.Ackonpush = value

		if value == "" {
			unsetFlag = true

			unset.Ackonpush = true
		}

	}
	if d.HasChange("mss") {
		updateFlag = true

		value := d.Get("mss").(int)
		update.Mss = value

		if value == 0 {
			unsetFlag = true

			unset.Mss = true
		}

	}
	if d.HasChange("maxburst") {
		updateFlag = true

		value := d.Get("maxburst").(int)
		update.Maxburst = value

		if value == 0 {
			unsetFlag = true

			unset.Maxburst = true
		}

	}
	if d.HasChange("initialcwnd") {
		updateFlag = true

		value := d.Get("initialcwnd").(int)
		update.Initialcwnd = value

		if value == 0 {
			unsetFlag = true

			unset.Initialcwnd = true
		}

	}
	if d.HasChange("delayedack") {
		updateFlag = true

		value := d.Get("delayedack").(int)
		update.Delayedack = value

		if value == 0 {
			unsetFlag = true

			unset.Delayedack = true
		}

	}
	if d.HasChange("oooqsize") {
		updateFlag = true

		value := d.Get("oooqsize").(int)
		update.Oooqsize = value

		if value == 0 {
			unsetFlag = true

			unset.Oooqsize = true
		}

	}
	if d.HasChange("maxpktpermss") {
		updateFlag = true

		value := d.Get("maxpktpermss").(int)
		update.Maxpktpermss = value

		if value == 0 {
			unsetFlag = true

			unset.Maxpktpermss = true
		}

	}
	if d.HasChange("pktperretx") {
		updateFlag = true

		value := d.Get("pktperretx").(int)
		update.Pktperretx = value

		if value == 0 {
			unsetFlag = true

			unset.Pktperretx = true
		}

	}
	if d.HasChange("minrto") {
		updateFlag = true

		value := d.Get("minrto").(int)
		update.Minrto = value

		if value == 0 {
			unsetFlag = true

			unset.Minrto = true
		}

	}
	if d.HasChange("slowstartincr") {
		updateFlag = true

		value := d.Get("slowstartincr").(int)
		update.Slowstartincr = value

		if value == 0 {
			unsetFlag = true

			unset.Slowstartincr = true
		}

	}
	if d.HasChange("buffersize") {
		updateFlag = true

		value := d.Get("buffersize").(int)
		update.Buffersize = value

		if value == 0 {
			unsetFlag = true

			unset.Buffersize = true
		}

	}
	if d.HasChange("syncookie") {
		updateFlag = true

		value := d.Get("syncookie").(string)
		update.Syncookie = value

		if value == "" {
			unsetFlag = true

			unset.Syncookie = true
		}

	}
	if d.HasChange("kaprobeupdatelastactivity") {
		updateFlag = true

		value := d.Get("kaprobeupdatelastactivity").(string)
		update.Kaprobeupdatelastactivity = value

		if value == "" {
			unsetFlag = true

			unset.Kaprobeupdatelastactivity = true
		}

	}
	if d.HasChange("flavor") {
		updateFlag = true

		value := d.Get("flavor").(string)
		update.Flavor = value

		if value == "" {
			unsetFlag = true

			unset.Flavor = true
		}

	}
	if d.HasChange("dynamicreceivebuffering") {
		updateFlag = true

		value := d.Get("dynamicreceivebuffering").(string)
		update.Dynamicreceivebuffering = value

		if value == "" {
			unsetFlag = true

			unset.Dynamicreceivebuffering = true
		}

	}
	if d.HasChange("ka") {
		updateFlag = true

		value := d.Get("ka").(string)
		update.Ka = value

		if value == "" {
			unsetFlag = true

			unset.Ka = true
		}

	}
	if d.HasChange("kaconnidletime") {
		updateFlag = true

		value := d.Get("kaconnidletime").(int)
		update.Kaconnidletime = value

		if value == 0 {
			unsetFlag = true

			unset.Kaconnidletime = true
		}

	}
	if d.HasChange("kamaxprobes") {
		updateFlag = true

		value := d.Get("kamaxprobes").(int)
		update.Kamaxprobes = value

		if value == 0 {
			unsetFlag = true

			unset.Kamaxprobes = true
		}

	}
	if d.HasChange("kaprobeinterval") {
		updateFlag = true

		value := d.Get("kaprobeinterval").(int)
		update.Kaprobeinterval = value

		if value == 0 {
			unsetFlag = true

			unset.Kaprobeinterval = true
		}

	}
	if d.HasChange("sendbuffsize") {
		updateFlag = true

		value := d.Get("sendbuffsize").(int)
		update.Sendbuffsize = value

		if value == 0 {
			unsetFlag = true

			unset.Sendbuffsize = true
		}

	}
	if d.HasChange("mptcp") {
		updateFlag = true

		value := d.Get("mptcp").(string)
		update.Mptcp = value

		if value == "" {
			unsetFlag = true

			unset.Mptcp = true
		}

	}
	if d.HasChange("establishclientconn") {
		updateFlag = true

		value := d.Get("establishclientconn").(string)
		update.Establishclientconn = value

		if value == "" {
			unsetFlag = true

			unset.Establishclientconn = true
		}

	}
	if d.HasChange("tcpsegoffload") {
		updateFlag = true

		value := d.Get("tcpsegoffload").(string)
		update.Tcpsegoffload = value

		if value == "" {
			unsetFlag = true

			unset.Tcpsegoffload = true
		}

	}
	if d.HasChange("rstwindowattenuate") {
		updateFlag = true

		value := d.Get("rstwindowattenuate").(string)
		update.Rstwindowattenuate = value

		if value == "" {
			unsetFlag = true

			unset.Rstwindowattenuate = true
		}

	}
	if d.HasChange("rstmaxack") {
		updateFlag = true

		value := d.Get("rstmaxack").(string)
		update.Rstmaxack = value

		if value == "" {
			unsetFlag = true

			unset.Rstmaxack = true
		}

	}
	if d.HasChange("spoofsyndrop") {
		updateFlag = true

		value := d.Get("spoofsyndrop").(string)
		update.Spoofsyndrop = value

		if value == "" {
			unsetFlag = true

			unset.Spoofsyndrop = true
		}

	}
	if d.HasChange("ecn") {
		updateFlag = true

		value := d.Get("ecn").(string)
		update.Ecn = value

		if value == "" {
			unsetFlag = true

			unset.Ecn = true
		}

	}
	if d.HasChange("mptcpdropdataonpreestsf") {
		updateFlag = true

		value := d.Get("mptcpdropdataonpreestsf").(string)
		update.Mptcpdropdataonpreestsf = value

		if value == "" {
			unsetFlag = true

			unset.Mptcpdropdataonpreestsf = true
		}

	}
	if d.HasChange("mptcpfastopen") {
		updateFlag = true

		value := d.Get("mptcpfastopen").(string)
		update.Mptcpfastopen = value

		if value == "" {
			unsetFlag = true

			unset.Mptcpfastopen = true
		}

	}
	if d.HasChange("mptcpsessiontimeout") {
		updateFlag = true

		value := d.Get("mptcpsessiontimeout").(int)
		update.Mptcpsessiontimeout = value

		if value == 0 {
			unsetFlag = true

			unset.Mptcpsessiontimeout = true
		}

	}
	if d.HasChange("timestamp") {
		updateFlag = true

		value := d.Get("timestamp").(string)
		update.Timestamp = value

		if value == "" {
			unsetFlag = true

			unset.Timestamp = true
		}

	}
	if d.HasChange("dsack") {
		updateFlag = true

		value := d.Get("dsack").(string)
		update.Dsack = value

		if value == "" {
			unsetFlag = true

			unset.Dsack = true
		}

	}
	if d.HasChange("ackaggregation") {
		updateFlag = true

		value := d.Get("ackaggregation").(string)
		update.Ackaggregation = value

		if value == "" {
			unsetFlag = true

			unset.Ackaggregation = true
		}

	}
	if d.HasChange("frto") {
		updateFlag = true

		value := d.Get("frto").(string)
		update.Frto = value

		if value == "" {
			unsetFlag = true

			unset.Frto = true
		}

	}
	if d.HasChange("maxcwnd") {
		updateFlag = true

		value := d.Get("maxcwnd").(int)
		update.Maxcwnd = value

		if value == 0 {
			unsetFlag = true

			unset.Maxcwnd = true
		}

	}
	if d.HasChange("fack") {
		updateFlag = true

		value := d.Get("fack").(string)
		update.Fack = value

		if value == "" {
			unsetFlag = true

			unset.Fack = true
		}

	}
	if d.HasChange("tcpmode") {
		updateFlag = true

		value := d.Get("tcpmode").(string)
		update.Tcpmode = value

		if value == "" {
			unsetFlag = true

			unset.Tcpmode = true
		}

	}
	if d.HasChange("tcpfastopen") {
		updateFlag = true

		value := d.Get("tcpfastopen").(string)
		update.Tcpfastopen = value

		if value == "" {
			unsetFlag = true

			unset.Tcpfastopen = true
		}

	}
	if d.HasChange("hystart") {
		updateFlag = true

		value := d.Get("hystart").(string)
		update.Hystart = value

		if value == "" {
			unsetFlag = true

			unset.Hystart = true
		}

	}
	if d.HasChange("dupackthresh") {
		updateFlag = true

		value := d.Get("dupackthresh").(int)
		update.Dupackthresh = value

		if value == 0 {
			unsetFlag = true

			unset.Dupackthresh = true
		}

	}
	if d.HasChange("burstratecontrol") {
		updateFlag = true

		value := d.Get("burstratecontrol").(string)
		update.Burstratecontrol = value

		if value == "" {
			unsetFlag = true

			unset.Burstratecontrol = true
		}

	}
	if d.HasChange("tcprate") {
		updateFlag = true

		value := d.Get("tcprate").(int)
		update.Tcprate = value

		if value == 0 {
			unsetFlag = true

			unset.Tcprate = true
		}

	}
	if d.HasChange("rateqmax") {
		updateFlag = true

		value := d.Get("rateqmax").(int)
		update.Rateqmax = value

		if value == 0 {
			unsetFlag = true

			unset.Rateqmax = true
		}

	}
	if d.HasChange("drophalfclosedconnontimeout") {
		updateFlag = true

		value := d.Get("drophalfclosedconnontimeout").(string)
		update.Drophalfclosedconnontimeout = value

		if value == "" {
			unsetFlag = true

			unset.Drophalfclosedconnontimeout = true
		}

	}
	if d.HasChange("dropestconnontimeout") {
		updateFlag = true

		value := d.Get("dropestconnontimeout").(string)
		update.Dropestconnontimeout = value

		if value == "" {
			unsetFlag = true

			unset.Dropestconnontimeout = true
		}

	}
	key := get_nstcpprofile_key(d)

	if updateFlag {
		if err := client.UpdateNstcpprofile(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetNstcpprofile(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetNstcpprofile(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_nstcpprofile(d, resource)
	}

	return nil
}

func delete_nstcpprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_nstcpprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_nstcpprofile(d)
	key := resource.ToKey()

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
