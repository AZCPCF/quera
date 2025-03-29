def find_numbers(m, s):
    
    if s == 0:
        return "-1 -1" if m > 1 else "0 0"
    if s > 9 * m:
        return "-1 -1"

    
    smallest = []
    sum_remaining = s
    for i in range(m):
        for digit in range(0 if i > 0 else 1, 10):
            if sum_remaining - digit <= 9 * (m - i - 1):
                smallest.append(digit)
                sum_remaining -= digit
                break

    
    largest = []
    sum_remaining = s
    for i in range(m):
        for digit in range(9, -1 if i > 0 else 0, -1):
            if sum_remaining - digit >= 0:
                largest.append(digit)
                sum_remaining -= digit
                break

    return "".join(map(str, smallest)) + " " + "".join(map(str, largest))



m, s = map(int, input().split())


result = find_numbers(m, s)
print(result)
