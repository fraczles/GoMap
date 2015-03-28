
===============================================================================================================
Michelle Hochman(mhochma1@binghamton.edu), Alex Fraczak(afracza1@binghamton.edu)

Code was tested on a 64-bit Dual core i7 MacBook Pro, and MacBook Air.

We compiled our code via:
go build emerging.go cmap.go
===============================================================================================================


===============================================================================================================
We ran out code via:
./emerging -readers=A -askers=B -askdelay=10 -reducedelay=100 -infiles="pg1041.txt,pg1103.txt,pg1107.txt,pg1112.txt,pg1120.txt,pg1128.txt,pg1129.txt,pg1514.txt,pg1524.txt,pg2235.txt,pg2240.txt,pg2242.txt,pg2243.txt,pg2264.txt,pg2265.txt,pg2267.txt" > Aread_Bask.txt
WHERE A is the number of readers, and B is the number of askers
One thing to note: We moved all of the data BESIDES ask.txt into the same directory as the executable. 
===============================================================================================================


===============================================================================================================
Findings:
We observed that the number of readers does not significantly affect the execution time of this program.
This is because each reader gets their own file to process concurrently.


The number of askers however, greatly increases the execution time. Because readers cause all other threads to block, there is more competition over the shared resource. 

You can see the exact output for each individual test run in it's associated .txt, i.e
1 readers 1 askers --> 2read_2ask.txt
4 readers 2 askers --> 2read_4ask.txt
...
...
16 readers 32 askers -->16_read_32ask.txt
===============================================================================================================
