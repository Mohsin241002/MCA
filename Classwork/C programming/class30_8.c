#include <stdio.h>


// int main(){
// union UnionName{
//     int int_member;
//     float float_member;
//     char char_member;
// };

// union Data {
//     int i;
//     float f;
//     char str[20];
// };

// union Data data;
// data.i = 10;
// data.f = 220.5;
// printf("%f", data.f);
// }


enum week {Mon, Tue, Wed, Thur, Fri, Sat, Sun};
int main(){
    enum week day; 
    day = Wed; 
    printf("%d", day);
    return 0;
}