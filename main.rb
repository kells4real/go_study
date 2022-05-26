require "./ruby/speed"
require "./ruby/functions"

speedTest()
puts fibonacci(10)

def looped(aList)
    for a in aList
        puts a
    end
end
#Basic loop
aList = ["GET", "POST", "PUT", "DELETE"]
looped(aList)

calculator()


