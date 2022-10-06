package main

import (
	"log"
	"strings"
	"testing"

	"github.com/tutividela/gophexcercises/HtmlParser/helpers"
	"golang.org/x/net/html"
)


func Test01OneLink(t *testing.T) {
	tag := "a"
	answers := make(map[string]string)
	answers["/dog"] ="Something in a span Text not in a span Bold text!"

	ex1 :=`<a href="/dog"><span>Something in a span</span>Text not in a span<b>Bold text!</b></a>`
	 
	docEx1, err := html.Parse(strings.NewReader(ex1))
	if err != nil {
		log.Fatal(err)
	}
	links := helpers.SearchHtmlLinks(docEx1,tag)
	
	if len(links) != len(answers) {
		t.Errorf("Expected lenght %d , got length %d",len(answers),len(links))
	}
	for _,v := range links {
		a,ok := answers[v.Href]
		if !ok {
			t.Errorf("Missing Href %s \n",v.Href)
		}
		if a != v.Text{
			t.Errorf("Expected %s , got %s",answers[v.Href],v.Text)
		} 
	}

}

func Test02LinkWithTagsUnderIt(t *testing.T){
	tag := "a"
	answers := make(map[string]string)
	answers["https://www.twitter.com/joncalhoun"] = "Check me out on twitter"
	answers["https://github.com/gophercises"] = "Gophercises is on Github !"

	ex2 :=`<html>
<head>
  <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
</head>
<body>
  <h1>Social stuffs</h1>
  <div>
    <a href="https://www.twitter.com/joncalhoun">
      Check me out on twitter
      <i class="fa fa-twitter" aria-hidden="true"></i>
    </a>
    <a href="https://github.com/gophercises">
      Gophercises is on <strong>Github</strong>!
    </a>
  </div>
</body>
</html>`
	 
	docEx2, err := html.Parse(strings.NewReader(ex2))
	if err != nil {
		log.Fatal(err)
	}
	links := helpers.SearchHtmlLinks(docEx2,tag)
	
	if len(links) != len(answers) {
		t.Errorf("Expected lenght %d , got length %d",len(answers),len(links))
	}
	for _,v := range links {
		a,ok := answers[v.Href]
		if !ok {
			t.Errorf("Missing Href %s \n",v.Href)
		}
		if a != v.Text{
			t.Errorf("Expected %s , got %s",answers[v.Href],v.Text)
		} 
	}
}

func Test03LinkInsideTags(t *testing.T){
	tag := "a"
	answers := make(map[string]string)
	answers["#"] = "Login"
	answers["/lost"] = "Lost? Need help?"
	answers["https://twitter.com/marcusolsson"] = "@marcusolsson"

	ex3 :=`<!DOCTYPE html>
<!--[if lt IE 7]> <html class="ie ie6 lt-ie9 lt-ie8 lt-ie7" lang="en"> <![endif]-->
<!--[if IE 7]>    <html class="ie ie7 lt-ie9 lt-ie8"        lang="en"> <![endif]-->
<!--[if IE 8]>    <html class="ie ie8 lt-ie9"               lang="en"> <![endif]-->
<!--[if IE 9]>    <html class="ie ie9"                      lang="en"> <![endif]-->
<!--[if !IE]><!-->
<html lang="en" class="no-ie">
<!--<![endif]-->

<head>
  <title>Gophercises - Coding exercises for budding gophers</title>
</head>

<body>
  <section class="header-section">
    <div class="jumbo-content">
      <div class="pull-right login-section">
        Already have an account?
        <a href="#" class="btn btn-login">Login <i class="fa fa-sign-in" aria-hidden="true"></i></a>
      </div>
      <center>
        <img src="https://gophercises.com/img/gophercises_logo.png" style="max-width: 85%; z-index: 3;">
        <h1>coding exercises for budding gophers</h1>
        <br/>
        <form action="/do-stuff" method="post">
          <div class="input-group">
            <input type="email" id="drip-email" name="fields[email]" class="btn-input" placeholder="Email Address" required>
            <button class="btn btn-success btn-lg" type="submit">Sign me up!</button>
            <a href="/lost">Lost? Need help?</a>
          </div>
        </form>
        <p class="disclaimer disclaimer-box">Gophercises is 100% FREE, but is currently in beta. There will be bugs, and things will be changing significantly over the coming weeks.</p>
      </center>
    </div>
  </section>
  <section class="footer-section">
    <div class="row">
      <div class="col-md-6 col-md-offset-1 vcenter">
        <div class="quote">
          "Success is no accident. It is hard work, perseverance, learning, studying, sacrifice and most of all, love of what you are doing or learning to do." - Pele
        </div>
      </div>
      <div class="col-md-4 col-md-offset-0 vcenter">
        <center>
          <img src="https://gophercises.com/img/gophercises_lifting.gif" style="width: 80%">
          <br/>
          <br/>
        </center>
      </div>
    </div>
    <div class="row">
      <div class="col-md-10 col-md-offset-1">
        <center>
          <p class="disclaimer">
            Artwork created by Marcus Olsson (<a href="https://twitter.com/marcusolsson">@marcusolsson</a>), animated by Jon Calhoun (that's me!), and inspired by the original Go Gopher created by Renee French.
          </p>
        </center>
      </div>
    </div>
  </section>
</body>
</html>`
	 
	docEx3, err := html.Parse(strings.NewReader(ex3))
	if err != nil {
		log.Fatal(err)
	}
	links := helpers.SearchHtmlLinks(docEx3,tag)
	
	if len(links) != len(answers) {
		t.Errorf("Expected lenght %d , got length %d",len(answers),len(links))
	}
	for _,v := range links {
		a,ok := answers[v.Href]
		if !ok {
			t.Errorf("Missing Href %s \n",v.Href)
		}
		if a != v.Text{
			t.Errorf("Expected %s , got %s",answers[v.Href],v.Text)
		} 
	}
}

func Test04LinkWithCommentInside(t *testing.T) {
	tag := "a"
	answers := make(map[string]string)
	answers["/dog-cat"] = "dog cat"

	ex4 :=`<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>`
	 
	docEx4, err := html.Parse(strings.NewReader(ex4))
	if err != nil {
		log.Fatal(err)
	}
	links := helpers.SearchHtmlLinks(docEx4,tag)
	
	if len(links) != len(answers) {
		t.Errorf("Expected lenght %d , got length %d",len(answers),len(links))
	}
	for _,v := range links {
		a,ok := answers[v.Href]
		if !ok {
			t.Errorf("Missing Href %s \n",v.Href)
		}
		if a != v.Text{
			t.Errorf("Expected %s , got %s",answers[v.Href],v.Text)
		} 
	}
}

func Test05LinkWithHTMLTagAndTextInside(t *testing.T){
	tag := "a"
	answers := make(map[string]string)
	answers["/dog"] = "Something in a span Text not in a span Bold text!"

	ex5 :=`<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>`
	 
	docEx5, err := html.Parse(strings.NewReader(ex5))
	if err != nil {
		log.Fatal(err)
	}
	links := helpers.SearchHtmlLinks(docEx5,tag)
	
	if len(links) != len(answers) {
		t.Errorf("Expected lenght %d , got length %d",len(answers),len(links))
	}
	for _,v := range links {
		a,ok := answers[v.Href]
		if !ok {
			t.Errorf("Missing Href %s \n",v.Href)
		}
		if a != v.Text{
			t.Errorf("Expected %s , got %s",answers[v.Href],v.Text)
		} 
	}
}

func Test06LinkInsideLink(t *testing.T) {
	tag := "a"
	answers := make(map[string]string)
	answers["/dog"] =""
	answers["/im-inside"]="Hello world this is a B span!"

	exLL :=`<a href="/dog">
  <a href="/im-inside">
    <p>
        Hello world
    </p>
    <span>this is a <b>B</b> span! </span>
  </a>
</a>`
	 
	docExLL, err := html.Parse(strings.NewReader(exLL))
	if err != nil {
		log.Fatal(err)
	}
	links := helpers.SearchHtmlLinks(docExLL,tag)
	
	if len(links) != len(answers) {
		t.Errorf("Expected lenght %d , got length %d",len(answers),len(links))
	}
	for _,v := range links {
		a,ok := answers[v.Href]
		if !ok {
			t.Errorf("Missing Href %s \n",v.Href)
		}
		if a != v.Text{
			t.Errorf("Expected %s , got %s",answers[v.Href],v.Text)
		} 
	}

}