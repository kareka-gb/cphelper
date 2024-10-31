# How tools should work 

## Requirement
There are two main tools, I want to create:
- `cpgetter`: This takes contest url as input and creates a folder with `.cpp` file for each problem and their respective test cases.
    1. get contest name, platform name from contest link
	2. Based on platform, read the target folder path from configuration.
    3. what should be the name of the configuration file and where should it be present in the OS?
    4. If the config file is empty, what should be the target path
    5. based on the platform, find out the number of problems
	6. create a folder for the contest in the target folder
    7. Create sub folders for each sub problems
    8. Create a submission link for each problem
    9. Get and Create the specified testcases for each problem
- `cprun`: This takes solution file path as input, executes the test cases against the code. It tells us in detail the inputs, outputs, expected outputs of test cases.
    1. Compile the solution using solution file path
    2. Get the testcases's paths, run the compiled binary against these test cases.
    3. For each test case, it will run the binary and compare with expected output 
    4. Display prettily the outcomes of each test case
- `cpsubmit`: This takes solution file path as input, submits the solution to the judge.
    1. Where should I get the submission link?
    2. How should I let the judge know the owner of this solution?
    3. If the solution to the above problem is some kind of token based authentication system, how should I manage it for different kinds of judges?


## Probable Solution 
1. Storing credentials in configuration and creating a seperate session with Judge specific configuration for http calls and sessions, this might be troublesome because of companies like cloudflare.
2. Identify authentication system of each judge and create a judge specific configuration for http calls, this might be troublesome because we might need to constantly change the configuration
