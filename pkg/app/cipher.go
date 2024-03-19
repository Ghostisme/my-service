package app

import (
	"unicode"
)

// 加密解密接口
type Cipher interface {
	Encryption(string) string
	Decryption(string) string
}

// 密钥
type cipher []int

// 加密算法
func (c cipher) cipherAlgorithm(letters string, shift func(int, int) int) string {
	shiftedText := ""
	for _, letter := range letters {
		// fmt.Println("循环加密算法字符串每一项", letter, unicode.IsLetter(letter))
		if !unicode.IsLetter(letter) {
			continue
		}
		shiftDist := c[len(shiftedText)%len(c)]
		// fmt.Println("中间变量", shiftDist)
		s := shift(int(unicode.ToLower(letter)), shiftDist)
		// fmt.Println("编码转换中间", s)
		switch {
		case s < 'a':
			s += 'z' - 'a' + 1
		case 'z' < s:
			s -= 'z' - 'a' + 1
		}
		shiftedText += string(rune(s))
		// fmt.Println("结果", shiftedText)
	}
	return shiftedText
}

// 加密
func (c *cipher) Encryption(plainText string) string {
	return c.cipherAlgorithm(plainText, func(a, b int) int { return a + b })
}

// 解密
func (c *cipher) Decryption(cipherText string) string {
	return c.cipherAlgorithm(cipherText, func(a, b int) int { return a - b })
}

// 创建新的凯撒密码
func NewCaesar(key int) Cipher {
	return NewShift(key)
}

// 创建新的密码.
func NewShift(shift int) Cipher {
	if shift < -25 || 25 < shift || shift == 0 {
		return nil
	}
	c := cipher([]int{shift})
	return &c
}
