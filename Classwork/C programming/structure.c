#include<stdio.h>
#include<string.h>

// struct employee{
//     int employee_id;
//     char name[50];
//     float salary;
// };


// int main(){
//     struct employee e1;
//     struct employee e2;

//     e1.employee_id = 101;
//     strcpy(e1.name,"Mohsin");
//     e1.salary = 20000;

//     printf("\nemployee id %d", e1.employee_id);
//     printf("\nname : %s", e1.name);
//     printf("\nsalary : %f", e1.salary);

//     e2.employee_id = 102;
//     strcpy(e2.name,"Joel");
//     e2.salary = 20040;

//     printf("\nemployee id %d", e2.employee_id);
//     printf("\nname : %s", e2.name);
//     printf("\nsalary : %f", e2.salary);

//     return 0;
// }


struct employee{
    int employee_id;
    char name[50];
    float salary;
};
struct Organisation  
{ 
  char organisation_name[20]; 
  char org_number[20]; 
    
  struct employee emp;  
}; 

int main(){
    struct Organisation org;  

  printf("The size of structure organisation : %ld\n",  
          sizeof(org)); 
    
  org.emp.employee_id = 101;   
  strcpy(org.emp.name, "Robert"); 
  org.emp.salary = 400000; 
  strcpy(org.organisation_name,  
         "GeeksforGeeks"); 
  strcpy(org.org_number, "GFG123768"); 
    
    
  // Printing the details 
  printf("Organisation Name : %s\n",  
          org.organisation_name);   
  printf("Organisation Number : %s\n",  
          org.org_number);   
  printf("Employee id : %d\n",  
          org.emp.employee_id);   
  printf("Employee name : %s\n",  
          org.emp.name);   
  printf("Employee Salary : %f\n",  
          org.emp.salary);   

    return 0;
}