/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"dts"}, `{"keywords":"","contains":[{"className":"class","begin":"/\\s*{","end":"};","relevance":10,"contains":[{"className":"variable","begin":"\\&[a-z\\d_]*\\b"},{"className":"meta-keyword","begin":"/[a-z][a-z\\d-]*/"},{"className":"symbol","begin":"^\\s*[a-zA-Z_][a-zA-Z\\d_]*:"},{"className":"class","begin":"[a-zA-Z_][a-zA-Z\\d_@]*\\s{","end":"[{;=]","returnBegin":true,"excludeEnd":true},{"className":"params","begin":"<","end":">","contains":[{"className":"number","variants":[{"begin":"\\b(\\d+(\\.\\d*)?|\\.\\d+)(u|U|l|L|ul|UL|f|F)"},{"begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)"}],"relevance":0},{"className":"variable","begin":"\\&[a-z\\d_]*\\b"}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"number","variants":[{"begin":"\\b(\\d+(\\.\\d*)?|\\.\\d+)(u|U|l|L|ul|UL|f|F)"},{"begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)"}],"relevance":0},{"className":"string","variants":[{"className":"string","begin":"((u8?|U)|L)?\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"begin":"(u8?|U)?R\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"'\\\\?.","end":"'","illegal":"."}]}]},{"className":"variable","begin":"\\&[a-z\\d_]*\\b"},{"className":"meta-keyword","begin":"/[a-z][a-z\\d-]*/"},{"className":"symbol","begin":"^\\s*[a-zA-Z_][a-zA-Z\\d_]*:"},{"className":"class","begin":"[a-zA-Z_][a-zA-Z\\d_@]*\\s{","end":"[{;=]","returnBegin":true,"excludeEnd":true},{"className":"params","begin":"<","end":">","contains":[{"className":"number","variants":[{"begin":"\\b(\\d+(\\.\\d*)?|\\.\\d+)(u|U|l|L|ul|UL|f|F)"},{"begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)"}],"relevance":0},{"className":"variable","begin":"\\&[a-z\\d_]*\\b"}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"number","variants":[{"begin":"\\b(\\d+(\\.\\d*)?|\\.\\d+)(u|U|l|L|ul|UL|f|F)"},{"begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)"}],"relevance":0},{"className":"string","variants":[{"className":"string","begin":"((u8?|U)|L)?\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"begin":"(u8?|U)?R\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"'\\\\?.","end":"'","illegal":"."}]},{"className":"meta","begin":"#","end":"$","keywords":{"meta-keyword":"if else elif endif define undef ifdef ifndef"},"contains":[{"begin":"\\\\\\n","relevance":0},{"beginKeywords":"include","end":"$","keywords":{"meta-keyword":"include"},"contains":[{"className":"meta-string","variants":[{"className":"string","begin":"((u8?|U)|L)?\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"begin":"(u8?|U)?R\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"'\\\\?.","end":"'","illegal":"."}]},{"className":"meta-string","begin":"<","end":">","illegal":"\\n"}]},{"className":"string","variants":[{"className":"string","begin":"((u8?|U)|L)?\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"begin":"(u8?|U)?R\"","end":"\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"begin":"'\\\\?.","end":"'","illegal":"."}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"[a-zA-Z]\\w*::","keywords":""}]}`)
}
