#include <iostream>
using namespace std;


int main() {
  int a = 123;  // a: 123
  int* pa = &a;
  *pa = 321;  // a: 321
  cout << pa << "   "<< &a << endl;

}