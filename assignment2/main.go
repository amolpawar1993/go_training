package main

import (
	"fmt"
	"math"
	"strings"
)

type salaryInfo struct {
	basic_amnt, variable_payout, hra float64
}
type contactInfo struct {
	city, mobile, email string
}

type Employee struct {
	emp_id , leave_balance int
	emp_fname, emp_lnane, emp_dob, rating string
	contactInfo
	salaryInfo
}

func main()  {

	emp := createEmployee("amol", "pawar", "16-May-1993", "B+", 80036,
		8, 24000, 10000, 12000, "pune", "9890262528", "amolpawar@calsoftinc.com")
	fmt.Println("################### Read Data #####################")
	fmt.Println("Read Employee Details:\n")
	emp.read_emp_data()
	fmt.Println("#####################################################")
	fmt.Println("################### Update Data #####################")
	fmt.Println("Update Employee Detail:\n")
	emp.update_emp_data(16000)
	fmt.Println("Read Employee Details after Update in variable payout:\n")
	emp.read_emp_data()
	fmt.Println("#####################################################")
	fmt.Println("################### Calculate net salary #####################")
	fmt.Println("Calculate Net Salary")
	net_salary := emp.calculateNetSalary(8)
	fmt.Println("Net Salary = ", net_salary)
	fmt.Println("#####################################################")
	fmt.Println("################### Appraisal Cycle #####################")
	fmt.Println("\nAfter Appraisal Net Salary")
	new_salary := emp.calculateAppraisal()
    fmt.Println("Net Salary = ", new_salary)
	fmt.Println("#####################################################")
	fmt.Println("################### Emergency Data #####################")
	fmt.Printf("\n Emergency Details: \n City = %s \n Mobile = %s \n", emp.contactInfo.city, emp.contactInfo.mobile)
	fmt.Println("#####################################################")
	fmt.Println("################   Map Functionality ##############")
	fmt.Println("List of all close friends")
	frieds := get_all_friends()
	for id, v := range frieds{
		fmt.Printf("EmpID: %d --- Name: %s \n", id, v)
	}
	fmt.Println("#####################################################")
	fmt.Println("################### Delete Data #####################")
	fmt.Printf("\n Fight with %s and %s \n", frieds[80042],frieds[80041])
	frieds_remove := []int{80042, 80041}
	fmt.Println("\nRemoved from close friends list")
	fights_with_friends(frieds, frieds_remove)
	fmt.Println("Now close friend list: ")
	for id, v := range frieds{
		fmt.Printf("EmpID: %d --- Name: %s \n", id, v)
	}
	fmt.Println("#####################################################")
}

func createEmployee(emp_fname, emp_lnane, emp_dob, rating string,
    emp_id, leave_balance int,
	basic_amnt, variable_payout, hra float64,
	city, mobile, email string) (Employee){
     emp := Employee{
	 emp_id:        emp_id,
	 leave_balance: leave_balance,
	 emp_fname:     emp_fname,
	 emp_lnane:     emp_lnane,
	 emp_dob:       emp_dob,
	 rating:        rating,
	 contactInfo:   contactInfo{city:city, mobile:mobile, email:email},
	 salaryInfo:    salaryInfo{basic_amnt:basic_amnt, variable_payout:variable_payout, hra:hra},
 }
     return emp
}

func (e Employee) read_emp_data(){
	fmt.Printf("Employee Detail Data: %+v \n", e)
}

func (e *Employee) update_emp_data(variable_payout float64){
	e.variable_payout = variable_payout
}

func (e *Employee) calculateNetSalary (leaves float64) float64{
	lb:= float64(e.leave_balance) - leaves
	newlb:= math.Abs(lb)
	var newBasic float64
	if (lb < 0) {
		lb = 0
		newBasic = e.basic_amnt/30 * float64(30.0 - newlb)
	} else {
		newBasic = e.basic_amnt
	}
	e.leave_balance = int(lb)
	totalSalary:= e.variable_payout + newBasic + e.hra
	netSalary:= totalSalary - 200
	return netSalary
}
func (e Employee) calculateAppraisal () float64{
	var newBasic float64
	if (e.rating == "A"){
		newBasic = e.basic_amnt + e.basic_amnt * 0.3
	} else if (e.rating == "B+") {
		newBasic = e.basic_amnt + e.basic_amnt * 0.2
	} else if (e.rating == "B") {
		newBasic = e.basic_amnt + e.basic_amnt * 0.1
	}
	newSalary:= newBasic + e.variable_payout + e.hra
	return newSalary
}

func get_all_friends()  map[int]string{
	emp1 := createEmployee("boby", "mohanty", "16-May-1996", "B+", 80038,
		8, 21000, 11000, 10000, "pune", "9890262555", "bm@calsoftinc.com")
	emp2 := createEmployee("chandan", "barole", "16-June-1990", "B", 80039,
		10, 18000, 9000, 10000, "pune", "9890299088", "cb@calsoftinc.com")
	emp3 := createEmployee("girish", "warrier", "18-Jun-1989", "A", 80040,
		19, 25000, 15000, 13000, "pune", "9890223322", "gw@calsoftinc.com")
	emp4 := createEmployee("shahrin", "shaikh", "23-Apl-1997", "B", 80041,
		5, 14000, 8000, 9000, "pune", "9880909043", "ss@calsoftinc.com")
	emp5 := createEmployee("nikhil", "jathar", "17-May-1994", "B", 80042,
		11, 21000, 17000, 11000, "pune", "9844664455", "nj@calsoftinc.com")

	friendsDetails := map[int]string{
		emp1.emp_id: strings.Join([]string{emp1.emp_fname, emp1.emp_lnane}," "),
		emp2.emp_id: strings.Join([]string{emp2.emp_fname, emp2.emp_lnane}," "),
		emp3.emp_id: strings.Join([]string{emp3.emp_fname, emp3.emp_lnane}," "),
		emp4.emp_id: strings.Join([]string{emp4.emp_fname, emp4.emp_lnane}," "),
		emp5.emp_id: strings.Join([]string{emp5.emp_fname, emp5.emp_lnane}," "),
	}
	return friendsDetails
}

func fights_with_friends(m map[int]string, ids []int){
	for _, value:= range ids{
		delete(m, value)
	}
}