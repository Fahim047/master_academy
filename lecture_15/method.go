package main

import "fmt"

type Email interface{
	Write(string,string,string)
	Send()
	Read()
}

type Test struct{
	To string
	From string
	Subject string
	Message string
}

func (t Test) Write(to string,from string,subject string){
	fmt.Println(to,from,subject)
}
func (t Test) Send(){
	fmt.Println(t.To, "email sent")
}
func (t Test) Read(){
	fmt.Println(t.From,"receieved from")
}



type Doctor struct{
	Name string
	Education string
	Age int
	Salary float32
}

func (d Doctor)Speak(){
	fmt.Println(d.Name,"can speak")
}
func (d Doctor)getName() string{
	return d.Name
}
func (d Doctor)getEducation() string{
	return d.Education
}
func (d Doctor)getAge() int{
	return d.Age
}
func (d Doctor)getSalary() float32{
	return d.Salary
}

func main(){
	//var d Doctor{"Biswajit", "MBBS",30, 50000.00} //error
	//var d Doctor{Name:"Biswajit", Education:"MBBS", Age:30, Salary:50000.00,} //error
	//var d = Doctor{Name:"Biswajit", Education:"MBBS", Age:30, Salary:50000.00}
	/*
	var d Doctor
	d.Name= "Biswajit"
	d.Education= "MBBS"
	d.Age= 30
	d.Salary= 50000.50
	fmt.Println(d.Name, d.Education, d.Age,d.Salary)
	d.Speak()
	fmt.Println(d.getName())
	fmt.Println(d.getSalary())
	*/

	//var e Email
	//fmt.Println(e)
	
	var tst Test
	tst.To="billahmdmostain@gmail.com"
	tst.From="fahimulislam58@gmail.com"
	tst.Subject="test email"
	tst.Message="this is a test email"
	tst.Write(tst.To, tst.From, tst.Subject)

}