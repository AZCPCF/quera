#include <iostream>

using namespace std;

int main(){

	int n=0;

	cin>>n;

	int a[n];

	int max=0;

	for(int i=0;i<n;i++){

		cin>>a[i];

	}

	for(int i=0;i<n;i++){

		max = max >= a[i] ? max : a[i];

	}

	cout<<max;

	return 0;

}