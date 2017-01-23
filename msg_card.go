package card

import (
	"encoding/json"
	"strings"
)

type MessageBox struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func (m MessageBox) Json() string {
	bin, _ := json.Marshal(m)
	return string(bin)
}

func (m MessageBox) String(l int) string {
	return strings.Join(m.Box(l), "\n")
}

func (m MessageBox) Box(l int) []string {
	nl := getWidthSize(m.Name)
	if l < nl {
		l = nl
	}
	if l%2 != 0 {
		l++
	}
	lines := strings.Split(m.Message, "\n")
	newLines := []string{}

	for _, line := range lines {
		if getWidthSize(line) < l {
			newLines = append(newLines, line)
			continue
		}
		newLines = append(newLines, newLine(line, l)...)
	}

	for i, line := range newLines {
		if len(line) == 0 {
			continue
		}
		if len(line) < l {
			newLines[i] = line + nString(" ", l-len(line))
		}
	}

	top := "┌" + nString("─", l) + "┐"
	name := m.Name + nString(" ", l-nl)
	name = wall([]string{name})[0]
	middle := "├" + nString("─", l) + "┤"
	newLines = wall(newLines)
	bottom := "└" + nString("─", l) + "┘"

	ret := []string{}

	ret = append(ret, top)
	ret = append(ret, name)
	ret = append(ret, middle)
	ret = append(ret, newLines...)
	ret = append(ret, bottom)

	return ret
}

func wall(ss []string) []string {
	for i, s := range ss {
		ss[i] = "│" + s + "│"
	}
	return ss
}

func nString(s string, n int) string {
	var ret string
	for i := 0; i < n; i++ {
		ret += s
	}
	return ret
}

func newLine(s string, l int) []string {

	w := getWidthSize(s)
	if w < l {
		return []string{
			s,
		}
	}
	for i, _ := range []rune(s) {
		w := getWidthSize(string([]rune(s)[:i+1]))
		if w == l {
			ret := []string{string([]rune(s)[:i+1])}
			return append(ret, newLine(string([]rune(s)[i+1:]), l)...)
		}
	}
	return nil
}

func isMultiByteChar(r rune) bool {
	return 1 != len([]byte(string([]rune{r})))
}

func getWidthSize(s string) int {
	var counter int
	for _, r := range []rune(s) {
		if isMultiByteChar(r) {
			counter += 2
		} else {
			counter++
		}
	}
	return counter
}
