def fibonacci(n)
    if n <= 0
        return 0
    elsif n == 1 || n == 2
        return 1
    else 
        return fibonacci(n-1) + fibonacci(n-2)
    end
end


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
    puts ("Enter First Number")

end