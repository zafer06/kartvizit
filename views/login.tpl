{{ template "layout.tpl" . }}

{{ define "content" }}
    <p>{{.flash.error}}</p>

    <form action="/login" method="post">
        <p>
            <label>{{ i18n .Lang "email" }}:</label>
            <input type="text" name="email">
        </p>
        <p>
            <label>{{ i18n .Lang "password" }}:</label>
            <input type="text" name="password">
        </p>
        <p>
            <input type="submit" value='{{ i18n .Lang "signin" }}'>
        </p>
    </form>
{{ end }}
