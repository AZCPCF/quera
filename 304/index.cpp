#include <iostream>

#include <iomanip>

#include <cmath>

using namespace std;



int main() {

    long double Base;

    unsigned int exp;

    cin >> Base >> exp;



    long double res = pow(Base, exp);

    cout << fixed << setprecision(3) << res << endl;



    return 0;

}

