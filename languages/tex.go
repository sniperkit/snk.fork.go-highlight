/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"tex"}, `{"contains":[{"className":"tag","begin":"\\\\","relevance":0,"contains":[{"className":"name","variants":[{"begin":"[a-zA-Zа-яА-я]+[*]?"},{"begin":"[^a-zA-Zа-яА-я0-9]"}],"starts":{"endsWithParent":true,"relevance":0,"contains":[{"className":"string","variants":[{"begin":"\\[","end":"\\]"},{"begin":"\\{","end":"\\}"}]},{"begin":"\\s*=\\s*","endsWithParent":true,"relevance":0,"contains":[{"className":"number","begin":"-?\\d*\\.?\\d+(pt|pc|mm|cm|in|dd|cc|ex|em)?"}]}]}}]},{"className":"formula","contains":[{"className":"tag","begin":"\\\\","relevance":0,"contains":[{"className":"name","variants":[{"begin":"[a-zA-Zа-яА-я]+[*]?"},{"begin":"[^a-zA-Zа-яА-я0-9]"}],"starts":{"endsWithParent":true,"relevance":0,"contains":[{"className":"string","variants":[{"begin":"\\[","end":"\\]"},{"begin":"\\{","end":"\\}"}]},{"begin":"\\s*=\\s*","endsWithParent":true,"relevance":0,"contains":[{"className":"number","begin":"-?\\d*\\.?\\d+(pt|pc|mm|cm|in|dd|cc|ex|em)?"}]}]}}]}],"relevance":0,"variants":[{"begin":"\\$\\$","end":"\\$\\$"},{"begin":"\\$","end":"\\$"}]},{"className":"comment","begin":"%","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":0}]}`)
}
