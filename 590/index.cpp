#include <iostream>



using namespace std;



long long gcd(long long a, long long b) {

    while (b != 0) {

        long long temp = b;

        b = a % b;

        a = temp;

    }

    return a;

}



long long lcm(long long a, long long b) {

    return (a / gcd(a, b)) * b;

}



int main() {

    long long a, b;

    cin >> a >> b;



    long long gcd_value = gcd(a, b);

    long long lcm_value = lcm(a, b);

    cout<<gcd_value<<" "<<lcm_value;



    return 0;

}

