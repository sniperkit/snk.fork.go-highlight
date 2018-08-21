/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"erlang-repl"}, `{"keywords":{"built_in":"spawn spawn_link self","keyword":"after and andalso|10 band begin bnot bor bsl bsr bxor case catch cond div end fun if let not of or orelse|10 query receive rem try when xor"},"contains":[{"className":"meta","begin":"^[0-9]+> ","relevance":10},{"className":"comment","begin":"%","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"number","begin":"\\b(\\d+#[a-fA-F0-9]+|\\d+(\\.\\d+)?([eE][-+]?\\d+)?)","relevance":0},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"begin":"\\?(::)?([A-Z]\\w*(::)?)+"},{"begin":"->"},{"begin":"ok"},{"begin":"!"},{"begin":"(\\b[a-z'][a-zA-Z0-9_']*:[a-z'][a-zA-Z0-9_']*)|(\\b[a-z'][a-zA-Z0-9_']*)","relevance":0},{"begin":"[A-Z][a-zA-Z0-9_']*","relevance":0}]}`)
}
