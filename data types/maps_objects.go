package mainimport "fmt"func main() {      //  elements := make(map[string]string)      //  elements["H"] = "Hydrogen"      //  elements["He"] = "Helium"      //  elements["Li"] = "Lithium"      //  elements["Be"] = "Beryllium"      //  elements["B"] = "Boron"      //  elements["C"] = "Carbon"      //  elements["N"] = "Nitrogen"      //  elements["O"] = "Oxygen"      //  elements["F"] = "Fluorine"      //  elements["Ne"] = "Neon"      // shorter way      elements := map[string]string{           "H": "Hydrogen",           "He": "Helium",           "Li": "Lithium",           "Be": "Beryllium",           "B": "Boron",           "C": "Carbon",           "N": "Nitrogen",           "O": "Oxygen",           "F": "Fluorine",           "Ne": "Neon",      }      // Multidimentional objects      elementsExpanded := map[string]map[string]string{           "H": map[string]string{                "name":"Hydrogen",                "state":"gas",           },           "He": map[string]string{                "name":"Helium",                "state":"gas",           },           "Li": map[string]string{                "name":"Lithium",                "state":"solid",           },      }      fmt.Println(elements["Li"])      if elements["Un"] == "" {          // returns nothing if actually just printed          fmt.Println(elements["Un"])      }      // another way of doing the above is      name, ok := elements["Un"]      fmt.Println(name, ok)      // For multidimentional object checking      if el, ok := elements["Li"]; ok {          fmt.Println(el["name"], el["state"])      }}