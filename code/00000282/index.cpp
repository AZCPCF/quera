#include <iostream>

using namespace std;



int main() {

    int n;

	cin>>n;

	int c;

	for(int i=1;i<n;i++){

		if(n%i==0){

			c+=i;

		}

	}

	if(c==n){

		cout<<"YES";

   		return 0;

	}

	cout<<"NO";

	return 0;

}

