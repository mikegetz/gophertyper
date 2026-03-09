# gophertyper
A TUI typing game about gophers by gophers

![2026-03-08 00-12-34](https://github.com/user-attachments/assets/5626fa9e-f429-44b9-b409-4edd5b422235)

## Will you make it ten waves or will gopher see sky?
Waves get progressively harder. Game difficulty scales to terminal size.

![2026-03-08 01-12-36](https://github.com/user-attachments/assets/28f971f4-cd5e-4efc-82a0-f8c1745d8f4c)


## Install
### Homebrew
```
brew tap mikegetz/gophertyper
brew install gophertyper
```
### Shell
#### curl
```
sh -c "$(curl -fsSL https://raw.githubusercontent.com/mikegetz/gophertyper/main/tools/install.sh)"
```
#### wget
```
sh -c "$(wget -qO- https://raw.githubusercontent.com/mikegetz/gophertyper/main/tools/install.sh)"
```

## How to Start the Game
After installing, launch the game from your terminal:

```
gophertyper
```

- Make sure your terminal window is at least 100 characters wide for the best experience.
- Use the on-screen key hints: `[Esc]` to quit, `[Space]` to pause/resume.
- Type the first letter of a gopher's word to select it, then type the rest of the word to eliminate the gopher.
- Survive all 10 waves to win!

## Report
At the end of each game you'll see a report with the following stats:

- **Gophers Per Minute (GPM)** — The number of gophers you eliminated per minute of active play time (excluding pauses and wave transitions).
- **Words Per Minute (WPM)** — Your typing speed using the standard WPM formula: `(correct keypresses + completed words) / 5 / minutes`. Each completed word adds one character to account for the space separator, matching how typing tests like MonkeyType calculate WPM. A "word" in standard WPM is defined as 5 characters.
- **Accuracy** — The percentage of your keypresses that were correct: `correct keypresses / total keypresses`.
- **Correct Keypresses** — The number of correct keypresses out of total keypresses.
