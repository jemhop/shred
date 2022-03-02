# Shred
#### _Multifunctional CLI file deletion program_

Shredder is a Freedesktop trash standard compliant CLI program, created because in my opinon rm is not a very human usable tool. Accidentally running rm -r on the wrong directory is never a fun time, and in my opinon `rm -rI` is not a very memorable way to safely delete a directory. 

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
- Improve installation instructions

## Installation

Shredder has been verified to work on Arch Linux, running KDE Plasma. No prebuilt binaries are included, but it is a standard go module and can be built accordingly.




