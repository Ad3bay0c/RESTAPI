<html>
<head>
    <title>Go Web Testing</title>
</head>
<body>
<form action="/login" method="post">
    <div style="width: 40%">
        <div style="margin-bottom: 10px">
            <label for="username">Username:</label>
            <input type="text" placeholder="Username" name="username">
        </div>

        <div style="margin-bottom: 5px">
            <label for="password">Password:</label>
            <input type="password" placeholder="****" name="password"><br>
        </div>
        <input type="hidden" name="token" value="{{ . }}">
        <div style="width: 30%">
            <input type="submit" value="SUBMIT" style="background-color: aqua; display: block; float: right">
        </div>

    </div>

</form>
</body>
</html>