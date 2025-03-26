package main

func main() {
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Buy bread")
	todos.add("Buy tea")
	todos.toggle(0)
	todos.print()
}	
