/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"nix", "nixos"}, `{"aliases":["nixos"],"keywords":{"keyword":"rec with let in inherit assert if else then","literal":"true false or and null","built_in":"import abort baseNameOf dirOf isNull builtins map removeAttrs throw toString derivation"},"contains":[{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0},{"className":"comment","begin":"#","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"string","contains":[{"className":"subst","begin":"\\$\\{","end":"}","keywords":{"keyword":"rec with let in inherit assert if else then","literal":"true false or and null","built_in":"import abort baseNameOf dirOf isNull builtins map removeAttrs throw toString derivation"},"contains":[{"Ref":["contains"],"IsArray":true}]}],"variants":[{"begin":"''","end":"''"},{"begin":"\"","end":"\""}]},{"begin":"[a-zA-Z0-9-_]+(\\s*=)","returnBegin":true,"relevance":0,"contains":[{"className":"attr","begin":"\\S+"}]}]}`)
}
