/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"nginx", "nginxconf"}, `{"aliases":["nginxconf"],"contains":[{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"begin":"[a-zA-Z_]\\w*\\s+{","returnBegin":true,"end":"{","contains":[{"className":"section","begin":"[a-zA-Z_]\\w*"}],"relevance":0},{"begin":"[a-zA-Z_]\\w*\\s","end":";|{","returnBegin":true,"contains":[{"className":"attribute","begin":"[a-zA-Z_]\\w*","starts":{"endsWithParent":true,"lexemes":"[a-z/_]+","keywords":{"literal":"on off yes no true false none blocked debug info notice warn error crit select break last permanent redirect kqueue rtsig epoll poll /dev/poll"},"relevance":0,"illegal":"=>","contains":[{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"string","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"variable","variants":[{"begin":"\\$\\d+"},{"begin":"\\$\\{","end":"}"},{"begin":"[\\$\\@][a-zA-Z_]\\w*"}]}],"variants":[{"begin":"\"","end":"\""},{"begin":"'","end":"'"}]},{"begin":"([a-z]+):/","end":"\\s","endsWithParent":true,"excludeEnd":true,"contains":[{"className":"variable","variants":[{"begin":"\\$\\d+"},{"begin":"\\$\\{","end":"}"},{"begin":"[\\$\\@][a-zA-Z_]\\w*"}]}]},{"className":"regexp","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"variable","variants":[{"begin":"\\$\\d+"},{"begin":"\\$\\{","end":"}"},{"begin":"[\\$\\@][a-zA-Z_]\\w*"}]}],"variants":[{"begin":"\\s\\^","end":"\\s|{|;","returnEnd":true},{"begin":"~\\*?\\s+","end":"\\s|{|;","returnEnd":true},{"begin":"\\*(\\.[a-z\\-]+)+"},{"begin":"([a-z\\-]+\\.)+\\*"}]},{"className":"number","begin":"\\b\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}(:\\d{1,5})?\\b"},{"className":"number","begin":"\\b\\d+[kKmMgGdshdwy]*\\b","relevance":0},{"className":"variable","variants":[{"begin":"\\$\\d+"},{"begin":"\\$\\{","end":"}"},{"begin":"[\\$\\@][a-zA-Z_]\\w*"}]}]}}],"relevance":0}],"illegal":"[^\\s\\}]"}`)
}
