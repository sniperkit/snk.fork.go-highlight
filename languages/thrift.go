package languages
import "github.com/d4l3k/go-highlight/registry"
func init() {
  registry.Register("thrift", `{"keywords":{"keyword":"namespace const typedef struct enum service exception void oneway set list map required optional","built_in":"bool byte i16 i32 i64 double string binary","literal":"true false"},"contains":[{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"class","beginKeywords":"struct enum service exception","end":"\\{","illegal":"\\n","contains":[{"className":"title","begin":"[a-zA-Z]\\w*","relevance":0,"starts":{"endsWithParent":true,"excludeEnd":true}}]},{"begin":"\\b(set|list|map)\\s*<","end":">","keywords":"bool byte i16 i32 i64 double string binary","contains":["self"]}]}`)
}