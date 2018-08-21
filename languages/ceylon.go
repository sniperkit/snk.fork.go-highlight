/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"ceylon"}, `{"keywords":{"keyword":"assembly module package import alias class interface object given value assign void function new of extends satisfies abstracts in out return break continue throw assert dynamic if else switch case for while try catch finally then let this outer super is exists nonempty shared abstract formal default actual variable late native deprecatedfinal sealed annotation suppressWarnings small","meta":"doc by license see throws tagged"},"illegal":"\\$[^01]|#[^0-9a-fA-F]","contains":[{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"Ref":["contains","1"]},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"meta","begin":"@[a-z]\\w*(?:\\:\"[^\"]*\")?"},{"className":"string","begin":"\"\"\"","end":"\"\"\"","relevance":10},{"className":"string","begin":"\"","end":"\"","contains":[{"className":"subst","excludeBegin":true,"excludeEnd":true,"begin":"`+"`"+``+"`"+`","end":"`+"`"+``+"`"+`","keywords":"assembly module package import alias class interface object given value assign void function new of extends satisfies abstracts in out return break continue throw assert dynamic if else switch case for while try catch finally then let this outer super is exists nonempty","relevance":10,"contains":[{"className":"string","begin":"\"\"\"","end":"\"\"\"","relevance":10},{"Ref":["contains","4"]},{"className":"string","begin":"'","end":"'"},{"className":"number","begin":"#[0-9a-fA-F_]+|\\$[01_]+|[0-9_]+(?:\\.[0-9_](?:[eE][+-]?\\d+)?)?[kMGTPmunpf]?","relevance":0}]}]},{"className":"string","begin":"'","end":"'"},{"className":"number","begin":"#[0-9a-fA-F_]+|\\$[01_]+|[0-9_]+(?:\\.[0-9_](?:[eE][+-]?\\d+)?)?[kMGTPmunpf]?","relevance":0}]}`)
}
