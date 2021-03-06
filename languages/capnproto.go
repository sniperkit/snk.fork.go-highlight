/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"capnproto", "capnp"}, `{"aliases":["capnp"],"keywords":{"keyword":"struct enum interface union group import using const annotation extends in of on as with from fixed","built_in":"Void Bool Int8 Int16 Int32 Int64 UInt8 UInt16 UInt32 UInt64 Float32 Float64 Text Data AnyPointer AnyStruct Capability List","literal":"true false"},"contains":[{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"meta","begin":"@0x[\\w\\d]{16};","illegal":"\\n"},{"className":"symbol","begin":"@\\d+\\b"},{"className":"class","beginKeywords":"struct enum","end":"\\{","illegal":"\\n","contains":[{"className":"title","begin":"[a-zA-Z]\\w*","relevance":0,"starts":{"endsWithParent":true,"excludeEnd":true}}]},{"className":"class","beginKeywords":"interface","end":"\\{","illegal":"\\n","contains":[{"className":"title","begin":"[a-zA-Z]\\w*","relevance":0,"starts":{"endsWithParent":true,"excludeEnd":true}}]}]}`)
}
