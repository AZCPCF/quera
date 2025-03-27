#include <iostream>



using namespace std;







void printPascalsTriangle(int n) {



    for (int i = 0; i < n; i++) {



        int val = 1;



        for (int j = 0; j <= i; j++) {



            cout << val << " ";



            val = val * (i - j) / (j + 1);



        }



        cout << endl;



    }



}







int main() {



    int n;



    cin >> n;



    printPascalsTriangle(n);



    return 0;



}