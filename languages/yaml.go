/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"yaml", "yml", "YAML", "yaml"}, `{"case_insensitive":true,"aliases":["yml","YAML","yaml"],"contains":[{"className":"attr","variants":[{"begin":"^[ \\-]*[a-zA-Z_][\\w\\-]*:"},{"begin":"^[ \\-]*\"[a-zA-Z_][\\w\\-]*\":"},{"begin":"^[ \\-]*'[a-zA-Z_][\\w\\-]*':"}]},{"className":"meta","begin":"^---s*$","relevance":10},{"className":"string","begin":"[\\|>] *$","returnEnd":true,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"template-variable","variants":[{"begin":"{{","end":"}}"},{"begin":"%{","end":"}"}]}],"end":"^[ \\-]*[a-zA-Z_][\\w\\-]*:"},{"begin":"<%[%=-]?","end":"[%-]?%>","subLanguage":["ruby"],"excludeBegin":true,"excludeEnd":true,"relevance":0},{"className":"type","begin":"!![a-zA-Z_]\\w*"},{"className":"meta","begin":"&[a-zA-Z_]\\w*$"},{"className":"meta","begin":"\\*[a-zA-Z_]\\w*$"},{"className":"bullet","begin":"^ *-","relevance":0},{"className":"string","relevance":0,"variants":[{"begin":"'","end":"'"},{"begin":"\"","end":"\""}],"contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"template-variable","variants":[{"begin":"{{","end":"}}"},{"begin":"%{","end":"}"}]}]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0}],"keywords":{"literal":"{ } true false yes no Yes No True False null"}}`)
}
