package provider

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/sigurdkb/canvaslms-client-go"
)

func dataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRead,

		Schema: map[string]*schema.Schema{
			"base_url": {
				Type:     schema.TypeString,
				Required: true,
			},

			"token": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("CANVAS_AUTH_TOKEN", ""),
				Required:    true,
				Sensitive:   true,
			},

			"course_codes": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},

			"body": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	baseUrl := d.Get("base_url").(string)
	token := d.Get("token").(string)
	courseCodeI := d.Get("course_codes").([]interface{})

	courseCodes := make([]int, len(courseCodeI))
	for i, vI := range courseCodeI {
		courseCodes[i] = vI.(int)
	}

	client, err := canvaslms.NewClient(&baseUrl, &token)
	if err != nil {
		return append(diags, diag.Errorf("Error creating rest client: %s", err)...)
	}

	course, err := client.GetCourses(courseCodes)
	if err != nil {
		return append(diags, diag.Errorf("Error retreiving course: %s", err)...)
	}

	result, err := json.Marshal(course)
	if err != nil {
		return append(diags, diag.Errorf("error marshaling json: %s", err)...)
	}

	d.Set("body", string(result))

	// set ID as something more stable than time
	d.SetId(baseUrl)

	return diags
}
