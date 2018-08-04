package utils

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Convert_set_to_string_array(set *schema.Set) []string {
	var results []string

	for _, item := range set.List() {
		results = append(results, item.(string))
	}

	return results
}

func Convert_set_to_int_array(set *schema.Set) []int {
	var results []int

	for _, item := range set.List() {
		results = append(results, item.(int))
	}

	return results
}
