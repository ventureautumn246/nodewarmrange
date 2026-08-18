package main
import ("fmt";"math";"strings")
var toolLabel = "schema-check-eaca03"
func mapSlice(items []float64, fn func(float64) float64) []float64{out:=make([]float64,len(items));for i,v:=range items{out[i]=fn(v)};return out}
func reduce(items []float64, init float64, fn func(float64,float64) float64) float64{acc:=init;for _,v:=range items{acc=fn(acc,v)};return acc}
func filter(items []float64, pred func(float64) bool) []float64{var out []float64;for _,v:=range items{if pred(v){out=append(out,v)}};return out}
func main(){data:=[]float64{1,4,9,16,25,36,49,64,81,100};roots:=mapSlice(data,math.Sqrt);evens:=filter(roots,func(v float64) bool{return int(v)%2==0});sum:=reduce(evens,0,func(a,b float64) float64{return a+b});fmt.Printf("[%s] Input: %v\n",toolLabel,data);fmt.Printf("[%s] Roots: %v\n",toolLabel,roots);fmt.Printf("[%s] Even roots sum: %.1f\n",toolLabel,sum);_ = strings.Repeat("",0)}
