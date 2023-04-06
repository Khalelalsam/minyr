package main

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"strconv"
	"strings"
)


func main() {
   var input string 
 lines :=opnelese() 

 scanner := bufio.NewScanner(os.Stdin)  
 for scanner.Scan() {   input = scanner.Text()
   if input == "q" || input == "exit" { 
   fmt.Println("exit")
    os.Exit(0)

} else if input == "co" {
    fmt.Println("Konverterer alle målingene gitt i grader Celsius til grader Fahrenheit.")  
    fmt.Print("you want new file? yes or no")

	var uinput string 
	fmt.Scan(&uinput)
        scanner.Scan()
     if uinput == "yes" { newfile, err := os.Create("kjevik-temp-fahr-20220318-20230318.csv")     
      if err != nil { 
      log.Println(err)   

      }

      defer newfile.Close()
      writer := bufio.NewWriter(newfile)
      defer writer.Flush()


for i := 1; i <= 16754; i++ {  
    line := lines[i] 
    fields := strings.Split(line, ";")

temp, err := strconv.ParseFloat(fields[len(fields)-1], 64)  
     if err != nil {   
    //log.Println(err)   
      continue      }

fahrenheit := float64(temp*9/5 + 32)   
   ls := fmt.Sprintf("%2.1f\n", fahrenheit)
  
   lastIndex := strings.LastIndex(line, ";")
if lastIndex != -1 { 
      line = line[:lastIndex]
      line += ";"  

      ls2 := fmt.Sprint(line, ls)
      writer.WriteString(ls2)
      writer.Flush()       }

}     } else {     

    fmt.Print("skal ikke copiere")
    }} else if input == "avr" { 
    fmt.Print("c or f")
    var uinput string
    fmt.Scan(&uinput) 
    scanner.Scan() 
   if uinput == "c" {  
   sum := 0.0

    for i := 1; i <= 16754; i++ { 
     line := lines[i]  
     fields := strings.Split(line, ";")
       temp, err := strconv.ParseFloat(fields[len(fields)-1], 64)

if err != nil {
       log.Println(err)
       continue      }

       sum += temp     } 
       fmt.Printf("gjennomsnittstemperatur (C) er : %0.2f", sum/16754)
} else if uinput == "f" {
     sum2 := 0.0
     for i := 1; i <= 16754; i++ {
      line := lines[i]
      fields := strings.Split(line, ";")


temp, err := strconv.ParseFloat(fields[len(fields)-1], 64) 
      if err != nil {
      
      log.Println(err) 
      continue     

      }   



    fahrenheit := conv.CelsiusToFahrenheit(temp)

    sum2 += fahrenheit

}
      fmt.Printf("gjennomsnittstemperatur (F) er : %0.2f", sum2/16754)  
   } 
   } else { 
   fmt.Println("Venligst velg convert, average eller exit:")   }  }  }


func opnelese() []string {  
var lines []string
  fill, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv") 
 if err != nil {   log.Println(err)   }

defer fill.Close() 
  scanner := bufio.NewScanner(fill)
   for scanner.Scan() { 
  lines = append(lines, scanner.Text()) 
  }
   return lines
 }
