Genvars
=======

An enviroment workspace manager for Go(lang) apps, that allows you to manage all the enviroment variables for various apps on your dev machine, and on your production machines with a unified interface.

## How it works
Just prefix your variables with the appropriate app, and set an APP_ENV for each app on your dev and production machines, and your game!

### Dev Machine Enviroment Vars
```
APP1_MYSQL_URL = ...
APP1_APP_ENV = DEVELOPMENT

APP2_MYSQL_URL = ...
APP2_APP_ENV = DEVELOPNENT
etc ...
```

### Production Machine (APP1) Enviroment Vars
```
MYSQL_URL = ...        	<= (Notice the clean variables)
APP1_APP_ENV = PRODUCTION
```

### Usage
```
package main

import (
	"github.com/jsimnz/genvars"
)

func main() {
	m, err := genvars.NewManager("APP1")
    if err != nil {
        panic(err)
    }
    
	connectMySQL(m.Getenv("MYSQL_URL")) // Will use appropriate vars
	... etc ...
}
```

### Configuration
```
genvars.NewManager("APP", genvars.ManagerOptions{
    EnviromentTag: "APP_ENV",
    DevTagValue: "DEVELOPMENT",
    ProdTagValue: "PRODUCTION",
})
```
<table>
<thead>
    <tr><td><b><u>Option</u></b></td><td><b><u>Function</u></b></td></tr>
</thead>

<tbody>
    <tr>
        <td>EnviromentTag</td>
        <td>Which enviroment variable to use to detect the current mode, (ie. PRODUCTION, DEVELOPMENT). Must be prefixed on the machine, not in the code</td>
    </tr>
    <tr>
        <td>DevTagValue</td>
        <td>The value of the EnviromentTag that indicates currently on a development machine</td>
    </tr>
     <tr>
        <td>ProdTagValue</td>
        <td>The value of the EnviromentTag that indicates currently on a production machine</td>
    </tr>
</tbody>
</table>

## Installation
    go get github.com/jsimnz/genvars

## Why
If you're like most devs, you'll be developing a few apps on your machine at once. Each one of those apps require their own enviroment variables (database, API keys, etc...), so you go through the process of setting them up locally, and on your production machines, so you dont have to keep switching you code before pushing to production.

#### Oh Crap!
Now you have two apps that both use a MySQL database, so your dev machine has MySQL_URL configured twice - one for each app - and they've over written them selves. Next thing you know you've foolishly pushed to production, and you cause a meltdown at your Datacenter. Bad day all around.

## License
The MIT License (MIT)

Copyright (c) 2014 John-Alan Simmons

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
