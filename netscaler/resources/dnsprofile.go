package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/citrix-netscaler-terraform-provider/netscaler/utils"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
	"strings"
)

func NetscalerDnsprofile() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        create_dnsprofile,
		Read:          read_dnsprofile,
		Update:        update_dnsprofile,
		Delete:        delete_dnsprofile,
		Schema: map[string]*schema.Schema{
			"dnsprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cacheecsresponses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cachenegativeresponses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"cacherecords": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dnsanswerseclogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dnserrorlogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dnsextendedlogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dnsquerylogging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
			"dropmultiqueryrequest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},
		},
	}
}

func key_dnsprofile(d *schema.ResourceData) string {
	return d.Get("dnsprofilename").(string)
}

func get_dnsprofile(d *schema.ResourceData) nitro.Dnsprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Dnsprofile{
		Dnsprofilename:         d.Get("dnsprofilename").(string),
		Cacheecsresponses:      d.Get("cacheecsresponses").(string),
		Cachenegativeresponses: d.Get("cachenegativeresponses").(string),
		Cacherecords:           d.Get("cacherecords").(string),
		Dnsanswerseclogging:    d.Get("dnsanswerseclogging").(string),
		Dnserrorlogging:        d.Get("dnserrorlogging").(string),
		Dnsextendedlogging:     d.Get("dnsextendedlogging").(string),
		Dnsquerylogging:        d.Get("dnsquerylogging").(string),
		Dropmultiqueryrequest:  d.Get("dropmultiqueryrequest").(string),
	}

	return resource
}

func set_dnsprofile(d *schema.ResourceData, resource *nitro.Dnsprofile) {
	var _ = strconv.Itoa

	d.Set("dnsprofilename", resource.Dnsprofilename)
	d.Set("cacheecsresponses", resource.Cacheecsresponses)
	d.Set("cachenegativeresponses", resource.Cachenegativeresponses)
	d.Set("cacherecords", resource.Cacherecords)
	d.Set("dnsanswerseclogging", resource.Dnsanswerseclogging)
	d.Set("dnserrorlogging", resource.Dnserrorlogging)
	d.Set("dnsextendedlogging", resource.Dnsextendedlogging)
	d.Set("dnsquerylogging", resource.Dnsquerylogging)
	d.Set("dropmultiqueryrequest", resource.Dropmultiqueryrequest)

	var key []string

	key = append(key, resource.Dnsprofilename)
	d.SetId(strings.Join(key, "-"))
}

func create_dnsprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dnsprofile")

	client := meta.(*nitro.NitroClient)

	key := key_dnsprofile(d)

	exists, err := client.ExistsDnsprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDnsprofile(key)

		if err != nil {
			log.Print("Failed to get existing resource : ", err)

			return err
		}

		set_dnsprofile(d, resource)
	} else {
		err := client.AddDnsprofile(get_dnsprofile(d))

		if err != nil {
			log.Print("Failed to create resource : ", err)

			return err
		}

		resource, err := client.GetDnsprofile(key)

		if err != nil {
			log.Print("Failed to get created resource : ", err)

			return err
		}

		set_dnsprofile(d, resource)
	}

	return nil
}

func read_dnsprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In read_dnsprofile")

	client := meta.(*nitro.NitroClient)

	key := key_dnsprofile(d)

	exists, err := client.ExistsDnsprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		resource, err := client.GetDnsprofile(key)

		if err != nil {
			log.Print("Failed to get resource : ", err)

			return err
		}

		set_dnsprofile(d, resource)
	} else {
		d.SetId("")
	}

	return nil
}

func update_dnsprofile(d *schema.ResourceData, meta interface{}) error {
	log.Println("[DEBUG] netscaler-provider:  In update_dnsprofile")

	return nil
}

func delete_dnsprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dnsprofile")

	client := meta.(*nitro.NitroClient)

	key := key_dnsprofile(d)

	exists, err := client.ExistsDnsprofile(key)

	if err != nil {
		log.Print("Failed to check if resource exists : ", err)

		return err
	}

	if exists {
		err := client.DeleteDnsprofile(key)

		if err != nil {
			log.Print("Failed to delete resource : ", err)

			return err
		}
	}

	d.SetId("")

	return nil
}
