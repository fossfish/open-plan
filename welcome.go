package openplan

import (
	"net/http"
)

// A Welcome message with title, demonstrates passing data to a template
type Welcome struct {
	Title   string
	Message string
}

// A template taking a struct pointer (&message) containing data to render
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	message := Welcome{Title: "Open-Plan.io", Message: "Welcome to your template site"}

	// outerTheme refernces the template defined within theme.html
	templates["welcome.html"].ExecuteTemplate(w, "outerTheme", &message)
}
