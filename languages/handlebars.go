/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"handlebars", "hbs", "html.hbs", "html.handlebars"}, `{"aliases":["hbs","html.hbs","html.handlebars"],"case_insensitive":true,"subLanguage":["xml"],"contains":[{"className":"comment","begin":"{{!(--)?","end":"(--)?}}","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"template-tag","begin":"\\{\\{[#\\/]","end":"\\}\\}","contains":[{"className":"name","begin":"[a-zA-Z\\.-]+","keywords":{"builtin-name":"each in with if else unless bindattr action collection debugger log outlet template unbound view yield"},"starts":{"endsWithParent":true,"relevance":0,"contains":[{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]}]}}]},{"className":"template-variable","begin":"\\{\\{","end":"\\}\\}","keywords":{"builtin-name":"each in with if else unless bindattr action collection debugger log outlet template unbound view yield"}}]}`)
}
