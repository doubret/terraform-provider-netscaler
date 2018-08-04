package bindings

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerLbmonitorSslcertkeyBinding() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_lbmonitor_sslcertkey_binding_func,
		//                Read:          read_lbmonitor_sslcertkey_binding_func,
		//                Update:        update_lbmonitor_sslcertkey_binding_func,
		//                Delete:        delete_lbmonitor_sslcertkey_binding_func,
		Schema: map[string]*schema.Schema{
			"ca": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"certkeyname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"crlcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"monitorname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"ocspcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
