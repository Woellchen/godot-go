package boneattachment

import (
	"log"
	"reflect"

	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
This node must be the child of a [Skeleton] node. You can then select a bone for this node to attach to. The BoneAttachment node will copy the transform of the selected bone.
*/
type BoneAttachment struct {
	Spatial
}

func (o *BoneAttachment) BaseClass() string {
	return "BoneAttachment"
}

/*
   Undocumented
*/
func (o *BoneAttachment) GetBoneName() gdnative.String {
	log.Println("Calling BoneAttachment.GetBoneName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_bone_name", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *BoneAttachment) SetBoneName(boneName gdnative.String) {
	log.Println("Calling BoneAttachment.SetBoneName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(boneName)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_bone_name", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   BoneAttachmentImplementer is an interface for BoneAttachment objects.
*/
type BoneAttachmentImplementer interface {
	Class
}
