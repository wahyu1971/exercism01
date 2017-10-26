package twelve

const testVersion = 1

var count = map[int][]string{
	1:  {"first", "a Partridge"},
	2:  {"second", "two Turtle Doves, and "},
	3:  {"third", "three French Hens, "},
	4:  {"fourth", "four Calling Birds, "},
	5:  {"fifth", "five Gold Rings, "},
	6:  {"sixth", "six Geese-a-Laying, "},
	7:  {"seventh", "seven Swans-a-Swimming, "},
	8:  {"eighth", "eight Maids-a-Milking, "},
	9:  {"ninth", "nine Ladies Dancing, "},
	10: {"tenth", "ten Lords-a-Leaping, "},
	11: {"eleventh", "eleven Pipers Piping, "},
	12: {"twelfth", "twelve Drummers Drumming, "},
}

func Verse(input int) string {
	sum := ""
	for i := input; i >= 1; i-- {
		sum = sum + count[i][1]
	}
	return "On the " + count[input][0] + " day of Christmas my true love gave to me, " + sum + " in a Pear Tree."
}

func Song() (sum string) {
	for i := 1; i <= 12; i++ {
		sum = sum + Verse(i) + "\n"
	}
	return sum
}
