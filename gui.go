/*
 * Copyright (c) 2017 Alex Yatskov <alex@foosoft.net>
 * Author: Alex Yatskov <alex@foosoft.net>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"github.com/andlabs/ui"
)

func gui() error {
	return ui.Main(func() {
		pathEntry := ui.NewEntry()
		browseButton := ui.NewButton("Browse")
		pathBox := ui.NewHorizontalBox()
		pathBox.Append(pathEntry, true)
		pathBox.Append(browseButton, false)

		portSpin := ui.NewSpinbox(0, 65535)
		portSpin.SetValue(9876)

		formatCombo := ui.NewCombobox()
		formatCombo.Append("EPWING")
		formatCombo.Append("EDICT")
		formatCombo.Append("ENAMDICT")
		formatCombo.Append("KANJIDIC")
		formatCombo.SetSelected(0)

		titleEntry := ui.NewEntry()
		outputLabel := ui.NewLabel("")
		importButton := ui.NewButton("Import...")

		mainBox := ui.NewVerticalBox()
		mainBox.Append(ui.NewLabel("Path to dictionary source (CATALOGS file for EPWING):"), false)
		mainBox.Append(pathBox, false)
		mainBox.Append(ui.NewLabel("Network port for extension server:"), false)
		mainBox.Append(portSpin, false)
		mainBox.Append(ui.NewLabel("Dictionary title (leave blank for default):"), false)
		mainBox.Append(titleEntry, false)
		mainBox.Append(ui.NewLabel("Dictionary format:"), false)
		mainBox.Append(formatCombo, false)
		mainBox.Append(ui.NewLabel("Application output:"), false)
		mainBox.Append(outputLabel, true)
		mainBox.Append(importButton, false)

		window := ui.NewWindow("Yomichan Import", 640, 480, false)
		window.SetMargined(true)
		window.SetChild(mainBox)

		browseButton.OnClicked(func(*ui.Button) {
			if path := ui.OpenFile(window); len(path) > 0 {
				pathEntry.SetText(path)
			}
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
	})
}
