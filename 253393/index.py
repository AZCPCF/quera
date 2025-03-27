def count_live_neighbors(i, j, n, grid):

    directions = [(-1, -1), (-1, 0), (-1, 1),

                  (0, -1),         (0, 1),

                  (1, -1), (1, 0), (1, 1)]

    count = 0

    for di, dj in directions:

        ni, nj = i + di, j + dj

        if 0 <= ni < n and 0 <= nj < n:

            count += grid[ni][nj]

    return count



def game_of_life(n, grid):

    new_grid = [[0] * n for _ in range(n)]

    for i in range(n):

        for j in range(n):

            live_neighbors = count_live_neighbors(i, j, n, grid)

            if grid[i][j] == 1:

                if live_neighbors < 2 or live_neighbors > 3:

                    new_grid[i][j] = 0

                else:

                    new_grid[i][j] = 1

            else:

                if live_neighbors == 3:

                    new_grid[i][j] = 1

    return new_grid



n = int(input())

grid = [list(map(int, input().split())) for _ in range(n)]

new_grid = game_of_life(n, grid)

for row in new_grid:

    print(' '.join(map(str, row)))

