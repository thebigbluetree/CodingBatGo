package string1

// Given a string name, e.g. "Bob", return a greeting of the form "Hello Bob!".
func HelloName(s string) string {
	return "Hello " + s + "!"
}

// Given an "out" string length 4, such as "<<>>", and a word, return a new
// string where the word is in the middle of the out string, e.g. "<<word>>".
func MakeoutWord(out, s string) string {
	return out[:2] + s + out[2:]
}

// Given a string of even length, return the first half. So the string
// "WooHoo" yields "Woo".
func FirstHalf(s string) string {
	strlen := len(s)
	halflen := strlen / 2
	return s[:halflen]
}

// Given 2 strings, return their concatenation, except omit the first char of
// each. The strings will be at least length 1.
func NonStart(s1, s2 string) string {
	return s1[1:] + s2[1:]
}

// Given two strings, a and b, return the result of putting them together in
// the order abba, e.g. "Hi" and "Bye" returns "HiByeByeHi".
func MakeAbba(a, b string) string {
	return a + b + b + a
}

// Given a string, return a new string made of 3 copies of the last 2 chars of
// the original string. The string length will be at least 2.
func ExtraEnd(s string) string {

	return s[len(s)-2:] + s[len(s)-2:] + s[len(s)-2:]

}

// Given a string, return a version without the first and last char, so "Hello"
// yields "ell". The string length will be at least 2.
func WithoutEnd(s string) string {
	return s[1 : len(s)-1]
}

// Given a string, return a "rotated left 2" version where the first 2 chars are
// moved to the end. The string length will be at least 2.
func Left2(s string) string {
	return s[2:] + s[:2]
}

// The web is built with HTML strings like "<i>Yay</i>" which draws Yay as
// italic text. In this example, the "i" tag makes <i> and </i> which surround
// the word "Yay". Given tag and word strings, create the HTML string with tags
// around the word, e.g. "<i>Yay</i>".
func MakeTags(tag, word string) string {
	return "<" + tag + ">" + word + "</" + tag + ">"
}

// Given a string, return the string made of its first two chars, so the String
// "Hello" yields "He". If the string is shorter than length 2, return whatever
// there is, so "X" yields "X", and the empty string "" yields the empty string "".
func FirstTwo(s string) string {
	if len(s) < 2 {
		return s
	} else {
		return s[:2]
	}
}

// Given 2 strings, a and b, return a string of the form short+long+short, with
// the shorter string on the outside and the longer string on the inside. The
// strings will not be the same length, but they may be empty (length 0).
func ComboString(a, b string) string {
	if len(a) > len(b) {
		return b + a + b
	} else if len(b) > len(a) {
		return a + b + a
	} else {
		return ""
	}
}
