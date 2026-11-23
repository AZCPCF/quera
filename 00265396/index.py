n = int(input())
s = input()

if n != len(s):
    exit(1)

arr = [[0] * n for _ in range(3)]

for i in range(n):
    if s[i] == "1":
        arr[0][i] = ".*."
        arr[1][i] = ".*."
        arr[2][i] = ".*."
    else:
        arr[0][i] = "***"
        arr[1][i] = "*.*"
        arr[2][i] = "***"

for row in arr:
    print("".join(row))