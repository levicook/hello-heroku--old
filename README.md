hello-heroku-postgres
=====================

```bash
$ heroku create -b https://github.com/kr/heroku-buildpack-go.git
Creating enigmatic-shore-5230... done, stack is cedar
BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git
http://enigmatic-shore-5230.herokuapp.com/ | git@heroku.com:enigmatic-shore-5230.git
Git remote heroku added

$ git push heroku master
Initializing repository, done.
Counting objects: 76, done.
Delta compression using up to 8 threads.
Compressing objects: 100% (64/64), done.
Writing objects: 100% (76/76), 57.95 KiB | 0 bytes/s, done.
Total 76 (delta 9), reused 0 (delta 0)

-----> Fetching custom git buildpack... done
-----> Go app detected
-----> Installing go1.2... done
-----> Running: godep go install -tags heroku ./...
-----> Discovering process types
       Procfile declares types -> web

-----> Compressing... done, 1.9MB
-----> Launching... done, v4
       http://enigmatic-shore-5230.herokuapp.com/ deployed to Heroku

To git@heroku.com:enigmatic-shore-5230.git
 * [new branch]      master -> master

$ curl -i  http://enigmatic-shore-5230.herokuapp.com/
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Wed, 18 Jun 2014 20:17:14 GMT
Content-Length: 246
Connection: keep-alive

hello
BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git
DYNO=web.1
PATH=/usr/local/bin:/usr/bin:/bin:/app/bin
PWD=/app
PS1=\[\033[01;34m\]\w\[\033[00m\] \[\033[01;32m\]$ \[\033[00m\]
SHLVL=1
HOME=/app
PORT=15639
_=/app/bin/hello-heroku
```
