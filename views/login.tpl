{{ template "layout.tpl" . }}

{{ define "content" }}
    <p>{{.flash.error}}</p>

    <div class="login-register">
        <div>
            <form action="/login" method="post">
                <fieldset>
                    <legend>{{ i18n .Lang "login"}}:</legend>
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
                </fieldset>
            </form>
        </div>

        <div>
            <form action="/register" method="post">
                <fieldset>
                    <legend>{{ i18n .Lang "register"}}:</legend>
                    <p>
                        <label>{{ i18n .Lang "name" }}:</label>
                        <input type="text" name="name">
                    </p>
                    <p>
                        <label>{{ i18n .Lang "email" }}:</label>
                        <input type="text" name="email">
                    </p>
                    <p>
                        <label>{{ i18n .Lang "password" }}:</label>
                        <input type="text" name="password">
                    </p>
                    <p>
                        <input type="submit" value='{{ i18n .Lang "register" }}'>
                    </p>
                </fieldset>
            </form>
        </div>
    </div>
{{ end }}
