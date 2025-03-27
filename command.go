package main

import "flag"

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index value & specify the new todo title, id:new_title")
    flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete")	
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all the todos"  )	

	// Parse all the arguments
	 
	flag.Parse()

	return &cf
}