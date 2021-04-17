package conversion

//supported format will be stored as key in SupportedFormat map,
//its value is corresponding PicType
var SupportedFormat map[string]PicType = map[string]PicType{
	"jpeg": 1,
	"jpg":  1,
	"png":  2,
}

//convert PicType to corresponding type's string
var PicType2String map[PicType]string = map[PicType]string{
	1: "jpeg",
	2: "png",
}
