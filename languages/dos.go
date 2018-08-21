/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"dos", "bat", "cmd"}, `{"aliases":["bat","cmd"],"case_insensitive":true,"illegal":"\\/\\*","keywords":{"keyword":"if else goto for in do call exit not exist errorlevel defined equ neq lss leq gtr geq","built_in":"prn nul lpt3 lpt2 lpt1 con com4 com3 com2 com1 aux shift cd dir echo setlocal endlocal set pause copy append assoc at attrib break cacls cd chcp chdir chkdsk chkntfs cls cmd color comp compact convert date dir diskcomp diskcopy doskey erase fs find findstr format ftype graftabl help keyb label md mkdir mode more move path pause print popd pushd promt rd recover rem rename replace restore rmdir shiftsort start subst time title tree type ver verify vol ping net ipconfig taskkill xcopy ren del"},"contains":[{"className":"variable","begin":"%%[^ ]|%[^ ]+?%|![^ ]+?!"},{"className":"function","begin":"^\\s*[A-Za-z._?][A-Za-z0-9_$#@~.?]*(:|\\s+label)","end":"goto:eof","contains":[{"className":"title","begin":"([_a-zA-Z]\\w*\\.)*([_a-zA-Z]\\w*:)?[_a-zA-Z]\\w*","relevance":0},{"className":"comment","begin":"^\\s*@?rem\\b","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10}]},{"className":"number","begin":"\\b\\d+","relevance":0},{"className":"comment","begin":"^\\s*@?rem\\b","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10}]}`)
}
