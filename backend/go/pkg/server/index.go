package server

import (
	"fmt"
	"net/http"
)

const (
	APP_TITLE = "go+es"
	WELCOME_MESSAGE = "Welcome to <span style=\"color: #d67936;\">" + APP_TITLE + "</span> !"
	HOME_HTML = `
<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1">

	<title>%s</title>
	<meta name="description" content="%s">
</head>

<body>
	<div style="display: flex; justify-content: center;">
		<h1>%s</h1>
	</div>
</body>
</html>
`
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, HOME_HTML, APP_TITLE, APP_TITLE, WELCOME_MESSAGE)
}
