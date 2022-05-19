import imp
from python.calculator import calculator
from python.speed import speedTest
from python.functions import fibonacci

if "__main__" == __name__:
    speedTest()
    # calculator()
    print(fibonacci(10))