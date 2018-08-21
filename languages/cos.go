/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"cos", "cos", "cls"}, `{"case_insensitive":true,"aliases":["cos","cls"],"keywords":"property parameter class classmethod clientmethod extends as break catch close continue do d|0 else elseif for goto halt hang h|0 if job j|0 kill k|0 lock l|0 merge new open quit q|0 read r|0 return set s|0 tcommit throw trollback try tstart use view while write w|0 xecute x|0 zkill znspace zn ztrap zwrite zw zzdump zzwrite print zbreak zinsert zload zprint zremove zsave zzprint mv mvcall mvcrt mvdim mvprint zquit zsync ascii","contains":[{"className":"number","begin":"\\b(\\d+(\\.\\d*)?|\\.\\d+)","relevance":0},{"className":"string","variants":[{"begin":"\"","end":"\"","contains":[{"begin":"\"\"","relevance":0}]}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":";","end":"$","relevance":0},{"className":"built_in","begin":"(?:\\$\\$?|\\.\\.)\\^?[a-zA-Z]+"},{"className":"built_in","begin":"\\$\\$\\$[a-zA-Z]+"},{"className":"built_in","begin":"%[a-z]+(?:\\.[a-z]+)*"},{"className":"symbol","begin":"\\^%?[a-zA-Z][\\w]*"},{"className":"keyword","begin":"##class|##super|#define|#dim"},{"begin":"&sql\\(","end":"\\)","excludeBegin":true,"excludeEnd":true,"subLanguage":["sql"]},{"begin":"&(js|jscript|javascript)<","end":">","excludeBegin":true,"excludeEnd":true,"subLanguage":["javascript"]},{"begin":"&html<\\s*<","end":">\\s*>","subLanguage":["xml"]}]}`)
}
