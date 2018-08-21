/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"accesslog"}, `{"contains":[{"className":"number","begin":"\\b\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}(:\\d{1,5})?\\b"},{"className":"number","begin":"\\b\\d+\\b","relevance":0},{"className":"string","begin":"\"(GET|POST|HEAD|PUT|DELETE|CONNECT|OPTIONS|PATCH|TRACE)","end":"\"","keywords":"GET POST HEAD PUT DELETE CONNECT OPTIONS PATCH TRACE","illegal":"\\n","relevance":10},{"className":"string","begin":"\\[","end":"\\]","illegal":"\\n"},{"className":"string","begin":"\"","end":"\"","illegal":"\\n"}]}`)
}
