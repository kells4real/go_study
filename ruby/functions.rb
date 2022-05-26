def fibonacci(n)
    if n <= 0
        return 0
    elsif n == 1 || n == 2
        return 1
    else 
        return fibonacci(n-1) + fibonacci(n-2)
    end
end

# Using switch in ruby
def subCal(num1, num2, sign)
    result = 0
    case sign
    when "+"
        result = num1 + num2
    when "-"
        result = num1 - num2
    when "*", "X", "x"
        result = num1 * num2
    when "/"
        result = num1 / num2
    end
end


def calculator() 
    print "Enter First Number: "
    num1 = gets()
    print "Enter Second Number: "
    num2 = gets()

    result = subCal(num1, num2, sign)
    puts result

    while True
        print "Enter sign, e.g +, -, /: "
        sign = gets()
        if sign.lower() == "exit"
            break
        else
            num3 = int(input("Enter second number: "))
            result = subCal(result, num3, sign)
            print(result)
        end
    end

end