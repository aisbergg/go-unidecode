// Package table contains the transliteration tables and lookup logic.
package table

var tables [][]string

// Lookup returns the value of the table at the given index. An empty return
// value means that the rune cannot be transliterated.
func Lookup(r rune) string {
	section := r >> 8 // Chop off the last two hex digits
	var positionTbl []string
	switch {
	case section < 0x100:
		positionTbl = tables[section]
	case section >= 0x1d0 && section <= 0x1df:
		positionTbl = tables[section-0x1d0+0x100]
	case section >= 0x1f0 && section <= 0x1ff:
		positionTbl = tables[section-0x1f0+0x110]
	default:
		return ""
	}

	position := r % 256 // Last two hex digits
	if int(position) >= len(positionTbl) {
		return ""
	}
	trl := positionTbl[position]
	// an empty value means no transliteration available
	if trl == "" {
		return ""
	}
	return trl
}

var xXXX []string

func init() {
	// we try to have a table with contiguous sections that we can easily lookup by slice index
	tables = [][]string{
		x000, x001, x002, x003, x004, x005, x006, x007, xXXX, x009, x00a, x00b, x00c, x00d, x00e, x00f,
		x010, x011, x012, x013, x014, x015, x016, x017, x018, xXXX, xXXX, xXXX, xXXX, x01d, x01e, x01f,
		x020, x021, x022, x023, x024, x025, x026, x027, x028, x029, x02a, xXXX, x02c, xXXX, x02e, x02f,
		x030, x031, x032, x033, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX,
		xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, x04d, x04e, x04f,
		x050, x051, x052, x053, x054, x055, x056, x057, x058, x059, x05a, x05b, x05c, x05d, x05e, x05f,
		x060, x061, x062, x063, x064, x065, x066, x067, x068, x069, x06a, x06b, x06c, x06d, x06e, x06f,
		x070, x071, x072, x073, x074, x075, x076, x077, x078, x079, x07a, x07b, x07c, x07d, x07e, x07f,
		x080, x081, x082, x083, x084, x085, x086, x087, x088, x089, x08a, x08b, x08c, x08d, x08e, x08f,
		x090, x091, x092, x093, x094, x095, x096, x097, x098, x099, x09a, x09b, x09c, x09d, x09e, x09f,
		x0a0, x0a1, x0a2, x0a3, x0a4, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, x0ac, x0ad, x0ae, x0af,
		x0b0, x0b1, x0b2, x0b3, x0b4, x0b5, x0b6, x0b7, x0b8, x0b9, x0ba, x0bb, x0bc, x0bd, x0be, x0bf,
		x0c0, x0c1, x0c2, x0c3, x0c4, x0c5, x0c6, x0c7, x0c8, x0c9, x0ca, x0cb, x0cc, x0cd, x0ce, x0cf,
		x0d0, x0d1, x0d2, x0d3, x0d4, x0d5, x0d6, x0d7, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX,
		xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX,
		xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, x0f9, x0fa, x0fb, x0fc, x0fd, x0fe, x0ff,

		xXXX, xXXX, xXXX, xXXX, x1d4, x1d5, x1d6, x1d7, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX,

		xXXX, x1f1, xXXX, xXXX, xXXX, xXXX, x1f6, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX, xXXX,
	}
}
