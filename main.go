package main

import (
	"GoogleBingExtractor/FyneApp"
	"GoogleBingExtractor/Utils"
	"log"
)

func main()  {
	_, err := Utils.CreateMutex("GoogleBingExtractor")
	if err != nil{
		log.Println("Program prohibits multi-opening")
		return
	}
	fyne := FyneApp.NewFyneApp()
	err = fyne.InitApp()
	if err != nil{
		log.Panicln(err)
	}
	fyne.Run()
}
