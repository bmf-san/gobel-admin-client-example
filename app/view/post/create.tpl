{{define "body"}}
<div>
    <h1>/posts/create</h1>
    <form action="/posts" method="post">
        <div>
            <label for="category">category</label>
            <select name="category" id="category">
                {{range .Categories}}
                <option value="{{ .ID }}">{{ .Name }}</option>
                {{end}}
            </select>
        </div>
        <div>
            <label for="tag">tag</label>
            <input type="email" name="tag" autocomplete="on" list="tags" multiple>
            <datalist id="tags">
                {{range .Tags}}
                <option value="{{ .ID }}">{{ .Name }}</option>
                {{end}}
            </datalist><br>
        </div>
        <div>
            <label for="title">title</label>
            <input type="text" name="title">
        </div>
        <div>
            <label for="md_body">md_body</label>
            <textarea name="md_body" id="md_body" cols="30" rows="10"></textarea>
        </div>
        <div>
            <label for="status">status</label>
            <select name="status" id="status">
                {{range .Statuses}}
                <option value="{{ . }}">
                {{end}}
            </select>
        </div>
        <div>
            <input type="submit" value="submit">
        </div>
    </form>
</div>
{{end}}