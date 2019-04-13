{{ template "layout.tpl" . }}

{{ define "content" }}
      <h4>Pano Sayfası</h4>

      {{ if .IsLogin }}
            <p> Hoşgeldin, {{ .UserName }} </p>
      {{ end }}
{{ end }}
