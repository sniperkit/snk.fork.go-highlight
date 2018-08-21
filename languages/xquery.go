/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"xquery", "xpath", "xq"}, `{"aliases":["xpath","xq"],"case_insensitive":false,"lexemes":"[a-zA-Z\\$][a-zA-Z0-9_:\\-]*","illegal":"(proc)|(abstract)|(extends)|(until)|(#)","keywords":{"keyword":"for let if while then else return where group by xquery encoding versionmodule namespace boundary-space preserve strip default collation base-uri orderingcopy-namespaces order declare import schema namespace function option in allowing emptyat tumbling window sliding window start when only end when previous next stable ascendingdescending empty greatest least some every satisfies switch case typeswitch try catch andor to union intersect instance of treat as castable cast map array delete insert intoreplace value rename copy modify update","literal":"false true xs:string xs:integer element item xs:date xs:datetime xs:float xs:double xs:decimal QName xs:anyURI xs:long xs:int xs:short xs:byte attribute"},"contains":[{"begin":"\\$[a-zA-Z0-9\\-]+"},{"className":"string","variants":[{"begin":"\"","end":"\"","contains":[{"begin":"\"\"","relevance":0}]},{"begin":"'","end":"'","contains":[{"begin":"''","relevance":0}]}]},{"className":"number","begin":"(\\b0[0-7_]+)|(\\b0x[0-9a-fA-F_]+)|(\\b[1-9][0-9_]*(\\.[0-9_]+)?)|[0_]\\b","relevance":0},{"className":"comment","begin":"\\(:","end":":\\)","relevance":10,"contains":[{"className":"doctag","begin":"@\\w+"}]},{"className":"meta","begin":"%\\w+"},{"begin":"{","end":"}","contains":[{"Ref":["contains"],"IsArray":true}]}]}`)
}
