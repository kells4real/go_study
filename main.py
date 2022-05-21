from python.calculator import calculator
from python.speed import speedTest
from python.functions import fibonacci, arguments

def looped(aList):
    for i in aList:
        print(i)

if "__main__" == __name__:
    speedTest()
    # calculator()
    print(fibonacci(10))
    # var = 78.6
    # print(arguments(1, 2, 5, "Great", "ironic", var, True))
    aList = ["GET", "POST", "PUT", "DELETE"]
    looped(aList)