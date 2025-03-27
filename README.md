# GO-TODO-CLI

#### TODO app for CLI made using GO

## How to run?

### 1. List all the existing TODOs

```zsh
go run ./ -list
```

<img src="https://www.kushcreates.com/images/go-todo-cli/go%20run%20-list.png" />

### 2. Add a new TODO

```zsh
go run ./ - add "New todo"
```

<img src="https://www.kushcreates.com/images/go-todo-cli/go%20run%20-add.png" />

## 3. Edit an existing TODO

```zsh
go run ./ -edit "2: New todo edit"
```

<img src="https://www.kushcreates.com/images/go-todo-cli/go%20run%20-edit.png" />

## 4. Toggle TODO (Set it to completed: "✅" )

```zsh
go run ./ -toggle 3
```

<img src="https://www.kushcreates.com/images/go-todo-cli/go%20run%20-toggle.png" />

## 5. Delete an existing TODO

```zsh
go run ./ -del 2
```

<img src="https://www.kushcreates.com/images/go-todo-cli/go%20run%20-del.png" />