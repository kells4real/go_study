import time
import random as rand

# Speed test in python
def speedTest():
    start = time.time()
    myList = []

    # While loop
    i = 0
    # while i < 100000:
    #     i += 1
    #     myList.append(rand.randint(1, 100))
        
    # For loop, no visible difference in speed between this and the while loop
    for i in range(100000):
        myList.append(rand.randint(1, 100))
        
    myList.sort()
    print(f"{(time.time() - start) * 1000}ms")

if "__main__" == __name__:
    speedTest()