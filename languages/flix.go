/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"flix"}, `{"keywords":{"literal":"true false","keyword":"case class def else enum if impl import in lat rel index let match namespace switch type yield with"},"contains":[{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"string","begin":"'(.|\\\\[xXuU][a-zA-Z0-9]+)'"},{"className":"string","variants":[{"begin":"\"","end":"\""}]},{"className":"function","beginKeywords":"def","end":"[:={\\[(\\n;]","excludeEnd":true,"contains":[{"className":"title","begin":"[^0-9\\n\\t \"'(),.`+"`"+`{}\\[\\]:;][^\\n\\t \"'(),.`+"`"+`{}\\[\\]:;]+|[^0-9\\n\\t \"'(),.`+"`"+`{}\\[\\]:;=]"}]},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0}]}`)
}
