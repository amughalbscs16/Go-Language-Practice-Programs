package main

import fmt
func main()
{
	//Static Type of Data in Go.
	//var card string = "Ace of Space" (Bottom line means same)
	card := newCard() //:= Only for declaratio and initialization;
	fmt.Println(card)
}

func newCard() int  {
 	return 5
}