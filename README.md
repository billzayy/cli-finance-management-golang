# Financial Management GoLang CLI

**Author**: [Bill Zay](https://github.com/billzayy)
<br>
**Type : Personal Project**

## Description of Project : 
**Financial Management GoLang CLI** is created for learning how to build a CLI Application on GoLang with some simple libraries :

* `flag` : A standard Go Library is using for creating flag on GoLang
* `github.com/alexeyco/simpletable` : A library is using for present a table view of data

We store a sample data into `budget.json`

Use this command to download all packages of project :
```go
go mod download
```

To build this project, run this command:
```go
go build ./cmd/main.go
```

---
### CLI Command & Flags : 

* `--list` : List a table of finance data
    ```bash
    ./main --list
    ```

* `--add` : Add your new item name
* `--price` : Input the price of the new item with type **Float** (10.2, 20.1, ..v..v. )
    ```bash
    ./main --add="<name>" --price=<price>
    ```

* `--addBudget` : Add your new budget to the Table
    ```bash
    ./main --addBudget=<budget>
    ```

* `--delete` : Choose which item of the table will be deleted
    ```bash
    ./main --delete=<index>
    ```

* `--updateName"` : Type the name you want to update
* `--index` : Choose which item of table will be updated
    ```bash
    ./main --updatePrice="<update name>" --index=<index>
    ```

* `--updatePrice` : Input the price you want to update with the type is **Float** (10.0, 20.15, ..v..v)
* `--index` : Choose which item of table will be updated
    ```bash
    ./main --updateName=<update price> --index=<index>
    ```
---
$${\color{lightgreen}HAPPY \space CODING \space 😉 \space !!!}$$