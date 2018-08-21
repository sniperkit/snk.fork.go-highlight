/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"asciidoc", "adoc"}, `{"aliases":["adoc"],"contains":[{"className":"comment","begin":"^/{4,}\\n","end":"\\n/{4,}$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"className":"comment","begin":"^//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":0},{"className":"title","begin":"^\\.\\w.*$"},{"begin":"^[=\\*]{4,}\\n","end":"\\n^[=\\*]{4,}$","relevance":10},{"className":"section","relevance":10,"variants":[{"begin":"^(={1,5}) .+?( \\1)?$"},{"begin":"^[^\\[\\]\\n]+?\\n[=\\-~\\^\\+]{2,}$"}]},{"className":"meta","begin":"^:.+?:","end":"\\s","excludeEnd":true,"relevance":10},{"className":"meta","begin":"^\\[.+?\\]$","relevance":0},{"className":"quote","begin":"^_{4,}\\n","end":"\\n_{4,}$","relevance":10},{"className":"code","begin":"^[\\-\\.]{4,}\\n","end":"\\n[\\-\\.]{4,}$","relevance":10},{"begin":"^\\+{4,}\\n","end":"\\n\\+{4,}$","contains":[{"begin":"<","end":">","subLanguage":["xml"],"relevance":0}],"relevance":10},{"className":"bullet","begin":"^(\\*+|\\-+|\\.+|[^\\n]+?::)\\s+"},{"className":"symbol","begin":"^(NOTE|TIP|IMPORTANT|WARNING|CAUTION):\\s+","relevance":10},{"className":"strong","begin":"\\B\\*(?![\\*\\s])","end":"(\\n{2}|\\*)","contains":[{"begin":"\\\\*\\w","relevance":0}]},{"className":"emphasis","begin":"\\B'(?!['\\s])","end":"(\\n{2}|')","contains":[{"begin":"\\\\'\\w","relevance":0}],"relevance":0},{"className":"emphasis","begin":"_(?![_\\s])","end":"(\\n{2}|_)","relevance":0},{"className":"string","variants":[{"begin":"`+"`"+``+"`"+`.+?''"},{"begin":"`+"`"+`.+?'"}]},{"className":"code","begin":"(`+"`"+`.+?`+"`"+`|\\+.+?\\+)","relevance":0},{"className":"code","begin":"^[ \\t]","end":"$","relevance":0},{"begin":"^'{3,}[ \\t]*$","relevance":10},{"begin":"(link:)?(http|https|ftp|file|irc|image:?):\\S+\\[.*?\\]","returnBegin":true,"contains":[{"begin":"(link|image:?):","relevance":0},{"className":"link","begin":"\\w","end":"[^\\[]+","relevance":0},{"className":"string","begin":"\\[","end":"\\]","excludeBegin":true,"excludeEnd":true,"relevance":0}],"relevance":10}]}`)
}
