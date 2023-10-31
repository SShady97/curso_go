// Introduction
// In some English accents, when you say "two for" quickly, it sounds like "two fer". Two-for-one is a way of saying that if you buy one, you also get one for free. So the phrase "two-fer" often implies a two-for-one offer.

// Imagine a bakery that has a holiday offer where you can buy two cookies for the price of one ("two-fer one!"). You go for the offer and (very generously) decide to give the extra cookie to a friend.

// Instructions
// Your task is to determine what you will say as you give away the extra cookie.

// If your friend likes cookies, and is named Do-yun, then you will say:

// One for Do-yun, one for me.
// If your friend doesn't like cookies, you give the cookie to the next person in line at the bakery. Since you don't know their name, you will say you instead.

// One for you, one for me.
// Here are some examples:

// Name	Dialogue
// Alice	One for Alice, one for me.
// Bohdan	One for Bohdan, one for me.
// One for you, one for me.
// Zaphod	One for Zaphod, one for me.

package main

import (
	"fmt"
)

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	message := fmt.Sprintf("One for %s, one for me", name)
	if name == "" {
		message = "One for you, one for me"
	}
	return message
}

func main() {
	fmt.Println(ShareWith("Jose"))
	fmt.Println(ShareWith(""))
}
