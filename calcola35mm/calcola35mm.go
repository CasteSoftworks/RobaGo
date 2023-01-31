/*
calcola prezzo di X minuti Y secondi di film a 16fps con rotoli da 35mm lunghi 144 fotogrammi
*/

package calcola35mm

func Calcola(ore, min, sec int, prezzo float64) (float64, int){
  var secondiGir,rulliNec int

  secondiGir=soloSecondi(ore,min,sec)
  frameNec:=secondiGir*16

  rulliNec=(frameNec/144)
  if rulliNec%144!=0{
    rulliNec+=1
  }

  return float64(rulliNec)*prezzo, rulliNec

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
