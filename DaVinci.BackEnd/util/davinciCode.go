package util

import (
	"strconv"
	"strings"
)

//DavinciCode , Rutina de encriptacion simple
type DavinciCode struct {
	patternHash  string
	patternToken string
}

var universo = "0123456789abcdefghij"

/*

MODO DE USO

	dcode := util.DavinciCode{}
	x := dcode.Encript("El codigo de Davinci metodo de encriptacion Tentatedecimal")
	fmt.Println("encript =", x)
	y := dcode.Decript(x)
	fmt.Println("decript =", y)

*/

//hexToDec , Retorna el Decimal de un HEX
func (d *DavinciCode) hexToDec(hex string) int64 {
	var result int64
	hex = strings.ToLower(hex)
	for i := range hex {
		digitValue := strings.Index(universo, string(hex[i]))
		result = result*(int64(len(universo))) + int64(digitValue)
	}
	return result
}

//decToHex , Dec to Hex
func (d *DavinciCode) decToHex(dec int64) string {
	return strconv.FormatInt(dec, len(universo))
}

//toASCIINumber , convierte el codigo ascii en retorno representado por HEX
func (d *DavinciCode) toASCIINumber(toASCII string) string {
	numberASCII := ""
	for r := range toASCII {
		unicode := string(toASCII[r])
		goInt := int64([]rune(unicode)[0])
		goHex := d.decToHex(goInt)
		numberASCII += goHex
	}
	//HEX del Ascii
	return numberASCII
}

//toNumberASCII , retorna el codigo ascii
func (d *DavinciCode) toNumberASCII(number string) string {
	numberLen := 1
	maxLen := 2
	hexASCII := ""
	ASCII := ""
	for r := range number {
		c := number[r]
		hexASCII += string(c)
		if numberLen == maxLen {
			decValue := d.hexToDec(hexASCII)
			charASCII := string(decValue)
			ASCII += charASCII
			hexASCII = ""
			maxLen += 2
		}
		numberLen++
	}
	return ASCII
}

func (d *DavinciCode) generateToken(textPlain string) string {
	bof := 32
	eof := 126
	seq := ""
	xf := (len(textPlain) % (eof - bof))
	i := 0
	h := 0

	for unicode := bof; unicode <= eof; unicode++ {
		xar := string(unicode)
		seq += xar
		if (i - (eof - bof)) == xf {
			h = (i - (eof - bof))
		}
		i++
	}
	d.patternToken = ""
	d.patternHash = ""
	x := h
	max := eof - bof
	for y := 0; y <= max; y++ {
		d.patternToken += string(seq[x])
		d.patternHash += string(seq[max-x])
		x++
		if x > max {
			x = 0
		}
	}
	return d.patternHash
}

//Encript : Encripta
func (d *DavinciCode) Encript(textPlain string) string {
	d.generateToken(textPlain)
	result := ""
	index := 0
	len := len(textPlain)
	for i := 0; i < len; i++ {
		r1 := rune(textPlain[i])
		i1 := int(r1)
		result += d.encriptDavinci(i1, len, index)
		index++
	}
	toASCIINumber := d.toASCIINumber(result)
	return toASCIINumber
}

func (d *DavinciCode) encriptDavinci(xar int, size int, tokenIndex int) string {
	if strings.Index(d.patternToken, string(xar)) != -1 {
		indexRadTOR := (strings.Index(d.patternToken, string(xar)) + size + tokenIndex) % len(d.patternToken)
		return string(d.patternHash[indexRadTOR])
	}
	return string(xar)
}

//Decript : Desincripta
func (d *DavinciCode) Decript(textPlain string) string {
	textPlain = d.toNumberASCII(textPlain)
	d.generateToken(textPlain)
	len := len(textPlain)

	result := ""
	index := 0
	for i := 0; i < len; i++ {
		r1 := rune(textPlain[i])
		i1 := int(r1)
		result += d.decriptDavinci(string(i1), len, index)
		index++
	}
	return result
}

func (d *DavinciCode) decriptDavinci(xar string, size int, tokenIndex int) string {
	indexDavinci := 0
	index := strings.Index(d.patternHash, xar)
	lenPattern := len(d.patternHash)
	lenHash := len(d.patternHash)
	if index != -1 {
		if (index - size - tokenIndex) > 0 {
			indexDavinci = (index - size - tokenIndex) % lenPattern
		} else {
			indexDavinci = len(d.patternToken) + (index-size-tokenIndex)%lenPattern
		}
		indexDavinci = indexDavinci % lenHash
		ccc := string(d.patternToken[indexDavinci])
		return ccc
	}
	return ""
}
