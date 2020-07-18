{{define "body"}}
<h1>Login</h1>
{{ if .Flash }}
    {{ range $key, $value := .Flash.Messages}}
        <span style="color: red;">{{ $key }}:{{ $value }}</span>
    {{ end }}
{{ end }}
<form method="post" action="/signin">
    <input type="hidden" name="csrf_token" value="{{ .CSRFToken }}">

    <label for="email">Email</label>
    <input type="text" id="email" name="email" value="{{ .Old.email }}">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
{{end}}