## Sypher Example

This application madly prints credentials to console 😳🤓

### How did I create this project?

**Created project directory:**

    $ mkdir sypher-example && cd sypher-example

**Initialized new module:**

    $ go mod init github.com/sertangulveren/sypher-example

**Created ```main.go```:**

    $ touch main.go

**Edited ```main.go``` as follows:**
```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Sypher")
}
```

**Ran app and confirmed: It works fantastically:**
```sh
$ go run main.go 
Hello Sypher
```

**Installed sypher package:**

    $ go get -u github.com/sertangulveren/sypher

**Installed sypher cli:**

    $ go install github.com/sertangulveren/sypher/sypher@latest

**Generated a credential file named myconfig via sypher to store API keys to use in my project:**

```sh
$ sypher gen myconfig
Created: myconfig
Done!
```

**Checked the project directory:**

Files generated for myconfig in sypher folder.

* ```myconfig.enc``` encrypted file.
* ```myconfig.key``` key file.
* ```sypher.go``` to register encrypted file as embedded.
```sh
$ tree
.
├── README.md
├── go.mod
├── main.go
└── sypher
    ├── myconfig.enc
    ├── myconfig.key
    └── sypher.go

1 directory, 6 files
```

**Ignored key files:**

It just creates or modifies ```.gitignore``` file to ignore key files. You can add manually ```sypher/*.key``` to your .gitignore file instead of use the command.

    $ sypher gitignore

**Updated the credentials:**
It opens vim editor by default.

    $ sypher edit myconfig

*To use another editor like nano, code etc. set EDITOR variable. Example:*

    $ EDITOR=code sypher edit myconfig

**Changed and saved the content as below and closed the editor:**

```sh
AWS_ACCESS_KEY=niceaccesskey
AWS_ACCESS_SECRET_KEY=hereisthesecretk3y
HOLY_API_KEY=123
```

**Edited the ```main.go``` file again:**
```go
package main

import (
	"fmt"
	"github.com/sertangulveren/sypher"
	_ "github.com/sertangulveren/sypher-example/sypher"
)

func main() {
	sypher.Load(sypher.Config{
		Name: "myconfig",
	})

	awsKey := sypher.Get("AWS_ACCESS_KEY")
	awsSecret := sypher.Get("AWS_ACCESS_SECRET_KEY")
	holyKey := sypher.Get("HOLY_API_KEY")

	// print my environment variables
	fmt.Println(awsKey)
	fmt.Println(awsSecret)
	fmt.Println(holyKey)
}
```

**Ran app and confirmed again: It works!:**

```sh
$ go run main.go 
niceaccesskey
hereisthesecretk3y
123
```

**Built the application:**

    go build -o sypher-example main.go

**Ran it in another instance/server:**

    ./sypher-example

*Before running, I assigned the key value in the myconfig.key file to the SYPHER_MASTER_KEY variable. If you want to do this manually with the command:*

    SYPHER_MASTER_KEY=mykeycontent ./sypher-example

### you can clone this project:

**With SSH:**

    git clone git@github.com:sertangulveren/sypher-example.git

**OR HTTPS:**

    git clone https://github.com/sertangulveren/sypher-example.git

You can create ```sypher/myconfig.key``` file with content:

```RFbD67TI3smTyVsGd6Xav1yu00ZAMPTA```

Run: ```go run main.go```

---

Instead of creating key file, you can also run it by assigning the SYPHER_MASTER_KEY variable.
