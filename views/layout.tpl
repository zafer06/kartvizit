<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link href="/static/css/site.css" rel="stylesheet">
  <script type="text/javascript">
      function setLang(lang) {
          document.cookie = "kartvizit_lang="+lang+";expires=Thu, 01 Jan 2020 00:00:00 UTC; path=/;";
          window.location.reload();
          //alert(document.cookie);
      }
</script>
</head>

<body>
    <div class="grid-container">
        <div class="header">
            <div class="logo">
                <h1>kartvizit.site</h1>
                <i>Kartvizitlerinizle kendinizi sanal dünyaya tanıtın.</i>
            </div>
            <div style="text-align:right">
                <a href="javascript:void(0)" onclick="setLang('en-US')">EN</a>
                <a href="javascript:void(0)" onclick="setLang('tr-TR')">TR</a>
            </div>

        </div> <!-- header end -->

        <div class="navigation">
            <ul id="main-menu">
                <li><a href="/">{{ i18n .Lang "home" }}</a></li>
                <li><a href="/login">{{ i18n .Lang "login" }}</a></li>
                {{ if .IsLogin }}
                    <li><a href="/logout">{{ i18n .Lang "logout" }}</a></li>
                {{ end }}
            </ul>
        </div> <!-- navigation end -->

        <div class="content">
            {{ block "content" . }}{{ end }}
        </div> <!-- content end -->

        <div class="footer">
            <i>2019, Kartvizit.site<i>
        </div> <!-- <footer></footer> end -->
    </div>
</body>
</html>
