<!DOCTYPE html>
<html lang="zh-CN">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- 上述3个meta标签*必须*放在最前面，任何其他内容都*必须*跟随其后！ -->
    <meta name="description" content="">
    <meta name="author" content="">

    <title>Login Page</title>

    <!-- Bootstrap core CSS -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">

    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
    <link href="/static/css/ie10-viewport-bug-workaround.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <link href="/static/css/signin.css" rel="stylesheet">
    <link rel="shortcut icon" href="static/favicon.ico">
    <!-- Just for debugging purposes. Don't actually copy these 2 lines! -->
    <!--[if lt IE 9]><script src="../../assets/js/ie8-responsive-file-warning.js"></script><![endif]-->
    <script src="/static/js/ie-emulation-modes-warning.js"></script>

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
      <script src="https://cdn.bootcss.com/html5shiv/3.7.3/html5shiv.min.js"></script>
      <script src="https://cdn.bootcss.com/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
  </head>

  <body>
 <nav class="navbar navbar-default navbar-static-top">
  <div class="container">
    <ul class="nav navbar-nav navbar-right">
       <li><a href="https://github.com/freedomkk-qfeng/nginx-ldap-auth">GITHUB</a></li>
    </ul>
  </div>
</nav>
    <div class="container">
	    <form action="/login" method="post" class="form-signin">
		{{ .xsrfdata }}
		    <h2 class="text-center" class="form-signin-heading">Login Page</h2>
		    <div class="form-group">
				<label class="sr-only">Username</label>
		        <input type="text" name="username" class="form-control" placeholder="Username" required>
			</div>
			<div class="form-group">
				<label for="inputPassword" class="sr-only">Password</label>
		        <input type="password" name="password" class="form-control" placeholder="Password" required>
			</div>
			{{if .captcha}}
			<div class="form-group">
				<div class="row">
					<div class="col-md-6">
						<input  class="form-control" name="captcha"  type="text" placeholder="Captcha" required>
			 	    </div>
					<div class="col-md-6">
						{{create_captcha}}
					</div>
				</div>
			</div>
			{{end}}		
			<button class="btn btn-lg btn-primary btn-block" type="submit" value="Login">login</button>
			 <input type="hidden" name="target" value={{.target}}>
			{{.msg}}
			
		</form>
    </div>

    <footer class="footer">
      <div class="container" class="center-block">
        <p class="text-muted">Copyright © 2017 freedomkk-qfeng</p>
      </div>
    </footer>

    <!-- IE10 viewport hack for Surface/desktop Windows 8 bug -->
    <script src="/static/js/ie10-viewport-bug-workaround.js"></script>
  </body>
</html>
