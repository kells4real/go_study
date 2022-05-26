# Python equivalent to the Go calculator in the calculator package
# From Python 3.10 you can now use the switch statement
def subCal(num1, num2, sign):
    result = 0
    match sign:
        case "+":
            result = num1 + num2
        case "-":
            result = num1 - num2
        case "X" | "x" | "*":
            result = num1 * num2
        case "/":
            result = num1 / num2

    return result


def calculator():
    num1 = int(input("Enter first number: "))
    sign = input("Enter sign, e.g +, -, /: ")
    num2 = int(input("Enter second number: "))
    
    result = subCal(num1, num2, sign)
    print(result)
    
    while True:
        sign = input("Enter sign, e.g +, -, /: ")
        if sign.lower() == "exit":
            break
        else:
            num3 = int(input("Enter second number: "))
            result = subCal(result, num3, sign)
            print(result)
        
        
    
    