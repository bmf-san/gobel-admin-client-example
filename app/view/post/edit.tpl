{{define "body"}}
<!-- TODO: formにする -->
<div>
    <h1>/posts/:id/edit</h1>
    <p><span>Post ID</span>:{{ .ID }}</p>
    <p><span>Admin ID</span>:{{ .Admin.ID }}</p>
    <p><span>Category ID</span>:{{ .Category.ID }}</p>
    <p><span>Category Name</span>:{{ .Category.Name }}</p>
    {{range .Tags}}
    <p><span>Tag ID</span>:{{ .ID }}</p>
    <p><span>Tag Name</span>:{{ .Name }}</p>
    {{end}}
    <p><span>Title</span>:{{ .Title }}</p>
    <p><span>MD Body</span>:{{ .MDBody }}</p>
    <p><span>HTML Body</span>:{{ .HTMLBody }}</p>
    <p><span>Status</span>:{{ .Status }}</p>
    {{range .Comments}}
    <p><span>Comment ID</span>:{{ .ID }}</p>
    <p><span>Comment Name</span>:{{ .Body }}</p>
    {{end}}
    <p><span>CreateAt</span>:{{ .CreatedAt }}</p>
    <p><span>UpdatedAt</span>:{{ .UpdatedAt }}</p>
</div>
{{end}}