package globalvar

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var OciFrameworkDataSources []func() datasource.DataSource
var OciFrameworkResources []func() resource.Resource
