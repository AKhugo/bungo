package common

// declare constant
const (
	VERSION = "0.0.1"
	ERROR= "error encoding file"
	ERROR_ENDCODING_FILE = "encoding file"
	ERROR_DECODING_FILE = "decoding file"
)


var ExtMap = map[string]string{

	// document files
	"image/jpeg": ".jpg",
	"image/png":  ".png",
	"image/gif":  ".gif",
	"image/webp": ".webp",
	"image/tiff": ".tiff",
	"image/avif": ".avif",

	// text files
	"text/plain": ".txt",
	"text/html": ".html",
	"text/css": ".css",
	"text/csv": ".csv",

	"application/msword": ".doc",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": ".docx",


}


// var imagesExt = []string{"jpg", "jpeg", "png", "gif", "bmp", "svg", "webp", "tiff", "ico"};
// var documentEx = []string{"docx", "txt", "xlx", "xls", "xlsx", "csv"};

// throw error normalisation
