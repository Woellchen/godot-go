package visualscript

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
Undocumented
*/
type VisualScriptSelf struct {
	VisualScriptNode
}

func (o *VisualScriptSelf) BaseClass() string {
	return "VisualScriptSelf"
}

/*
   VisualScriptSelfImplementer is an interface for VisualScriptSelf objects.
*/
type VisualScriptSelfImplementer interface {
	Class
}
