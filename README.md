# Github Releases Packet Manager

Github Releases Packet Manager (grpm) is a simple tool to handle installed releases from github.

basic usage:

this tool can track your installed releases, update them, etc.

commands:

 - install <release> -> install a new releases
 - update <release> -> updates specific release
 - upgrade -> update all releases
 - list -> show installed releases
 - release -> get 
 - info -> repo information
 - search -> search specific package -> start with this
 - new -> check if the installed release in it is latest version

 /* plans */
 
 1 - sprint (predicted at: 27 jan 2026): [Implement the config plan]
- One: We need to generate config a toml file for general configuration (Done) (Refactor phase)
- Two: We need to use the generated config information (Done) (Refactor phase)
- Three: user can change config file as he/she needs (Done) (Refactor phase)

 2 - sprint (predicted at: 28 jan 2026): [Implement the search with JSON accept]
- One: We need to fetch data from github as JSON (Done) (Refactor phase)
- Two: Then we start by marshling/decoding the string into JSON (Done) (Refactor phase)
- Third: print the info into the terminal (Done) (Refactor phase)
- Fourth: add caching technique/and fetch data from db and filesystem (Done) (Refactor phase)
- Fifth: clear cache operation (Done) (Refactor phase)

 3 - sprint (predicted at: 31 jan 2026): [Start working with info search operation]
- One: open a specific repo (Done) (Refactor phase)
- Two: fetch specified info like (name,readme,url,etc) (Done) (Refactor phase)

4 - sprint (predicted at: 01 Feb 2026): [Start working on fetching latest releases]
- One: fetch a releases from specific repo (Done) (Refactor phase)
- Two: display releases to a user with release information (Done) (Refactor phase)
- Three: add URL address of the release (Done) (Refactor phase)

 5 - sprint (predicted at: 02 Feb 2026): [Start working with instllation of files]
- One: Install a file by writing its name like (hish22/grpm) (Done) (Refactor phase)
- Two: Make sure the file installed based on the installation proccess (declared when we start this    process) (Refactor phase)
- Three: make sure to include the db to track installed packets/releases

 6 - sprint (predicated at 03 Feb 2026): [Start working with listting of installed packets]
- One: Hit db to fetch list of installed files
- Two: display information
