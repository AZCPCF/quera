#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
    int a, b, c;
    cin >> a >> b >> c;

    vector<int> time(101, 0); 

    for (int i = 0; i < 3; i++) {
        int start, end;
        cin >> start >> end;
        for (int t = start; t < end; t++) {
            time[t]++;
        }
    }

    int cost = 0;
    for (int t = 1; t <= 100; t++) {
        if (time[t] == 1) {
            cost += a;
        } else if (time[t] == 2) {
            cost += 2 * b; 
        } else if (time[t] == 3) {
            cost += 3 * c; 
        }
    }

    cout << cost << endl;

    return 0;
}
