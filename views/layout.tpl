<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style type="text/css">
    .menu {
        margin-top:10px;
        border-top:1px solid #ddd;
        border-bottom:1px solid #ddd;
        padding: 5px;
    }
  </style>
  <script type="text/javascript">
  function setLang(lang) {
      document.cookie = "kartvizit_lang="+lang+";expires=Thu, 01 Jan 2020 00:00:00 UTC; path=/;";
      window.location.reload();
      //alert(document.cookie);
  }
</script>
</head>

<body>
  <header>
    <h1 class="logo">kartvizit.site</h1>
    <div>
        Kartvizitlerinizle kendinizi sanal dünyaya tanıtın.
        <a href="javascript:void(0)" onclick="setLang('en-US')">EN</a>
        <a href="javascript:void(0)" onclick="setLang('tr-TR')">TR</a>
    </div>
    <div class="menu">
        <a href="/">{{ i18n .Lang "home" }}</a>
        <a href="/login">{{ i18n .Lang "login" }}</a>
        {{ if .IsLogin }}
            <a href="/logout">{{ i18n .Lang "logout" }}</a>
        {{ end }}
        </div>
  </header>

  <div>{{ block "content" . }}{{ end }}</div>

  <footer>
      <i>2019, Kartvizit.site<i>
  </footer>
</body>
</html>
