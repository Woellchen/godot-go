package vscrollbar

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*

 */
type VScrollBar struct {
	ScrollBar
}

func (o *VScrollBar) BaseClass() string {
	return "VScrollBar"
}

/*
   VScrollBarImplementer is an interface for VScrollBar objects.
*/
type VScrollBarImplementer interface {
	Class
}
