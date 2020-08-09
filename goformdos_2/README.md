
## Goformdos

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pierelucas/goformdos) ![GitHub](https://img.shields.io/github/license/pierelucas/goformdos) ![GitHub last commit](https://img.shields.io/github/last-commit/pierelucas/goformdos)

### Golang Layer 7 HTTP GET/POST DOS TOOL

#
### Install

    go build

or

    go install

#
### Usage

+ GET

        ./goformdos -u https://example.com/ -m POST -t 100 -r 100 -hp headers.txt 

+ POST
 
        ./goformdos -u https://example.com/login.php -m POST -t 100 -r 100 -fp forms.txt -hp headers.txt 

+ POST/GET + Basic Authorization

        ./goformdos -u https://example.com/login.php -m POST -t 100 -r 100 -fp forms.txt -hp headers.txt -a USER:PASSWORD

+ POST/GET + Basic Authorization + custom Logfile

        ./goformdos -u https://example.com/login.php -m POST -t 100 -r 100 -fp forms.txt -hp headers.txt -a USER:PASSWORD -o log.txt

#
### How to write Configuration files

+ Generally, like a Hashmap/JSON
        
        Key:Value