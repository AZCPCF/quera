def gcd(a,b):

    while (b!=0):

        temp = b

        b = a%b

        a = temp

    return a

    

# end def

num1 = int(input())

num2 = int(input())



print(gcd(num1,num2))

