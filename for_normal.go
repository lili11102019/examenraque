package main
import"fmt"
func main(){
  Nombre, Apell_Pat, Apell_mat, Edad, Cuota, Acred := "Fernando","Diaz","OSorno",21,1500.00,true

  if Acred == true{
      if Cuota >2000.00 {
          fmt.Println(Nombre,Apell_Pat,Apell_mat,Edad,Cuota,Acred)
      }else{
          fmt.Println("Cuota incorrecta")
      }
    }else{
    fmt.Println("No credita")
  }

AumentoCuota:=Cuota

  for i:=2; i<=12; i++ {
    switch i {
      case 1:

      case 2:
          AumentoCuota=(Cuota*0.05)+Cuota

      case 3:
          AumentoCuota=(Cuota*0.05)+Cuota

      case 4:
          AumentoCuota=(Cuota*0.05)+Cuota

      case 5:
          AumentoCuota=(Cuota*0.05)+Cuota

      case 6:
          AumentoCuota=(Cuota*0.05)+Cuota

      case 7:
         AumentoCuota=(Cuota*0.05)+Cuota

      case 8:
          AumentoCuota=(Cuota*0.05)+Cuota

      case 9:
          AumentoCuota=(Cuota*0.05)+Cuota
          
      case 10:
         AumentoCuota=(Cuota*0.05)+Cuota
         
      case 11:
         AumentoCuota=(Cuota*0.05)+Cuota
         
      case 12:
         AumentoCuota=(Cuota*0.05)+Cuota
         
    }

}
}
