/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"stylus", "styl"}, `{"aliases":["styl"],"case_insensitive":false,"keywords":"if else for in","illegal":"(\\?|(\\bReturn\\b)|(\\bEnd\\b)|(\\bend\\b)|(\\bdef\\b)|;|#\\s|\\*\\s|===\\s|\\||%)","contains":[{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"comment","begin":"//","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]},{"className":"number","begin":"#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})"},{"begin":"\\.[a-zA-Z][a-zA-Z0-9_-]*[\\.\\s\\n\\[\\:,]","returnBegin":true,"contains":[{"className":"selector-class","begin":"\\.[a-zA-Z][a-zA-Z0-9_-]*"}]},{"begin":"\\#[a-zA-Z][a-zA-Z0-9_-]*[\\.\\s\\n\\[\\:,]","returnBegin":true,"contains":[{"className":"selector-id","begin":"\\#[a-zA-Z][a-zA-Z0-9_-]*"}]},{"begin":"\\b(a|abbr|address|article|aside|audio|b|blockquote|body|button|canvas|caption|cite|code|dd|del|details|dfn|div|dl|dt|em|fieldset|figcaption|figure|footer|form|h1|h2|h3|h4|h5|h6|header|hgroup|html|i|iframe|img|input|ins|kbd|label|legend|li|mark|menu|nav|object|ol|p|q|quote|samp|section|span|strong|summary|sup|table|tbody|td|textarea|tfoot|th|thead|time|tr|ul|var|video)[\\.\\s\\n\\[\\:,]","returnBegin":true,"contains":[{"className":"selector-tag","begin":"\\b[a-zA-Z][a-zA-Z0-9_-]*"}]},{"begin":"&?:?:\\b(after|before|first-letter|first-line|active|first-child|focus|hover|lang|link|visited)[\\.\\s\\n\\[\\:,]"},{"begin":"@(charset|css|debug|extend|font-face|for|import|include|media|mixin|page|warn|while)\\b"},{"className":"variable","begin":"\\$[a-zA-Z]\\w*"},{"className":"number","begin":"\\b\\d+(\\.\\d+)?(%|em|ex|ch|rem|vw|vh|vmin|vmax|cm|mm|in|pt|pc|px|deg|grad|rad|turn|s|ms|Hz|kHz|dpi|dpcm|dppx)?","relevance":0},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0},{"className":"function","begin":"^[a-zA-Z][a-zA-Z0-9_-]*\\(.*\\)","illegal":"[\\n]","returnBegin":true,"contains":[{"className":"title","begin":"\\b[a-zA-Z][a-zA-Z0-9_-]*"},{"className":"params","begin":"\\(","end":"\\)","contains":[{"className":"number","begin":"#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})"},{"className":"variable","begin":"\\$[a-zA-Z]\\w*"},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"number","begin":"\\b\\d+(\\.\\d+)?(%|em|ex|ch|rem|vw|vh|vmin|vmax|cm|mm|in|pt|pc|px|deg|grad|rad|turn|s|ms|Hz|kHz|dpi|dpcm|dppx)?","relevance":0},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]}]}]},{"className":"attribute","begin":"\\b(z-index|word-wrap|word-spacing|word-break|width|widows|white-space|visibility|vertical-align|unicode-bidi|transition-timing-function|transition-property|transition-duration|transition-delay|transition|transform-style|transform-origin|transform|top|text-underline-position|text-transform|text-shadow|text-rendering|text-overflow|text-indent|text-decoration-style|text-decoration-line|text-decoration-color|text-decoration|text-align-last|text-align|table-layout|tab-size|right|resize|quotes|position|pointer-events|perspective-origin|perspective|page-break-inside|page-break-before|page-break-after|padding-top|padding-right|padding-left|padding-bottom|padding|overflow-y|overflow-x|overflow-wrap|overflow|outline-width|outline-style|outline-offset|outline-color|outline|orphans|order|opacity|object-position|object-fit|normal|none|nav-up|nav-right|nav-left|nav-index|nav-down|min-width|min-height|max-width|max-height|mask|marks|margin-top|margin-right|margin-left|margin-bottom|margin|list-style-type|list-style-position|list-style-image|list-style|line-height|letter-spacing|left|justify-content|initial|inherit|ime-mode|image-resolution|image-rendering|image-orientation|icon|hyphens|height|font-weight|font-variant-ligatures|font-variant|font-style|font-stretch|font-size-adjust|font-size|font-language-override|font-kerning|font-feature-settings|font-family|font|float|flex-wrap|flex-shrink|flex-grow|flex-flow|flex-direction|flex-basis|flex|filter|empty-cells|display|direction|cursor|counter-reset|counter-increment|content|columns|column-width|column-span|column-rule-width|column-rule-style|column-rule-color|column-rule|column-gap|column-fill|column-count|color|clip-path|clip|clear|caption-side|break-inside|break-before|break-after|box-sizing|box-shadow|box-decoration-break|bottom|border-width|border-top-width|border-top-style|border-top-right-radius|border-top-left-radius|border-top-color|border-top|border-style|border-spacing|border-right-width|border-right-style|border-right-color|border-right|border-radius|border-left-width|border-left-style|border-left-color|border-left|border-image-width|border-image-source|border-image-slice|border-image-repeat|border-image-outset|border-image|border-color|border-collapse|border-bottom-width|border-bottom-style|border-bottom-right-radius|border-bottom-left-radius|border-bottom-color|border-bottom|border|background-size|background-repeat|background-position|background-origin|background-image|background-color|background-clip|background-attachment|background|backface-visibility|auto|animation-timing-function|animation-play-state|animation-name|animation-iteration-count|animation-fill-mode|animation-duration|animation-direction|animation-delay|animation|align-self|align-items|align-content)\\b","starts":{"end":";|$","contains":[{"className":"number","begin":"#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})"},{"className":"variable","begin":"\\$[a-zA-Z]\\w*"},{"className":"string","begin":"'","end":"'","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0}]},{"className":"string","begin":"\"","end":"\"","illegal":"\\n","contains":[{"begin":"\\\\[\\s\\S]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0},{"className":"subst","begin":"\\\\[abfnrtv]\\|\\\\x[0-9a-fA-F]*\\\\\\|%[-+# *.0-9]*[dioxXucsfeEgGp]","relevance":0}]},{"className":"number","begin":"\\b\\d+(\\.\\d+)?(%|em|ex|ch|rem|vw|vh|vmin|vmax|cm|mm|in|pt|pc|px|deg|grad|rad|turn|s|ms|Hz|kHz|dpi|dpcm|dppx)?","relevance":0},{"className":"number","begin":"\\b\\d+(\\.\\d+)?","relevance":0},{"className":"comment","begin":"/\\*","end":"\\*/","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}]}],"illegal":"\\.","relevance":0}}]}`)
}
