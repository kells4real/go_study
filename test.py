import time
import random as rand

# Speed test in python
def main():
    start = time.time()
    myList = []

    for i in range(100000):
        myList.append(rand.randint(1, 100))
        
    myList.sort()
    print(f"{(time.time() - start) * 1000}ms")

if "__main__" == __name__:
    main()