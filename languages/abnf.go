/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"abnf"}, `{"illegal":"[!@#$^&',?+~`+"`"+`|:]","keywords":"ALPHA BIT CHAR CR CRLF CTL DIGIT DQUOTE HEXDIG HTAB LF LWSP OCTET SP VCHAR WSP","contains":[{"begin":"^[a-zA-Z][a-zA-Z0-9-]*\\s*=","returnBegin":true,"end":"=","relevance":0,"contains":[{"className":"attribute","begin":"^[a-zA-Z][a-zA-Z0-9-]*"}]},{"className":"comment","begin":";","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"symbol","begin":"%b[0-1]+(-[0-1]+|(\\.[0-1]+)+){0,1}"},{"className":"symbol","begin":"%d[0-9]+(-[0-9]+|(\\.[0-9]+)+){0,1}"},{"className":"symbol","begin":"%x[0-9A-F]+(-[0-9A-F]+|(\\.[0-9A-F]+)+){0,1}"},{"className":"symbol","begin":"%[si]"},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0}]}`)
}
