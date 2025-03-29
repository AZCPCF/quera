def draw_hollow_square(a, b):

    if b >= a:

        print("Wrong order!")

        return

    if (a - b) % 2 != 0:

        print("Wrong difference!")

        return

    margin = (a - b) // 2

    for i in range(a):

        for j in range(a):

            if i < margin or i >= a - margin or j < margin or j >= a - margin:

                print('*', end=' ')

            else:

                print(' ', end=' ')

        print()

a=int(input())

b=int(input())

draw_hollow_square(a, b)

