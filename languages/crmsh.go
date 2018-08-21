/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"crmsh", "crm", "pcmk"}, `{"aliases":["crm","pcmk"],"case_insensitive":true,"keywords":{"keyword":"params meta operations op rule attributes utilization read write deny defined not_defined in_range date spec in ref reference attribute type xpath version and or lt gt tag lte gte eq ne \\ number string","literal":"Master Started Slave Stopped start promote demote stop monitor true false"},"contains":[{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"beginKeywords":"node","starts":{"end":"\\s*([\\w_-]+:)?","starts":{"className":"title","end":"\\s*[\\$\\w_][\\w_-]*"}}},{"beginKeywords":"primitive rsc_template","starts":{"className":"title","end":"\\s*[\\$\\w_][\\w_-]*","starts":{"end":"\\s*@?[\\w_][\\w_\\.:-]*"}}},{"begin":"\\b(group|clone|ms|master|location|colocation|order|fencing_topology|rsc_ticket|acl_target|acl_group|user|role|tag|xml)\\s+","keywords":"group clone ms master location colocation order fencing_topology rsc_ticket acl_target acl_group user role tag xml","starts":{"className":"title","end":"[\\$\\w_][\\w_-]*"}},{"beginKeywords":"property rsc_defaults op_defaults","starts":{"className":"title","end":"\\s*([\\w_-]+:)?"}},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"meta","begin":"(ocf|systemd|service|lsb):[\\w_:-]+","relevance":0},{"className":"number","begin":"\\b\\d+(\\.\\d+)?(ms|s|h|m)?","relevance":0},{"className":"literal","begin":"[-]?(infinity|inf)","relevance":0},{"className":"attr","begin":"([A-Za-z\\$_\\#][\\w_-]+)=","relevance":0},{"className":"tag","begin":"</?","end":"/?>","relevance":0}]}`)
}
