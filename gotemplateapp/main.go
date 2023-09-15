/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);


Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
Example 3:

Input: s = "A", numRows = 1
Output: "A"


Constraints:

1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000
*/

package main

import "fmt"

func init2dArr(numRows int) [][]string {
	twoDArr := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		twoDArr[i] = make([]string, 1000)
		for j := 0; j < 1000; j++ {
			twoDArr[i][j] = ""
		}
	}
	return twoDArr
}

func calcAnswer(arr [][]string, numRows int) string {
	answer := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < 1000; j++ {
			if arr[i][j] != "" {
				answer = answer + arr[i][j]
			}
		}
	}
    return answer
}

func convert(s string, numRows int) string {
	if len(s) == 1 || numRows == 1 {
		return s
	}
	arr := init2dArr(numRows)

	c := 0
	i := 0
	j := 0
	verticalMode := true
	done := false
	for {
		if verticalMode && !done { // vertical mode and we are not done
			for i < numRows {
				arr[i][j] = string(s[c])
				
				c++
				if c < len(s) {
					i++
				} else {
					done = true
					break
				}
			}
			// set up first zigzag
			i--
			i--
			j++
			verticalMode = false
		} else if !done { // zigzag mode and we are not done
			for i > 0 {
				arr[i][j] = string(s[c])

				c++
				if c < len(s) {
					// keep doing zigzag back up to i == 0
					i--
					j++
				} else {
					done = true
					break
				}
			}
			verticalMode = true
		} else { // finished draining string, exit
			break
		}
	}

	answer := calcAnswer(arr, numRows)
	return answer
}

func main() {
	fmt.Println("Hello World!")
}
