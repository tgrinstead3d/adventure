package main

import (
	"fmt"
	"strings"
)

var name string
var direction string
var action string

func main() {
	fmt.Println("You find yourself in a desolate mountain pass, surrounded by towering peaks and blanketed in heavy snowfall. Your group of adventurers presses forward, seeking refuge from the biting cold within the safety of a nearby cave. The wind howls, and the air is thick with the swirling snowflakes that obscure your vision.")
	fmt.Println("As you approach the cave, relief washes over you, anticipating warmth and shelter. The entrance looms ahead, a dark recess in the rock. However, just as you reach the safety of the cave, a deafening roar echoes through the mountains. The ground trembles beneath your feet, and you turn to witness a massive avalanche thundering down from the slopes above.")
	fmt.Println("The cascading wall of snow engulfs the mountain pass, burying everything in its path, including the entrance to the cave. Panic sets in as your heart races – you're now trapped inside the cave, cut off from your companions who were not able to make it in time. The once welcoming refuge now feels like a cold, isolated chamber. ")
	fmt.Println("What is your name?")
	fmt.Scanln(&name)
	fmt.Println(strings.ToUpper(name), "searches among their pockets and belongings for any items that may be of any use. They find only a pocket knife, a handful of gold coin, and two potions; one a health potion and the other a poison potion.")
	fmt.Println("WELCOME TO THE CAVE OF DOOM! - A text based adventure game by Tyler Grinstead.")
	fmt.Println("NOTE: Once prompted for your action, you can type in any of the capitalized words from the prompt to proceed.")
	fmt.Println(" ")
	fmt.Println("In the dim light filtering through the entrance, the cave reveals itself as a rocky chamber with moss-covered walls. The limited illumination casts long shadows, creating an atmosphere of mystery and uncertainty. The air is damp, and the temperature inside is noticeably colder than you expected.")
	fmt.Println("To your LEFT, a narrow tunnel stretches into the darkness, its path winding through the cold stone. STRAIGHT ahead, another passage beckons, its depths hidden from view. The mossy rock formations block any further exploration in other directions, enclosing you in an eerie silence.")
	fmt.Println("As you survey the cave, you notice a suspicious lump of dirt on the ground near your feet. The lump seems out of place in the rocky environment. Without realizing the significance, you decide to ignore the pile of dirt and you head further into the cave. Which direction do you choose?")
	fmt.Println("Please choose left or straight")
	fmt.Scanln(&direction)
	direction = strings.ToLower(direction)
	for {
		switch direction {
		case "left":
			fmt.Println("*****")
			fmt.Println("Curiosity leading the way, you decide to venture into the tunnel to the left. As you proceed, the sound of running water becomes more pronounced, echoing off the walls of the narrow passage. The subtle vibration beneath your feet and along the cave walls suggests a subterranean river coursing through the depths. The ambient noise creates an eerie symphony, combining the rushing water with the natural hum of the cave. The path continues to wind, inviting you to explore the secrets hidden within the subterranean labyrinth.")
			fmt.Println("Go STRAIGHT or BACK?")
			fmt.Scanln(&direction)
			direction = strings.ToLower(direction)

			switch direction {
			case "straight":
				fmt.Println("*****")
				fmt.Println("You chose straight")
				break
			case "back":
				fmt.Println("*****")
				fmt.Println("You chose back")
				break
			default:
				fmt.Println("Please choose STRAIGHT or BACK")
			}

		case "straight":
			fmt.Println("*****")
			fmt.Println("As you cautiously proceed straight into the dark path, the air becomes even colder, and the oppressive darkness seems to intensify. The narrow corridor winds deeper into the cavern, and your senses heighten as you navigate the unknown. Suddenly, without warning, the floor beneath you gives way, and you find yourself tumbling into a tight pit.")
			fmt.Println("After a moment of disorientation, you realize you're not alone in this cramped space. Through a small crack in the pit, a dim light flickers, resembling the glow of a fire. The confined area restricts your movement, and the hum of what sounds like many large insects fills the air, creating an eerie ambiance.")
			fmt.Println("You assess your options in this predicament. The crack in the pit offers a glimpse into an unknown space beyond, illuminated by the mysterious light. What will you do? EXPLORE the source of the light and the insect-like sounds, or attempt to find a way to CLIMB out and continue your journey through the cavern?")
			fmt.Scanln(&action)
			action = strings.ToLower(action)

			switch action {
			case "explore":
				fmt.Println("*****")
				fmt.Println("You chose to slide through the small crack and investigate the room ahead.")
				break
			case "climb":
				fmt.Println("*****")
				fmt.Println("You try to climb back up out of the pit, but the walls are too slick and you are unable to gain any traction. You decide to explore the room ahead.")
				break
			default:
				fmt.Println("Please choose straight or back")
			}
		}
	}
}
