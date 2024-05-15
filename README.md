### Dark-Terminal

Ever got bored at work? and needed a quick break from the day and wanted to play dark souls? Well dont do that... thats a terrible idea. But what if you could run dark souls from your terminal?

...

Well you cant...

But I made this clicker version of Dark Souls for your terminal.

Please give it a try!

### Rules

- When open it runs like a normal clicker except `spacebar` is your click.
- Navigate the terminal ui with your arrow keys, along with enter to buy or go through menus
- Time is continually counted in `dark-terminal`, but the dont worry it is not continuously running.

### How time is handled

When the game is live and open time is handled on a basic clock, no tricks. Its honestly very very simple. For this game I wanted it to still run in the background though but i didnt want the game to hog any memory from the cpu as I want this game to be a continual play. To do this I log the exact time of the close for the application, then on opening it again it loads the previous time and the current and calculates the time inbetween in secs before applying a score to your player. Pretty neat right!

### Installation

TBD...

- homebrew
- winget
- go
