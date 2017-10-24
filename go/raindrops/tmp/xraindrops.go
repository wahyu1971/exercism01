// Package raindrops is an INCREDIBLY useful package which helps convert numbers of raindrops to sounds
package raindrops

import (
	"bytes"
	"strconv"
)

const testVersion = 3

/* Convert a number to a string, the contents of which depend on the number's factors.
- If the number has 3 as a factor, output 'Pling'.
- If the number has 5 as a factor, output 'Plang'.
- If the number has 7 as a factor, output 'Plong'.
- If the number does not have 3, 5, or 7 as a factor,
  just pass the number's digits straight through.

## Examples

- 28's factors are 1, 2, 4, **7**, 14, 28.
  - In raindrop-speak, this would be a simple "Plong".
- 30's factors are 1, 2, **3**, **5**, 6, 10, 15, 30.
  - In raindrop-speak, this would be a "PlingPlang".
- 34 has four factors: 1, 2, 17, and 34.
  - In raindrop-speak, this would be "34".
*/
func Convert(i int) string {
	var buffer bytes.Buffer
	if 0 == i%3 {
		buffer.WriteString("Pling")
	}
	if 0 == i%5 {
		buffer.WriteString("Plang")
	}
	if 0 == i%7 {
		buffer.WriteString("Plong")
	}
	if 0 == buffer.Len() {
		buffer.WriteString(strconv.Itoa(i))
	}
	return buffer.String()
}
