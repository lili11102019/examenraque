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
          fmt.Println("Enero "+fmt.Sprintf("%f",Cuota))
      case 2:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println("Febrero "+fmt.Sprintf("%f",AumentoCuota))
      case 3:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println("Marzo "+fmt.Sprintf("%f",AumentoCuota))
      case 4:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println("Abril "+fmt.Sprintf("%f",AumentoCuota))
      case 5:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println("Mayo "+fmt.Sprintf("%f",AumentoCuota))
      case 6:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println("Junio "+fmt.Sprintf("%f",AumentoCuota))
      case 7:
         AumentoCuota=(Cuota*0.05)+Cuota
         fmt.Println("Julio "+fmt.Sprintf("%f",AumentoCuota))
      case 8:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println("Agosto "+fmt.Sprintf("%f",AumentoCuota))
      case 9:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println("Septiembre "+fmt.Sprintf("%f",AumentoCuota))
      case 10:
         AumentoCuota=(Cuota*0.05)+Cuota
         fmt.Println("Octubre "+fmt.Sprintf("%f",AumentoCuota))
      case 11:
         AumentoCuota=(Cuota*0.05)+Cuota
         fmt.Println("Noviembre "+fmt.Sprintf("%f",AumentoCuota))
      case 12:
         AumentoCuota=(Cuota*0.05)+Cuota
         fmt.Println("Diciembre "+fmt.Sprintf("%f",AumentoCuota))
    }

}
}
