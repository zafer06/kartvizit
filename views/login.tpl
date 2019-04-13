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
                        <label>{{ i18n .Lang "firstName" }}:</label>
                        <input type="text" name="firstName">
                    </p>
                    <p>
                        <label>{{ i18n .Lang "lastName" }}:</label>
                        <input type="text" name="lastName">
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
                        <label>{{ i18n .Lang "confirmPassword" }}:</label>
                        <input type="text" name="confirmPassword">
                    </p>
                    <p>
                        <input type="submit" value='{{ i18n .Lang "register" }}'>
                    </p>
                </fieldset>
            </form>
        </div>
    </div>
{{ end }}
