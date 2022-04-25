# A CLI app to calculate the whole day difference between two given dates

This command line application takes two dates from the user and calculates the 
whole day difference between the two without using the `time` package in Go.

### Usage
- Download the binary file applicable to your operating system from the bin folder.
- Execute the binary file without any arguments. (ex: ./date-diff-linux)
- Input the dates when prompted and press Enter. Correct date constraints are given below.

### Date input constraints
- The date format should be of form (dd|d)/(mm|m)/yyyy. For example if the date is 21st of March 1993, 
it should be given as 21/5/1993 or 21/05/1993.
- The date should be between 01/01/1900 and 31/12/2999.

### CI/CD
GitHub actions is configured to run on pushes and pulls from the master branch. It will run the tests
and build the artifacts for different operating systems.