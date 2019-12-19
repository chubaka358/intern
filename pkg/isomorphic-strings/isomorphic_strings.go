package isomorphic_strings

import "unicode/utf8"

// IsIsomorphic returns true if two strings are isomorphic
// Otherwise IsIsomorphic returns false
// Two strings are isomorphic if the characters in s can be replaced to get t.
func IsIsomorphic(s string, t string) bool {
	// If two lines have different lengths, then they are not isomorphic
	if utf8.RuneCountInString(s) != utf8.RuneCountInString(t){
		return false
	}
	// characters stores mapping character in line s to character in line t
	characters := make(map[uint8]uint8)
	// Iterate over the characters of the first line
	for i := 0; i < utf8.RuneCountInString(s); i++ {
		// Check if character contains in characters
		_, ok := characters[s[i]]
		// If not, then create a mapping of the character in string s to the character in string t
		if !ok{
			characters[s[i]] = t[i]
		// Else check whether the old mapping matches the new
		// If not, then these strings are not isomorphic
		} else if characters[s[i]] != t[i]{
			return false
		}
	}
	// uniqueValues allows us to check the uniqueness of a value
	uniqueValues := make(map[uint8]bool)
	// Iterate over characters map by value
	for _, val := range characters{
		// Checks if there is already a mapping to this character
		// If yes, then strins are not isomorphic
		if _, ok := uniqueValues[val]; ok{
			return false
		}
		// Marks symbol as used
		uniqueValues[val] = true
	}
	// Strings are isomorphic
	return true
}
