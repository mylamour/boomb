
<p align="center">
  <img src="./boomb.png" />
</p>


> Burp force authority like boomb boomb boomb


This is my golang learning tour, logo generator from https://gopherize.me/

# Build

run `make build` to build project, then you will get binaray in `bin` folder
also you can use `make test` to run all unit tests, but may be you
need to configure your environment


# Useage

* `boomb --target http://127.0.0.1:8080 --user user --pass test/pass.list`
* `boomb --target ssh://127.0.0.1:2222 --user test/user.list --pass test/pass.list`
* `boomb --target redis://127.0.0.1:6379 --user " " --pass test/pass.list`

# ToDo

[ ] work poll
[ ] schedular task
[ ] retry 
[ ] queue
[x] goroutoin
[x] hasring or arrang slic?