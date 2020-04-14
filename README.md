# Map-Sync
Tool to synchronise local map directory with community maps repository.

Supports MacOS and Windows 10.

## Download
Download the appropriate binary for your system from the pre-built releases.
[Releases](https://github.com/infinity-tts-community/map-sync/releases)

Unzip and place the binary somewhere convenient for you to run.

## Usage
For Windows it's possbile to double click the .exe to run the tool.
It's recommended that if you wish to view the output that you run
from within the command prompt.

*Windows*
```
# 'cd' to the location of the downloaded tool.
# Initial run will download a copy of all maps to "Saved Objects/Community Maps"
>map-sync
Maps directory not found
Cloning maps repository to C:\Users\Rick Sanchez\Documents\My Games\Tabletop Simulator\Saves\Saved Objects\Community Maps...
Enumerating objects: 34, done.
Counting objects: 100% (34/34), done.
Compressing objects: 100% (28/28), done.
Total 34 (delta 10), reused 21 (delta 4), pack-reused 0
Done

# Periodically run map-sync again to grab latest contributions.
>map-sync
Updating maps...
Already up to date
```

*MacOS*
```
```
