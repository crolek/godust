godust
======

For server side rendering of dustjs using https://github.com/robertkrimen/otto

Dustjs has the nice ability of being able to be rendered on the server or client side. This package is a lightweight wrapper to help perform that rendering.

## Features
Loads a compiled template file, selects the proper template (based on name), and applies the data. It returns the rendered content.

#####future
- better error response (it's kind of lacking at the moment)
- Have godust make the rest call for you
- ability to swap from a dev to live data-url
- better examples of how to use this

##Use


#### func RenderDustjs
```(templatePath string, templateName string, jsonDataString string) string```

Returns the rendered content of a template.

A quick example (/example/godust-main.go)

```
package main

import (
	"fmt"

	"github.com/crolek/godust"
)

func main() {
	value := godust.RenderDustjs("../assets/test-template.js", "test", `{"name": "Chuck"}`)

	fmt.Println("value: ")
	fmt.Println(value)

}
```
That returns:
```Hello, Chuck
```

## Questions
Ping me on twitter [@crolek](http://twitter.com/crolek)

##Warning
This is very much an in-development/hacky-weekend project, so please keep that in mind.


## License

The MIT License (MIT)

Copyright (c) 2014 Chuck Rolek

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
