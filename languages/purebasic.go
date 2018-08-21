/*
Sniperkit-Bot
- Status: analyzed
*/

package languages

import (
	"github.com/sniperkit/snk.fork.go-highlight/registry"
)

func init() {
	registry.Register([]string{"purebasic", "pb", "pbi"}, `{"aliases":["pb","pbi"],"keywords":"And As Break CallDebugger Case CompilerCase CompilerDefault CompilerElse CompilerEndIf CompilerEndSelect CompilerError CompilerIf CompilerSelect Continue Data DataSection EndDataSection Debug DebugLevel Default Define Dim DisableASM DisableDebugger DisableExplicit Else ElseIf EnableASM EnableDebugger EnableExplicit End EndEnumeration EndIf EndImport EndInterface EndMacro EndProcedure EndSelect EndStructure EndStructureUnion EndWith Enumeration Extends FakeReturn For Next ForEach ForEver Global Gosub Goto If Import ImportC IncludeBinary IncludeFile IncludePath Interface Macro NewList Not Or ProcedureReturn Protected Prototype PrototypeC Read ReDim Repeat Until Restore Return Select Shared Static Step Structure StructureUnion Swap To Wend While With XIncludeFile XOr Procedure ProcedureC ProcedureCDLL ProcedureDLL Declare DeclareC DeclareCDLL DeclareDLL","contains":[{"className":"comment","begin":";","end":"$","contains":[{"begin":"\\b(a|an|the|are|I'm|isn't|don't|doesn't|won't|but|just|should|pretty|simply|enough|gonna|going|wtf|so|such|will|you|your|like)\\b"},{"className":"doctag","begin":"(?:TODO|FIXME|NOTE|BUG|XXX):","relevance":0}],"relevance":0},{"className":"function","begin":"\\b(Procedure|Declare)(C|CDLL|DLL)?\\b","end":"\\(","excludeEnd":true,"returnBegin":true,"contains":[{"className":"keyword","begin":"(Procedure|Declare)(C|CDLL|DLL)?","excludeEnd":true},{"className":"type","begin":"\\.\\w*"},{"className":"title","begin":"[a-zA-Z_]\\w*","relevance":0}]},{"className":"string","begin":"(~)?\"","end":"\"","illegal":"\\n"},{"className":"symbol","begin":"#[a-zA-Z_]\\w*\\$?"}]}`)
}
