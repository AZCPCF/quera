#include <iostream>
using namespace std;

void ShowFibNth(long int current, long int next) {
    
    cout << current << endl;

    
    if (current <= 0) {
        return;
    }

    
    ShowFibNth(next, current - next);
}

int main() {
    long int n, n_plus_1;
    cin >> n >> n_plus_1;

    
    if (n < 0 || n_plus_1 < 0 || n_plus_1 < n || n >= 1000000 || n_plus_1 >= 1000000) {
        cerr << "Invalid inputs. Please ensure 0 <= n < n+1 < 1,000,000." << endl;
        return 1;
    }

    
    ShowFibNth(n, n_plus_1);

    return 0;
}
