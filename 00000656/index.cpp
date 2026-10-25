#include <iostream>
using namespace std;

long long determinant(int a[10][10], int n) {
    if (n == 1)
        return a[0][0];

    long long det = 0;
    int temp[10][10];

    for (int p = 0; p < n; p++) {
        int row = 0, col = 0;

        for (int i = 1; i < n; i++) {
            col = 0;

            for (int j = 0; j < n; j++) {
                if (j == p)
                    continue;

                temp[row][col] = a[i][j];
                col++;
            }

            row++;
        }

        long long sign = (p % 2 == 0) ? 1 : -1;
        det += sign * a[0][p] * determinant(temp, n - 1);
    }

    return det;
}

int main() {
    int n;
    cin >> n;

    int a[10][10], b[10][10];

    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
            cin >> a[i][j];

    for (int i = 0; i < n; i++)
        for (int j = 0; j < n; j++)
            cin >> b[i][j];

    long long detA = determinant(a, n);
    long long detB = determinant(b, n);

    if (detA % 2 != 0 && detB % 2 != 0)
        cout << "Danial" << endl;
    else
        cout << "Farzad" << endl;

    return 0;
}