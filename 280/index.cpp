#include <iostream>
#include <algorithm>
using namespace std;

bool isPythagoreanTriple(int a, int b, int c) {

    int sides[] = {a, b, c};

    sort(sides, sides + 3);

    return (sides[0] * sides[0] + sides[1] * sides[1] == sides[2] * sides[2]);

}

int main() {

    int a, b, c;
    cin >> a >> b >> c;
    if (isPythagoreanTriple(a, b, c)) {
        cout << "YES" << endl;
    } else {
        cout << "NO" << endl;
    }
    return 0;
}

