# DCIM-Importer
Imports and organizes your Digital Camera IMages (DCIM) directory to selected destination
# Digital Camera Image Importer
The worst part of photography is having to import and organize your images. I had a back-log of 10k+ images on a number of SD cards, so I made a command-line program to efficiently copy files into an organized file structure. 

```
E:.
├───2019
│   ├───9 - September
│   ├───10 - October
│   ├───11 - November
│   └───12 - December
├───2020
│   ├───1 - January
│   ├───2 - February
│   ├───3 - March
│   ├───4 - April
│   ├───5 - May
│   ├───6 - June
│   └───7 - July
...
```
## Installation
For now, you can build the code yourself.
```
git clone https://github.com/Vrandus/DCIM-Importer.git
cd DCIM-Importer
go build import.go
```
## Usage
```
import <DCIM_src> <dst>
```
## Todo
* Integrate with Windows Task Scheduler 
*  Create GUI with Qt 
