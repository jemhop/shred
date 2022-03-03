# Shred
#### _Multifunctional CLI file deletion program_

Shred is a Freedesktop trash compliant CLI program, created because in my opinon rm is not a very human usable tool. Accidentally running rm -r on the wrong directory is never a fun time, and in my opinon `rm -rI` is not a very memorable way to safely delete a directory. 

## Features
- Interacts with your trash directory intuitively (list, move to, recover, and permanently delete files in trash direcotry)
- Comes with the ability to overwrite files with output from /dev/urandom 5 times before deleting to avoid drive recovery ('shredding')
- Stylized and helpful output. No log spam.
- Ability to chain commands certain commands (for example, you can shred files in your trash bin)*

_*Planned, not yet implemented_

## Goals 
- Implement better argument handling to allow for command chaining 
- Benchmark performance vs dd/rm, try to match 
- Refactor code to reduce duplicate calls to stat
- Generally refactor code
- Set up AUR repo
- Improve installation instructions

## Installation

Shred has been verified to work on Arch Linux, running KDE Plasma. No prebuilt binaries are included, but it is a standard go module and can be built accordingly.


## Notes
This project is a work in progress. I would not reccomend letting 
your important data anywhere near it. I have barely started thinking about how it might handle files that recquire root etc. 
I'm gonna keep working on it, but for now, be careful. 




