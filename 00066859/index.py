def convert_to_base(n, b):
    if n == 0:
        return "0"
    digits = []
    while n > 0:
        remainder = n % b
        if remainder < 10:
            digits.append(str(remainder))
        else:
            digits.append(chr(remainder - 10 + ord('A')))
        n //= b
    return ''.join(digits[::-1])


n, b = map(int, input().split())


result = convert_to_base(n, b)
print(result)
