# Test 3

This is a test of the american broad cast system.

Lorem **ipsum** dolor *sit amet*, consectetur ***adipiscing elit***, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ac orci phasellus egestas tellus rutrum tellus pellentesque eu tincidunt. In vitae turpis massa sed elementum tempus egestas sed sed. [Erat nam at lectus](https://duckduckgo.com) urna duis. Dis parturient montes nascetur ridiculus mus mauris. Iaculis urna id volutpat lacus. Amet justo donec enim diam vulputate ut pharetra sit amet. Enim nec dui nunc mattis enim ut tellus. Adipiscing elit pellentesque habitant morbi tristique senectus et netus. Sapien nec sagittis aliquam malesuada bibendum arcu. Volutpat diam ut venenatis tellus in metus vulputate. Nisl nisi scelerisque eu ultrices vitae auctor eu augue ut. `Libero enim` sed faucibus turpis.

I really like markdown.    
There is a line break.

## WOW

Here is a blockquote:

> Dorothy followed her through many of the beautiful rooms in her castle.

And here is another:

> Dorothy followed her through many of the beautiful rooms in her castle.
>
> The Witch bade her clean the pots and kettles and sweep the floor and keep the fire fed with wood.

### But now We Need To Use Lists

Such as a list of languages.

- Python
- Golang
- Java
- C#

But we can also use...

#### Ordered Lists of Languages

1. Python
2. Golang
3. Java
4. C#

________________

## Lastly

To test some code.

```go
func AboutHandler(w http.ResponseWriter, r *http.Request, t *template.Template) {
	err := t.ExecuteTemplate(w, "about.gohtml", AboutData{Page: "/about"})

	if err != nil {
		log.Print(err)
		return
	}
}
```
