package resources

import (
	"github.com/doubret/citrix-netscaler-nitro-go-client/nitro"
	"github.com/doubret/terraform-provider-netscaler/netscaler/utils"
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
			"dnsprofilename": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
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

func get_dnsprofile(d *schema.ResourceData) nitro.Dnsprofile {
	var _ = utils.Convert_set_to_string_array

	resource := nitro.Dnsprofile{
		Cacheecsresponses:      d.Get("cacheecsresponses").(string),
		Cachenegativeresponses: d.Get("cachenegativeresponses").(string),
		Cacherecords:           d.Get("cacherecords").(string),
		Dnsanswerseclogging:    d.Get("dnsanswerseclogging").(string),
		Dnserrorlogging:        d.Get("dnserrorlogging").(string),
		Dnsextendedlogging:     d.Get("dnsextendedlogging").(string),
		Dnsprofilename:         d.Get("dnsprofilename").(string),
		Dnsquerylogging:        d.Get("dnsquerylogging").(string),
		Dropmultiqueryrequest:  d.Get("dropmultiqueryrequest").(string),
	}

	return resource
}

func set_dnsprofile(d *schema.ResourceData, resource *nitro.Dnsprofile) {
	var _ = strconv.Itoa
	var _ = strconv.FormatBool

	d.Set("cacheecsresponses", resource.Cacheecsresponses)
	d.Set("cachenegativeresponses", resource.Cachenegativeresponses)
	d.Set("cacherecords", resource.Cacherecords)
	d.Set("dnsanswerseclogging", resource.Dnsanswerseclogging)
	d.Set("dnserrorlogging", resource.Dnserrorlogging)
	d.Set("dnsextendedlogging", resource.Dnsextendedlogging)
	d.Set("dnsprofilename", resource.Dnsprofilename)
	d.Set("dnsquerylogging", resource.Dnsquerylogging)
	d.Set("dropmultiqueryrequest", resource.Dropmultiqueryrequest)

	var key []string

	key = append(key, resource.Dnsprofilename)
	d.SetId(strings.Join(key, "-"))
}

func get_dnsprofile_key(d *schema.ResourceData) nitro.DnsprofileKey {

	key := nitro.DnsprofileKey{
		d.Get("dnsprofilename").(string),
	}
	return key
}

func create_dnsprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In create_dnsprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_dnsprofile(d)
	key := resource.ToKey()

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

	resource := get_dnsprofile(d)
	key := resource.ToKey()

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

	client := meta.(*nitro.NitroClient)

	update := nitro.DnsprofileUpdate{}
	unset := nitro.DnsprofileUnset{}

	updateFlag := false
	unsetFlag := false

	update.Dnsprofilename = d.Get("dnsprofilename").(string)
	unset.Dnsprofilename = d.Get("dnsprofilename").(string)

	if d.HasChange("dnsquerylogging") {
		updateFlag = true

		value := d.Get("dnsquerylogging").(string)
		update.Dnsquerylogging = value

		if value == "" {
			unsetFlag = true

			unset.Dnsquerylogging = true
		}

	}
	if d.HasChange("dnsanswerseclogging") {
		updateFlag = true

		value := d.Get("dnsanswerseclogging").(string)
		update.Dnsanswerseclogging = value

		if value == "" {
			unsetFlag = true

			unset.Dnsanswerseclogging = true
		}

	}
	if d.HasChange("dnsextendedlogging") {
		updateFlag = true

		value := d.Get("dnsextendedlogging").(string)
		update.Dnsextendedlogging = value

		if value == "" {
			unsetFlag = true

			unset.Dnsextendedlogging = true
		}

	}
	if d.HasChange("dnserrorlogging") {
		updateFlag = true

		value := d.Get("dnserrorlogging").(string)
		update.Dnserrorlogging = value

		if value == "" {
			unsetFlag = true

			unset.Dnserrorlogging = true
		}

	}
	if d.HasChange("cacherecords") {
		updateFlag = true

		value := d.Get("cacherecords").(string)
		update.Cacherecords = value

		if value == "" {
			unsetFlag = true

			unset.Cacherecords = true
		}

	}
	if d.HasChange("cachenegativeresponses") {
		updateFlag = true

		value := d.Get("cachenegativeresponses").(string)
		update.Cachenegativeresponses = value

		if value == "" {
			unsetFlag = true

			unset.Cachenegativeresponses = true
		}

	}
	if d.HasChange("dropmultiqueryrequest") {
		updateFlag = true

		value := d.Get("dropmultiqueryrequest").(string)
		update.Dropmultiqueryrequest = value

		if value == "" {
			unsetFlag = true

			unset.Dropmultiqueryrequest = true
		}

	}
	if d.HasChange("cacheecsresponses") {
		updateFlag = true

		value := d.Get("cacheecsresponses").(string)
		update.Cacheecsresponses = value

		if value == "" {
			unsetFlag = true

			unset.Cacheecsresponses = true
		}

	}
	key := get_dnsprofile_key(d)

	if updateFlag {
		if err := client.UpdateDnsprofile(update); err != nil {
			log.Print("Failed to update resource : ", err)

			return err
		}
	}

	if unsetFlag {
		if err := client.UnsetDnsprofile(unset); err != nil {
			log.Print("Failed to unset resource : ", err)

			return err
		}
	}

	if resource, err := client.GetDnsprofile(key); err != nil {
		log.Print("Failed to get resource : ", err)

		return err
	} else {
		set_dnsprofile(d, resource)
	}

	return nil
}

func delete_dnsprofile(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG]  netscaler-provider: In delete_dnsprofile")

	client := meta.(*nitro.NitroClient)

	resource := get_dnsprofile(d)
	key := resource.ToKey()

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
