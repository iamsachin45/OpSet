# OpSet 

An optimzation of built in set for unsigned interger. 

It consist of basic operation like Add, remove, find. 

# How it works : 
let's assume the OpSet is empty and starts with a single uint64 which is initialized to 0:

OpSet: [00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000]

To insert the number 6, we need to set the 7th bit (since bit positions start from 0) in the first uint64.
OpSet: [00000000 00000000 00000000 00000000 00000000 00000000 00000000 01000000]
