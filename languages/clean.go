/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"clean", "clean", "icl", "dcl"}, `{"aliases":["clean","icl","dcl"],"keywords":{"keyword":"if let in with where case of class instance otherwise implementation definition system module from import qualified as special code inline foreign export ccall stdcall generic derive infix infixl infixr","literal":"True False"},"contains":[{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"begin":"->|<-[|:]?|::|#!?|>>=|\\{\\||\\|\\}|:==|=:|\\.\\.|<>|`+"`"+`"}]}`)
}
