/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"openscad", "scad"}, `{"aliases":["scad"],"keywords":{"keyword":"function module include use for intersection_for if else \\%","literal":"false true PI undef","built_in":"circle square polygon text sphere cube cylinder polyhedron translate rotate scale resize mirror multmatrix color offset hull minkowski union difference intersection abs sign sin cos tan acos asin atan atan2 floor round ceil ln log pow sqrt exp rands min max concat lookup str chr search version version_num norm cross parent_module echo import import_dxf dxf_linear_extrude linear_extrude rotate_extrude surface projection render children dxf_cross dxf_dim let assign"},"contains":[{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"number","begin":"\\b\\d+(\\.\\d+)?(e-?\\d+)?","relevance":0},{"className":"meta","keywords":{"meta-keyword":"include use"},"begin":"include|use <","end":">"},{"className":"string","begin":"\"","end":"\"","illegal":null,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"keyword","begin":"\\$(f[asn]|t|vp[rtd]|children)"},{"begin":"[*!#%]","relevance":0},{"className":"function","beginKeywords":"module function","end":"\\=|\\{","contains":[{"className":"params","begin":"\\(","end":"\\)","contains":[{"Ref":["contains","7","contains","0"]},{"className":"number","begin":"\\b\\d+(\\.\\d+)?(e-?\\d+)?","relevance":0},{"className":"string","begin":"\"","end":"\"","illegal":null,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"keyword","begin":"\\$(f[asn]|t|vp[rtd]|children)"},{"className":"literal","begin":"false|true|PI|undef"}]},{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0}]}]}`)
}
