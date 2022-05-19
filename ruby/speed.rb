
def speedTest()
    start = Time.now()
    myList = []
    rand = Random.new()
    for i in 1..10000000 do
        myList.append(rand.rand(100))
    end

    # While Loop. This seems to be a bit faster than the For Loop,
    # The only language I've seen that
    # i = 0
    # while i < 100000 do
    #     i += 1
    #     myList.append(rand.rand(100))
    # end

    result = myList.sort()
    puts (((Time.now() - start).to_f * 1000)).round(4).to_s + "ms"

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


# def calculator() 
#     puts ("Enter First Number")
#     num1 = get
# end