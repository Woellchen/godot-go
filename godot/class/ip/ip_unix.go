package ip

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
Unix-specific implementation of IP support functions. See [IP].
*/
type IP_Unix struct {
	ip
}

func (o *IP_Unix) BaseClass() string {
	return "IP_Unix"
}

/*
   IP_UnixImplementer is an interface for IP_Unix objects.
*/
type IP_UnixImplementer interface {
	Class
}
