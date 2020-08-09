package renderer

import (
	"context"
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

	if len(o) == 0 {
		t.Error("it should have rendered a template")
	}

	if !strings.Contains(string(o), "/auth/login") {
		t.Error("expected the url to be rendered out for the form post location")
	}
}
