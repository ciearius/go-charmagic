package unicode

var (
	UTF32_BE_BOM = []byte{0x00, 0x00, 0xFE, 0xFF}
	UTF32_LE_BOM = []byte{0xFF, 0xFE, 0x00, 0x00}
)

// TODO: add support for utf-32
