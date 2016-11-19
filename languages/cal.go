package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("cal", `{"case_insensitive":true,"keywords":{"keyword":"div mod in and or not xor asserterror begin case do downto else end exit for if of repeat then to until while with var","literal":"false true"},"illegal":"\\/\\*","contains":[{"className":"string","begin":"'","end":"'","contains":[{"begin":"''"}]},{"className":"string","begin":"(#\\d+)+"},{"className":"number","begin":"\\b\\d+(\\.\\d+)?(DT|D|T)","relevance":0},{"className":"string","begin":"\"","end":"\""},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0},{"className":"class","begin":"OBJECT (Table|Form|Report|Dataport|Codeunit|XMLport|MenuSuite|Page|Query) (\\d+) ([^\\r\\n]+)","returnBegin":true,"contains":[{"className":"title","begin":"[a-zA-Z]\\w*","relevance":0},{"className":"function","beginKeywords":"procedure","end":"[:;]","keywords":"procedure|10","contains":[{"className":"title","begin":"[a-zA-Z]\\w*","relevance":0},{"className":"params","begin":"\\(","end":"\\)","keywords":"div mod in and or not xor asserterror begin case do downto else end exit for if of repeat then to until while with var","contains":[{"className":"string","begin":"'","end":"'","contains":[{"begin":"''"}]},{"className":"string","begin":"(#\\d+)+"}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"\\{","end":"\\}","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":0},{"className":"comment","begin":"\\(\\*","end":"\\*\\)","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10}]}]},{"className":"function","beginKeywords":"procedure","end":"[:;]","keywords":"procedure|10","contains":[{"className":"title","begin":"[a-zA-Z]\\w*","relevance":0},{"className":"params","begin":"\\(","end":"\\)","keywords":"div mod in and or not xor asserterror begin case do downto else end exit for if of repeat then to until while with var","contains":[{"className":"string","begin":"'","end":"'","contains":[{"begin":"''"}]},{"className":"string","begin":"(#\\d+)+"}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"\\{","end":"\\}","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":0},{"className":"comment","begin":"\\(\\*","end":"\\*\\)","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10}]}]}`)
}