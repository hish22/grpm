# Github Releases Packet Manager

Github Releases Packet Manager (grpm) is a simple tool to handle installed releases from github.

basic usage:

this tool can track your installed releases, update them, etc.

commands:

 - install <release> -> install a new releases
 - update <release> -> updates specific release
 - upgrade -> update all releases
 - list -> show installed releases
 - find -> check if release is install
 - info -> repo information
 - search -> search specific package -> start with this
 - new -> check if the installed release in it is latest version

 /* plans */
 
 1 - sprint (predicted at: 27 jan 2026): [Implement the config plan]
- One: We need to generate config a toml file for general configuration (Done) (Refactor phase)
- Two: We need to use the generated config information (Done) (Refactor phase)
- Three: user can change config file as he/she needs (Done) (Refactor phase)

 2 - sprint (predicted at: 28 jan 2026): [Start working with info search operation]
- One: open a specific repo
- Two: fetch specified info like (name,readme,url,etc)

 3 - sprint (predicted at: 29 jan 2026): [Start working with instllation of files]
- One: Install a file by writing its name like (hish22/grpm)
- Two: Make sure the file installed based on the installation proccess (declared when we start this    process)
- Three: make sure to include the db to track installed packets/releases
