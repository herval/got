package main

import "math/rand"

var quotes = []string{
	"Bran thought about it. 'Can a man still be brave if he's afraid?' 'That is the only time a man can be brave,' his father told him.",
	"... a mind needs books as a sword needs a whetstone, if it is to keep its edge.",
	"Never forget what you are, for surely the world will not. Make it your strength. Then it can never be your weakness. Armour yourself in it, and it will never be used to hurt you.",
	"Fear cuts deeper than swords.",
	"Some old wounds never truly heal, and bleed again at the slightest word.",
	"Winter is coming.",
	"When you play a game of thrones you win or you die.",
	"The man who passes the sentence should swing the sword. If you would take a man's life, you owe it to him to look into his eyes and hear his final words. And if you cannot bear to do that, then perhaps the man does not deserve to die.",
	"Most men would rather deny a hard truth than face it.",
	"The things we love destroy us every time, lad. Remember that.",
	"Death is so terribly final, while life is full of possibilities.",
	"What is honor compared to a woman's love? What is duty against the feel of a newborn son in your arms . . . or the memory of a brother's smile? Wind and words. Wind and words. We are only human, and the gods have fashioned us for love. That is our great glory, and our great tragedy.",
	"Why is it that when one man builds a wall, the next man immediately needs to know what's on the other side?",
	"When the snows fall and the white winds blow, the lone wolf dies but the pack survives.",
	"And I have a tender spot in my heart for cripples and bastards and broken things.",
	"If I look back I am lost.",
	"Nothing burns like the cold.",
	"Once you’ve accepted your flaws, no one can use them against you.",
	"Laughter is poison to fear.",
	"Every flight begins with a fall.",
	"My brother has his sword, King Robert has his warhammer and I have my mind...and a mind needs books as a sword needs a whetstone if it is to keep its edge. That's why I read so much Jon Snow.",
	"'Oh, my sweet summer child', Old Nan said quietly, 'what do you know of fear? Fear is for the winter, my little lord, when the snows fall a hundred feet deep and the ice wind comes howling out of the north. Fear is for the long night, when the sun hides its face for years at a time, and little children are born and live and die all in darkness while the direwolves grow gaunt and hungry, and the white walkers move through the woods",
	"Different roads sometimes lead to the same castle.",
	"'What do we say to the Lord of Death?' 'Not today.'",
	"Life is not a song, sweetling. Someday you may learn that, to your sorrow.",
	"Give me honorable enemies rather than ambitious ones, and I'll sleep more easily by night.",
	"The things I do for love.",
	"You are your mother's trueborn son of Lannister.",
	"Summer will end soon enough, and childhood as well.",
	"A bruise is a lesson... and each lesson makes us better.",
}

func RandomQuote() string {
	return "\"" + quotes[rand.Intn(len(quotes)-1)] + "\""
}
