package resources

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func NetscalerRewritepolicylabel() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//                Create:        create_rewritepolicylabel_func,
		//                Read:          read_rewritepolicylabel_func,
		//                Update:        update_rewritepolicylabel_func,
		//                Delete:        delete_rewritepolicylabel_func,
		Schema: map[string]*schema.Schema{
			"labelname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"transform": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}
