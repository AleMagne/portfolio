# Portfolio
My first portfolio website.

The idea is to create a functional portfolio site and learn Go in the mean time. All the backend will be coded in Go. I think HTMX might be involved in some capacity, as a wise man once said "we are true ecologists, the less JS is run the better for the enviroment" especially if the JS is written by me. 

**Note**: English is not my first language, I can understand it, I can-*ish* speak it but my writing skills might not be perfect. If you have any problem with this just contact me, I'll make sure to show you everithing I learned while I was doing my SQL exam on paper (you will recive very creative insults, blasphemy and all kinds of obscenities).

## The concept
For this project I'd like to make a meme-ish portfolio site. I would like to scare off non-technical people by creating a CLI web-application (because it makes no sense and I love it). So all the informations needed are accessible via a command like `resume` starts the download of my resume, `contacts` lists all my contact info, ecc... Plus some shenanigans like the implementation of the infamous `sl` command and a theme changer.

The site won't use any cookies, and it won't store any information just to be on the clear with european regulations and the GDPR (I don't have the time or the will power to be compliant so f*** it, I won't store any user data). 

## Roadmap
First of all I need to learn Go:
* simple command parser
* I/O components for the parser
* network interface and web server
* some kind of persistent memory

Then I need a view for my application:
* HTML and CSS
* the minimum amount of vanilla JS to make the website look at least functional
* HTMX playgorund

At this point I will need to learn HTMX:
* kinda-RESTfull API for trasporing HTML from the backend to the frontend

And then and only then:
* a funcional CLI web-application (it makes me "WHEEZE eheheheh" every time I think about how stupid this idea is)
