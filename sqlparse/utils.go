package sqlparse

import (
	"fmt"
	"regexp"
	"strings"
)

type alterQuery struct {
	Table  string
	Action string
}

// ParseAlterQuery make alter query to struct
func ParseAlterQuery(str string) (bool, alterQuery) {
	// Remove comment --
	re := regexp.MustCompile("(?s)--.*?\n|/\\*.*?\\*/")
	str = re.ReplaceAllString(str, "")

	// Remove new lines char '\n' to let it become one long time.
	str = strings.Replace(str, "\n", " ", -1)
	str = strings.Replace(str, ";", ";\n", -1)

	re = regexp.MustCompile(`(?m)^(?:.*)ALTER\sTABLE\s(\S*)\s(.*);$`)

	match := re.FindStringSubmatch(str)
	if match != nil {
		return true, alterQuery{
			Table:  strings.Replace(match[1], "`", "", -1),
			Action: match[2],
		}
	}

	return false, alterQuery{}
}

// ConverToJT convert query to pt-online-schema-change comment
func ConverToJT(q alterQuery) string {
	return fmt.Sprintf("--alter \"%s\" t=%s", q.Action, q.Table)
}
