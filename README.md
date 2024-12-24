# dnd-cli

`dnd-cli` is a command-line interface designed for Dungeon Masters to streamline the management of their Dungeons & Dragons campaigns. It offers a range of tools to handle dice rolling, campaigns, combat, players, and inventory.

---

## Features

### ✅ Dice Rolling
- Roll standard dice (d4, d6, d8, d10, d12, d20) in NdX format (e.g., `2d6`, `1d20`).
- Supports advanced options:
  - Percentage rolls (`2d10 -p`).
  - Display expected averages (`-a`).

### ✅ Root Command
- Provides help and guidance for using the CLI.
- Lists all subcommands and their descriptions.

---

## Roadmap

### 🚧 Campaign Management (`campaign.go`)
- Create, update, list, and delete campaigns.
- Persist campaigns in structured data (e.g., JSON or YAML).

### 🚧 Combat Tracker (`combat.go`)
- Manage combat rounds.
- Track initiative and turn progression.
- Provide commands for starting and ending combats.

### 🚧 Player Management (`player.go`)
- Create, update, list, and delete player characters.
- Link players to campaigns and manage their stats.

### 🚧 Inventory Management (`inventory.go`)
- Add, remove, and list items for campaigns and players.
- Provide a central inventory repository.

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/avstudio1/dnd-cli.git
   cd dnd-cli

