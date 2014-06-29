package app

import (
	"html/template"
)

func add(x, y int) int {
	return x + y
}

var templateFuncs = template.FuncMap{"add": add}

var guestbookTemplate = template.Must(template.New("book").Funcs(templateFuncs).Parse(`
<html>
  <head>
    <title>Go Guestbook</title>
  </head>
  <body>
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
    {{range $index, $results := .}}
      {{with .Author}}
        <p>{{add $index 1}}: <b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
  </body>
</html>
`))
