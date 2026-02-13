# Pokedex CLI

A command-line Pokedex application built in Go that lets you explore the world of Pokemon right from your terminal.

## Overview

This project is a REPL (Read-Eval-Print Loop) application that interacts with the [PokéAPI](https://pokeapi.co/) to fetch and display information about Pokemon. Look up stats, types, and other details about your favorite Pokemon without ever leaving the command line.

## Features

- **Interactive CLI**: Navigate through the Pokedex using simple text commands
- **Live Data**: Fetches real-time Pokemon data from the PokéAPI
- **Caching**: Implements caching to improve performance and reduce API calls
- **Pokemon Lookup**: Search for Pokemon by name and view their stats, types, and more

## Tech Stack

- **Language**: Go
- **External API**: [PokéAPI](https://pokeapi.co/)
- **Data Format**: JSON

## What I Learned

- Parsing JSON in Go
- Making HTTP requests in Go
- Building interactive CLI tools
- Implementing caching strategies for API optimization
- Local Go development and tooling best practices

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (latest version recommended)

### Installation

```bash
git clone https://github.com/yourusername/pokedex-cli.git
cd pokedex-cli
go build -o pokedex
./pokedex
