package main
import"fmt"
func main(){
  Nombre, Apell_Pat, Apell_mat, Edad, Cuota, Acred := "Fernando","Diaz","Osorno",21,2100.00,true

  if Acred == true{
      if Cuota >2000.00 {
          fmt.Println(Nombre,Apell_Pat,Apell_mat,Edad,Cuota,Acred)
      }else{
          fmt.Println("Cuota incorecta")
      }
    }else{
    fmt.Println("No credita")
  }

  AumentoCuota:=Cuota

  MES := [12] string {"Enero","Febrero","Marzo","Abril","Mayo","Junio","Julio","Agosto","Septiembre","Octubre","Noviembre","Diciembre"}
  //tamano := len(MES)

  fmt.Println(MES[0:11])


  for  i:=1; i<=len(MES); i++ {
    if i==1{
      AumentoCuota=Cuota
    }else{
      AumentoCuota=(AumentoCuota*0.05)+AumentoCuota
    }
    switch i {
      case 1:
          fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",Cuota))
      case 2:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",AumentoCuota))
      case 3:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",AumentoCuota))
      case 4:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",AumentoCuota))
      case 5:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",AumentoCuota))
      case 6:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",AumentoCuota))
      case 7:
         AumentoCuota=(Cuota*0.05)+Cuota
         fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",AumentoCuota))
      case 8:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",AumentoCuota))
      case 9:
          AumentoCuota=(Cuota*0.05)+Cuota
          fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",AumentoCuota))
      case 10:
         AumentoCuota=(Cuota*0.05)+Cuota
         fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",AumentoCuota))
      case 11:
         AumentoCuota=(Cuota*0.05)+Cuota
         fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",AumentoCuota))
      case 12:
         AumentoCuota=(Cuota*0.05)+Cuota
         fmt.Println(MES[i-1]+fmt.Sprintf("|%6.2f|",AumentoCuota))
    }
  }
}
