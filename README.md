# Welcome to TREMIS

Tremis is an in memory database system that uses RAM to store data unlike the relational databases that stores in the disk.
Tremis was created keeping in mind the well known REDIS database and I wanted to build something of my own .

Since Redis is no more open-source my sole aim was to make Tremis open source and also add in functionalities that Redis doesn't have. Overall Tremis can be indeed used for setting up small databases and can be much efficient since that is what is in-memory databases are known for

This project is still under development and shall function with much more improved functionalities and also cool new features. 

Overall,  Tremis does the following for now : 

1. Basics Key- Value Pair Commands  -> SET, GET, DEL
2. Increment and Decrement
3. Lists, Sets
4. PUB/SUB
5. Disk Storage (additional)

The reason for adding disk storage was solely because people should not miss out the jest of disk storage capabilities and those who are new to database should experience that as well.

There are several more additions to be done to the project but before that I must create a proper website that serves as a platform to download Tremis locally. 

BTW, if you want to contribute this project locally on your machine just use `netcat` and connect it to the port specified in the code (`6379`) 

I am soon going to create a setup wizard where users would be able to access TREMIS locally on their host machines and level up things slowly as time progresses

Until Then,
Signing Off,
DumbPrism


