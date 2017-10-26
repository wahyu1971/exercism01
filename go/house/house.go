package house

const testVersion = 1

var mverse = []struct {
	noun string
	verb string
}{
	{"house that Jack built.", ""},
	{"malt", "lay in"},
	{"rat", "ate"},
	{"cat", "killed"},
	{"dog", "worried"},
	{"cow with the crumpled horn", "tossed"},
	{"maiden all forlorn", "milked"},
	{"man all tattered and torn", "kissed"},
	{"priest all shaven and shorn", "married"},
	{"rooster that crowed in the morn", "woke"},
	{"farmer sowing his corn", "kept"},
	{"horse and the hound and the horn", "belonged to"},
}

func recursive2(n int) (result string) {
	result = "\nthat " + mverse[n].verb + " the " + mverse[n-1].noun
	if n == 1 {
		return
	}
	result = result + recursive2(n-1)
	return
}

func recursive(n int) (result string) {
	result = "This is the " + mverse[n].noun
	if n == 0 {
		return
	}
	return result + recursive2(n)
}

func Verse(n int) string {
	if n < 1 || n > len(mverse) {
		return ""
	}
	n-- // 1 -> 0
	return recursive(n)
}

func rec_song(n int) (result string) {
	if n == len(mverse) {
		result = Verse(n)
	} else {
		result = Verse(n) + "\n\n" + rec_song(n+1)
	}
	return
}

func Song() (result string) {
	return rec_song(1)
}
