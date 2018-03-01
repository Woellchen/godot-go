package image

import (
	"log"
	"reflect"
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
Native image datatype. Contains image data, which can be converted to a [Texture], and several functions to interact with it. The maximum width and height for an [code]Image[/code] is 16384 pixels.
*/
type Image struct {
	Resource
}

func (o *Image) BaseClass() string {
	return "Image"
}

/*
   Undocumented
*/
func (o *Image) X_GetData() *Dictionary {
	log.Println("Calling Image.X_GetData()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "_get_data", goArguments, "*Dictionary")

	returnValue := goRet.Interface().(*Dictionary)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *Image) X_SetData(data *Dictionary) {
	log.Println("Calling Image.X_SetData()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(data)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_set_data", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Alpha-blends [code]src_rect[/code] from [code]src[/code] image to this image at coordinates [code]dest[/code].
*/
func (o *Image) BlendRect(src *Image, srcRect *Rect2, dst *Vector2) {
	log.Println("Calling Image.BlendRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(src)
	goArguments[1] = reflect.ValueOf(srcRect)
	goArguments[2] = reflect.ValueOf(dst)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "blend_rect", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Alpha-blends [code]src_rect[/code] from [code]src[/code] image to this image using [code]mask[/code] image at coordinates [code]dst[/code]. Alpha channels are required for both [code]src[/code] and [code]mask[/code]. [code]dst[/code] pixels and [code]src[/code] pixels will blend if the corresponding mask pixel's alpha value is not 0. [code]src[/code] image and [code]mask[/code] image [b]must[/b] have the same size (width and height) but they can have different formats.
*/
func (o *Image) BlendRectMask(src *Image, mask *Image, srcRect *Rect2, dst *Vector2) {
	log.Println("Calling Image.BlendRectMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(src)
	goArguments[1] = reflect.ValueOf(mask)
	goArguments[2] = reflect.ValueOf(srcRect)
	goArguments[3] = reflect.ValueOf(dst)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "blend_rect_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Copies [code]src_rect[/code] from [code]src[/code] image to this image at coordinates [code]dst[/code].
*/
func (o *Image) BlitRect(src *Image, srcRect *Rect2, dst *Vector2) {
	log.Println("Calling Image.BlitRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(src)
	goArguments[1] = reflect.ValueOf(srcRect)
	goArguments[2] = reflect.ValueOf(dst)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "blit_rect", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Blits [code]src_rect[/code] area from [code]src[/code] image to this image at the coordinates given by [code]dst[/code]. [code]src[/code] pixel is copied onto [code]dst[/code] if the corresponding [code]mask[/code] pixel's alpha value is not 0. [code]src[/code] image and [code]mask[/code] image [b]must[/b] have the same size (width and height) but they can have different formats.
*/
func (o *Image) BlitRectMask(src *Image, mask *Image, srcRect *Rect2, dst *Vector2) {
	log.Println("Calling Image.BlitRectMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(src)
	goArguments[1] = reflect.ValueOf(mask)
	goArguments[2] = reflect.ValueOf(srcRect)
	goArguments[3] = reflect.ValueOf(dst)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "blit_rect_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes the image's mipmaps.
*/
func (o *Image) ClearMipmaps() {
	log.Println("Calling Image.ClearMipmaps()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "clear_mipmaps", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Compresses the image to use less memory. Can not directly access pixel data while the image is compressed. Returns error if the chosen compression mode is not available. See [code]COMPRESS_*[/code] constants.
*/
func (o *Image) Compress(mode gdnative.Int, source gdnative.Int, lossyQuality gdnative.Float) gdnative.Int {
	log.Println("Calling Image.Compress()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(mode)
	goArguments[1] = reflect.ValueOf(source)
	goArguments[2] = reflect.ValueOf(lossyQuality)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "compress", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Converts the image's format. See [code]FORMAT_*[/code] constants.
*/
func (o *Image) Convert(format gdnative.Int) {
	log.Println("Calling Image.Convert()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(format)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "convert", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Copies [code]src[/code] image to this image.
*/
func (o *Image) CopyFrom(src *Image) {
	log.Println("Calling Image.CopyFrom()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(src)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "copy_from", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Creates an empty image of given size and format. See [code]FORMAT_*[/code] constants. If [code]use_mipmaps[/code] is true then generate mipmaps for this image. See the [code]generate_mipmaps[/code] method.
*/
func (o *Image) Create(width gdnative.Int, height gdnative.Int, useMipmaps gdnative.Bool, format gdnative.Int) {
	log.Println("Calling Image.Create()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(width)
	goArguments[1] = reflect.ValueOf(height)
	goArguments[2] = reflect.ValueOf(useMipmaps)
	goArguments[3] = reflect.ValueOf(format)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "create", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Creates a new image of given size and format. See [code]FORMAT_*[/code] constants. Fills the image with the given raw data. If [code]use_mipmaps[/code] is true then generate mipmaps for this image. See the [code]generate_mipmaps[/code] method.
*/
func (o *Image) CreateFromData(width gdnative.Int, height gdnative.Int, useMipmaps gdnative.Bool, format gdnative.Int, data *PoolByteArray) {
	log.Println("Calling Image.CreateFromData()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 5, 5)
	goArguments[0] = reflect.ValueOf(width)
	goArguments[1] = reflect.ValueOf(height)
	goArguments[2] = reflect.ValueOf(useMipmaps)
	goArguments[3] = reflect.ValueOf(format)
	goArguments[4] = reflect.ValueOf(data)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "create_from_data", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Crops the image to the given [code]width[/code] and [code]height[/code]. If the specified size is larger than the current size, the extra area is filled with black pixels.
*/
func (o *Image) Crop(width gdnative.Int, height gdnative.Int) {
	log.Println("Calling Image.Crop()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(width)
	goArguments[1] = reflect.ValueOf(height)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "crop", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Decompresses the image if it is compressed. Returns an error if decompress function is not available.
*/
func (o *Image) Decompress() gdnative.Int {
	log.Println("Calling Image.Decompress()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "decompress", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns ALPHA_BLEND if the image has data for alpha values. Returns ALPHA_BIT if all the alpha values are below a certain threshold or the maximum value. Returns ALPHA_NONE if no data for alpha values is found.
*/
func (o *Image) DetectAlpha() gdnative.Int {
	log.Println("Calling Image.DetectAlpha()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "detect_alpha", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Stretches the image and enlarges it by a factor of 2. No interpolation is done.
*/
func (o *Image) ExpandX2Hq2X() {
	log.Println("Calling Image.ExpandX2Hq2X()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "expand_x2_hq2x", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Fills the image with a given [Color].
*/
func (o *Image) Fill(color *Color) {
	log.Println("Calling Image.Fill()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "fill", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Blends low-alpha pixels with nearby pixels.
*/
func (o *Image) FixAlphaEdges() {
	log.Println("Calling Image.FixAlphaEdges()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "fix_alpha_edges", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Flips the image horizontally.
*/
func (o *Image) FlipX() {
	log.Println("Calling Image.FlipX()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "flip_x", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Flips the image vertically.
*/
func (o *Image) FlipY() {
	log.Println("Calling Image.FlipY()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "flip_y", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Generates mipmaps for the image. Mipmaps are pre-calculated and lower resolution copies of the image. Mipmaps are automatically used if the image needs to be scaled down when rendered. This improves image quality and the performance of the rendering. Returns an error if the image is compressed, in a custom format or if the image's width/height is 0.
*/
func (o *Image) GenerateMipmaps() gdnative.Int {
	log.Println("Calling Image.GenerateMipmaps()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "generate_mipmaps", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the image's raw data.
*/
func (o *Image) GetData() *PoolByteArray {
	log.Println("Calling Image.GetData()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_data", goArguments, "*PoolByteArray")

	returnValue := goRet.Interface().(*PoolByteArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the image’s format. See [code]FORMAT_*[/code] constants.
*/
func (o *Image) GetFormat() gdnative.Int {
	log.Println("Calling Image.GetFormat()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_format", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the image's height.
*/
func (o *Image) GetHeight() gdnative.Int {
	log.Println("Calling Image.GetHeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_height", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the offset where the image's mipmap with index [code]mipmap[/code] is stored in the [code]data[/code] dictionary.
*/
func (o *Image) GetMipmapOffset(mipmap gdnative.Int) gdnative.Int {
	log.Println("Calling Image.GetMipmapOffset()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mipmap)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mipmap_offset", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the color of the pixel at [code](x, y)[/code] if the image is locked. If the image is unlocked it always returns a [Color] with the value [code](0, 0, 0, 1.0)[/code].
*/
func (o *Image) GetPixel(x gdnative.Int, y gdnative.Int) *Color {
	log.Println("Calling Image.GetPixel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(x)
	goArguments[1] = reflect.ValueOf(y)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_pixel", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a new image that is a copy of the image's area specified with [code]rect[/code].
*/
func (o *Image) GetRect(rect *Rect2) *Image {
	log.Println("Calling Image.GetRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rect)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_rect", goArguments, "*Image")

	returnValue := goRet.Interface().(*Image)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the image's size (width and height).
*/
func (o *Image) GetSize() *Vector2 {
	log.Println("Calling Image.GetSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_size", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a [Rect2] enclosing the visible portion of the image.
*/
func (o *Image) GetUsedRect() *Rect2 {
	log.Println("Calling Image.GetUsedRect()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_used_rect", goArguments, "*Rect2")

	returnValue := goRet.Interface().(*Rect2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the image's width.
*/
func (o *Image) GetWidth() gdnative.Int {
	log.Println("Calling Image.GetWidth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_width", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the image has generated mipmaps.
*/
func (o *Image) HasMipmaps() gdnative.Bool {
	log.Println("Calling Image.HasMipmaps()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "has_mipmaps", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the image is compressed.
*/
func (o *Image) IsCompressed() gdnative.Bool {
	log.Println("Calling Image.IsCompressed()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_compressed", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if the image has no data.
*/
func (o *Image) IsEmpty() gdnative.Bool {
	log.Println("Calling Image.IsEmpty()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_empty", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns [code]true[/code] if all the image's pixels have an alpha value of 0. Returns [code]false[/code] if any pixel has an alpha value higher than 0.
*/
func (o *Image) IsInvisible() gdnative.Bool {
	log.Println("Calling Image.IsInvisible()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_invisible", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Loads an image from file [code]path[/code].
*/
func (o *Image) Load(path gdnative.String) gdnative.Int {
	log.Println("Calling Image.Load()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(path)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "load", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *Image) LoadJpgFromBuffer(buffer *PoolByteArray) gdnative.Int {
	log.Println("Calling Image.LoadJpgFromBuffer()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(buffer)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "load_jpg_from_buffer", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *Image) LoadPngFromBuffer(buffer *PoolByteArray) gdnative.Int {
	log.Println("Calling Image.LoadPngFromBuffer()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(buffer)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "load_png_from_buffer", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Locks the data for writing access.
*/
func (o *Image) Lock() {
	log.Println("Calling Image.Lock()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "lock", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Converts the image's data to represent coordinates on a 3D plane. This is used when the image represents a normalmap. A normalmap can add lots of detail to a 3D surface without increasing the polygon count.
*/
func (o *Image) NormalmapToXy() {
	log.Println("Calling Image.NormalmapToXy()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "normalmap_to_xy", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Multiplies color values with alpha values. Resulting color values for a pixel are [code](color * alpha)/256[/code].
*/
func (o *Image) PremultiplyAlpha() {
	log.Println("Calling Image.PremultiplyAlpha()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "premultiply_alpha", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Resizes the image to the given [code]width[/code] and [code]height[/code]. New pixels are calculated using [code]interpolation[/code]. See [code]interpolation[/code] constants.
*/
func (o *Image) Resize(width gdnative.Int, height gdnative.Int, interpolation gdnative.Int) {
	log.Println("Calling Image.Resize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(width)
	goArguments[1] = reflect.ValueOf(height)
	goArguments[2] = reflect.ValueOf(interpolation)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "resize", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Resizes the image to the nearest power of 2 for the width and height. If [code]square[/code] is [code]true[/code] then set width and height to be the same.
*/
func (o *Image) ResizeToPo2(square gdnative.Bool) {
	log.Println("Calling Image.ResizeToPo2()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(square)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "resize_to_po2", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Saves the image as a PNG file to [code]path[/code].
*/
func (o *Image) SavePng(path gdnative.String) gdnative.Int {
	log.Println("Calling Image.SavePng()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(path)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "save_png", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Sets the [Color] of the pixel at [code](x, y)[/code] if the image is locked. Example: [codeblock] var img = Image.new() img.create(img_width, img_height, false, Image.FORMAT_RGBA8) img.lock() img.set_pixel(x, y, color) # Works img.unlock() img.set_pixel(x, y, color) # Does not have an effect [/codeblock]
*/
func (o *Image) SetPixel(x gdnative.Int, y gdnative.Int, color *Color) {
	log.Println("Calling Image.SetPixel()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(x)
	goArguments[1] = reflect.ValueOf(y)
	goArguments[2] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_pixel", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Shrinks the image by a factor of 2.
*/
func (o *Image) ShrinkX2() {
	log.Println("Calling Image.ShrinkX2()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "shrink_x2", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Converts the raw data from the sRGB colorspace to a linear scale.
*/
func (o *Image) SrgbToLinear() {
	log.Println("Calling Image.SrgbToLinear()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "srgb_to_linear", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Unlocks the data and prevents changes.
*/
func (o *Image) Unlock() {
	log.Println("Calling Image.Unlock()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "unlock", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   ImageImplementer is an interface for Image objects.
*/
type ImageImplementer interface {
	Class
}
