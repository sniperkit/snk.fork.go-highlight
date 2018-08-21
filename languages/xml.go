/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"xml", "html", "xhtml", "rss", "atom", "xjb", "xsd", "xsl", "plist"}, `{"aliases":["html","xhtml","rss","atom","xjb","xsd","xsl","plist"],"case_insensitive":true,"contains":[{"className":"meta","begin":"<!DOCTYPE","end":">","relevance":10,"contains":[{"begin":"\\[","end":"\\]"}]},{"className":"comment","begin":"<!--","end":"-->","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":10},{"begin":"<\\!\\[CDATA\\[","end":"\\]\\]>","relevance":10},{"begin":"<\\?(php)?","end":"\\?>","subLanguage":["php"],"contains":[{"begin":"/\\*","end":"\\*/","skip":true}]},{"className":"tag","begin":"<style(?=\\s|>|$)","end":">","keywords":{"name":"style"},"contains":[{"endsWithParent":true,"illegal":"<","relevance":0,"contains":[{"className":"attr","begin":"[A-Za-z0-9\\._:-]+","relevance":0},{"begin":"=\\s*","relevance":0,"contains":[{"className":"string","endsParent":true,"variants":[{"begin":"\"","end":"\""},{"begin":"'","end":"'"},{"begin":"[^\\s\"'=<>`+"`"+`]+"}]}]}]}],"starts":{"end":"</style>","returnEnd":true,"subLanguage":["css","xml"]}},{"className":"tag","begin":"<script(?=\\s|>|$)","end":">","keywords":{"name":"script"},"contains":[{"endsWithParent":true,"illegal":"<","relevance":0,"contains":[{"className":"attr","begin":"[A-Za-z0-9\\._:-]+","relevance":0},{"begin":"=\\s*","relevance":0,"contains":[{"className":"string","endsParent":true,"variants":[{"begin":"\"","end":"\""},{"begin":"'","end":"'"},{"begin":"[^\\s\"'=<>`+"`"+`]+"}]}]}]}],"starts":{"end":"</script>","returnEnd":true,"subLanguage":["actionscript","javascript","handlebars","xml"]}},{"className":"meta","variants":[{"begin":"<\\?xml","end":"\\?>","relevance":10},{"begin":"<\\?\\w+","end":"\\?>"}]},{"className":"tag","begin":"</?","end":"/?>","contains":[{"className":"name","begin":"[^\\/><\\s]+","relevance":0},{"endsWithParent":true,"illegal":"<","relevance":0,"contains":[{"className":"attr","begin":"[A-Za-z0-9\\._:-]+","relevance":0},{"begin":"=\\s*","relevance":0,"contains":[{"className":"string","endsParent":true,"variants":[{"begin":"\"","end":"\""},{"begin":"'","end":"'"},{"begin":"[^\\s\"'=<>`+"`"+`]+"}]}]}]}]}]}`)
}
