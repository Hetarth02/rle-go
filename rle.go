package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Pair struct {
	count int
	char  string
}

// Stores as array of Structs
func encode(str string) []Pair {
	var rle_str []Pair
	char_count := 0
	var prev_char byte
	var curr_char byte

	for index := 1; index < len(str); index++ {
		char_count += 1
		prev_char = str[index-1]
		curr_char = str[index]

		if prev_char != curr_char {
			rle_str = append(rle_str, Pair{char_count, string(prev_char)})
			char_count = 0
		}
	}

	char_count += 1
	rle_str = append(rle_str, Pair{char_count, string(curr_char)})

	return rle_str
}

func decode(pair_arr []Pair) string {
	var og_str strings.Builder
	for _, value := range pair_arr {
		og_str.WriteString(strings.Repeat(value.char, value.count))
	}
	return og_str.String()
}

// Stores as an encoded string
func str_encode(str string) string {
	var rle_str strings.Builder
	char_count := 0
	var prev_char byte
	var curr_char byte

	for index := 1; index < len(str); index++ {
		char_count += 1
		prev_char = str[index-1]
		curr_char = str[index]

		if prev_char != curr_char || char_count == 9 {
			write_str := fmt.Sprintf("%d%s", char_count, string(prev_char))
			rle_str.WriteString(write_str)
			char_count = 0
		}
	}

	char_count += 1
	write_str := fmt.Sprintf("%d%s", char_count, string(curr_char))
	rle_str.WriteString(write_str)

	return rle_str.String()
}

func str_decode(str string) string {
	var og_str strings.Builder
	var count int64
	var char string
	var ok error

	for index := 0; index < len(str); index += 2 {
		count, ok = strconv.ParseInt(string(str[index]), 10, 0)
		char = string(str[index+1])

		if ok != nil {
			panic("Invalid int count of character")
		}

		og_str.WriteString(strings.Repeat(char, int(count)))
	}

	return og_str.String()
}

// Stores as an encoded string with most repeating frequency stripped off
func str_encode_without_one(str string) string {
	var rle_str strings.Builder
	char_count := 0
	var prev_char byte
	var curr_char byte

	for index := 1; index < len(str); index++ {
		char_count += 1
		prev_char = str[index-1]
		curr_char = str[index]

		if prev_char != curr_char || char_count == 9 {
			write_str := fmt.Sprintf("%d%s", char_count, string(prev_char))
			if char_count == 1 {
				write_str = fmt.Sprintf("%s", string(prev_char))
			}
			rle_str.WriteString(write_str)
			char_count = 0
		}
	}

	char_count += 1
	write_str := fmt.Sprintf("%d%s", char_count, string(prev_char))
	if char_count == 1 {
		write_str = fmt.Sprintf("%s", string(prev_char))
	}
	rle_str.WriteString(write_str)

	return rle_str.String()
}

func str_decode_without_one(str string) string {
	var og_str strings.Builder
	var count int64
	var ok error

	for index := 0; index < len(str); index++ {
		count, ok = strconv.ParseInt(string(str[index]), 10, 0)
		if ok != nil {
			og_str.WriteString(string(str[index]))
		} else {
			index += 1
			og_str.WriteString(strings.Repeat(string(str[index]), int(count)))
		}
	}

	return og_str.String()
}

// Determine maximum occuring number from encoded string
func determine_max_frequency(str string) int {
	// Since we only have buckets of 9 we can have an fixed array
	frequency_map := [10]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	key := 0
	largest := 0

	for index := 0; index < len(str); index += 2 {
		count, ok := strconv.ParseInt(string(str[index]), 10, 8)

		if ok != nil {
			panic("Invalid int count of character")
		}

		frequency_map[count] += 1
	}

	for k, v := range frequency_map[1:] {
		if v >= largest {
			key = k
			largest = v
		}
	}

	key++

	return key
}

// Stores as an encoded string with most repeating frequency stripped off
func str_dynamic_encode(str string) (string, int) {
	var rle_str strings.Builder
	encoded_str := str_encode(str)
	repeating_value := determine_max_frequency(encoded_str)

	for index := 0; index < len(encoded_str); index += 2 {
		character := encoded_str[index+1]
		freq, ok := strconv.ParseInt(string(encoded_str[index]), 10, 8)

		if ok != nil {
			panic("Invalid int count of character")
		}

		write_str := fmt.Sprintf("%d%s", freq, string(character))

		if freq == int64(repeating_value) {
			write_str = fmt.Sprintf("%s", string(character))
		}

		rle_str.WriteString(write_str)
	}

	return rle_str.String(), repeating_value
}

func str_dynamic_decode(str string, max_repeat_count int) string {
	var og_str strings.Builder
	var count int64
	var ok error

	for index := 0; index < len(str); index++ {
		count, ok = strconv.ParseInt(string(str[index]), 10, 0)
		if ok != nil {
			og_str.WriteString(strings.Repeat(string(str[index]), max_repeat_count))
		} else {
			// We shift the index to character
			index += 1
			og_str.WriteString(strings.Repeat(string(str[index]), int(count)))
		}
	}

	return og_str.String()
}

func main() {
	test_data := [16]string{
		"AAAAAAAAAAABBBBBBBBBBCCCCCCCCCCDDDDDDDDDDEEEEEEEEEE",
		"AABBAABBAABBAABBAABB",
		"MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"abcabcabcabcabcabcabcabcabcabc",
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"aAaAaAaAaAaAaAaAaAaAaAaAaAaAaAaAaAaAaAaAaAaAaAaAaA",
		"ZZaZZaZZaZZaZZaZZa",
		"QWERTYUIOPASDFGHJKLZXCVBNM",
		"TheQuickBrownFoxJumpsOverTheLazyDog",
		"LLLLLLLLLLMMMMMMMMMMNNNNNNNNNNOOOOOOOOOOPPPPPPPPPP",
		"HELLOHELLOHELLOHELLOHELLO",
		"AbCdEfGhIjKlMnOpQrStUvWxYz",
		"aaaaabbbbbcccccdddddeeeeefffffggggghhhhh",
		"AABCCCCCCCCCCCCDDDDDEEEEEFGHIJK",
		"abwwwwwwwwwwwwwqqwweerrttyyuubbhhiiuuttyywwxxbbbbbbbbbbjjjjjsaaaaaaaaaaaasoooooooodsgjhgdddddddddddddddddd",
	}

	for index := 0; index < len(test_data); index++ {
		og_str := test_data[index]

		fmt.Println(strings.Repeat("-", 100))

		fmt.Println("Normal: Stores as array of structs")
		fmt.Println(encode(og_str))
		fmt.Println("Do the original and decompressed strings match: ", decode(encode(og_str)) == og_str)

		fmt.Println("")

		fmt.Println("Standard RLE")
		fmt.Println(str_encode(og_str))
		fmt.Println("Do the original and decompressed strings match: ", str_decode(str_encode(og_str)) == og_str)
		fmt.Println("Original Size: ", len(og_str), "bytes", " | ", "Compressed Size: ", len(str_encode(og_str)), "bytes")

		fmt.Println("")

		fmt.Println("Encode without one")
		fmt.Println(str_encode_without_one(og_str))
		fmt.Println("Do the original and decompressed strings match: ", str_decode_without_one(str_encode_without_one(og_str)) == og_str)
		fmt.Println("Original Size: ", len(og_str), "bytes", " | ", "Compressed Size: ", len(str_encode_without_one(og_str)), "bytes")

		fmt.Println("")

		fmt.Println("Dynamic Encode")
		encoded_str, repeating_value := str_dynamic_encode(og_str)
		fmt.Println(encoded_str)
		fmt.Println("Do the original and decompressed strings match: ", str_dynamic_decode(encoded_str, repeating_value) == og_str)
		fmt.Println("Original Size: ", len(og_str), "bytes", " | ", "Compressed Size: ", len(encoded_str), "bytes")

		fmt.Println(strings.Repeat("-", 100))

		fmt.Println("")
	}
}
