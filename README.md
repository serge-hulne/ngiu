# ngiu
A SwiftUI inspired GUI library for go, based on giu
![Screenshot 2024-11-15 at 11 57 03](https://github.com/user-attachments/assets/430cdc61-afe6-4da2-a0b9-71cf24464039)

# Example of use:

```
// ############################################
// ### Custom FormData with data input handling
// ############################################

package main

import (
	. "github.com/serge-hulne/ngiu"

	"github.com/AllenDang/giu"
)

// #######
// Main()
// #######

var formData FormData

func main() {
	wnd := giu.NewMasterWindow("Password encrypted wallet", 600, 400, 0)
	wnd.Run(loop)
}

func loop() {
	giu.SingleWindow().Layout(
		VStack(
			giu.Label("Enter a new service / password combination"),

			// Main widget
			NewWidget(formData.FormDataLayout()).
				PaddingTop(10).
				PaddingLeft(50),

			// Flexible spacer to push items below it to the bottom
			giu.Custom(func() {
				_, height := giu.GetAvailableRegion()
				giu.Dummy(0, height-30).Build() // Adjust height to control space
			}),

			// Status label at the bottom
			NewWidget(giu.Label("Status: OK")).
				BackgroundColor(NewColorByName(ColorNames.Basic.Blue)).
				ForegroundColor(NewColorByName(ColorNames.Vivid.Lime)).
				FontWeight("bold").
				PaddingTop(10),
		),
	)
}

// ###################
// ## FormData Object
// ###################

type FormData struct {
	Service string
	Pwd     string
	Output  string
}

func (f *FormData) FormDataLayout() giu.Layout {
	return giu.Layout{
		// Service
		giu.Label("Service:"),
		giu.InputText(&f.Service).Hint("Enter a service name"),
		giu.Dummy(0, 10),
		// Password
		giu.Label("Password:"),
		giu.InputText(&f.Pwd).Hint("Enter a password"),
		// Action
		giu.Dummy(0, 10),
		giu.Button("Add new password").OnClick(func() {
			f.FormDataSubmit()
		}),
		giu.Dummy(0, 20),
		// Output
		giu.Label("FormData output: " + f.Output),
	}
}

func (f *FormData) FormDataSubmit() {
	if f.Service == "" {
		f.Output = "Please enter a service name."
	} else if f.Pwd == "" {
		f.Output = fmt.Sprintf("Please enter a password for %s", f.Service)
	} else {
		f.Output = fmt.Sprintf("Submitted! Service: %s, Password: %s", f.Service, f.Pwd)
	}
	giu.Update()
}
```

