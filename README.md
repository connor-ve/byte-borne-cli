### Dark-Terminal >> _NEW_ ByteBorne: Idle Ascension

Ever got bored at work? and needed a quick break from the day and wanted to play dark souls? Well dont do that... thats a terrible idea. But what if you could run dark souls from your terminal?

...

Well you cant...

But I made this clicker version of Dark Souls for your terminal.

Please give it a try!

A souls-borne theme game to run in your terminal.

You will use AFK style leveling to enter center arenas or areas to level up your chacter at center character levels or areas bosses will appear. If they do you `sometimes` have the option to fight or move on.

There are different modes for your character :

- Training
- Exploring (Only one that actually progress your to boss)
- Resting (When healing or respawning)
- Sales
  `not selectable`
- upgrading
- boss fights, handled with terminal commands

Boss Fights will be semi interactive.

- you have a bank of skills you have learned or collected
- the boss will have between 1-10 attacks, after the boss has attacked or visa versa you can select one of your attacks
- You can select, or create a new attach based on the ancient scrolls. (if you have collected an ancient scroll this is possible)

Exploring

- If your character is set to explore, while you are away you collect items and battle lesser enemies some enemies make items or drop and so on. this is handled again by randomization.
- behind the scenes you collect up to a certain amount of afk time. When the users runs a command the code chackes the mode and if in exploring mode will calculate how many encounters, damage recieved, items, and xp the character has achieved. then you recieve the menu data in a print.

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

## Ideas for the future

- ~~Json based mod files~~
  - YAML was used for ease of reading
- ~~Moving character sprite in the terminal~~
  - Still potentially viable but not for now, maybe puzzles in future
- Color themes
  - Added to prints for now
- ~~Screen needs stay on constantly~~
  - Decided this is no longer needed for the game to have a purpose as it is Idle
