<html>
<head>
<title>{{ .Title }}</title>
</head>
<script src="https://cdn.jsdelivr.net/npm/vue"></script>
<body>
<h1>{{ .Title }}</h1>
<form action="/search" method="post">
    <p>
    <input type="checkbox" name="enemy" value="エウロペ">エウロペ
    <input type="checkbox" name="enemy" value="ワムデュス">ワムデュス
    </p>
    <input type="submit" value="検索">
</form>
<div id="app-1">{{ message }}</div>
<script>
var app1 = new Vue({
    el: '#app-1',
    data: { message: 'Hello world!' }
})
</script>
</body>
</html>