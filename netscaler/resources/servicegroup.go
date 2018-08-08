package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerServicegroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_servicegroup,
		Read:          read_servicegroup,
		Update:        update_servicegroup,
		Delete:        delete_servicegroup,
		Schema: map[string]*schema.Schema{
			"appflowlog": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"autoscale": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cacheable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cachetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"cip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cipheader": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cka": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"clttimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cmp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"downstateflush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"healthmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"httpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxbandwidth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"maxclient": &schema.Schema{
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
			"memberport": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"monconnectionclose": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"monthreshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"netprofile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"pathmonitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"pathmonitorindv": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"rtspsessionidremap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"sc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"servicegroupname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"servicetype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"svrtimeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"tcpb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"tcpprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"td": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"useproxyport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"usip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func get_servicegroup(d *schema.ResourceData) nitro.Servicegroup {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Servicegroup{
		Appflowlog:         d.Get("appflowlog").(string),
		Autoscale:          d.Get("autoscale").(string),
		Cacheable:          d.Get("cacheable").(string),
		Cachetype:          d.Get("cachetype").(string),
		Cip:                d.Get("cip").(string),
		Cipheader:          d.Get("cipheader").(string),
		Cka:                d.Get("cka").(string),
		Clttimeout:         d.Get("clttimeout").(int),
		Cmp:                d.Get("cmp").(string),
		Comment:            d.Get("comment").(string),
		Downstateflush:     d.Get("downstateflush").(string),
		Healthmonitor:      d.Get("healthmonitor").(string),
		Httpprofilename:    d.Get("httpprofilename").(string),
		Maxbandwidth:       d.Get("maxbandwidth").(int),
		Maxclient:          d.Get("maxclient").(int),
		Maxreq:             d.Get("maxreq").(int),
		Memberport:         d.Get("memberport").(int),
		Monconnectionclose: d.Get("monconnectionclose").(string),
		Monthreshold:       d.Get("monthreshold").(int),
		Netprofile:         d.Get("netprofile").(string),
		Pathmonitor:        d.Get("pathmonitor").(string),
		Pathmonitorindv:    d.Get("pathmonitorindv").(string),
		Rtspsessionidremap: d.Get("rtspsessionidremap").(string),
		Sc:                 d.Get("sc").(string),
		Servicegroupname:   d.Get("servicegroupname").(string),
		Servicetype:        d.Get("servicetype").(string),
		Sp:                 d.Get("sp").(string),
		State:              d.Get("state").(string),
		Svrtimeout:         d.Get("svrtimeout").(int),
		Tcpb:               d.Get("tcpb").(string),
		Tcpprofilename:     d.Get("tcpprofilename").(string),
		Td:                 d.Get("td").(int),
		Useproxyport:       d.Get("useproxyport").(string),
		Usip:               d.Get("usip").(string),
	}

	return resource
}

func set_servicegroup(d *schema.ResourceData, resource *nitro.Servicegroup) {
	var _ = strconv.Itoa

	d.Set("appflowlog", resource.Appflowlog)
	d.Set("autoscale", resource.Autoscale)
	d.Set("cacheable", resource.Cacheable)
	d.Set("cachetype", resource.Cachetype)
	d.Set("cip", resource.Cip)
	d.Set("cipheader", resource.Cipheader)
	d.Set("cka", resource.Cka)
	d.Set("clttimeout", resource.Clttimeout)
	d.Set("cmp", resource.Cmp)
	d.Set("comment", resource.Comment)
	d.Set("downstateflush", resource.Downstateflush)
	d.Set("healthmonitor", resource.Healthmonitor)
	d.Set("httpprofilename", resource.Httpprofilename)
	d.Set("maxbandwidth", resource.Maxbandwidth)
	d.Set("maxclient", resource.Maxclient)
	d.Set("maxreq", resource.Maxreq)
	d.Set("memberport", resource.Memberport)
	d.Set("monconnectionclose", resource.Monconnectionclose)
	d.Set("monthreshold", resource.Monthreshold)
	d.Set("netprofile", resource.Netprofile)
	d.Set("pathmonitor", resource.Pathmonitor)
	d.Set("pathmonitorindv", resource.Pathmonitorindv)
	d.Set("rtspsessionidremap", resource.Rtspsessionidremap)
	d.Set("sc", resource.Sc)
	d.Set("servicegroupname", resource.Servicegroupname)
	d.Set("servicetype", resource.Servicetype)
	d.Set("sp", resource.Sp)
	d.Set("state", resource.State)
	d.Set("svrtimeout", resource.Svrtimeout)
	d.Set("tcpb", resource.Tcpb)
	d.Set("tcpprofilename", resource.Tcpprofilename)
	d.Set("td", resource.Td)
	d.Set("useproxyport", resource.Useproxyport)
	d.Set("usip", resource.Usip)

	var key []string

	key = append(key, resource.Servicegroupname)
	d.SetId(strings.Join(key, "-"))
}

func create_servicegroup(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_servicegroup")

	client := meta.(*nitro.NitroClient)

	resource := get_servicegroup(d)
	key := resource.ToKey()

	exists, err := client.ExistsServicegroup(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServicegroup(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_servicegroup(d, resource)
	} else {
		err := client.AddServicegroup(get_servicegroup(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetServicegroup(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_servicegroup(d, resource)
	}

	return nil
}

func read_servicegroup(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_servicegroup")

	client := meta.(*nitro.NitroClient)

	resource := get_servicegroup(d)
	key := resource.ToKey()

	exists, err := client.ExistsServicegroup(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetServicegroup(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_servicegroup(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_servicegroup(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_servicegroup")

	// TODO
	// client := meta.(*nitro.NitroClient)

	// err := client.UpdateServicegroup(get_servicegroup(d))

	// if err != nil {
	//       return err
	// }

	return nil
}

func delete_servicegroup(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_servicegroup")

	client := meta.(*nitro.NitroClient)

	resource := get_servicegroup(d)
	key := resource.ToKey()

	exists, err := client.ExistsServicegroup(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteServicegroup(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
