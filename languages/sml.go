/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"sml", "ml"}, `{"aliases":["ml"],"keywords":{"keyword":"abstype and andalso as case datatype do else end eqtype exception fn fun functor handle if in include infix infixr let local nonfix of op open orelse raise rec sharing sig signature struct structure then type val with withtype where while","built_in":"array bool char exn int list option order real ref string substring vector unit word","literal":"true false NONE SOME LESS EQUAL GREATER nil"},"illegal":"\\/\\/|>>","lexemes":"[a-z_]\\w*!?","contains":[{"className":"literal","begin":"\\[(\\|\\|)?\\]|\\(\\)","relevance":0},{"className":"comment","begin":"\\(\\*","end":"\\*\\)","contains":[{"Ref":["contains","1"]},{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"symbol","begin":"'[A-Za-z_](?!')[\\w']*"},{"className":"type","begin":"`+"`"+`[A-Z][\\w']*"},{"className":"type","begin":"\\b[A-Z][\\w']*","relevance":0},{"begin":"[a-z_]\\w*'[\\w']*"},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}],"relevance":0},{"className":"string","begin":"\"","end":"\"","illegal":null,"contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"number","begin":"\\b(0[xX][a-fA-F0-9_]+[Lln]?|0[oO][0-7_]+[Lln]?|0[bB][01_]+[Lln]?|[0-9][0-9_]*([Lln]|(\\.[0-9_]*)?([eE][-+]?[0-9_]+)?)?)","relevance":0},{"begin":"[-=]>"}]}`)
}
