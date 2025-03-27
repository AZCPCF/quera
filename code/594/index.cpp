#include <iostream>

#include <vector>

#include <cmath>



using namespace std;



vector<int> convertToBase(int number, int base) {

    vector<int> digits;

    while (number > 0) {

        digits.push_back(number % base);

        number /= base;

    }

    return digits;

}

bool isAlternatingSumZero(const vector<int>& digits) {

    int sum1 = 0, sum2 = 0;

    for (size_t i = 0; i < digits.size(); ++i) {

        if (i % 2 == 0) {

            sum1 += digits[i];

        } else {

            sum2 += digits[i];

        }

    }

    return sum1 == sum2;

}



int main() {

    int a, b;

    cin >> a >> b;

    vector<int> digits = convertToBase(a, b);



    if (isAlternatingSumZero(digits)) {

        cout << "Yes" << endl;

    } else {

        cout << "No" << endl;

    }



    return 0;

}

