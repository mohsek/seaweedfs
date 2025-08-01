// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.906
package app

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/seaweedfs/seaweedfs/weed/admin/dash"
)

func ClusterFilers(data dash.ClusterFilersData) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom\"><h1 class=\"h2\"><i class=\"fas fa-folder-open me-2\"></i>Filers</h1><div class=\"btn-toolbar mb-2 mb-md-0\"><div class=\"btn-group me-2\"><button type=\"button\" class=\"btn btn-sm btn-outline-primary\" onclick=\"exportFilers()\"><i class=\"fas fa-download me-1\"></i>Export</button></div></div></div><div id=\"filers-content\"><!-- Summary Cards --><div class=\"row mb-4\"><div class=\"col-xl-12 col-md-12 mb-4\"><div class=\"card border-left-primary shadow h-100 py-2\"><div class=\"card-body\"><div class=\"row no-gutters align-items-center\"><div class=\"col mr-2\"><div class=\"text-xs font-weight-bold text-primary text-uppercase mb-1\">Total Filers</div><div class=\"h5 mb-0 font-weight-bold text-gray-800\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", data.TotalFilers))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/app/cluster_filers.templ`, Line: 34, Col: 46}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div></div><div class=\"col-auto\"><i class=\"fas fa-folder-open fa-2x text-gray-300\"></i></div></div></div></div></div></div><!-- Filers Table --><div class=\"card shadow mb-4\"><div class=\"card-header py-3\"><h6 class=\"m-0 font-weight-bold text-primary\"><i class=\"fas fa-folder-open me-2\"></i>Filers</h6></div><div class=\"card-body\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(data.Filers) > 0 {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<div class=\"table-responsive\"><table class=\"table table-hover\" id=\"filersTable\"><thead><tr><th>Address</th><th>Version</th><th>Data Center</th><th>Rack</th><th>Created At</th><th>Actions</th></tr></thead> <tbody>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			for _, filer := range data.Filers {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<tr><td><a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 templ.SafeURL
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinURLErrs(templ.SafeURL(fmt.Sprintf("http://%s", filer.Address)))
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/app/cluster_filers.templ`, Line: 71, Col: 75}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\" target=\"_blank\" class=\"text-decoration-none\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var4 string
				templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(filer.Address)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/app/cluster_filers.templ`, Line: 72, Col: 27}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, " <i class=\"fas fa-external-link-alt ms-1 text-muted\"></i></a></td><td><span class=\"badge bg-light text-dark\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var5 string
				templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(filer.Version)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/app/cluster_filers.templ`, Line: 77, Col: 65}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "</span></td><td><span class=\"badge bg-light text-dark\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var6 string
				templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(filer.DataCenter)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/app/cluster_filers.templ`, Line: 80, Col: 68}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</span></td><td><span class=\"badge bg-light text-dark\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var7 string
				templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(filer.Rack)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/app/cluster_filers.templ`, Line: 83, Col: 62}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "</span></td><td>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				if !filer.CreatedAt.IsZero() {
					var templ_7745c5c3_Var8 string
					templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(filer.CreatedAt.Format("2006-01-02 15:04:05"))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/app/cluster_filers.templ`, Line: 87, Col: 59}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				} else {
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "<span class=\"text-muted\">N/A</span>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "</td><td><div class=\"btn-group btn-group-sm\" role=\"group\"><button type=\"button\" class=\"btn btn-outline-secondary btn-sm\" title=\"File Browser\" data-action=\"open-filer\" data-address=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var9 string
				templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(filer.Address)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/app/cluster_filers.templ`, Line: 94, Col: 149}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\"><i class=\"fas fa-folder-open\"></i></button></div></td></tr>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "</tbody></table></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "<div class=\"text-center py-5\"><i class=\"fas fa-folder-open fa-3x text-muted mb-3\"></i><h5 class=\"text-muted\">No Filers Found</h5><p class=\"text-muted\">No filer servers are currently available in the cluster.</p></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "</div></div><!-- Last Updated --><div class=\"row\"><div class=\"col-12\"><small class=\"text-muted\"><i class=\"fas fa-clock me-1\"></i> Last updated: ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(data.LastUpdated.Format("2006-01-02 15:04:05"))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `view/app/cluster_filers.templ`, Line: 119, Col: 67}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</small></div></div></div><!-- JavaScript for cluster filers functionality --><script>\n\tdocument.addEventListener('DOMContentLoaded', function() {\n\t\t// Handle filer action buttons\n\t\tdocument.addEventListener('click', function(e) {\n\t\t\tconst button = e.target.closest('[data-action]');\n\t\t\tif (!button) return;\n\t\t\t\n\t\t\tconst action = button.getAttribute('data-action');\n\t\t\tconst address = button.getAttribute('data-address');\n\t\t\t\n\t\t\tif (!address) return;\n\t\t\t\n\t\t\tswitch(action) {\n\t\t\t\tcase 'open-filer':\n\t\t\t\t\topenFilerBrowser(address);\n\t\t\t\t\tbreak;\n\t\t\t}\n\t\t});\n\t});\n\t\n\tfunction openFilerBrowser(address) {\n\t\t// Open file browser for specific filer\n\t\twindow.open('/files?filer=' + encodeURIComponent(address), '_blank');\n\t}\n\t\n\tfunction exportFilers() {\n\t\t// Simple CSV export of filers list\n\t\tconst rows = Array.from(document.querySelectorAll('#filersTable tbody tr')).map(row => {\n\t\t\tconst cells = row.querySelectorAll('td');\n\t\t\tif (cells.length > 1) {\n\t\t\t\treturn {\n\t\t\t\t\taddress: cells[0].textContent.trim(),\n\t\t\t\t\tversion: cells[1].textContent.trim(),\n\t\t\t\t\tdatacenter: cells[2].textContent.trim(),\n\t\t\t\t\track: cells[3].textContent.trim(),\n\t\t\t\t\tcreated: cells[4].textContent.trim()\n\t\t\t\t};\n\t\t\t}\n\t\t\treturn null;\n\t\t}).filter(row => row !== null);\n\t\t\n\t\tconst csvContent = \"data:text/csv;charset=utf-8,\" + \n\t\t\t\"Address,Version,Data Center,Rack,Created At\\n\" +\n\t\t\trows.map(r => '\"' + r.address + '\",\"' + r.version + '\",\"' + r.datacenter + '\",\"' + r.rack + '\",\"' + r.created + '\"').join(\"\\n\");\n\t\t\n\t\tconst encodedUri = encodeURI(csvContent);\n\t\tconst link = document.createElement(\"a\");\n\t\tlink.setAttribute(\"href\", encodedUri);\n\t\tlink.setAttribute(\"download\", \"filers.csv\");\n\t\tdocument.body.appendChild(link);\n\t\tlink.click();\n\t\tdocument.body.removeChild(link);\n\t}\n\t</script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
