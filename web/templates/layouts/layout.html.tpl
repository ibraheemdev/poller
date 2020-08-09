{{ define "base" }}
    <!DOCTYPE html>
    <html>
        <head>
            <title>Poller</title>
        </head>
        <body>
            {{ template "content" . }}
        </body>
    </html>
{{ end }}