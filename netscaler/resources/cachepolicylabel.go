package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerCachepolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_cachepolicylabel_func,
		//                Read:          read_cachepolicylabel_func,
		//                Update:        update_cachepolicylabel_func,
		//                Delete:        delete_cachepolicylabel_func,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"evaluates": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
