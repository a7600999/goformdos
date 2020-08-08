- Golang Layer 7 HTTP GET/POST DOS TOOL

Install:

+ go build

or

+ go install

Usage:

GET
+ ./goformdos -u https://example.com/ -m POST -t 100 -r 100 -hp headers.txt 

POST
+ ./goformdos -u https://example.com/login.php -m POST -t 100 -r 100 -fp forms.txt -hp headers.txt 

POST/GET + Basic Authorization
+ ./goformdos -u https://example.com/login.php -m POST -t 100 -r 100 -fp forms.txt -hp headers.txt -a USER:PASSWORD

POST/GET + Basic Authorization + custom Logfile
+ ./goformdos -u https://example.com/login.php -m POST -t 100 -r 100 -fp forms.txt -hp headers.txt -a USER:PASSWORD -o log.txt

How to write Configuration files:

+ KEY:VALUE

Generally, like a Hashmap