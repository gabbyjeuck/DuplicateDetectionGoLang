# DuplicateDetectionGoLang

Data structure of files must be in .csv format and comma separated.

CSV Format:
The program expects barcode,code,YearWeek as headers.
- Where barcode (0), code (1), YearWeek (2).
- Specifically, the code column (1) is the only one being used to validate duplicates and MUST be within column (1).

Folder Format:
The program expects your data files to live in the "data" directory which should live at the root of the project.
- It will only detect files ending in .csv

Running the application:
You can either run the main.exe file OR
You can choose to use go commands.
- go run main.go

Altering the application:
If you choose to alter the application you should recompile the build.
- go build main.go 
- go run main.go OR run the main.exe

