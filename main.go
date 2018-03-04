package main

import (
	"log"
	"os"
	"text/template/parse"
)

func SkipToken(b []byte) []byte {
	return b[6:]
}

func main() {
	forest, err := parse.Parse("twsql", `
SELECT
	*
FROM
	foo
WHERE
	username = /* .UserName */'john'
`, "/*", "*/")
	if err != nil {
		log.Fatal(err)
	}

	tree := forest["twsql"]
	skip := false
	for _, node := range tree.Root.Nodes {
		switch node := node.(type) {
		case *parse.TextNode:
			t := node.Text
			if skip {
				t = SkipToken(t)
				skip = false
			}
			os.Stdout.Write(t)
		case *parse.ActionNode:
			skip = true
			os.Stdout.Write([]byte{'?'})
		}
	}
	// pp.Println(tree.Root.Nodes)
}
