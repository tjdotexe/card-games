# Card games
A number of card games played in the terminal.


## Ninety-Nine

### How to play

The game uses one standard deck of 52 cards. It uses 2 decks if there are four or more players.
This version will use 3 players.

Each player is given 5 tokens (3 if there are 4 or more players) and they are each dealt 3 cards from the deck. 
The remainder is used as a face-down stock pile. The face-up pile (center) always has a value of zero when starting.

At the beginning of each round, the current player will use one of their 3 cards and place it in the center.
The center pile will be added up, displayed then the current player will draw from the top of the stock pile
marking the end of their turn.

If a player cannot play a card without taking the face-up pile (center) over 99, their turn is over and the
player must place a token in the center. When a player loses all their tokens, they are out of the game.
The game continues until only one player remains.

### Special deck values

- Ace - increase the pile by 1 or 11, player's choice
- Four - pile value remains unchanged but direction reverses
- Nine - pile remains unchanged, and it's the next player's turn
- Ten - increases or decreases pile by 10, player's choice
- King - pile value becomes 99
- Jacks and Queens - increases the value by 10

### Rules

1. when there are only 2 players left, the four card's only special effect is that it counts as zero.
The reverse direction no longer applies and it's the other player's turn.
2. if the face-down stock pile is empty, the top card of the face-up pile is kept, but the remainders 
are shuffled and used for the new face-down stock pile.
