package lib

import (
	"fmt"
	"os"
	"os/exec"
)

type Dot struct{
	name string
	edges []struct{from fmt.Stringer; to fmt.Stringer}
}

func NewDot(name string)*Dot{
	d := &Dot{name: name}
	d.edges = make([]struct{from fmt.Stringer; to fmt.Stringer}, 0, 1)
	return d
}

func (d *Dot)AppendEdge(from fmt.Stringer, to fmt.Stringer){
	d.edges = append(d.edges, struct {
		from fmt.Stringer
		to   fmt.Stringer
	}{from: from, to: to})
}

func (d *Dot)Layout(){
	fp, err := os.Create(d.name+".dot")
	if err != nil{
		panic("open file error")
	}
	_, err = fp.WriteString("digraph "+d.name+"{\n")
	if err != nil{
		panic("write file error")
	}
	// all edges
	for _, edge := range d.edges{
		_, err = fp.WriteString("\t"+edge.from.String()+" -> "+edge.to.String()+";\n")
		if err !=nil{
			panic("write error")
		}
	}

	_, err = fp.WriteString("\n}\n\n")
	if err != nil{
		panic("write file error")
	}
	err = fp.Close()
	if err != nil{
		panic("error")
	}
	// draw it
	cmd := exec.Command("dot", "-O", "-Tjpeg", d.name+".dot")
	err = cmd.Start()
	if err != nil{
		panic(err.Error())
	}
	err = cmd.Wait()
	if err != nil{
		panic(err.Error())
	}
}



