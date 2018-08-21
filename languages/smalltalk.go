/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"smalltalk", "st"}, `{"aliases":["st"],"keywords":"self super nil true false thisContext","contains":[{"className":"comment","begin":"\"","end":"\"","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"type","begin":"\\b[A-Z][A-Za-z0-9_]*","relevance":0},{"begin":"[a-z][a-zA-Z0-9_]*:","relevance":0},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"className":"symbol","begin":"#[a-zA-Z_]\\w*"},{"className":"string","begin":"\\$.{1}"},{"begin":"\\|[ ]*[a-z][a-zA-Z0-9_]*([ ]+[a-z][a-zA-Z0-9_]*)*[ ]*\\|","returnBegin":true,"end":"\\|","illegal":"\\S","contains":[{"begin":"(\\|[ ]*)?[a-z][a-zA-Z0-9_]*"}]},{"begin":"\\#\\(","end":"\\)","contains":[{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\\$.{1}"},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"className":"symbol","begin":"#[a-zA-Z_]\\w*"}]}]}`)
}
