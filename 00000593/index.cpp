#include <iostream>

using namespace std;



bool isPrime(int x) {

    if (x <= 1) return false;

    for (int i = 2; i * i <= x; i++) {

        if (x % i == 0) return false;

    }

    return true;

}



int main() {

    int n, sum = 0, b = 0;

    cin >> n;



    int temp = n;

    while (temp > 0) {

        sum += temp % 10;

        temp /= 10;

    }



    while (b < sum) {

        n++;

        if (isPrime(n)) {

            b++;

        }

    }



    cout << n << endl;

    return 0;

}

