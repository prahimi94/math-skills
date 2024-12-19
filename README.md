# math-skills
**Objectives**

The purpose of this project is to calculate the following:

Average
Median
Variance
Standard Deviation
Instructions

The program reads from a file and prints the result of each statistic mentioned above. The data in the file will be presented as the following example:

189
113
121
114
145
110
...

**This data represents a statistical population: each line contains one value.**

after running the program it prints the results in the following manner (the following numbers are only examples):


Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
Please note that the values are rounded integers.

## To run the program

1. **Clone the Repository:**
    ```bash
    git clone https://01.gritlab.ax/git/prahimi/math-skills.git

## Running the Program

1. Navigate to the math-skills Folder:

    ```bash
    cd math-skills

2. Run the Program:

    ```bash
    go run main.go data.txt


## Testing the Program
1. Ensure Docker is Installed And Running.
2. Download the [Test File](https://assets.01-edu.org/stats-projects/stat-bin-dockerized.zip).
3. Run the Command Twice For First Time:
    ```bash
    ./run.sh math-skills
        
It will create a data.txt file with random numbers and return the average, median, variance, and standard deviation for  the numbers in the data.txt file.

4. You should then use this data.txt file in your math-skills folder to run the program based on that and compare the results with those from the test case.
