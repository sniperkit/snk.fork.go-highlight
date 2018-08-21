/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"ini", "toml"}, `{"aliases":["toml"],"case_insensitive":true,"illegal":"\\S","contains":[{"className":"comment","begin":";","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"section","begin":"^\\s*\\[+","end":"\\]+"},{"begin":"^[a-z0-9\\[\\]_-]+\\s*=\\s*","end":"$","returnBegin":true,"contains":[{"className":"attr","begin":"[a-z0-9\\[\\]_-]+"},{"begin":"=","endsWithParent":true,"relevance":0,"contains":[{"className":"literal","begin":"\\bon|off|true|false|yes|no\\b"},{"className":"variable","variants":[{"begin":"\\$[\\w\\d\"][\\w\\d_]*"},{"begin":"\\$\\{(.*?)}"}]},{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}],"variants":[{"begin":"'''","end":"'''","relevance":10},{"begin":"\"\"\"","end":"\"\"\"","relevance":10},{"begin":"\"","end":"\""},{"begin":"'","end":"'"}]},{"className":"number","begin":"([\\+\\-]+)?[\\d]+_[\\d_]+"},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0}]}]}]}`)
}
