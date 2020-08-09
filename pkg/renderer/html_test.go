// package renderer

// import (
// 	"context"
// 	"log"
// 	"strings"
// 	"testing"
// )

// func TestHTMLRenderer(t *testing.T) {
// 	r := NewHTML("/auth", "../../../internal/users/templates/html")

// 	err := r.Load("login")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	o, content, err := r.Render(context.Background(), "login", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if content != "text/html" {
// 		t.Error("context type not set properly")
// 	}
// 	log.Println(string(o))

// 	if len(o) == 0 {
// 		t.Error("it should have rendered a template")
// 	}

// 	if !strings.Contains(string(o), "/auth/login") {
// 		t.Error("expected the url to be rendered out for the form post location")
// 	}
// }

package renderer

import (
	"context"
	"log"
	"strings"
	"testing"
)

func TestHTMLRenderer(t *testing.T) {
	r := newTemplateRenderer("/auth", "../../web/templates/layouts/*.tpl", "../../web/templates/authboss/*.tpl")

	o, content, err := r.Render(context.Background(), "login.html.tpl", nil)
	if err != nil {
		t.Fatal(err)
	}

	if content != "text/html" {
		t.Error("context type not set properly")
	}
	log.Println("asdasdasd")
	log.Println(string(o))

	if len(o) == 0 {
		t.Error("it should have rendered a template")
	}

	if !strings.Contains(string(o), "/auth/login") {
		t.Error("expected the url to be rendered out for the form post location")
	}
}
