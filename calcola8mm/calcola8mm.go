/*
calcola prezzo di X minuti Y secondi di film a 16fps con rotoli da 8mm lunghi 66ft FOMA
*/

package calcola8mm

func Calcola(ore, min, sec int, prezzo float64) (float64,int){
  var secondiGir,rulliNec int

  secondiGir=soloSecondi(ore,min,sec)

  rulliNec=(secondiGir/317)
  if secondiGir%317!=0{
    rulliNec+=1
  }

  return float64(rulliNec)*prezzo,rulliNec

}

func soloSecondi(ore, min, sec int) (secTot int){
  secTot+=sec
  if ore>0{
    secTot+=(ore*60)*60
  }
  if min>0{
    secTot+=min*60
  }
  return
}
