{{define "body"}}
<h1>Home</h1>
<p>User ID:{{ .UserID }}</p>
<form method="post" action="/signout">
    <input type="hidden" name="csrf_token" value="{{ .CSRFToken }}">
    <button type="submit">Logout</button>
</form>
{{end}}