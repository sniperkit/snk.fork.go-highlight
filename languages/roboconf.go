/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"roboconf", "graph", "instances"}, `{"aliases":["graph","instances"],"case_insensitive":true,"keywords":"import","contains":[{"begin":"^facet [a-zA-Z-_][^\\n{]+\\{","end":"}","keywords":"facet","contains":[{"className":"attribute","begin":"[a-zA-Z-_]+","end":"\\s*:","excludeEnd":true,"starts":{"end":";","relevance":0,"contains":[{"className":"variable","begin":"\\.[a-zA-Z-_]+"},{"className":"keyword","begin":"\\(optional\\)"}]}},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"^\\s*instance of [a-zA-Z-_][^\\n{]+\\{","end":"}","keywords":"name count channels instance-data instance-state instance of","illegal":"\\S","contains":[{"Ref":["contains","1"]},{"className":"attribute","begin":"[a-zA-Z-_]+","end":"\\s*:","excludeEnd":true,"starts":{"end":";","relevance":0,"contains":[{"className":"variable","begin":"\\.[a-zA-Z-_]+"},{"className":"keyword","begin":"\\(optional\\)"}]}},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"begin":"^[a-zA-Z-_][^\\n{]+\\{","end":"}","contains":[{"className":"attribute","begin":"[a-zA-Z-_]+","end":"\\s*:","excludeEnd":true,"starts":{"end":";","relevance":0,"contains":[{"className":"variable","begin":"\\.[a-zA-Z-_]+"},{"className":"keyword","begin":"\\(optional\\)"}]}},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}]}`)
}
