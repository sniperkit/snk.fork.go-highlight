/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"scilab", "sci"}, `{"aliases":["sci"],"lexemes":"%?\\w+","keywords":{"keyword":"abort break case clear catch continue do elseif else endfunction end for function global if pause return resume select try then while","literal":"%f %F %t %T %pi %eps %inf %nan %e %i %z %s","built_in":"abs and acos asin atan ceil cd chdir clearglobal cosh cos cumprod deff disp error exec execstr exists exp eye gettext floor fprintf fread fsolve imag isdef isempty isinfisnan isvector lasterror length load linspace list listfiles log10 log2 log max min msprintf mclose mopen ones or pathconvert poly printf prod pwd rand real round sinh sin size gsort sprintf sqrt strcat strcmps tring sum system tanh tan type typename warning zeros matrix"},"illegal":"(\"|#|/\\*|\\s+/\\w+)","contains":[{"className":"function","beginKeywords":"function","end":"$","contains":[{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0},{"className":"params","begin":"\\(","end":"\\)"}]},{"begin":"[a-zA-Z_][a-zA-Z_0-9]*('+[\\.']*|[\\.']+)","end":"","relevance":0},{"begin":"\\[","end":"\\]'*[\\.']*","relevance":0,"contains":[{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"className":"string","begin":"'|\"","end":"'|\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"begin":"''"}]}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"number","begin":"(-?)(\\b0[xX][a-fA-F0-9]+|(\\b\\d+(\\.\\d*)?|\\.\\d+)([eE][-+]?\\d+)?)","relevance":0},{"className":"string","begin":"'|\"","end":"'|\"","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"begin":"''"}]}]}`)
}
