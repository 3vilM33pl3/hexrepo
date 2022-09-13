![hexcloud](https://github.com/3vilM33pl3/hexrepo/actions/workflows/build-hexcloud.yml/badge.svg)

![hexcom](https://github.com/3vilM33pl3/hexrepo/actions/workflows/build-hexcom.yml/badge.svg)

![hexcom](https://github.com/3vilM33pl3/hexrepo/actions/workflows/deploy-hexcloud.yml/badge.svg)

# Overview
This is the monorepo for [hexcloud](./src/hexcloud/README.md), [hexcli](./src/hexcli/README.md) and [hexcom](./src/hexcom/README.md). 

The Unreal Engine plugin [HexWorld](https://github.com/3vilM33pl3/HexWorld) is kept seperate (for now) because of the size of the 3D content. 

## [hexcloud](./src/hexcloud/README.md)
Backend server and main access point via gRPC API for the game.

## [hexcli](./src/hexcli/README.md)
Command line interface to manipulate game content

## [hexcom](./src/hexcom/README.md) 
Hexcom is the client library which can be used to communicate with the hexcloud backend


