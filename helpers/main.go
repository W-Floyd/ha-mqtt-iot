package main

func main() {

	devices := DevicesInit()

	for _, d := range devices {
		f := GoFile{
			Name: d.Name,
		}
		f.Write()
	}

}
